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
type ZKService struct {
	txMgr    *blockchain.TxManager
	registry *bindings.ZKEnrollmentRegistry
	log      *zap.Logger
}
func NewZKService(txMgr *blockchain.TxManager, registry *bindings.ZKEnrollmentRegistry, log *zap.Logger) *ZKService {
	return &ZKService{txMgr: txMgr, registry: registry, log: log}
}
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
func (s *ZKService) GetCurrentRoot(ctx context.Context) ([32]byte, error) {
	root, err := s.registry.EnrollmentMerkleRoot(&bind.CallOpts{Context: ctx})
	if err != nil {
		return [32]byte{}, fmt.Errorf("get root: %w", err)
	}
	return root, nil
}
func (s *ZKService) IsEnrollmentVerified(ctx context.Context, nullifierHash [32]byte) (common.Address, error) {
	addr, err := s.registry.IsEnrollmentVerified(&bind.CallOpts{Context: ctx}, nullifierHash)
	if err != nil {
		return common.Address{}, fmt.Errorf("isEnrollmentVerified: %w", err)
	}
	return addr, nil
}
func ComputeCommitment(studentID string) [32]byte {
	hash := crypto.Keccak256Hash([]byte(studentID))
	return hash
}
