package service

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"

	"github.com/cabon-tech/unitreasury-chain/backend/internal/blockchain"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/blockchain/bindings"
)

// ZKService handles ZK enrollment proof operations.
type ZKService struct {
	txMgr    *blockchain.TxManager
	registry *bindings.ZKEnrollmentRegistry
	log      *zap.Logger
}

func NewZKService(txMgr *blockchain.TxManager, registry *bindings.ZKEnrollmentRegistry, log *zap.Logger) *ZKService {
	return &ZKService{txMgr: txMgr, registry: registry, log: log}
}

// UpdateMerkleRoot updates the on-chain enrollment Merkle root (admin only).
func (s *ZKService) UpdateMerkleRoot(ctx context.Context, newRoot [32]byte) (string, error) {
	opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)
	if err != nil {
		return "", err
	}

	tx, err := s.registry.UpdateMerkleRoot(opts, newRoot)
	if err != nil {
		rollback()
		return "", fmt.Errorf("updateMerkleRoot tx: %w", err)
	}
	confirm()

	s.log.Info("merkle root updated", zap.String("tx", tx.Hash().Hex()))
	return tx.Hash().Hex(), nil
}

// ProveEnrollment submits a ZK proof on-chain to verify a student is enrolled.
func (s *ZKService) ProveEnrollment(ctx context.Context, proof []byte, nullifierHash [32]byte, claimedRoot [32]byte) (string, error) {
	opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)
	if err != nil {
		return "", err
	}

	tx, err := s.registry.ProveEnrollment(opts, proof, nullifierHash, claimedRoot)
	if err != nil {
		rollback()
		return "", fmt.Errorf("proveEnrollment tx: %w", err)
	}
	confirm()

	s.log.Info("enrollment proof submitted", zap.String("nullifier", fmt.Sprintf("%x", nullifierHash)))
	return tx.Hash().Hex(), nil
}

// GetCurrentRoot returns the current on-chain Merkle root.
func (s *ZKService) GetCurrentRoot(ctx context.Context) ([32]byte, error) {
	root, err := s.registry.EnrollmentMerkleRoot(&bind.CallOpts{Context: ctx})
	if err != nil {
		return [32]byte{}, fmt.Errorf("get root: %w", err)
	}
	return root, nil
}

// IsEnrollmentVerified checks if a nullifier has been used to prove enrollment.
func (s *ZKService) IsEnrollmentVerified(ctx context.Context, nullifierHash [32]byte) (common.Address, error) {
	addr, err := s.registry.IsEnrollmentVerified(&bind.CallOpts{Context: ctx}, nullifierHash)
	if err != nil {
		return common.Address{}, fmt.Errorf("isEnrollmentVerified: %w", err)
	}
	return addr, nil
}

// ComputeCommitment computes a keccak256 commitment of a student ID string (off-chain helper).
// In production this would be done inside a ZK circuit.
func ComputeCommitment(studentID string) [32]byte {
	hash := crypto.Keccak256Hash([]byte(studentID))
	return hash
}
