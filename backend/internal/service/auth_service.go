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

// BootstrapDefaultUsers ensures admin, finance, and a test student user exist.
func (s *AuthService) BootstrapDefaultUsers(ctx context.Context) error {
	users := []domain.User{
		{Username: "admin", Role: string(auth.RoleAdmin)},
		{Username: "finance", Role: string(auth.RoleFinance)},
		{Username: "student1", Role: string(auth.RoleStudent), StudentID: "CS/001/2021"},
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	passStr := string(hash)

	for _, u := range users {
		existing, err := s.repo.GetByUsername(ctx, u.Username)
		if err != nil {
			return err
		}
		if existing == nil {
			u.PasswordHash = passStr
			if err := s.repo.Create(ctx, &u); err != nil {
				return err
			}
			s.log.Info("bootstrapped user", zap.String("username", u.Username))
		}
	}
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
