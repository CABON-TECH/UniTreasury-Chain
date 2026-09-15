// Package domain defines the core domain models and repository interfaces
// for UniTreasury Chain. Services depend on these interfaces; Postgres
// implementations live in internal/repository/postgres.
package domain

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// ── Student ───────────────────────────────────────────────────────────────────

// Student represents a registered university student.
// StudentID is the canonical university identifier (never sent on-chain).
// Hash is keccak256(StudentID) — used for on-chain privacy-preserving records.
type Student struct {
	ID        int64     `json:"id"`
	StudentID string    `json:"student_id"` // university identifier, e.g. "CS/001/2021"
	Hash      string    `json:"hash"`       // hex keccak256 of StudentID
	Name      string    `json:"name"`
	Program   string    `json:"program"`
	Year      int       `json:"year"`
	Credits   int       `json:"credits"`
	GPA         float64   `json:"gpa"`
	KYCVerified bool      `json:"kyc_verified"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ── Fee Structures ────────────────────────────────────────────────────────────

// FeeStructure defines the fee schedule for a given academic semester.
type FeeStructure struct {
	ID         int64     `json:"id"`
	Semester   int       `json:"semester"`  // e.g. 20241 = 2024 Semester 1
	TuitionFee uint64    `json:"tuition_fee"` // in wei
	HostelFee  uint64    `json:"hostel_fee"`
	ExamFee    uint64    `json:"exam_fee"`
	Active     bool      `json:"active"`
	CreatedAt  time.Time `json:"created_at"`
}

// ── Payments ──────────────────────────────────────────────────────────────────

// PaymentStatus tracks the lifecycle of an on-chain payment submission.
type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusSubmitted PaymentStatus = "submitted"
	PaymentStatusConfirmed PaymentStatus = "confirmed"
	PaymentStatusFailed    PaymentStatus = "failed"
)

// FeeType bitmask values — mirrors the on-chain feeType encoding.
const (
	FeeTypeTuition uint64 = 1
	FeeTypeHostel  uint64 = 2
	FeeTypeExam    uint64 = 4
)

// Payment represents a student fee payment submitted to the FeeRegistry contract.
type Payment struct {
	ID          int64         `json:"id"`
	StudentID   int64         `json:"student_id"`   // FK to students.id
	StudentHash string        `json:"student_hash"`  // keccak256 hash (on-chain)
	ReceiptHash string        `json:"receipt_hash"`  // unique receipt identifier
	Amount      uint64        `json:"amount"`        // in wei
	Semester    int           `json:"semester"`
	FeeType     uint64        `json:"fee_type"`      // bitmask
	Status      PaymentStatus `json:"status"`
	TxHash      string        `json:"tx_hash"`       // populated after submission
	BlockNumber uint64        `json:"block_number"`  // populated after confirmation
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

// ── Treasury ──────────────────────────────────────────────────────────────────

// ProposalStatus mirrors the on-chain enum.
type ProposalStatus string

const (
	ProposalStatusPending   ProposalStatus = "pending"
	ProposalStatusExecuted  ProposalStatus = "executed"
	ProposalStatusCancelled ProposalStatus = "cancelled"
)

// WithdrawalProposal mirrors the on-chain proposal struct, cached in the DB.
type WithdrawalProposal struct {
	ID              int64          `json:"id"`
	OnChainID       uint64         `json:"on_chain_id"`    // proposalId from contract
	Proposer        common.Address `json:"proposer"`
	Recipient       common.Address `json:"recipient"`
	Amount          uint64         `json:"amount"`         // in wei
	Purpose         string         `json:"purpose"`
	Status          ProposalStatus `json:"status"`
	ApprovalCount   int            `json:"approval_count"`
	TxHash          string         `json:"tx_hash"`
	BlockNumber     uint64         `json:"block_number"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

// ── Scholarship Funds ─────────────────────────────────────────────────────────

// ScholarshipFund represents an on-chain escrow fund.
type ScholarshipFund struct {
	ID            int64     `json:"id"`
	OnChainID     uint64    `json:"on_chain_id"`
	Sponsor       string    `json:"sponsor"`
	TotalAmount   uint64    `json:"total_amount"`  // in wei
	ReleasedAmount uint64   `json:"released_amount"`
	TrancheCount  int       `json:"tranche_count"`
	TrancheAmount uint64    `json:"tranche_amount"`
	Paused        bool      `json:"paused"`
	CreatedAt     time.Time `json:"created_at"`
}

// TrancheRelease records each released scholarship tranche.
type TrancheRelease struct {
	ID           int64     `json:"id"`
	FundID       int64     `json:"fund_id"`        // FK to scholarship_funds.id
	OnChainFundID uint64   `json:"on_chain_fund_id"`
	StudentHash  string    `json:"student_hash"`
	TrancheIndex int       `json:"tranche_index"`
	Amount       uint64    `json:"amount"`
	Recipient    string    `json:"recipient"`
	TxHash       string    `json:"tx_hash"`
	BlockNumber  uint64    `json:"block_number"`
	ReleasedAt   time.Time `json:"released_at"`
}

// ── Audit Events ─────────────────────────────────────────────────────────────

// AuditEvent is an indexed on-chain event captured by the event indexer.
type AuditEvent struct {
	ID          int64     `json:"id"`
	Contract    string    `json:"contract"`     // e.g. "treasury", "feeregistry"
	EventName   string    `json:"event_name"`
	TxHash      string    `json:"tx_hash"`
	BlockNumber uint64    `json:"block_number"`
	LogIndex    uint      `json:"log_index"`
	Payload     []byte    `json:"payload"`      // JSON-encoded event args (JSONB in PG)
	IndexedAt   time.Time `json:"indexed_at"`
}

// TrancheReportRow is used for PDF reporting.
type TrancheReportRow struct {
	StudentName  string
	StudentGPA   float64
	TrancheIndex int
	Amount       uint64
	TxHash       string
	ReleasedAt   time.Time
}
