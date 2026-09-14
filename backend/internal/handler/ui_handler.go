package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cabon-tech/unitreasury-chain/backend/internal/auth"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/service"
)

type UIHandler struct {
	studentSvc *service.StudentService
}

func NewUIHandler(studentSvc *service.StudentService) *UIHandler {
	return &UIHandler{studentSvc: studentSvc}
}

func (h *UIHandler) Index(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")

	// If token exists, redirect to dashboard, else login
	tokenStr, _ := c.Cookie("token")
	if tokenStr != "" {
		c.Redirect(http.StatusFound, "/dashboard")
		return
	}
	c.Redirect(http.StatusFound, "/sign-in")
}

func (h *UIHandler) Login(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (h *UIHandler) Logout(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, "/sign-in")
}

func (h *UIHandler) Dashboard(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	claims := auth.GetClaims(c)
	if claims == nil {
		c.Redirect(http.StatusFound, "/sign-in")
		return
	}

	switch claims.Role {
	case auth.RoleAdmin:
		c.HTML(http.StatusOK, "admin.html", gin.H{"Role": claims.Role, "StudentID": claims.StudentID})
	case auth.RoleFinance:
		c.HTML(http.StatusOK, "finance.html", gin.H{"Role": claims.Role, "StudentID": claims.StudentID})
	case auth.RoleStudent:
		if claims.StudentID == "" {
			c.Redirect(http.StatusFound, "/logout")
			return
		}
		credits := 0
		if st, err := h.studentSvc.GetByStudentID(c.Request.Context(), claims.StudentID); err == nil && st != nil {
			credits = st.Credits
		}
		c.HTML(http.StatusOK, "student.html", gin.H{"Role": claims.Role, "StudentID": claims.StudentID, "Credits": credits})
	default:
		c.String(http.StatusInternalServerError, "Unknown role")
	}
}
