// Command api is the HTTP API server for UniTreasury Chain.
package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/cabon-tech/unitreasury-chain/backend/internal/auth"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/blockchain"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/blockchain/bindings"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/handler"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/repository/postgres"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/service"
	"github.com/cabon-tech/unitreasury-chain/backend/pkg/config"
	"github.com/cabon-tech/unitreasury-chain/backend/pkg/logger"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "config error:", err)
		os.Exit(1)
	}

	log := logger.Must(cfg.Env)
	defer log.Sync() //nolint:errcheck

	if cfg.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	ctx := context.Background()

	// ── Database ──────────────────────────────────────────────────────────────────
	pool, err := postgres.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal("database connection failed", zap.Error(err))
	}
	defer pool.Close()
	log.Info("database connected")

	// ── Repositories ──────────────────────────────────────────────────────────────
	studentRepo := postgres.NewStudentRepo(pool)
	paymentRepo := postgres.NewPaymentRepo(pool)
	_ = postgres.NewAuditRepo(pool) // used by event indexer in worker binary

	// ── Blockchain client ─────────────────────────────────────────────────────────
	ethClient, err := blockchain.NewClient(ctx, cfg.RPCURL, cfg.ChainID, log)
	if err != nil {
		log.Warn("blockchain client unavailable (running without on-chain features)", zap.Error(err))
		ethClient = nil
	}

	// ── Contract bindings ──────────────────────────────────────────────────────────
	var feeRegistry *bindings.FeeRegistryContract
	var txMgr *blockchain.TxManager

	if ethClient != nil {
		feeRegistryAddr := common.HexToAddress(cfg.FeeRegistryAddress)
		feeRegistry, err = bindings.NewFeeRegistryContract(feeRegistryAddr, ethClient.Inner())
		if err != nil {
			log.Fatal("fee registry binding failed", zap.Error(err))
		}

		attestorKey, err := blockchain.ParsePrivateKey(cfg.AttestorPrivateKey)
		if err != nil {
			log.Fatal("parse attestor key", zap.Error(err))
		}
		txMgr, err = blockchain.NewTxManager(ctx, ethClient, attestorKey, log)
		if err != nil {
			log.Fatal("tx manager init", zap.Error(err))
		}
	}

	// ── Services ──────────────────────────────────────────────────────────────────
	jwtMgr := auth.NewJWTManager(cfg.JWTSecret, cfg.JWTExpiryHours)
	studentSvc := service.NewStudentService(studentRepo, log)

	var feeSvc *service.FeeService
	if ethClient != nil && feeRegistry != nil && txMgr != nil {
		feeSvc = service.NewFeeService(paymentRepo, studentRepo, txMgr, feeRegistry, ethClient, log)
	}

	// ── Handlers ──────────────────────────────────────────────────────────────────
	studentH := handler.NewStudentHandler(studentSvc, log)
	var paymentH *handler.PaymentHandler
	if feeSvc != nil {
		paymentH = handler.NewPaymentHandler(feeSvc, log)
	}

	// ── Router ────────────────────────────────────────────────────────────────────
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(requestLogger(log))

	r.GET("/health", func(c *gin.Context) {
		dbStatus := "ok"
		if err := pool.Ping(c.Request.Context()); err != nil {
			dbStatus = "degraded: " + err.Error()
		}
		chainStatus := "disconnected"
		if ethClient != nil {
			chainStatus = "connected"
		}
		c.JSON(http.StatusOK, gin.H{
			"status":       "ok",
			"version":      "0.2.0",
			"chain_id":     cfg.ChainID,
			"db":           dbStatus,
			"blockchain":   chainStatus,
		})
	})

	v1 := r.Group("/api/v1")
	v1.Use(auth.Authenticate(jwtMgr))

	// Students
	students := v1.Group("/students")
	students.POST("", auth.RequireRole(auth.RoleAdmin), studentH.Create)
	students.GET("", auth.RequireRole(auth.RoleAdmin, auth.RoleFinance), studentH.List)
	students.GET("/:id", auth.RequireRole(auth.RoleAdmin, auth.RoleFinance, auth.RoleStudent), studentH.Get)

	// Payments
	payments := v1.Group("/payments")
	payments.Use(auth.RequireRole(auth.RoleAdmin, auth.RoleFinance))
	if paymentH != nil {
		payments.POST("/csv", paymentH.UploadCSV)
		payments.GET("", paymentH.List)
		payments.GET("/:id", paymentH.Get)
	} else {
		payments.POST("/csv", unavailable("blockchain not connected"))
		payments.GET("", unavailable("blockchain not connected"))
		payments.GET("/:id", unavailable("blockchain not connected"))
	}

	// Treasury (Sprint 3)
	treasury := v1.Group("/treasury")
	treasury.GET("/proposals", auth.RequireRole(auth.RoleAdmin, auth.RoleFinance), placeholder("list proposals — Sprint 3"))
	treasury.POST("/proposals", auth.RequireRole(auth.RoleAdmin, auth.RoleFinance), placeholder("propose withdrawal — Sprint 3"))
	treasury.POST("/proposals/:id/approve", auth.RequireRole(auth.RoleAdmin), placeholder("approve withdrawal — Sprint 3"))
	treasury.POST("/proposals/:id/execute", auth.RequireRole(auth.RoleAdmin), placeholder("execute withdrawal — Sprint 3"))
	treasury.POST("/proposals/:id/cancel", auth.RequireRole(auth.RoleAdmin), placeholder("cancel withdrawal — Sprint 3"))

	// Audit
	audit := v1.Group("/audit")
	audit.Use(auth.RequireRole(auth.RoleAdmin))
	audit.GET("/events", placeholder("list audit events — Sprint 3"))

	// Auth (no middleware — public)
	r.POST("/api/v1/auth/token", devTokenHandler(jwtMgr))

	// ── HTTP server ───────────────────────────────────────────────────────────────
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Info("API server starting", zap.String("addr", srv.Addr), zap.String("version", "0.2.0"))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("server error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("shutting down...")
	shutCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutCtx); err != nil {
		log.Error("forced shutdown", zap.Error(err))
	}
	log.Info("server stopped")
}

// devTokenHandler issues JWT tokens for development — NOT for production use.
// POST /api/v1/auth/token  body: {"role":"admin"|"finance"|"student","student_id":"..."}
func devTokenHandler(jwtMgr *auth.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Role      string `json:"role"`
			StudentID string `json:"student_id"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		role := auth.Role(req.Role)
		if role != auth.RoleAdmin && role != auth.RoleFinance && role != auth.RoleStudent {
			c.JSON(http.StatusBadRequest, gin.H{"error": "role must be admin, finance, or student"})
			return
		}
		token, err := jwtMgr.Generate(1, role, req.StudentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token, "role": role})
	}
}

func placeholder(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{"note": name})
	}
}

func unavailable(reason string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": reason})
	}
}

func requestLogger(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Info("request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", time.Since(start)),
			zap.String("ip", c.ClientIP()),
		)
	}
}
