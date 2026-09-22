package sis
import "context"
type StudentInfo struct {
	StudentID string
	Name      string
	Program   string
	Year      int
	Credits   int      
	GPA       float64
	Active    bool     
}
type EligibilityResult struct {
	Eligible bool
	Reason   string 
}
type Adapter interface {
	GetStudentInfo(ctx context.Context, studentID string) (*StudentInfo, error)
	CheckScholarshipEligibility(ctx context.Context, studentID string, fundOnChainID uint64) (*EligibilityResult, error)
	ListEligibleStudents(ctx context.Context, fundOnChainID uint64) ([]string, error)
}
