package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cabon-tech/unitreasury-chain/backend/internal/domain"
)

type UserRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{pool: pool}
}

func (r *UserRepo) Create(ctx context.Context, u *domain.User) error {
	const q = `
		INSERT INTO users (username, password_hash, role, student_id, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	
	var studentID sql.NullString
	if u.StudentID != "" {
		studentID = sql.NullString{String: u.StudentID, Valid: true}
	}
	
	u.CreatedAt = time.Now()
	err := r.pool.QueryRow(ctx, q, u.Username, u.PasswordHash, u.Role, studentID, u.CreatedAt).Scan(&u.ID)
	if err != nil {
		return fmt.Errorf("user_repo: create: %w", err)
	}
	return nil
}

func (r *UserRepo) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	const q = `SELECT id, username, password_hash, role, student_id, created_at FROM users WHERE username = $1`
	
	u := &domain.User{}
	var studentID sql.NullString
	
	err := r.pool.QueryRow(ctx, q, username).Scan(&u.ID, &u.Username, &u.PasswordHash, &u.Role, &studentID, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, fmt.Errorf("user_repo: get: %w", err)
	}
	
	if studentID.Valid {
		u.StudentID = studentID.String
	}
	return u, nil
}
