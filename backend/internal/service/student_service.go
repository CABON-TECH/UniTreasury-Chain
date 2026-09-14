// Package service contains the application business logic layer.
package service

import (
	"context"
	"fmt"

	"golang.org/x/crypto/sha3"
	"go.uber.org/zap"

	"github.com/cabon-tech/unitreasury-chain/backend/internal/domain"
)

// StudentService handles student registration and querying.
type StudentService struct {
	repo domain.StudentRepository
	log  *zap.Logger
}

// NewStudentService creates a StudentService.
func NewStudentService(repo domain.StudentRepository, log *zap.Logger) *StudentService {
	return &StudentService{repo: repo, log: log}
}

// CreateStudentInput holds the data required to register a new student.
type CreateStudentInput struct {
	StudentID string `json:"student_id" binding:"required"`
	Name      string `json:"name"       binding:"required"`
	Program   string `json:"program"    binding:"required"`
	Year      int    `json:"year"       binding:"required,min=1,max=8"`
}

// Create registers a new student, computing the keccak256 hash for on-chain use.
func (s *StudentService) Create(ctx context.Context, in CreateStudentInput) (*domain.Student, error) {
	existing, err := s.repo.GetByStudentID(ctx, in.StudentID)
	if err != nil {
		return nil, fmt.Errorf("student_service: check duplicate: %w", err)
	}
	if existing != nil {
		return nil, fmt.Errorf("student_service: student ID %q already registered", in.StudentID)
	}

	student := &domain.Student{
		StudentID: in.StudentID,
		Hash:      keccak256Hex(in.StudentID),
		Name:      in.Name,
		Program:   in.Program,
		Year:      in.Year,
		Credits:   0,
	}

	if err := s.repo.Create(ctx, student); err != nil {
		return nil, fmt.Errorf("student_service: create: %w", err)
	}

	s.log.Info("student registered",
		zap.Int64("id", student.ID),
		zap.String("student_id", student.StudentID),
		zap.String("hash", student.Hash),
	)
	return student, nil
}

// GetByID fetches a student by primary key.
func (s *StudentService) GetByID(ctx context.Context, id int64) (*domain.Student, error) {
	student, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("student_service: get: %w", err)
	}
	if student == nil {
		return nil, fmt.Errorf("student_service: student %d not found", id)
	}
	return student, nil
}

// List returns a paginated list of students.
func (s *StudentService) List(ctx context.Context, page, pageSize int) ([]*domain.Student, int64, error) {
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	return s.repo.List(ctx, offset, pageSize)
}

// ── Helpers ───────────────────────────────────────────────────────────────────

// keccak256Hex computes keccak256(input) and returns the 64-char hex string.
// Mirrors: keccak256(abi.encodePacked(studentID)) used in Solidity.
func keccak256Hex(input string) string {
	h := sha3.NewLegacyKeccak256()
	h.Write([]byte(input))
	return fmt.Sprintf("%x", h.Sum(nil))
}
