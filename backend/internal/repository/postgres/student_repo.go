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

// StudentRepo implements domain.StudentRepository against PostgreSQL.
type StudentRepo struct {
	pool *pgxpool.Pool
}

// NewStudentRepo creates a StudentRepo.
func NewStudentRepo(pool *pgxpool.Pool) *StudentRepo {
	return &StudentRepo{pool: pool}
}

func (r *StudentRepo) Create(ctx context.Context, s *domain.Student) error {
	const q = `
		INSERT INTO students (student_id, hash, name, program, year, credits, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	now := time.Now()
	s.CreatedAt = now
	s.UpdatedAt = now

	return r.pool.QueryRow(ctx, q,
		s.StudentID, s.Hash, s.Name, s.Program, s.Year, s.Credits, s.CreatedAt, s.UpdatedAt,
	).Scan(&s.ID)
}

func (r *StudentRepo) GetByID(ctx context.Context, id int64) (*domain.Student, error) {
	const q = `SELECT id, student_id, hash, name, program, year, credits, created_at, updated_at
	           FROM students WHERE id = $1`
	return r.scanStudent(r.pool.QueryRow(ctx, q, id))
}

func (r *StudentRepo) GetByStudentID(ctx context.Context, studentID string) (*domain.Student, error) {
	const q = `SELECT id, student_id, hash, name, program, year, credits, created_at, updated_at
	           FROM students WHERE student_id = $1`
	return r.scanStudent(r.pool.QueryRow(ctx, q, studentID))
}

func (r *StudentRepo) GetByHash(ctx context.Context, hash string) (*domain.Student, error) {
	const q = `SELECT id, student_id, hash, name, program, year, credits, created_at, updated_at
	           FROM students WHERE hash = $1`
	return r.scanStudent(r.pool.QueryRow(ctx, q, hash))
}

func (r *StudentRepo) List(ctx context.Context, offset, limit int) ([]*domain.Student, int64, error) {
	const countQ = `SELECT COUNT(*) FROM students`
	var total int64
	if err := r.pool.QueryRow(ctx, countQ).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("student_repo: count: %w", err)
	}

	const q = `SELECT id, student_id, hash, name, program, year, credits, created_at, updated_at
	           FROM students ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := r.pool.Query(ctx, q, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("student_repo: list: %w", err)
	}
	defer rows.Close()

	var students []*domain.Student
	for rows.Next() {
		s := &domain.Student{}
		if err := rows.Scan(&s.ID, &s.StudentID, &s.Hash, &s.Name, &s.Program,
			&s.Year, &s.Credits, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, 0, fmt.Errorf("student_repo: scan: %w", err)
		}
		students = append(students, s)
	}
	return students, total, rows.Err()
}

func (r *StudentRepo) UpdateCredits(ctx context.Context, id int64, credits int) error {
	const q = `UPDATE students SET credits = $1, updated_at = NOW() WHERE id = $2`
	tag, err := r.pool.Exec(ctx, q, credits, id)
	if err != nil {
		return fmt.Errorf("student_repo: update credits: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("student_repo: student %d not found", id)
	}
	return nil
}

// ── Internal ──────────────────────────────────────────────────────────────────

func (r *StudentRepo) scanStudent(row pgx.Row) (*domain.Student, error) {
	s := &domain.Student{}
	err := row.Scan(&s.ID, &s.StudentID, &s.Hash, &s.Name, &s.Program,
		&s.Year, &s.Credits, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // caller checks for nil
		}
		return nil, fmt.Errorf("student_repo: scan: %w", err)
	}
	return s, nil
}

func (r *StudentRepo) UpdateCreditsByHash(ctx context.Context, hash string, credits int) error {
	const q = `UPDATE students SET credits = $1, updated_at = NOW() WHERE hash = $2`
	_, err := r.pool.Exec(ctx, q, credits, hash)
	return err
}
