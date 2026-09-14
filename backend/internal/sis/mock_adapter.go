package sis

import (
	"context"
	"fmt"
)

// MockAdapter is a deterministic SIS adapter for development and testing.
// It serves pre-seeded student data and uses a simple credit-hour threshold
// for eligibility.
//
// Replace with a real HTTP adapter in production.
type MockAdapter struct {
	students map[string]*StudentInfo
	// fundEligibilityCriteria maps fundOnChainID → minimum credits required
	fundCriteria map[uint64]int
}

// NewMockAdapter creates a MockAdapter pre-loaded with synthetic student data.
func NewMockAdapter() *MockAdapter {
	students := map[string]*StudentInfo{
		"CS/001/2021": {StudentID: "CS/001/2021", Name: "Alice Mwangi", Program: "Computer Science", Year: 3, Credits: 90, GPA: 3.8, Active: true},
		"CS/002/2021": {StudentID: "CS/002/2021", Name: "Bob Otieno", Program: "Computer Science", Year: 3, Credits: 75, GPA: 3.2, Active: true},
		"EE/001/2022": {StudentID: "EE/001/2022", Name: "Carol Njeri", Program: "Electrical Engineering", Year: 2, Credits: 45, GPA: 3.5, Active: true},
		"EE/002/2022": {StudentID: "EE/002/2022", Name: "David Kamau", Program: "Electrical Engineering", Year: 2, Credits: 28, GPA: 2.9, Active: true},
		"CS/003/2023": {StudentID: "CS/003/2023", Name: "Eve Achieng", Program: "Computer Science", Year: 1, Credits: 10, GPA: 3.1, Active: true},
	}

	// Fund 1 requires >= 30 credits; Fund 2 requires >= 60 credits
	fundCriteria := map[uint64]int{
		1: 30,
		2: 60,
	}

	return &MockAdapter{students: students, fundCriteria: fundCriteria}
}

func (m *MockAdapter) GetStudentInfo(_ context.Context, studentID string) (*StudentInfo, error) {
	s, ok := m.students[studentID]
	if !ok {
		return nil, fmt.Errorf("sis: student %q not found", studentID)
	}
	return s, nil
}

func (m *MockAdapter) CheckScholarshipEligibility(_ context.Context, studentID string, fundOnChainID uint64) (*EligibilityResult, error) {
	s, ok := m.students[studentID]
	if !ok {
		return nil, fmt.Errorf("sis: student %q not found", studentID)
	}
	if !s.Active {
		return &EligibilityResult{Eligible: false, Reason: "student not actively enrolled"}, nil
	}

	minCredits, ok := m.fundCriteria[fundOnChainID]
	if !ok {
		return &EligibilityResult{Eligible: false, Reason: fmt.Sprintf("no eligibility criteria for fund %d", fundOnChainID)}, nil
	}

	if s.Credits < minCredits {
		return &EligibilityResult{
			Eligible: false,
			Reason:   fmt.Sprintf("credits %d < required %d", s.Credits, minCredits),
		}, nil
	}

	return &EligibilityResult{Eligible: true, Reason: fmt.Sprintf("credits %d >= required %d", s.Credits, minCredits)}, nil
}

func (m *MockAdapter) ListEligibleStudents(_ context.Context, fundOnChainID uint64) ([]string, error) {
	minCredits, ok := m.fundCriteria[fundOnChainID]
	if !ok {
		return nil, fmt.Errorf("sis: no eligibility criteria for fund %d", fundOnChainID)
	}

	var eligible []string
	for _, s := range m.students {
		if s.Active && s.Credits >= minCredits {
			eligible = append(eligible, s.StudentID)
		}
	}
	return eligible, nil
}
