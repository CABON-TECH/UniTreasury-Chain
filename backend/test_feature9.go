package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cabon-tech/unitreasury-chain/backend/internal/domain"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/repository/postgres"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func main() {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "postgres://unitreasury:password@localhost:5433/unitreasury?sslmode=disable"
	}
	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		panic(err)
	}

	log, _ := zap.NewDevelopment()
	scholarshipRepo := postgres.NewScholarshipRepo(pool)
	reportSvc := service.NewReportService(scholarshipRepo, log)

	ctx := context.Background()

	// 1. Create a dummy fund
	fund := &domain.ScholarshipFund{
		Sponsor:        "Google Foundation",
		TotalAmount:    50000_000000,
		ReleasedAmount: 0,
		TrancheCount:   2,
		TrancheAmount:  1000_000000,
		OnChainID:      99,
		Paused:         false,
	}
	err = scholarshipRepo.CreateFund(ctx, fund)
	if err != nil {
		fmt.Printf("Fund creation failed (maybe already exists), continuing...\n")
	}

	// Make sure we have the fund ID. 
	funds, _ := scholarshipRepo.ListActiveFunds(ctx)
	if len(funds) == 0 {
		panic("no funds available")
	}
	targetFund := funds[len(funds)-1]

	// 2. Insert some dummy students directly if they don't exist
	studentHash := "0xabc1234567890abcdef"
	pool.Exec(ctx, "INSERT INTO students (student_id, hash, name, program, year, credits, gpa, created_at, updated_at) VALUES ('TEST-PDF-1', $1, 'Jane Doe', 'CS', 3, 45, 3.9, NOW(), NOW()) ON CONFLICT DO NOTHING", studentHash)
	
	// 3. Create a dummy tranche release
	release := &domain.TrancheRelease{
		FundID:       targetFund.ID,
		OnChainFundID: targetFund.OnChainID,
		StudentHash:  studentHash,
		TrancheIndex: 0,
		Amount:       1200_000000, // 20% bonus!
		Recipient:    "0x111122223333444455556666777788889999aaaa",
		TxHash:       "0x9999888877776666555544443333222211110000",
		BlockNumber:  42,
		ReleasedAt:   time.Now(),
	}
	scholarshipRepo.CreateTrancheRelease(ctx, release)

	// Update the fund released amount to match
	scholarshipRepo.UpdateFundReleasedAmount(ctx, targetFund.ID, 1200_000000)

	// 4. Generate the PDF!
	fileName := "sponsor_report_demo.pdf"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = reportSvc.GenerateFundReportPDF(ctx, targetFund.ID, file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nSUCCESS! I generated a PDF report for Fund #%d.\n", targetFund.ID)
	fmt.Printf("The file is saved at: backend/%s\n", fileName)
	fmt.Printf("Open it in your file explorer to see what it looks like!\n\n")
}
