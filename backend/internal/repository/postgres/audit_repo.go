package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cabon-tech/unitreasury-chain/backend/internal/domain"
)

// AuditRepo implements domain.AuditRepository against PostgreSQL.
type AuditRepo struct {
	pool *pgxpool.Pool
}

// NewAuditRepo creates an AuditRepo.
func NewAuditRepo(pool *pgxpool.Pool) *AuditRepo {
	return &AuditRepo{pool: pool}
}

// Insert persists an audit event. Uses ON CONFLICT DO NOTHING on (tx_hash, log_index)
// making it safe to call repeatedly (idempotent — indexer reruns are safe).
func (r *AuditRepo) Insert(ctx context.Context, e *domain.AuditEvent) error {
	const q = `
		INSERT INTO audit_events (contract, event_name, tx_hash, block_number, log_index, payload, indexed_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (tx_hash, log_index) DO NOTHING`

	e.IndexedAt = time.Now()

	// Ensure payload is valid JSON for the JSONB column
	if !json.Valid(e.Payload) {
		e.Payload = []byte(`{}`)
	}

	_, err := r.pool.Exec(ctx, q,
		e.Contract, e.EventName, e.TxHash, e.BlockNumber, e.LogIndex, e.Payload, e.IndexedAt,
	)
	if err != nil {
		return fmt.Errorf("audit_repo: insert: %w", err)
	}
	return nil
}

// GetLastIndexedBlock returns the highest block successfully indexed for a contract.
// Returns 0 if no checkpoint exists (fresh start).
func (r *AuditRepo) GetLastIndexedBlock(ctx context.Context, contract string) (uint64, error) {
	const q = `SELECT last_block FROM indexer_checkpoints WHERE contract = $1`
	var block uint64
	err := r.pool.QueryRow(ctx, q, contract).Scan(&block)
	if err != nil {
		return 0, fmt.Errorf("audit_repo: get checkpoint for %s: %w", contract, err)
	}
	return block, nil
}

// SetLastIndexedBlock updates the indexer checkpoint for a contract.
func (r *AuditRepo) SetLastIndexedBlock(ctx context.Context, contract string, block uint64) error {
	const q = `
		INSERT INTO indexer_checkpoints (contract, last_block, updated_at)
		VALUES ($1, $2, NOW())
		ON CONFLICT (contract) DO UPDATE SET last_block = $2, updated_at = NOW()`
	_, err := r.pool.Exec(ctx, q, contract, block)
	if err != nil {
		return fmt.Errorf("audit_repo: set checkpoint for %s: %w", contract, err)
	}
	return nil
}

// List returns paginated audit events for a contract, newest first.
func (r *AuditRepo) List(ctx context.Context, contract string, offset, limit int) ([]*domain.AuditEvent, int64, error) {
	const countQ = `SELECT COUNT(*) FROM audit_events WHERE ($1 = '' OR contract = $1)`
	var total int64
	if err := r.pool.QueryRow(ctx, countQ, contract).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("audit_repo: count: %w", err)
	}

	const q = `
		SELECT id, contract, event_name, tx_hash, block_number, log_index, payload, indexed_at
		FROM audit_events
		WHERE ($1 = '' OR contract = $1)
		ORDER BY block_number DESC, log_index DESC
		LIMIT $2 OFFSET $3`

	rows, err := r.pool.Query(ctx, q, contract, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("audit_repo: list: %w", err)
	}
	defer rows.Close()

	var events []*domain.AuditEvent
	for rows.Next() {
		e := &domain.AuditEvent{}
		if err := rows.Scan(&e.ID, &e.Contract, &e.EventName, &e.TxHash,
			&e.BlockNumber, &e.LogIndex, &e.Payload, &e.IndexedAt); err != nil {
			return nil, 0, fmt.Errorf("audit_repo: scan: %w", err)
		}
		events = append(events, e)
	}
	return events, total, rows.Err()
}
