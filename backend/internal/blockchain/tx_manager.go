package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

// TxManager serializes all outbound transactions from one signer address,
// managing nonce sequencing to prevent the classic multi-goroutine nonce race.
// All contract interactions (recordPayment, approveWithdrawal, releaseTranche)
// must go through TxManager to avoid nonce conflicts.
type TxManager struct {
	client     *Client
	privateKey *ecdsa.PrivateKey
	address    common.Address
	mu         sync.Mutex // guards nonce
	nonce      uint64
	log        *zap.Logger
}

// NewTxManager creates a TxManager, fetching the current nonce from the chain.
func NewTxManager(ctx context.Context, client *Client, privateKey *ecdsa.PrivateKey, log *zap.Logger) (*TxManager, error) {
	addr := AddressFromKey(privateKey)
	nonce, err := client.inner.PendingNonceAt(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("tx_manager: fetch nonce for %s: %w", addr.Hex(), err)
	}

	log.Info("TxManager initialized",
		zap.String("address", addr.Hex()),
		zap.Uint64("starting_nonce", nonce),
	)

	return &TxManager{
		client:     client,
		privateKey: privateKey,
		address:    addr,
		nonce:      nonce,
		log:        log,
	}, nil
}

// TransactOpts returns a TransactOpts with the next available nonce.
// The caller MUST invoke ConfirmNonce or RollbackNonce after the transaction
// is submitted (or fails), otherwise the nonce sequence will be corrupted.
//
// The mutex is held between AcquireNonce and Confirm/Rollback — this means
// only one transaction can be in-flight at a time per TxManager instance,
// which is intentional for a research prototype (avoids gap-nonce complexity).
func (m *TxManager) TransactOpts(ctx context.Context) (*bind.TransactOpts, func(), func(), error) {
	m.mu.Lock()

	opts, err := bind.NewKeyedTransactorWithChainID(m.privateKey, m.client.chainID)
	if err != nil {
		m.mu.Unlock()
		return nil, nil, nil, fmt.Errorf("tx_manager: create transactor: %w", err)
	}
	opts.Context = ctx
	opts.Nonce = big.NewInt(int64(m.nonce))

	confirm := func() {
		m.nonce++
		m.mu.Unlock()
		m.log.Debug("nonce confirmed", zap.Uint64("new_nonce", m.nonce))
	}

	rollback := func() {
		m.mu.Unlock()
		m.log.Warn("nonce rolled back — transaction was not submitted")
	}

	return opts, confirm, rollback, nil
}

// SyncNonce re-fetches the pending nonce from the chain.
// Call this if you suspect nonce drift (e.g. after a process restart).
func (m *TxManager) SyncNonce(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	nonce, err := m.client.inner.PendingNonceAt(ctx, m.address)
	if err != nil {
		return fmt.Errorf("tx_manager: sync nonce: %w", err)
	}
	m.log.Info("nonce synced",
		zap.Uint64("old_nonce", m.nonce),
		zap.Uint64("new_nonce", nonce),
	)
	m.nonce = nonce
	return nil
}

// WaitMined polls until a transaction is included in a block.
// Returns the receipt or an error if the context is cancelled.
func (m *TxManager) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	return bind.WaitMined(ctx, m.client.inner, tx)
}

// Address returns the signer address managed by this TxManager.
func (m *TxManager) Address() common.Address {
	return m.address
}
