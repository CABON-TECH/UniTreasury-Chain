package service
import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/blockchain"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/blockchain/bindings"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/domain"
)
type FeeService struct {
	paymentRepo domain.PaymentRepository
	studentRepo domain.StudentRepository
	txMgr       *blockchain.TxManager
	registry    *bindings.FeeRegistryContract
	ethClient   *blockchain.Client
	log         *zap.Logger
}
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
type CSVRow struct {
	StudentID   string
	ReceiptHash string
	AmountWei   uint64
	Semester    int
	FeeType     uint64
}
type IngestResult struct {
	Total     int               `json:"total"`
	Submitted int               `json:"submitted"`
	Skipped   int               `json:"skipped"`
	Errors    []RowError        `json:"errors,omitempty"`
	Payments  []*domain.Payment `json:"payments"`
}
type RowError struct {
	Row     int    `json:"row"`
	Message string `json:"message"`
}
func (s *FeeService) IngestCSV(ctx context.Context, r io.Reader) (*IngestResult, error) {
	rows, err := parseCSV(r)
	if err != nil {
		return nil, fmt.Errorf("fee_service: parse CSV: %w", err)
	}
	result := &IngestResult{Total: len(rows)}
	for i, row := range rows {
		rowNum := i + 2 
		payment, skipped, err := s.processRow(ctx, row, rowNum)
		if err != nil {
			result.Errors = append(result.Errors, RowError{Row: rowNum, Message: err.Error()})
			continue
		}
		if skipped {
			result.Skipped++
			continue
		}
		result.Submitted++
		result.Payments = append(result.Payments, payment)
	}
	return result, nil
}
func (s *FeeService) processRow(ctx context.Context, row CSVRow, rowNum int) (payment *domain.Payment, skipped bool, err error) {
	student, err := s.studentRepo.GetByStudentID(ctx, row.StudentID)
	if err != nil {
		return nil, false, fmt.Errorf("resolve student %q: %w", row.StudentID, err)
	}
	if student == nil {
		return nil, false, fmt.Errorf("student %q not found — register student first", row.StudentID)
	}
	existing, err := s.paymentRepo.GetByReceiptHash(ctx, row.ReceiptHash)
	if err != nil {
		return nil, false, fmt.Errorf("check receipt: %w", err)
	}
	if existing != nil {
		s.log.Info("skipping duplicate receipt", zap.String("receipt", row.ReceiptHash))
		return nil, true, nil
	}
	p := &domain.Payment{
		StudentID:   student.ID,
		StudentHash: student.Hash,
		ReceiptHash: row.ReceiptHash,
		Amount:      row.AmountWei,
		Semester:    row.Semester,
		FeeType:     row.FeeType,
	}
	if err := s.paymentRepo.Create(ctx, p); err != nil {
		return nil, false, fmt.Errorf("persist payment: %w", err)
	}
	var studentHash [32]byte
	copy(studentHash[:], common.FromHex(student.Hash))
	_ = keccak256Hex(row.ReceiptHash)
	amount := new(big.Int).SetUint64(row.AmountWei)
	semester := new(big.Int).SetInt64(int64(row.Semester))
	opts, confirm, rollback, txErr := s.txMgr.TransactOpts(ctx)
	if txErr != nil {
		_ = s.paymentRepo.UpdateStatus(ctx, p.ID, domain.PaymentStatusFailed, "")
		return nil, false, fmt.Errorf("transact opts: %w", txErr)
	}
	tx, txErr := s.registry.FeeRegistryContractTransactor.RecordFee(
		opts, studentHash, semester, amount,
	)
	if txErr != nil {
		rollback()
		_ = s.paymentRepo.UpdateStatus(ctx, p.ID, domain.PaymentStatusFailed, "")
		return nil, false, fmt.Errorf("recordPayment tx: %w", txErr)
	}
	confirm()
	if err := s.paymentRepo.UpdateStatus(ctx, p.ID, domain.PaymentStatusSubmitted, tx.Hash().Hex()); err != nil {
		s.log.Error("failed to update payment status to submitted", zap.Error(err))
	}
	p.TxHash = tx.Hash().Hex()
	p.Status = domain.PaymentStatusSubmitted
	s.log.Info("payment submitted on-chain",
		zap.Int64("payment_id", p.ID),
		zap.String("tx_hash", p.TxHash),
		zap.String("receipt", row.ReceiptHash),
	)
	go s.awaitConfirmation(p.ID, tx)
	return p, false, nil
}
func (s *FeeService) awaitConfirmation(paymentID int64, tx *types.Transaction) {
	ctx := context.Background()
	receipt, err := s.txMgr.WaitMinedWithBump(ctx, tx)
	if err != nil {
		s.log.Error("await confirmation failed",
			zap.Int64("payment_id", paymentID),
			zap.String("tx_hash", tx.Hash().Hex()),
			zap.Error(err),
		)
		_ = s.paymentRepo.UpdateStatus(ctx, paymentID, domain.PaymentStatusFailed, tx.Hash().Hex())
		return
	}
	if err := s.paymentRepo.UpdateConfirmed(ctx, paymentID, receipt.BlockNumber.Uint64()); err != nil {
		s.log.Error("update confirmed failed", zap.Int64("payment_id", paymentID), zap.Error(err))
		return
	}
	s.log.Info("payment confirmed",
		zap.Int64("payment_id", paymentID),
		zap.Uint64("block", receipt.BlockNumber.Uint64()),
		zap.String("tx_hash", tx.Hash().Hex()),
	)
}
func parseCSV(r io.Reader) ([]CSVRow, error) {
	cr := csv.NewReader(r)
	cr.TrimLeadingSpace = true
	header, err := cr.Read()
	if err != nil {
		return nil, fmt.Errorf("read header: %w", err)
	}
	expected := []string{"student_id", "receipt_hash", "amount_wei", "semester", "fee_type"}
	for i, col := range expected {
		if i >= len(header) || strings.TrimSpace(strings.ToLower(header[i])) != col {
			return nil, fmt.Errorf("invalid header: expected %v, got %v", expected, header)
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
			return nil, fmt.Errorf("line %d: invalid fee_type %q (must be 1–7 bitmask)", lineNum, record[4])
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
