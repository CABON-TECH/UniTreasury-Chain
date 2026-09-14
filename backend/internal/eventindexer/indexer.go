// Package eventindexer polls the chain for contract events, decodes them, and
// persists them to the audit_events table via AuditRepository.
// It is idempotent (uses (tx_hash, log_index) as the unique key) and resumable
// (persists the last indexed block to the DB via checkpoint.go).
package eventindexer

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"

	"github.com/cabon-tech/unitreasury-chain/backend/internal/blockchain"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/domain"
)

// Indexer polls for events from one or more contracts and writes them to the DB.
type Indexer struct {
	client       *blockchain.Client
	auditRepo    domain.AuditRepository
	contracts    []IndexedContract
	pollInterval time.Duration
	batchSize    uint64
	log          *zap.Logger
}

// IndexedContract describes a contract to index: its address, name (used as
// the "contract" label in audit_events), and the set of event topic hashes to
// capture.
type IndexedContract struct {
	Name     string
	Address  common.Address
	Topics   []common.Hash // event signature hashes to filter on
}

// New creates an Indexer.
func New(
	client *blockchain.Client,
	auditRepo domain.AuditRepository,
	contracts []IndexedContract,
	pollInterval time.Duration,
	batchSize uint64,
	log *zap.Logger,
) *Indexer {
	return &Indexer{
		client:       client,
		auditRepo:    auditRepo,
		contracts:    contracts,
		pollInterval: pollInterval,
		batchSize:    batchSize,
		log:          log,
	}
}

// Run starts the indexing loop. It blocks until ctx is cancelled.
// Call this in a goroutine from the worker binary.
func (idx *Indexer) Run(ctx context.Context) error {
	ticker := time.NewTicker(idx.pollInterval)
	defer ticker.Stop()

	idx.log.Info("event indexer started",
		zap.Duration("poll_interval", idx.pollInterval),
		zap.Uint64("batch_size", idx.batchSize),
	)

	for {
		select {
		case <-ctx.Done():
			idx.log.Info("event indexer stopping")
			return ctx.Err()
		case <-ticker.C:
			if err := idx.poll(ctx); err != nil {
				idx.log.Error("indexer poll error", zap.Error(err))
				// Continue — transient RPC errors should not kill the loop
			}
		}
	}
}

func (idx *Indexer) poll(ctx context.Context) error {
	head, err := idx.client.CurrentBlock(ctx)
	if err != nil {
		return fmt.Errorf("indexer: get current block: %w", err)
	}

	for _, c := range idx.contracts {
		if err := idx.indexContract(ctx, c, head); err != nil {
			idx.log.Error("indexer: contract poll error",
				zap.String("contract", c.Name),
				zap.Error(err),
			)
		}
	}
	return nil
}

func (idx *Indexer) indexContract(ctx context.Context, c IndexedContract, head uint64) error {
	lastIndexed, err := idx.auditRepo.GetLastIndexedBlock(ctx, c.Name)
	if err != nil {
		return fmt.Errorf("indexer: get last indexed block for %s: %w", c.Name, err)
	}
	
	// If starting fresh on a deployed chain, do not start at 0 to avoid pruned nodes.
	if lastIndexed == 0 && head > 1000 {
		lastIndexed = head - 1000
	}

	from := lastIndexed + 1
	if from > head {
		return nil // already up to date
	}

	to := min(from+idx.batchSize-1, head)

	idx.log.Debug("indexing contract events",
		zap.String("contract", c.Name),
		zap.Uint64("from", from),
		zap.Uint64("to", to),
	)

	logs, err := idx.client.Inner().FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: numberToHex(from),
		ToBlock:   numberToHex(to),
		Addresses: []common.Address{c.Address},
		Topics:    [][]common.Hash{c.Topics},
	})
	if err != nil {
		return fmt.Errorf("indexer: filter logs %s [%d-%d]: %w", c.Name, from, to, err)
	}

	for _, l := range logs {
		if err := idx.processLog(ctx, c.Name, l); err != nil {
			idx.log.Error("indexer: process log error",
				zap.String("tx", l.TxHash.Hex()),
				zap.Error(err),
			)
		}
	}

	if err := idx.auditRepo.SetLastIndexedBlock(ctx, c.Name, to); err != nil {
		return fmt.Errorf("indexer: set checkpoint for %s: %w", c.Name, err)
	}

	return nil
}

func (idx *Indexer) processLog(ctx context.Context, contractName string, l types.Log) error {
	// Derive event name from the first topic (event signature hash)
	eventName := l.Topics[0].Hex() // In Sprint 2 this will be resolved via ABI

	// Payload is the raw log data — the Postgres JSONB column stores hex-encoded topics + data
	payload := encodeLogPayload(l)

	event := &domain.AuditEvent{
		Contract:    contractName,
		EventName:   eventName,
		TxHash:      l.TxHash.Hex(),
		BlockNumber: l.BlockNumber,
		LogIndex:    l.Index,
		Payload:     payload,
		IndexedAt:   time.Now(),
	}

	return idx.auditRepo.Insert(ctx, event)
}

// encodeLogPayload produces a simple JSON-compatible byte slice from a log.
func encodeLogPayload(l types.Log) []byte {
	return []byte(fmt.Sprintf(
		`{"address":%q,"topics":%s,"data":%q,"block_number":%d,"tx_hash":%q,"log_index":%d}`,
		l.Address.Hex(),
		topicsJSON(l.Topics),
		common.Bytes2Hex(l.Data),
		l.BlockNumber,
		l.TxHash.Hex(),
		l.Index,
	))
}

func topicsJSON(topics []common.Hash) string {
	out := "["
	for i, t := range topics {
		if i > 0 {
			out += ","
		}
		out += fmt.Sprintf("%q", t.Hex())
	}
	return out + "]"
}

func numberToHex(n uint64) *big.Int {
	return new(big.Int).SetUint64(n)
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
