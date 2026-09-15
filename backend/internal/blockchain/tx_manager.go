package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sync"
	"time"

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
	nonce, err := client.PendingNonceAt(ctx, addr)
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

	nonce, err := m.client.PendingNonceAt(ctx, m.address)
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
	return bind.WaitMined(ctx, m.client, tx)
}

// Address returns the signer address managed by this TxManager.
func (m *TxManager) Address() common.Address {
	return m.address
}

// WaitMinedWithBump polls until a transaction is included in a block.
// If it takes more than 15 seconds, it fetches new gas prices and resubmits
// a bumped transaction to replace the stuck one.
func (m *TxManager) WaitMinedWithBump(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	currentTx := tx

	bumpTicker := time.NewTicker(15 * time.Second)
	defer bumpTicker.Stop()

	pollTicker := time.NewTicker(2 * time.Second)
	defer pollTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()

		case <-pollTicker.C:
			receipt, err := m.client.TransactionReceipt(ctx, currentTx.Hash())
			if err == nil && receipt != nil {
				m.log.Info("transaction mined", zap.String("hash", currentTx.Hash().Hex()))
				return receipt, nil
			}

		case <-bumpTicker.C:
			m.log.Info("transaction pending too long, attempting gas bump", zap.String("hash", currentTx.Hash().Hex()))

			signer := types.LatestSignerForChainID(m.client.chainID)

			if currentTx.Type() == types.LegacyTxType {
				gasPrice, err := m.client.SuggestGasPrice(ctx)
				if err != nil {
					m.log.Warn("failed to suggest gas price", zap.Error(err))
					continue
				}

				oldGasPrice := currentTx.GasPrice()
				minRequiredGasPrice := new(big.Int).Mul(oldGasPrice, big.NewInt(120))
				minRequiredGasPrice.Div(minRequiredGasPrice, big.NewInt(100))

				if gasPrice.Cmp(minRequiredGasPrice) < 0 {
					gasPrice = minRequiredGasPrice
				}

				newTx := types.NewTx(&types.LegacyTx{
					Nonce:    currentTx.Nonce(),
					To:       currentTx.To(),
					Value:    currentTx.Value(),
					Gas:      currentTx.Gas(),
					GasPrice: gasPrice,
					Data:     currentTx.Data(),
				})

				signedTx, err := types.SignTx(newTx, signer, m.privateKey)
				if err != nil {
					m.log.Error("failed to sign bumped tx", zap.Error(err))
					continue
				}

				err = m.client.SendTransaction(ctx, signedTx)
				if err != nil {
					m.log.Error("failed to send bumped tx", zap.Error(err))
					continue
				}

				m.log.Info("successfully bumped legacy transaction",
					zap.String("old_hash", currentTx.Hash().Hex()),
					zap.String("new_hash", signedTx.Hash().Hex()),
				)
				currentTx = signedTx

			} else if currentTx.Type() == types.DynamicFeeTxType {
				tip, err := m.client.SuggestGasTipCap(ctx)
				if err != nil {
					m.log.Warn("failed to suggest gas tip cap", zap.Error(err))
					continue
				}

				head, err := m.client.HeaderByNumber(ctx, nil)
				if err != nil {
					m.log.Warn("failed to get header", zap.Error(err))
					continue
				}

				baseFee := head.BaseFee
				if baseFee == nil {
					baseFee = big.NewInt(0)
				}
				fee := new(big.Int).Add(tip, new(big.Int).Mul(baseFee, big.NewInt(2)))

				oldTip := currentTx.GasTipCap()
				oldFee := currentTx.GasFeeCap()

				minRequiredTip := new(big.Int).Mul(oldTip, big.NewInt(120))
				minRequiredTip.Div(minRequiredTip, big.NewInt(100))

				minRequiredFee := new(big.Int).Mul(oldFee, big.NewInt(120))
				minRequiredFee.Div(minRequiredFee, big.NewInt(100))

				if tip.Cmp(minRequiredTip) < 0 {
					tip = minRequiredTip
				}
				if fee.Cmp(minRequiredFee) < 0 {
					fee = minRequiredFee
				}

				newTx := types.NewTx(&types.DynamicFeeTx{
					ChainID:   m.client.chainID,
					Nonce:     currentTx.Nonce(),
					To:        currentTx.To(),
					Value:     currentTx.Value(),
					Gas:       currentTx.Gas(),
					GasTipCap: tip,
					GasFeeCap: fee,
					Data:      currentTx.Data(),
				})

				signedTx, err := types.SignTx(newTx, signer, m.privateKey)
				if err != nil {
					m.log.Error("failed to sign bumped tx", zap.Error(err))
					continue
				}

				err = m.client.SendTransaction(ctx, signedTx)
				if err != nil {
					m.log.Error("failed to send bumped tx", zap.Error(err))
					continue
				}

				m.log.Info("successfully bumped dynamic transaction",
					zap.String("old_hash", currentTx.Hash().Hex()),
					zap.String("new_hash", signedTx.Hash().Hex()),
				)
				currentTx = signedTx
			} else {
				m.log.Warn("unsupported tx type for bumping", zap.Uint8("type", currentTx.Type()))
			}
		}
	}
}
