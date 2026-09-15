package handler

import (
	"net/http"
	"strconv"

	"github.com/cabon-tech/unitreasury-chain/backend/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ReportHandler struct {
	svc *service.ReportService
	log *zap.Logger
}

func NewReportHandler(svc *service.ReportService, log *zap.Logger) *ReportHandler {
	return &ReportHandler{
		svc: svc,
		log: log,
	}
}

func (h *ReportHandler) DownloadFundReport(c *gin.Context) {
	fundIDStr := c.Param("id")
	fundID, err := strconv.ParseInt(fundIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid fund ID"})
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=fund_"+fundIDStr+"_report.pdf")

	if err := h.svc.GenerateFundReportPDF(c.Request.Context(), fundID, c.Writer); err != nil {
		h.log.Error("failed to generate pdf", zap.Error(err))
		// Note: header is already written, but we can log it.
	}
}
