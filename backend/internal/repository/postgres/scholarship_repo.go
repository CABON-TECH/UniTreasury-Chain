package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cabon-tech/unitreasury-chain/backend/internal/domain"
)

type ScholarshipRepo struct {
	pool *pgxpool.Pool
}

func NewScholarshipRepo(pool *pgxpool.Pool) *ScholarshipRepo {
	return &ScholarshipRepo{pool: pool}
}

func (r *ScholarshipRepo) CreateFund(ctx context.Context, f *domain.ScholarshipFund) error {
	const q = `
		INSERT INTO scholarship_funds
		  (on_chain_id, sponsor, total_amount, released_amount, tranche_count, tranche_amount, paused, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`
	f.CreatedAt = time.Now()
	return r.pool.QueryRow(ctx, q,
		f.OnChainID, f.Sponsor, f.TotalAmount, f.ReleasedAmount, f.TrancheCount, f.TrancheAmount, f.Paused, f.CreatedAt,
	).Scan(&f.ID)
}

func (r *ScholarshipRepo) GetFundByID(ctx context.Context, id int64) (*domain.ScholarshipFund, error) {
	const q = `SELECT id, on_chain_id, sponsor, total_amount, released_amount, tranche_count, tranche_amount, paused, created_at FROM scholarship_funds WHERE id = $1`
	return r.scanFund(r.pool.QueryRow(ctx, q, id))
}

func (r *ScholarshipRepo) GetFundByOnChainID(ctx context.Context, onChainID uint64) (*domain.ScholarshipFund, error) {
	const q = `SELECT id, on_chain_id, sponsor, total_amount, released_amount, tranche_count, tranche_amount, paused, created_at FROM scholarship_funds WHERE on_chain_id = $1`
	return r.scanFund(r.pool.QueryRow(ctx, q, onChainID))
}

func (r *ScholarshipRepo) ListActiveFunds(ctx context.Context) ([]*domain.ScholarshipFund, error) {
	const q = `SELECT id, on_chain_id, sponsor, total_amount, released_amount, tranche_count, tranche_amount, paused, created_at FROM scholarship_funds WHERE paused = false AND released_amount < total_amount`
	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var funds []*domain.ScholarshipFund
	for rows.Next() {
		f, err := r.scanFund(rows)
		if err != nil {
			return nil, err
		}
		funds = append(funds, f)
	}
	return funds, nil
}

func (r *ScholarshipRepo) UpdateFundReleasedAmount(ctx context.Context, fundID int64, releasedAmount uint64) error {
	const q = `UPDATE scholarship_funds SET released_amount = $1 WHERE id = $2`
	_, err := r.pool.Exec(ctx, q, releasedAmount, fundID)
	return err
}

func (r *ScholarshipRepo) CreateTrancheRelease(ctx context.Context, t *domain.TrancheRelease) error {
	const q = `
		INSERT INTO tranche_releases
		  (fund_id, on_chain_fund_id, student_hash, tranche_index, amount, recipient, tx_hash, block_number, released_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`
	t.ReleasedAt = time.Now()
	var txHash *string
	if t.TxHash != "" {
		txHash = &t.TxHash
	}
	var blockNum *uint64
	if t.BlockNumber > 0 {
		blockNum = &t.BlockNumber
	}
	return r.pool.QueryRow(ctx, q,
		t.FundID, t.OnChainFundID, t.StudentHash, t.TrancheIndex, t.Amount, t.Recipient, txHash, blockNum, t.ReleasedAt,
	).Scan(&t.ID)
}

func (r *ScholarshipRepo) HasReleased(ctx context.Context, fundID int64, studentHash string, trancheIndex int) (bool, error) {
	const q = `SELECT 1 FROM tranche_releases WHERE fund_id = $1 AND student_hash = $2 AND tranche_index = $3`
	var dummy int
	err := r.pool.QueryRow(ctx, q, fundID, studentHash, trancheIndex).Scan(&dummy)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *ScholarshipRepo) scanFund(row pgx.Row) (*domain.ScholarshipFund, error) {
	f := &domain.ScholarshipFund{}
	var onChainID *uint64
	err := row.Scan(
		&f.ID, &onChainID, &f.Sponsor, &f.TotalAmount, &f.ReleasedAmount,
		&f.TrancheCount, &f.TrancheAmount, &f.Paused, &f.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if onChainID != nil {
		f.OnChainID = *onChainID
	}
	return f, nil
}
