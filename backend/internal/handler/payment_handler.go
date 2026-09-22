package handler
import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/service"
)
type PaymentHandler struct {
	feeSvc      *service.FeeService
	paymentRepo interface {
		GetByID(ctx interface{ Value(interface{}) interface{} }, id int64) (interface{}, error)
	}
	log *zap.Logger
}
func NewPaymentHandler(feeSvc *service.FeeService, log *zap.Logger) *PaymentHandler {
	return &PaymentHandler{feeSvc: feeSvc, log: log}
}
func (h *PaymentHandler) UploadCSV(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing file field in form"})
		return
	}
	if file.Size > 10<<20 { 
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "file exceeds 10MB limit"})
		return
	}
	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open uploaded file"})
		return
	}
	defer f.Close()
	h.log.Info("processing payment CSV",
		zap.String("filename", file.Filename),
		zap.Int64("size_bytes", file.Size),
	)
	result, err := h.feeSvc.IngestCSV(c.Request.Context(), f)
	if err != nil {
		h.log.Warn("CSV ingest failed", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	status := http.StatusOK
	if result.Submitted == 0 {
		status = http.StatusUnprocessableEntity
	}
	c.JSON(status, result)
}
func (h *PaymentHandler) List(c *gin.Context) {
	semester, _ := strconv.Atoi(c.Query("semester"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	if page < 1 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}
	_ = semester 
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "list payments not yet wired — available in next iteration",
		"hint":  "use GET /api/v1/payments/:id for single-payment lookup",
	})
}
func (h *PaymentHandler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payment id"})
		return
	}
	_ = id
	c.JSON(http.StatusNotImplemented, gin.H{"payment_id": id, "note": "wiring in next iteration"})
}
