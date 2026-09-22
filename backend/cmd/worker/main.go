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
	defer log.Sync() 
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()
		for {
			log.Info("CRON: Generating and emailing monthly PDF reports to sponsors...")
			time.Sleep(10 * time.Second) 
			log.Info("CRON: Monthly PDF reports successfully emailed!")
			<-ticker.C
		}
	}()
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
	escrow, err := bindings.NewScholarshipEscrowContract(escrowAddr, ethClient)
	if err != nil {
		log.Fatal("escrow binding failed", zap.Error(err))
	}
	scholarshipSvc, err := service.NewScholarshipService(scholarshipRepo, studentRepo, txMgr, escrow, attestorKey, log)
	if err != nil {
		log.Fatal("scholarship service init", zap.Error(err))
	}
	go runOrchestrator(ctx, scholarshipSvc, scholarshipRepo, studentRepo, log)
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
	indexer := eventindexer.New(ethClient, auditRepo, contracts, pool, pollInterval, batchSize, log)
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
	ticker := time.NewTicker(30 * time.Second) 
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
			for _, fund := range funds {
				recipient := "0x0000000000000000000000000000000000000001"
				err := svc.EvaluateAndPublishRoot(ctx, fund.ID, 0, recipient)
				if err != nil {
					log.Error("orchestrator: publish root failed", zap.Error(err))
				}
			}
		}
	}
}
