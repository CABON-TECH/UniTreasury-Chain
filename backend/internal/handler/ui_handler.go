package handler
import (
	"net/http"
	"os"
	"os/exec"
	"fmt"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/auth"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/service"
)
type UIHandler struct {
	studentSvc *service.StudentService
	pool       *pgxpool.Pool
}
func NewUIHandler(studentSvc *service.StudentService, pool *pgxpool.Pool) *UIHandler {
	return &UIHandler{studentSvc: studentSvc, pool: pool}
}
func (h *UIHandler) Index(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
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
		type AuditEvent struct {
			EventName string
			Contract  string
			TxHash    string
			Date      string
			Payload   string
		}
		var auditEvents []AuditEvent
		rows, err := h.pool.Query(c.Request.Context(), "SELECT event_name, contract, tx_hash, payload, indexed_at FROM audit_events ORDER BY indexed_at DESC LIMIT 50")
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				var evt AuditEvent
				rows.Scan(&evt.EventName, &evt.Contract, &evt.TxHash, &evt.Payload, &evt.Date)
				auditEvents = append(auditEvents, evt)
			}
		}
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"Role": claims.Role, 
			"StudentID": claims.StudentID,
			"AuditEvents": auditEvents,
		})
	case auth.RoleProfessor:
		c.HTML(http.StatusOK, "professor.html", gin.H{"Role": claims.Role, "StudentID": claims.StudentID})
	case auth.RoleFinance:
		c.HTML(http.StatusOK, "finance.html", gin.H{"Role": claims.Role, "StudentID": claims.StudentID})
	case auth.RoleStudent:
		if claims.StudentID == "" {
			c.Redirect(http.StatusFound, "/logout")
			return
		}
		credits := 0
		kyc := false
		var id int64
		if st, err := h.studentSvc.GetByStudentID(c.Request.Context(), claims.StudentID); err == nil && st != nil {
			credits = st.Credits
			kyc = st.KYCVerified
			id = st.ID
		}
		c.HTML(http.StatusOK, "student.html", gin.H{
			"Role": claims.Role, 
			"StudentID": claims.StudentID, 
			"Credits": credits,
			"KYCVerified": kyc,
			"ID": id,
		})
	default:
		c.String(http.StatusInternalServerError, "Unknown role")
	}
}
func (h *UIHandler) Treasury(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	treasuryAbi, _ := ioutil.ReadFile("../contracts/out/TreasuryContract.sol/TreasuryContract.json")
	c.HTML(http.StatusOK, "treasury.html", gin.H{
		"Role": "admin", 
		"TreasuryAddress": os.Getenv("TREASURY_CONTRACT_ADDRESS"),
		"TreasuryABI": string(treasuryAbi),
	})
}
func (h *UIHandler) DemoFundAndRole(c *gin.Context) {
	var req struct {
		Address string `json:"address"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
    pk := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
    treasury := os.Getenv("TREASURY_CONTRACT_ADDRESS")
    cmd1 := exec.Command("/home/cabon-tech/.foundry/bin/cast", "send", req.Address, "--value", "10ether", "--private-key", pk, "--rpc-url", "http://127.0.0.1:8545")
    out1, err1 := cmd1.CombinedOutput(); fmt.Println(string(out1), err1)
    cmd2 := exec.Command("/home/cabon-tech/.foundry/bin/cast", "send", treasury, "grantRole(bytes32,address)", "0xb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc1", req.Address, "--private-key", pk, "--rpc-url", "http://127.0.0.1:8545")
    out2, _ := cmd2.CombinedOutput(); fmt.Println(string(out2))
    cmd3 := exec.Command("/home/cabon-tech/.foundry/bin/cast", "send", treasury, "grantRole(bytes32,address)", "0xe463e2730303383a1253a69a239ed7e6cb9bebd1452485c2c7b508f7ceecfbbf", req.Address, "--private-key", pk, "--rpc-url", "http://127.0.0.1:8545")
    out3, _ := cmd3.CombinedOutput(); fmt.Println(string(out3))
	c.JSON(200, gin.H{"status": "ok"})
}
func (h *UIHandler) PublicPortal(c *gin.Context) {
	ctx := context.Background()
	var totalFunds uint64
	var totalReleased uint64
	var studentCount int
	_ = h.pool.QueryRow(ctx, "SELECT COALESCE(SUM(total_amount),0), COALESCE(SUM(released_amount),0) FROM scholarship_funds").Scan(&totalFunds, &totalReleased)
	_ = h.pool.QueryRow(ctx, "SELECT COUNT(*) FROM students").Scan(&studentCount)
	type AuditEvent struct {
		EventName string
		TxHash    string
		Date      string
	}
	var recentEvents []AuditEvent
	rows, err := h.pool.Query(ctx, "SELECT event_name, tx_hash, indexed_at FROM audit_events ORDER BY indexed_at DESC LIMIT 5")
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var evt AuditEvent
			rows.Scan(&evt.EventName, &evt.TxHash, &evt.Date)
			recentEvents = append(recentEvents, evt)
		}
	}
	c.HTML(http.StatusOK, "portal.html", gin.H{
		"TotalFunds":    totalFunds / 1_000000, 
		"TotalReleased": totalReleased / 1_000000,
		"StudentCount":  studentCount,
		"RecentEvents":  recentEvents,
	})
}
