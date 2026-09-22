package eventindexer
import (
	"context"
	"fmt"
	"math/big"
	"encoding/json"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/blockchain"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/domain"
)
type Indexer struct {
	client       *blockchain.Client
	auditRepo    domain.AuditRepository
	contracts    []IndexedContract
	pool         *pgxpool.Pool
	pollInterval time.Duration
	batchSize    uint64
	log          *zap.Logger
}
type IndexedContract struct {
	Name     string
	Address  common.Address
	Topics   []common.Hash 
}
func New(
	client *blockchain.Client,
	auditRepo domain.AuditRepository,
	contracts []IndexedContract,
	pool *pgxpool.Pool,
	pollInterval time.Duration,
	batchSize uint64,
	log *zap.Logger,
) *Indexer {
	return &Indexer{
		client:       client,
		auditRepo:    auditRepo,
		contracts:    contracts,
		pool:         pool,
		pollInterval: pollInterval,
		batchSize:    batchSize,
		log:          log,
	}
}
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
			}
		}
	}
}
func (idx *Indexer) poll(ctx context.Context) error {
	head, err := idx.client.BlockNumber(ctx)
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
	if lastIndexed == 0 && head > 1000 {
		lastIndexed = head - 1000
	}
	from := lastIndexed + 1
	if from > head {
		return nil 
	}
	to := min(from+idx.batchSize-1, head)
	idx.log.Debug("indexing contract events",
		zap.String("contract", c.Name),
		zap.Uint64("from", from),
		zap.Uint64("to", to),
	)
	logs, err := idx.client.FilterLogs(ctx, ethereum.FilterQuery{
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
	eventName := l.Topics[0].Hex() 
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
	err := idx.auditRepo.Insert(ctx, event)
	if err == nil && idx.pool != nil {
		payloadMap := map[string]interface{}{
			"type": "BlockchainEvent",
			"contract": contractName,
			"tx_hash": l.TxHash.Hex(),
		}
		bMsg, _ := json.Marshal(payloadMap)
		idx.pool.Exec(ctx, "NOTIFY ws_events, '" + string(bMsg) + "'")
	}
	return err
}
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
