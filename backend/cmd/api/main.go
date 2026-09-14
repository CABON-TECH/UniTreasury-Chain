// Command api is the HTTP API server for UniTreasury Chain.
// It exposes REST endpoints for students, payments, treasury proposals, and audit logs.
// Authentication is JWT-based with role-based access control (admin / finance / student).
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
	"github.com/cabon-tech/unitreasury-chain/backend/pkg/config"
	"github.com/cabon-tech/unitreasury-chain/backend/pkg/logger"
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

	// ── Wire dependencies ────────────────────────────────────────────────────────
	jwtMgr := auth.NewJWTManager(cfg.JWTSecret, cfg.JWTExpiryHours)

	// TODO Sprint 2: initialise DB pool, repositories, services, handlers here
	_ = jwtMgr

	// ── Router setup ─────────────────────────────────────────────────────────────
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(requestLogger(log))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"version": "0.1.0",
			"chain":   cfg.ChainID,
		})
	})

	// API v1 group
	v1 := r.Group("/api/v1")
	v1.Use(auth.Authenticate(jwtMgr))

	// Student routes (admin + finance can read; student can read own)
	students := v1.Group("/students")
	students.Use(auth.RequireRole(auth.RoleAdmin, auth.RoleFinance, auth.RoleStudent))
	students.GET("", auth.RequireRole(auth.RoleAdmin, auth.RoleFinance), placeholder("list students"))
	students.POST("", auth.RequireRole(auth.RoleAdmin), placeholder("create student"))
	students.GET("/:id", placeholder("get student"))

	// Payment routes
	payments := v1.Group("/payments")
	payments.Use(auth.RequireRole(auth.RoleAdmin, auth.RoleFinance))
	payments.GET("", placeholder("list payments"))
	payments.POST("/csv", placeholder("upload payment CSV"))
	payments.GET("/:id", placeholder("get payment"))

	// Treasury routes (admin only for write, finance can read)
	treasury := v1.Group("/treasury")
	treasury.GET("/proposals", auth.RequireRole(auth.RoleAdmin, auth.RoleFinance), placeholder("list proposals"))
	treasury.POST("/proposals", auth.RequireRole(auth.RoleAdmin, auth.RoleFinance), placeholder("propose withdrawal"))
	treasury.POST("/proposals/:id/approve", auth.RequireRole(auth.RoleAdmin), placeholder("approve withdrawal"))
	treasury.POST("/proposals/:id/execute", auth.RequireRole(auth.RoleAdmin), placeholder("execute withdrawal"))
	treasury.POST("/proposals/:id/cancel", auth.RequireRole(auth.RoleAdmin), placeholder("cancel withdrawal"))

	// Audit log routes
	audit := v1.Group("/audit")
	audit.Use(auth.RequireRole(auth.RoleAdmin))
	audit.GET("/events", placeholder("list audit events"))

	// ── HTTP server with graceful shutdown ────────────────────────────────────────
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Info("API server starting", zap.String("addr", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("server error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("shutting down API server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("forced shutdown", zap.Error(err))
	}
	log.Info("API server stopped")
}

// placeholder returns a handler that responds with 501 and a "not yet implemented" message.
// These are replaced with real handlers in Sprint 2+.
func placeholder(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"error":   "not_implemented",
			"handler": name,
			"note":    "This endpoint will be implemented in Sprint 2",
		})
	}
}

// requestLogger is a structured Gin middleware using zap.
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
