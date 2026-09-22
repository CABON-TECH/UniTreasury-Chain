package handler
import (
	"github.com/ethereum/go-ethereum/common"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/auth"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/service"
)
type StudentHandler struct {
	svc *service.StudentService
	scholarshipService *service.ScholarshipService
	log *zap.Logger
}
func NewStudentHandler(svc *service.StudentService, scholarshipService *service.ScholarshipService, log *zap.Logger) *StudentHandler {
	return &StudentHandler{svc: svc, scholarshipService: scholarshipService, log: log}
}
func (h *StudentHandler) Create(c *gin.Context) {
	var in service.CreateStudentInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	student, err := h.svc.Create(c.Request.Context(), in)
	if err != nil {
		h.log.Warn("create student failed", zap.Error(err))
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, student)
}
func (h *StudentHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	if page < 1 {
		page = 1
	}
	students, total, err := h.svc.List(c.Request.Context(), page, pageSize)
	if err != nil {
		h.log.Error("list students", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list students"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":      students,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}
func (h *StudentHandler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student id"})
		return
	}
	student, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	claims := auth.GetClaims(c)
	if claims != nil && claims.Role == auth.RoleStudent {
		if student.StudentID != claims.StudentID {
			c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
			return
		}
	}
	c.JSON(http.StatusOK, student)
}
func (h *StudentHandler) VerifyKYC(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	if err := h.svc.VerifyKYC(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "kyc verified"})
}
func (h *StudentHandler) ClaimCrossChain(c *gin.Context) {
	studentID := c.Param("id")
	var req struct {
		FundID       int64  `json:"fund_id"`
		TrancheIndex int    `json:"tranche_index"`
		Recipient    string `json:"recipient"`
		DstChainId   uint16 `json:"dst_chain_id"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	err := h.scholarshipService.SimulateCrossChainClaim(c, req.FundID, studentID, req.TrancheIndex, req.Recipient, req.DstChainId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
func (h *StudentHandler) ClaimGasless(c *gin.Context) {
	studentID := c.Param("id")
	var req struct {
		FundID       int64          `json:"fund_id"`
		TrancheIndex int            `json:"tranche_index"`
		Recipient    common.Address `json:"recipient"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.scholarshipService.SimulateGaslessClaim(c.Request.Context(), req.FundID, studentID, req.TrancheIndex, req.Recipient)
	if err != nil {
		h.log.Error("Failed to claim gasless", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
