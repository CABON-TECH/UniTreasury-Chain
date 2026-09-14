package service

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"github.com/cabon-tech/unitreasury-chain/backend/internal/auth"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/domain"
)

type AuthService struct {
	repo   domain.UserRepository
	jwtMgr *auth.JWTManager
	log    *zap.Logger
}

func NewAuthService(repo domain.UserRepository, jwtMgr *auth.JWTManager, log *zap.Logger) *AuthService {
	return &AuthService{
		repo:   repo,
		jwtMgr: jwtMgr,
		log:    log,
	}
}

// BootstrapAdmin ensures an admin user exists.
func (s *AuthService) BootstrapAdmin(ctx context.Context, defaultPass string) error {
	existing, err := s.repo.GetByUsername(ctx, "admin")
	if err != nil {
		return err
	}
	if existing != nil {
		return nil // already exists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(defaultPass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin := &domain.User{
		Username:     "admin",
		PasswordHash: string(hash),
		Role:         string(auth.RoleAdmin),
	}

	if err := s.repo.Create(ctx, admin); err != nil {
		return err
	}
	s.log.Info("bootstrapped default admin user")
	return nil
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	u, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return "", fmt.Errorf("auth_service: %w", err)
	}
	if u == nil {
		return "", fmt.Errorf("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	role := auth.Role(u.Role)
	token, err := s.jwtMgr.Generate(u.ID, role, u.StudentID)
	if err != nil {
		return "", fmt.Errorf("auth_service: generate token: %w", err)
	}

	return token, nil
}
