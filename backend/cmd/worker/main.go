// Command worker runs background jobs for UniTreasury Chain:
//   - Event indexer: polls chain for contract events → writes to audit_events
//   - Scholarship orchestrator: periodically checks eligibility → releases tranches
//   - Reconciliation job: compares DB totals with on-chain state
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"



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

	ctx, cancel := context.WithCancel(context.Background())

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		log.Info("worker: shutdown signal received")
		cancel()
	}()

	log.Info("worker starting")

	// TODO Sprint 2: wire real indexer
	// TODO Sprint 3: wire scholarship orchestrator + reconciliation job

	// Placeholder: demonstrate goroutine lifecycle
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		log.Info("placeholder worker loop started (Sprint 2 will replace this)")
		for {
			select {
			case <-ctx.Done():
				log.Info("placeholder worker stopped")
				return
			case <-ticker.C:
				log.Debug("worker heartbeat")
			}
		}
	}()

	<-ctx.Done()
	log.Info("worker stopped")
}
