// Package handler contains Gin HTTP handlers. Each handler is a thin layer:
// parse/validate request → call service → return JSON response.
// Business logic lives entirely in internal/service.
package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/cabon-tech/unitreasury-chain/backend/internal/auth"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/service"
)

// StudentHandler handles student-related HTTP endpoints.
type StudentHandler struct {
	svc *service.StudentService
	log *zap.Logger
}

// NewStudentHandler creates a StudentHandler.
func NewStudentHandler(svc *service.StudentService, log *zap.Logger) *StudentHandler {
	return &StudentHandler{svc: svc, log: log}
}

// Create handles POST /api/v1/students
// Role: admin only
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

// List handles GET /api/v1/students?page=1&page_size=20
// Role: admin, finance
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

// Get handles GET /api/v1/students/:id
// Role: admin, finance (any); student (own record only)
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

	// Students can only view their own record
	claims := auth.GetClaims(c)
	if claims != nil && claims.Role == auth.RoleStudent {
		if student.StudentID != claims.StudentID {
			c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
			return
		}
	}

	c.JSON(http.StatusOK, student)
}
