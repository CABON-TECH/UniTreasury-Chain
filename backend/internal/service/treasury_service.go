package service

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"

	"github.com/cabon-tech/unitreasury-chain/backend/internal/blockchain"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/blockchain/bindings"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/domain"
)

type TreasuryService struct {
	repo      domain.ProposalRepository
	txMgr     *blockchain.TxManager
	treasury  *bindings.TreasuryContract
	log       *zap.Logger
}

func NewTreasuryService(
	repo domain.ProposalRepository,
	txMgr *blockchain.TxManager,
	treasury *bindings.TreasuryContract,
	log *zap.Logger,
) *TreasuryService {
	return &TreasuryService{
		repo:     repo,
		txMgr:    txMgr,
		treasury: treasury,
		log:      log,
	}
}

type ProposeInput struct {
	Proposer  string
	Recipient string
	Amount    uint64
	Purpose   string
}

func (s *TreasuryService) ProposeWithdrawal(ctx context.Context, in ProposeInput) (*domain.WithdrawalProposal, error) {
	p := &domain.WithdrawalProposal{
		Proposer:  common.HexToAddress(in.Proposer),
		Recipient: common.HexToAddress(in.Recipient),
		Amount:    in.Amount,
		Purpose:   in.Purpose,
		Status:    domain.ProposalStatusPending,
	}
	if err := s.repo.Create(ctx, p); err != nil {
		return nil, fmt.Errorf("treasury_service: create proposal: %w", err)
	}

	opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)
	if err != nil {
		return nil, fmt.Errorf("transact opts: %w", err)
	}

	tx, err := s.treasury.TreasuryContractTransactor.ProposeWithdrawal(
		opts,
		common.HexToAddress(in.Recipient),
		new(big.Int).SetUint64(in.Amount),
		in.Purpose,
	)
	if err != nil {
		rollback()
		return nil, fmt.Errorf("propose tx: %w", err)
	}
	confirm()
	
	p.TxHash = tx.Hash().Hex()
	_ = s.repo.UpdateStatus(ctx, p.ID, domain.ProposalStatusPending, p.TxHash)

	s.log.Info("withdrawal proposed", zap.Int64("proposal_id", p.ID), zap.String("tx_hash", p.TxHash))
	return p, nil
}

func (s *TreasuryService) ApproveWithdrawal(ctx context.Context, proposalID int64) error {
	p, err := s.repo.GetByID(ctx, proposalID)
	if err != nil || p == nil {
		return fmt.Errorf("proposal not found")
	}
	if p.OnChainID == 0 {
		return fmt.Errorf("proposal not yet indexed on-chain")
	}

	opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)
	if err != nil {
		return err
	}

	_, err = s.treasury.TreasuryContractTransactor.ApproveWithdrawal(opts, new(big.Int).SetUint64(p.OnChainID))
	if err != nil {
		rollback()
		return fmt.Errorf("approve tx: %w", err)
	}
	confirm()
	
	s.log.Info("withdrawal approved", zap.Int64("proposal_id", p.ID))
	return nil
}

func (s *TreasuryService) ExecuteWithdrawal(ctx context.Context, proposalID int64) error {
	p, err := s.repo.GetByID(ctx, proposalID)
	if err != nil || p == nil {
		return fmt.Errorf("proposal not found")
	}
	if p.OnChainID == 0 {
		return fmt.Errorf("proposal not yet indexed on-chain")
	}

	opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)
	if err != nil {
		return err
	}

	tx, err := s.treasury.TreasuryContractTransactor.ExecuteWithdrawal(opts, new(big.Int).SetUint64(p.OnChainID))
	if err != nil {
		rollback()
		return fmt.Errorf("execute tx: %w", err)
	}
	confirm()
	
	_ = s.repo.UpdateStatus(ctx, p.ID, domain.ProposalStatusExecuted, tx.Hash().Hex())
	s.log.Info("withdrawal executed", zap.Int64("proposal_id", p.ID))
	return nil
}

func (s *TreasuryService) CancelWithdrawal(ctx context.Context, proposalID int64) error {
	p, err := s.repo.GetByID(ctx, proposalID)
	if err != nil || p == nil {
		return fmt.Errorf("proposal not found")
	}
	if p.OnChainID == 0 {
		return fmt.Errorf("proposal not yet indexed on-chain")
	}

	opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)
	if err != nil {
		return err
	}

	tx, err := s.treasury.TreasuryContractTransactor.CancelWithdrawal(opts, new(big.Int).SetUint64(p.OnChainID))
	if err != nil {
		rollback()
		return fmt.Errorf("cancel tx: %w", err)
	}
	confirm()
	
	_ = s.repo.UpdateStatus(ctx, p.ID, domain.ProposalStatusCancelled, tx.Hash().Hex())
	s.log.Info("withdrawal cancelled", zap.Int64("proposal_id", p.ID))
	return nil
}

func (s *TreasuryService) ListProposals(ctx context.Context, page, pageSize int) ([]*domain.WithdrawalProposal, int64, error) {
	if page < 1 { page = 1 }
	if pageSize <= 0 || pageSize > 100 { pageSize = 20 }
	return s.repo.List(ctx, (page-1)*pageSize, pageSize)
}
