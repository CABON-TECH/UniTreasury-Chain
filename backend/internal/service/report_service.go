package service
import (
	"context"
	"fmt"
	"io"
	"time"
	"strconv"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/domain"
	"github.com/jung-kurt/gofpdf"
	"go.uber.org/zap"
)
type ReportService struct {
	repo domain.ScholarshipRepository
	log  *zap.Logger
}
func NewReportService(repo domain.ScholarshipRepository, log *zap.Logger) *ReportService {
	return &ReportService{
		repo: repo,
		log:  log,
	}
}
func (s *ReportService) GenerateFundReportPDF(ctx context.Context, fundID int64, w io.Writer) error {
	fund, err := s.repo.GetFundByID(ctx, fundID)
	if err != nil {
		return fmt.Errorf("get fund: %w", err)
	}
	rows, err := s.repo.GetReportRows(ctx, fundID)
	if err != nil {
		return fmt.Errorf("get report rows: %w", err)
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Scholarship Fund Report")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Fund ID: %d (On-Chain ID: %d)", fund.ID, fund.OnChainID))
	pdf.Ln(6)
	pdf.Cell(40, 10, fmt.Sprintf("Sponsor: %s", fund.Sponsor))
	pdf.Ln(6)
	pdf.Cell(40, 10, fmt.Sprintf("Total Amount: %d USDC", fund.TotalAmount/1000000))
	pdf.Ln(6)
	pdf.Cell(40, 10, fmt.Sprintf("Released Amount: %d USDC", fund.ReleasedAmount/1000000))
	pdf.Ln(15)
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(40, 8, "Student", "1", 0, "L", false, 0, "")
	pdf.CellFormat(15, 8, "GPA", "1", 0, "C", false, 0, "")
	pdf.CellFormat(25, 8, "Tranche", "1", 0, "C", false, 0, "")
	pdf.CellFormat(25, 8, "Amount", "1", 0, "C", false, 0, "")
	pdf.CellFormat(75, 8, "TxHash", "1", 0, "L", false, 0, "")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 9)
	for _, r := range rows {
		pdf.CellFormat(40, 8, r.StudentName, "1", 0, "L", false, 0, "")
		pdf.CellFormat(15, 8, fmt.Sprintf("%.2f", r.StudentGPA), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 8, strconv.Itoa(r.TrancheIndex), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 8, fmt.Sprintf("%d", r.Amount/1000000), "1", 0, "C", false, 0, "")
		shortHash := r.TxHash
		if len(shortHash) > 40 {
			shortHash = shortHash[:40] + "..."
		}
		pdf.CellFormat(75, 8, shortHash, "1", 0, "L", false, 0, "")
		pdf.Ln(8)
	}
	pdf.Ln(10)
	pdf.SetFont("Arial", "I", 8)
	pdf.Cell(40, 10, fmt.Sprintf("Report generated on %s", time.Now().Format(time.RFC1123)))
	return pdf.Output(w)
}
