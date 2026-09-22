package postgres
import (
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/domain"
)
type PaymentRepo struct {
	pool *pgxpool.Pool
}
func NewPaymentRepo(pool *pgxpool.Pool) *PaymentRepo {
	return &PaymentRepo{pool: pool}
}
func (r *PaymentRepo) Create(ctx context.Context, p *domain.Payment) error {
	const q = `
		INSERT INTO payments
		  (student_id, student_hash, receipt_hash, amount, semester, fee_type, status, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING id`
	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = now
	p.Status = domain.PaymentStatusPending
	return r.pool.QueryRow(ctx, q,
		p.StudentID, p.StudentHash, p.ReceiptHash,
		p.Amount, p.Semester, p.FeeType, p.Status,
		p.CreatedAt, p.UpdatedAt,
	).Scan(&p.ID)
}
func (r *PaymentRepo) GetByID(ctx context.Context, id int64) (*domain.Payment, error) {
	const q = `SELECT id, student_id, student_hash, receipt_hash, amount, semester,
	                  fee_type, status, tx_hash, block_number, created_at, updated_at
	           FROM payments WHERE id = $1`
	return r.scanPayment(r.pool.QueryRow(ctx, q, id))
}
func (r *PaymentRepo) GetByReceiptHash(ctx context.Context, receiptHash string) (*domain.Payment, error) {
	const q = `SELECT id, student_id, student_hash, receipt_hash, amount, semester,
	                  fee_type, status, tx_hash, block_number, created_at, updated_at
	           FROM payments WHERE receipt_hash = $1`
	return r.scanPayment(r.pool.QueryRow(ctx, q, receiptHash))
}
func (r *PaymentRepo) ListByStudent(ctx context.Context, studentID int64, offset, limit int) ([]*domain.Payment, int64, error) {
	const countQ = `SELECT COUNT(*) FROM payments WHERE student_id = $1`
	var total int64
	if err := r.pool.QueryRow(ctx, countQ, studentID).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("payment_repo: count by student: %w", err)
	}
	const q = `SELECT id, student_id, student_hash, receipt_hash, amount, semester,
	                  fee_type, status, tx_hash, block_number, created_at, updated_at
	           FROM payments WHERE student_id = $1 ORDER BY id DESC LIMIT $2 OFFSET $3`
	return r.listPayments(ctx, q, studentID, limit, offset)
}
func (r *PaymentRepo) ListBySemester(ctx context.Context, semester int, offset, limit int) ([]*domain.Payment, int64, error) {
	const countQ = `SELECT COUNT(*) FROM payments WHERE semester = $1`
	var total int64
	if err := r.pool.QueryRow(ctx, countQ, semester).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("payment_repo: count by semester: %w", err)
	}
	const q = `SELECT id, student_id, student_hash, receipt_hash, amount, semester,
	                  fee_type, status, tx_hash, block_number, created_at, updated_at
	           FROM payments WHERE semester = $1 ORDER BY id DESC LIMIT $2 OFFSET $3`
	return r.listPayments(ctx, q, semester, limit, offset)
}
func (r *PaymentRepo) UpdateStatus(ctx context.Context, id int64, status domain.PaymentStatus, txHash string) error {
	const q = `UPDATE payments SET status = $1, tx_hash = $2, updated_at = NOW() WHERE id = $3`
	_, err := r.pool.Exec(ctx, q, status, txHash, id)
	if err != nil {
		return fmt.Errorf("payment_repo: update status: %w", err)
	}
	return nil
}
func (r *PaymentRepo) UpdateConfirmed(ctx context.Context, id int64, blockNumber uint64) error {
	const q = `UPDATE payments SET status = 'confirmed', block_number = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.pool.Exec(ctx, q, blockNumber, id)
	if err != nil {
		return fmt.Errorf("payment_repo: update confirmed: %w", err)
	}
	return nil
}
func (r *PaymentRepo) GetTotalVolumeForSemester(ctx context.Context, semester int) (uint64, error) {
	const q = `SELECT COALESCE(SUM(amount), 0) FROM payments WHERE semester = $1 AND status = 'confirmed'`
	var vol uint64
	err := r.pool.QueryRow(ctx, q, semester).Scan(&vol)
	if err != nil {
		return 0, fmt.Errorf("payment_repo: total volume: %w", err)
	}
	return vol, nil
}
func (r *PaymentRepo) scanPayment(row pgx.Row) (*domain.Payment, error) {
	p := &domain.Payment{}
	var txHash *string
	var blockNumber *uint64
	err := row.Scan(
		&p.ID, &p.StudentID, &p.StudentHash, &p.ReceiptHash,
		&p.Amount, &p.Semester, &p.FeeType, &p.Status,
		&txHash, &blockNumber, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("payment_repo: scan: %w", err)
	}
	if txHash != nil {
		p.TxHash = *txHash
	}
	if blockNumber != nil {
		p.BlockNumber = *blockNumber
	}
	return p, nil
}
func (r *PaymentRepo) listPayments(ctx context.Context, q string, args ...any) ([]*domain.Payment, int64, error) {
	rows, err := r.pool.Query(ctx, q, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("payment_repo: list: %w", err)
	}
	defer rows.Close()
	var payments []*domain.Payment
	for rows.Next() {
		p := &domain.Payment{}
		var txHash *string
		var blockNumber *uint64
		if err := rows.Scan(
			&p.ID, &p.StudentID, &p.StudentHash, &p.ReceiptHash,
			&p.Amount, &p.Semester, &p.FeeType, &p.Status,
			&txHash, &blockNumber, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("payment_repo: scan row: %w", err)
		}
		if txHash != nil {
			p.TxHash = *txHash
		}
		if blockNumber != nil {
			p.BlockNumber = *blockNumber
		}
		payments = append(payments, p)
	}
	return payments, 0, rows.Err()
}
