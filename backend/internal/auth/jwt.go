// Package auth provides JWT token generation/validation and RBAC middleware.
package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Role represents an access-control role in the system.
type Role string

const (
	RoleAdmin   Role = "admin"
	RoleFinance Role = "finance"
	RoleStudent Role = "student"
)

// Claims is the JWT payload.
type Claims struct {
	UserID    int64  `json:"user_id"`
	StudentID string `json:"student_id,omitempty"` // only for student role
	Role      Role   `json:"role"`
	jwt.RegisteredClaims
}

// JWTManager handles token signing and validation.
type JWTManager struct {
	secret      []byte
	expiryHours int
}

// NewJWTManager creates a JWTManager with the given signing secret.
func NewJWTManager(secret string, expiryHours int) *JWTManager {
	return &JWTManager{
		secret:      []byte(secret),
		expiryHours: expiryHours,
	}
}

// Generate creates a signed JWT for the given user.
func (m *JWTManager) Generate(userID int64, role Role, studentID string) (string, error) {
	claims := Claims{
		UserID:    userID,
		StudentID: studentID,
		Role:      role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(m.expiryHours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(m.secret)
	if err != nil {
		return "", fmt.Errorf("auth: sign token: %w", err)
	}
	return signed, nil
}

// Validate parses and validates a JWT, returning the Claims on success.
func (m *JWTManager) Validate(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("auth: unexpected signing method: %v", t.Header["alg"])
		}
		return m.secret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("auth: parse token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("auth: invalid token claims")
	}
	return claims, nil
}
