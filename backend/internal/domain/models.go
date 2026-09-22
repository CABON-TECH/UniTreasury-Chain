package domain
import (
	"time"
	"github.com/ethereum/go-ethereum/common"
)
type Student struct {
	ID        int64     `json:"id"`
	StudentID string    `json:"student_id"` 
	Hash      string    `json:"hash"`       
	Name      string    `json:"name"`
	Program   string    `json:"program"`
	Year      int       `json:"year"`
	Credits   int       `json:"credits"`
	GPA         float64   `json:"gpa"`
	KYCVerified bool      `json:"kyc_verified"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type FeeStructure struct {
	ID         int64     `json:"id"`
	Semester   int       `json:"semester"`  
	TuitionFee uint64    `json:"tuition_fee"` 
	HostelFee  uint64    `json:"hostel_fee"`
	ExamFee    uint64    `json:"exam_fee"`
	Active     bool      `json:"active"`
	CreatedAt  time.Time `json:"created_at"`
}
type PaymentStatus string
const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusSubmitted PaymentStatus = "submitted"
	PaymentStatusConfirmed PaymentStatus = "confirmed"
	PaymentStatusFailed    PaymentStatus = "failed"
)
const (
	FeeTypeTuition uint64 = 1
	FeeTypeHostel  uint64 = 2
	FeeTypeExam    uint64 = 4
)
type Payment struct {
	ID          int64         `json:"id"`
	StudentID   int64         `json:"student_id"`   
	StudentHash string        `json:"student_hash"`  
	ReceiptHash string        `json:"receipt_hash"`  
	Amount      uint64        `json:"amount"`        
	Semester    int           `json:"semester"`
	FeeType     uint64        `json:"fee_type"`      
	Status      PaymentStatus `json:"status"`
	TxHash      string        `json:"tx_hash"`       
	BlockNumber uint64        `json:"block_number"`  
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}
type ProposalStatus string
const (
	ProposalStatusPending   ProposalStatus = "pending"
	ProposalStatusExecuted  ProposalStatus = "executed"
	ProposalStatusCancelled ProposalStatus = "cancelled"
)
type WithdrawalProposal struct {
	ID              int64          `json:"id"`
	OnChainID       uint64         `json:"on_chain_id"`    
	Proposer        common.Address `json:"proposer"`
	Recipient       common.Address `json:"recipient"`
	Amount          uint64         `json:"amount"`         
	Purpose         string         `json:"purpose"`
	Status          ProposalStatus `json:"status"`
	ApprovalCount   int            `json:"approval_count"`
	TxHash          string         `json:"tx_hash"`
	BlockNumber     uint64         `json:"block_number"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}
type ScholarshipFund struct {
	ID            int64     `json:"id"`
	OnChainID     uint64    `json:"on_chain_id"`
	Sponsor       string    `json:"sponsor"`
	TotalAmount   uint64    `json:"total_amount"`  
	ReleasedAmount uint64   `json:"released_amount"`
	TrancheCount  int       `json:"tranche_count"`
	TrancheAmount uint64    `json:"tranche_amount"`
	Paused        bool      `json:"paused"`
	CreatedAt     time.Time `json:"created_at"`
}
type TrancheRelease struct {
	ID           int64     `json:"id"`
	FundID       int64     `json:"fund_id"`        
	OnChainFundID uint64   `json:"on_chain_fund_id"`
	StudentHash  string    `json:"student_hash"`
	TrancheIndex int       `json:"tranche_index"`
	Amount       uint64    `json:"amount"`
	Recipient    string    `json:"recipient"`
	TxHash       string    `json:"tx_hash"`
	BlockNumber  uint64    `json:"block_number"`
	ReleasedAt   time.Time `json:"released_at"`
}
type AuditEvent struct {
	ID          int64     `json:"id"`
	Contract    string    `json:"contract"`     
	EventName   string    `json:"event_name"`
	TxHash      string    `json:"tx_hash"`
	BlockNumber uint64    `json:"block_number"`
	LogIndex    uint      `json:"log_index"`
	Payload     []byte    `json:"payload"`      
	IndexedAt   time.Time `json:"indexed_at"`
}
type TrancheReportRow struct {
	StudentName  string
	StudentGPA   float64
	TrancheIndex int
	Amount       uint64
	TxHash       string
	ReleasedAt   time.Time
}
