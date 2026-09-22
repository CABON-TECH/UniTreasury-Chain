package handler
import (
	"encoding/hex"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/service"
)
type ZKHandler struct {
	svc *service.ZKService
	log *zap.Logger
}
func NewZKHandler(svc *service.ZKService, log *zap.Logger) *ZKHandler {
	return &ZKHandler{svc: svc, log: log}
}
func (h *ZKHandler) UpdateMerkleRoot(c *gin.Context) {
	var req struct {
		MerkleRoot string `json:"merkle_root"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	rootBytes, err := hex.DecodeString(strings.TrimPrefix(req.MerkleRoot, "0x"))
	if err != nil || len(rootBytes) != 32 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "merkle_root must be 32-byte hex"})
		return
	}
	var root [32]byte
	copy(root[:], rootBytes)
	txHash, err := h.svc.UpdateMerkleRoot(c.Request.Context(), root)
	if err != nil {
		h.log.Error("update merkle root failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tx_hash": txHash, "message": "merkle root updated"})
}
func (h *ZKHandler) ProveEnrollment(c *gin.Context) {
	var req struct {
		Proof         string `json:"proof"`          
		NullifierHash string `json:"nullifier_hash"` 
		ClaimedRoot   string `json:"claimed_root"`   
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	proofBytes, err := hex.DecodeString(strings.TrimPrefix(req.Proof, "0x"))
	if err != nil || len(proofBytes) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "proof must be non-empty hex"})
		return
	}
	nullBytes, err := hex.DecodeString(strings.TrimPrefix(req.NullifierHash, "0x"))
	if err != nil || len(nullBytes) != 32 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nullifier_hash must be 32-byte hex"})
		return
	}
	rootBytes, err := hex.DecodeString(strings.TrimPrefix(req.ClaimedRoot, "0x"))
	if err != nil || len(rootBytes) != 32 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "claimed_root must be 32-byte hex"})
		return
	}
	var nullifier [32]byte
	var claimedRoot [32]byte
	copy(nullifier[:], nullBytes)
	copy(claimedRoot[:], rootBytes)
	txHash, err := h.svc.ProveEnrollment(c.Request.Context(), proofBytes, nullifier, claimedRoot)
	if err != nil {
		h.log.Error("prove enrollment failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tx_hash": txHash, "message": "enrollment proved anonymously"})
}
func (h *ZKHandler) GetCurrentRoot(c *gin.Context) {
	root, err := h.svc.GetCurrentRoot(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"merkle_root": hex.EncodeToString(root[:])})
}
func (h *ZKHandler) CheckNullifier(c *gin.Context) {
	nullifierStr := c.Param("nullifier")
	nullBytes, err := hex.DecodeString(strings.TrimPrefix(nullifierStr, "0x"))
	if err != nil || len(nullBytes) != 32 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nullifier must be 32-byte hex"})
		return
	}
	var nullifier [32]byte
	copy(nullifier[:], nullBytes)
	addr, err := h.svc.IsEnrollmentVerified(c.Request.Context(), nullifier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	used := addr.Hex() != "0x0000000000000000000000000000000000000000"
	c.JSON(http.StatusOK, gin.H{
		"nullifier": nullifierStr,
		"used":      used,
		"student":   addr.Hex(),
	})
}
