package domain

import "context"

// ── Student Repository ────────────────────────────────────────────────────────

// StudentRepository is the port for student persistence.
type StudentRepository interface {
	Create(ctx context.Context, s *Student) error
	GetByID(ctx context.Context, id int64) (*Student, error)
	GetByStudentID(ctx context.Context, studentID string) (*Student, error)
	GetByHash(ctx context.Context, hash string) (*Student, error)
	ListAll(ctx context.Context) ([]*Student, error)
	List(ctx context.Context, offset, limit int) ([]*Student, int64, error)
	UpdateCredits(ctx context.Context, id int64, credits int) error
	UpdateCreditsByHash(ctx context.Context, hash string, credits int) error
	VerifyKYC(ctx context.Context, id int64) error
}

// ── Payment Repository ────────────────────────────────────────────────────────

// PaymentRepository is the port for payment persistence.
type PaymentRepository interface {
	Create(ctx context.Context, p *Payment) error
	GetByID(ctx context.Context, id int64) (*Payment, error)
	GetByReceiptHash(ctx context.Context, receiptHash string) (*Payment, error)
	ListByStudent(ctx context.Context, studentID int64, offset, limit int) ([]*Payment, int64, error)
	ListBySemester(ctx context.Context, semester int, offset, limit int) ([]*Payment, int64, error)
	UpdateStatus(ctx context.Context, id int64, status PaymentStatus, txHash string) error
	UpdateConfirmed(ctx context.Context, id int64, blockNumber uint64) error
	GetTotalVolumeForSemester(ctx context.Context, semester int) (uint64, error)
}

// ── Treasury Repository ────────────────────────────────────────────────────────

// ProposalRepository is the port for withdrawal proposal persistence.
type ProposalRepository interface {
	Create(ctx context.Context, p *WithdrawalProposal) error
	GetByID(ctx context.Context, id int64) (*WithdrawalProposal, error)
	GetByOnChainID(ctx context.Context, onChainID uint64) (*WithdrawalProposal, error)
	List(ctx context.Context, offset, limit int) ([]*WithdrawalProposal, int64, error)
	UpdateStatus(ctx context.Context, id int64, status ProposalStatus, txHash string) error
	UpdateApprovalCount(ctx context.Context, onChainID uint64, count int) error
	UpdateOnChainID(ctx context.Context, id int64, onChainID uint64) error
}

// ── Scholarship Repository ─────────────────────────────────────────────────────

// ScholarshipRepository is the port for scholarship fund and tranche persistence.
type ScholarshipRepository interface {
	CreateFund(ctx context.Context, f *ScholarshipFund) error
	GetFundByID(ctx context.Context, id int64) (*ScholarshipFund, error)
	GetFundByOnChainID(ctx context.Context, onChainID uint64) (*ScholarshipFund, error)
	ListActiveFunds(ctx context.Context) ([]*ScholarshipFund, error)
	UpdateFundReleasedAmount(ctx context.Context, fundID int64, releasedAmount uint64) error

	CreateTrancheRelease(ctx context.Context, t *TrancheRelease) error
	DeleteTrancheRelease(ctx context.Context, fundID int64, studentHash string, trancheIndex int) error
	HasReleased(ctx context.Context, fundID int64, studentHash string, trancheIndex int) (bool, error)
	GetReportRows(ctx context.Context, fundID int64) ([]*TrancheReportRow, error)
}

// ── Audit Repository ──────────────────────────────────────────────────────────

// AuditRepository is the port for event log persistence.
type AuditRepository interface {
	Insert(ctx context.Context, e *AuditEvent) error
	GetLastIndexedBlock(ctx context.Context, contract string) (uint64, error)
	SetLastIndexedBlock(ctx context.Context, contract string, block uint64) error
	List(ctx context.Context, contract string, offset, limit int) ([]*AuditEvent, int64, error)
}
