package service

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"

	"github.com/cabon-tech/unitreasury-chain/backend/internal/blockchain"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/blockchain/bindings"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/domain"
)

// FeeService handles payment CSV ingestion and on-chain submission.
type FeeService struct {
	paymentRepo domain.PaymentRepository
	studentRepo domain.StudentRepository
	txMgr       *blockchain.TxManager
	registry    *bindings.FeeRegistryContract
	ethClient   *blockchain.Client
	log         *zap.Logger
}

// NewFeeService creates a FeeService.
func NewFeeService(
	paymentRepo domain.PaymentRepository,
	studentRepo domain.StudentRepository,
	txMgr *blockchain.TxManager,
	registry *bindings.FeeRegistryContract,
	ethClient *blockchain.Client,
	log *zap.Logger,
) *FeeService {
	return &FeeService{
		paymentRepo: paymentRepo,
		studentRepo: studentRepo,
		txMgr:       txMgr,
		registry:    registry,
		ethClient:   ethClient,
		log:         log,
	}
}

// CSVRow represents one validated row from the payment CSV.
type CSVRow struct {
	StudentID   string
	ReceiptHash string
	AmountWei   uint64
	Semester    int
	FeeType     uint64
}

// IngestResult summarises a CSV upload.
type IngestResult struct {
	Total     int            `json:"total"`
	Submitted int            `json:"submitted"`
	Skipped   int            `json:"skipped"`   // duplicate receipt hashes
	Errors    []RowError     `json:"errors,omitempty"`
	Payments  []*domain.Payment `json:"payments"`
}

// RowError records a per-row failure.
type RowError struct {
	Row     int    `json:"row"`
	Message string `json:"message"`
}

// IngestCSV parses a payment CSV, validates each row, and submits each payment
// on-chain via FeeRegistryContract.recordPayment. Submission is synchronous per
// row — for large batches a queue would be preferable, but this keeps Sprint 2
// simple and traceable.
//
// CSV format (header required):
//   student_id,receipt_hash,amount_wei,semester,fee_type
func (s *FeeService) IngestCSV(ctx context.Context, r io.Reader) (*IngestResult, error) {
	rows, err := parseCSV(r)
	if err != nil {
		return nil, fmt.Errorf("fee_service: parse CSV: %w", err)
	}

	result := &IngestResult{Total: len(rows)}

	for i, row := range rows {
		rowNum := i + 2 // 1-indexed + header

		payment, err := s.processRow(ctx, row, rowNum)
		if err != nil {
			result.Errors = append(result.Errors, RowError{Row: rowNum, Message: err.Error()})
			continue
		}
		if payment == nil {
			result.Skipped++
			continue
		}

		result.Submitted++
		result.Payments = append(result.Payments, payment)
	}

	return result, nil
}

func (s *FeeService) processRow(ctx context.Context, row CSVRow, rowNum int) (*domain.Payment, error) {
	// Resolve student
	student, err := s.studentRepo.GetByStudentID(ctx, row.StudentID)
	if err != nil {
		return nil, fmt.Errorf("row %d: resolve student %q: %w", rowNum, row.StudentID, err)
	}
	if student == nil {
		return nil, fmt.Errorf("row %d: student %q not found — register student first", rowNum, row.StudentID)
	}

	// Idempotency: skip if receipt already exists
	existing, err := s.paymentRepo.GetByReceiptHash(ctx, row.ReceiptHash)
	if err != nil {
		return nil, fmt.Errorf("row %d: check receipt: %w", rowNum, err)
	}
	if existing != nil {
		s.log.Info("skipping duplicate receipt", zap.String("receipt", row.ReceiptHash))
		return nil, nil // nil, nil = skipped
	}

	// Persist as pending before touching the chain
	payment := &domain.Payment{
		StudentID:   student.ID,
		StudentHash: student.Hash,
		ReceiptHash: row.ReceiptHash,
		Amount:      row.AmountWei,
		Semester:    row.Semester,
		FeeType:     row.FeeType,
	}
	if err := s.paymentRepo.Create(ctx, payment); err != nil {
		return nil, fmt.Errorf("row %d: persist payment: %w", rowNum, err)
	}

	// Build on-chain args
	var studentHash [32]byte
	copy(studentHash[:], common.FromHex(student.Hash))

	var receiptHash [32]byte
	receiptBytes := keccak256Hex(row.ReceiptHash) // hash the receipt ID for on-chain dedup
	copy(receiptHash[:], common.FromHex(receiptBytes))

	amount := new(big.Int).SetUint64(row.AmountWei)
	semester := new(big.Int).SetInt64(int64(row.Semester))
	feeType := new(big.Int).SetUint64(row.FeeType)

	// Submit transaction (nonce-sequenced via TxManager)
	opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)
	if err != nil {
		_ = s.paymentRepo.UpdateStatus(ctx, payment.ID, domain.PaymentStatusFailed, "")
		return nil, fmt.Errorf("row %d: transact opts: %w", rowNum, err)
	}

	tx, err := s.registry.FeeRegistryContractTransactor.RecordPayment(
		opts, studentHash, receiptHash, amount, semester, feeType,
	)
	if err != nil {
		rollback()
		_ = s.paymentRepo.UpdateStatus(ctx, payment.ID, domain.PaymentStatusFailed, "")
		return nil, fmt.Errorf("row %d: recordPayment tx: %w", rowNum, err)
	}
	confirm()

	// Mark submitted
	if err := s.paymentRepo.UpdateStatus(ctx, payment.ID, domain.PaymentStatusSubmitted, tx.Hash().Hex()); err != nil {
		s.log.Error("failed to update payment status", zap.Error(err))
	}
	payment.TxHash = tx.Hash().Hex()
	payment.Status = domain.PaymentStatusSubmitted

	s.log.Info("payment submitted",
		zap.Int64("payment_id", payment.ID),
		zap.String("tx_hash", payment.TxHash),
		zap.String("receipt", row.ReceiptHash),
	)

	// Async: wait for confirmation and update block_number
	go s.awaitConfirmation(payment.ID, tx)

	return payment, nil
}

// awaitConfirmation waits for a tx receipt and updates the DB record.
// Runs in a goroutine — failures are logged, not fatal.
func (s *FeeService) awaitConfirmation(paymentID int64, tx interface{ Hash() common.Hash }) {
	ctx := context.Background()

	type mined interface {
		Hash() common.Hash
	}
	typedTx, ok := tx.(interface {
		Hash() common.Hash
	})
	if !ok {
		return
	}

	// Use ethclient directly for WaitMined
	receipt, err := bind.WaitMined(ctx, s.ethClient.Inner(), typedTx.(interface {
		Hash() common.Hash
		// ethclient expects *types.Transaction — FeeService uses TxManager which returns *types.Transaction
		// This cast is safe because RecordPayment returns *types.Transaction
	}))
	if err != nil {
		s.log.Error("await confirmation failed", zap.Int64("payment_id", paymentID), zap.Error(err))
		return
	}

	if err := s.paymentRepo.UpdateConfirmed(ctx, paymentID, receipt.BlockNumber.Uint64()); err != nil {
		s.log.Error("update confirmed failed", zap.Int64("payment_id", paymentID), zap.Error(err))
		return
	}

	s.log.Info("payment confirmed",
		zap.Int64("payment_id", paymentID),
		zap.Uint64("block", receipt.BlockNumber.Uint64()),
	)
}

// ── CSV parsing ───────────────────────────────────────────────────────────────

// parseCSV reads and validates all rows from a payment CSV.
// Expected header: student_id,receipt_hash,amount_wei,semester,fee_type
func parseCSV(r io.Reader) ([]CSVRow, error) {
	cr := csv.NewReader(r)
	cr.TrimLeadingSpace = true

	header, err := cr.Read()
	if err != nil {
		return nil, fmt.Errorf("read header: %w", err)
	}

	expectedCols := []string{"student_id", "receipt_hash", "amount_wei", "semester", "fee_type"}
	for i, col := range expectedCols {
		if i >= len(header) || strings.TrimSpace(strings.ToLower(header[i])) != col {
			return nil, fmt.Errorf("invalid header: expected %v, got %v", expectedCols, header)
		}
	}

	var rows []CSVRow
	lineNum := 1
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", lineNum+1, err)
		}
		lineNum++

		if len(record) < 5 {
			return nil, fmt.Errorf("line %d: expected 5 columns, got %d", lineNum, len(record))
		}

		amountWei, err := strconv.ParseUint(strings.TrimSpace(record[2]), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("line %d: invalid amount_wei %q: %w", lineNum, record[2], err)
		}

		semester, err := strconv.Atoi(strings.TrimSpace(record[3]))
		if err != nil || semester <= 0 {
			return nil, fmt.Errorf("line %d: invalid semester %q", lineNum, record[3])
		}

		feeType, err := strconv.ParseUint(strings.TrimSpace(record[4]), 10, 64)
		if err != nil || feeType == 0 || feeType > 7 {
			return nil, fmt.Errorf("line %d: invalid fee_type %q (must be 1-7 bitmask)", lineNum, record[4])
		}

		rows = append(rows, CSVRow{
			StudentID:   strings.TrimSpace(record[0]),
			ReceiptHash: strings.TrimSpace(record[1]),
			AmountWei:   amountWei,
			Semester:    semester,
			FeeType:     feeType,
		})
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("CSV has no data rows")
	}
	return rows, nil
}
