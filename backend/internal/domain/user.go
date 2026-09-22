package domain
import (
	"context"
	"time"
)
type User struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`
	StudentID    string    `json:"student_id,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}
type UserRepository interface {
	Create(ctx context.Context, u *User) error
	GetByUsername(ctx context.Context, username string) (*User, error)
}
