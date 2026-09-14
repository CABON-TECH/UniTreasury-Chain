package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cabon-tech/unitreasury-chain/backend/internal/domain"
	"github.com/ethereum/go-ethereum/common"
)

type ProposalRepo struct {
	pool *pgxpool.Pool
}

func NewProposalRepo(pool *pgxpool.Pool) *ProposalRepo {
	return &ProposalRepo{pool: pool}
}

func (r *ProposalRepo) Create(ctx context.Context, p *domain.WithdrawalProposal) error {
	const q = `
		INSERT INTO withdrawal_proposals
		  (on_chain_id, proposer, recipient, amount, purpose, status, approval_count, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`
	
	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = now

	return r.pool.QueryRow(ctx, q,
		p.OnChainID, p.Proposer.Hex(), p.Recipient.Hex(), p.Amount, p.Purpose,
		p.Status, p.ApprovalCount, p.CreatedAt, p.UpdatedAt,
	).Scan(&p.ID)
}

func (r *ProposalRepo) GetByID(ctx context.Context, id int64) (*domain.WithdrawalProposal, error) {
	const q = `
		SELECT id, on_chain_id, proposer, recipient, amount, purpose, status, approval_count, tx_hash, block_number, created_at, updated_at
		FROM withdrawal_proposals WHERE id = $1`
	return r.scanProposal(r.pool.QueryRow(ctx, q, id))
}

func (r *ProposalRepo) GetByOnChainID(ctx context.Context, onChainID uint64) (*domain.WithdrawalProposal, error) {
	const q = `
		SELECT id, on_chain_id, proposer, recipient, amount, purpose, status, approval_count, tx_hash, block_number, created_at, updated_at
		FROM withdrawal_proposals WHERE on_chain_id = $1`
	return r.scanProposal(r.pool.QueryRow(ctx, q, onChainID))
}

func (r *ProposalRepo) List(ctx context.Context, offset, limit int) ([]*domain.WithdrawalProposal, int64, error) {
	var total int64
	if err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM withdrawal_proposals`).Scan(&total); err != nil {
		return nil, 0, err
	}

	const q = `
		SELECT id, on_chain_id, proposer, recipient, amount, purpose, status, approval_count, tx_hash, block_number, created_at, updated_at
		FROM withdrawal_proposals ORDER BY id DESC LIMIT $1 OFFSET $2`
	
	rows, err := r.pool.Query(ctx, q, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var proposals []*domain.WithdrawalProposal
	for rows.Next() {
		p, err := r.scanProposal(rows)
		if err != nil {
			return nil, 0, err
		}
		proposals = append(proposals, p)
	}
	return proposals, total, rows.Err()
}

func (r *ProposalRepo) UpdateStatus(ctx context.Context, id int64, status domain.ProposalStatus, txHash string) error {
	const q = `UPDATE withdrawal_proposals SET status = $1, tx_hash = $2, updated_at = NOW() WHERE id = $3`
	_, err := r.pool.Exec(ctx, q, status, txHash, id)
	return err
}

func (r *ProposalRepo) UpdateApprovalCount(ctx context.Context, onChainID uint64, count int) error {
	const q = `UPDATE withdrawal_proposals SET approval_count = $1, updated_at = NOW() WHERE on_chain_id = $2`
	_, err := r.pool.Exec(ctx, q, count, onChainID)
	return err
}

func (r *ProposalRepo) scanProposal(row pgx.Row) (*domain.WithdrawalProposal, error) {
	p := &domain.WithdrawalProposal{}
	var txHash *string
	var blockNumber *uint64
	var onChainID *uint64
	var proposer string
	var recipient string
	err := row.Scan(
		&p.ID, &onChainID, &proposer, &recipient, &p.Amount, &p.Purpose,
		&p.Status, &p.ApprovalCount, &txHash, &blockNumber, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if onChainID != nil {
		p.OnChainID = *onChainID
	}
	if txHash != nil {
		p.TxHash = *txHash
	}
	if blockNumber != nil {
		p.BlockNumber = *blockNumber
	}
	// We convert string addresses back to common.Address (but for DB struct mapping this is annoying, let's keep it simple for Sprint 3. The struct models.go defines Proposer as common.Address but let's just cheat and not assign it for now or implement a quick wrapper).
	// Actually models.go defines Proposer and Recipient as common.Address.
	// We need to import "github.com/ethereum/go-ethereum/common" and convert it.
	p.Proposer = common.HexToAddress(proposer)
	p.Recipient = common.HexToAddress(recipient)
	return p, nil
}
