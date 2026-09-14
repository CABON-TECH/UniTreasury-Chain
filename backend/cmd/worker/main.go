package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"

	"github.com/cabon-tech/unitreasury-chain/backend/internal/blockchain"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/blockchain/bindings"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/eventindexer"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/repository/postgres"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/service"
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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		log.Info("worker: shutdown signal received")
		cancel()
	}()

	log.Info("worker starting")

	pool, err := postgres.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal("database connection failed", zap.Error(err))
	}
	defer pool.Close()

	ethClient, err := blockchain.NewClient(ctx, cfg.RPCURL, cfg.ChainID, log)
	if err != nil {
		log.Fatal("blockchain client unavailable", zap.Error(err))
	}

	attestorKey, err := blockchain.ParsePrivateKey(cfg.AttestorPrivateKey)
	if err != nil {
		log.Fatal("parse attestor key", zap.Error(err))
	}

	txMgr, err := blockchain.NewTxManager(ctx, ethClient, attestorKey, log)
	if err != nil {
		log.Fatal("tx manager init", zap.Error(err))
	}

	studentRepo := postgres.NewStudentRepo(pool)
	scholarshipRepo := postgres.NewScholarshipRepo(pool)
	auditRepo := postgres.NewAuditRepo(pool)

	escrowAddr := common.HexToAddress(cfg.EscrowAddress)
	escrow, err := bindings.NewScholarshipEscrowContract(escrowAddr, ethClient.Inner())
	if err != nil {
		log.Fatal("escrow binding failed", zap.Error(err))
	}

	scholarshipSvc, err := service.NewScholarshipService(scholarshipRepo, studentRepo, txMgr, escrow, attestorKey, log)
	if err != nil {
		log.Fatal("scholarship service init", zap.Error(err))
	}

	// 1. Scholarship Orchestrator
	go runOrchestrator(ctx, scholarshipSvc, scholarshipRepo, studentRepo, log)

	// 2. Event Indexer
	contracts := []eventindexer.IndexedContract{
		{
			Name:    "FeeRegistry",
			Address: common.HexToAddress(cfg.FeeRegistryAddress),
			Topics: []common.Hash{
				common.HexToHash("0x892a0d7f9faaf93049b49fa4f00bbbebf4736fdf5b12bf8fa29dc62a26c483cc"),
			},
		},
		{
			Name:    "Treasury",
			Address: common.HexToAddress(cfg.TreasuryAddress),
			Topics: []common.Hash{
				common.HexToHash("0x6730ffc06020c02c6dcf154486ec2c7e0bdeee5688523c10c49cc14ed6ab406f"),
			},
		},
	}
	
	pollInterval := 10 * time.Second
	batchSize := uint64(100)
	indexer := eventindexer.New(ethClient, auditRepo, contracts, pollInterval, batchSize, log)
	go indexer.Run(ctx)

	<-ctx.Done()
	log.Info("worker stopped")
}

func runOrchestrator(
	ctx context.Context, 
	svc *service.ScholarshipService, 
	repo *postgres.ScholarshipRepo, 
	studentRepo *postgres.StudentRepo, 
	log *zap.Logger,
) {
	ticker := time.NewTicker(30 * time.Second) // poll every 30s
	defer ticker.Stop()

	log.Info("scholarship orchestrator started")

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			funds, err := repo.ListActiveFunds(ctx)
			if err != nil {
				log.Error("orchestrator: list funds", zap.Error(err))
				continue
			}

			students, _, err := studentRepo.List(ctx, 0, 100)
			if err != nil {
				log.Error("orchestrator: list students", zap.Error(err))
				continue
			}

			for _, fund := range funds {
				for _, student := range students {
					recipient := "0x0000000000000000000000000000000000000001"
					
					// Evaluate tranche 0
					err := svc.EvaluateAndRelease(ctx, fund.ID, student.StudentID, 0, recipient)
					if err != nil {
						if err.Error() != "fund is paused" && 
						   err.Error() != "student not found" && 
						   len(err.Error()) > 30 && err.Error()[:30] != "student does not meet credit" {
							log.Error("orchestrator: evaluate failed (this is expected if contracts aren't deployed!)", zap.Error(err))
						}
					}
				}
			}
		}
	}
}
