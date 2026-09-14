// Package sis defines the Student Information System adapter interface
// and a mock implementation for development/testing.
package sis

import "context"

// StudentInfo is the subset of SIS data the scholarship orchestrator needs.
type StudentInfo struct {
	StudentID string
	Name      string
	Program   string
	Year      int
	Credits   int      // academic credit hours completed
	GPA       float64
	Active    bool     // enrolled and in good standing
}

// EligibilityResult captures the outcome of an eligibility check.
type EligibilityResult struct {
	Eligible bool
	Reason   string // human-readable explanation (logged, not sent on-chain)
}

// Adapter is the port for the Student Information System integration.
// The real implementation would call the university's SIS REST API or read
// from a sync'd DB view. The mock below is used for Sprint 1-3 development.
type Adapter interface {
	// GetStudentInfo returns current academic information for a student.
	GetStudentInfo(ctx context.Context, studentID string) (*StudentInfo, error)

	// CheckScholarshipEligibility evaluates whether a student meets the
	// eligibility criteria for a given scholarship fund.
	CheckScholarshipEligibility(ctx context.Context, studentID string, fundOnChainID uint64) (*EligibilityResult, error)

	// ListEligibleStudents returns all student IDs that currently meet the
	// eligibility criteria for a given fund (used by the scholarship orchestrator).
	ListEligibleStudents(ctx context.Context, fundOnChainID uint64) ([]string, error)
}
