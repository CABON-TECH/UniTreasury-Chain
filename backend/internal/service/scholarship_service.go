package service

import (
	"time"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	
	"github.com/ethereum/go-ethereum/common"
	
	"go.uber.org/zap"

	"github.com/cabon-tech/unitreasury-chain/backend/internal/blockchain"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/blockchain/bindings"
	"github.com/cabon-tech/unitreasury-chain/backend/internal/domain"
)

type ScholarshipService struct {
	repo         domain.ScholarshipRepository
	studentRepo  domain.StudentRepository
	txMgr        *blockchain.TxManager
	escrow       *bindings.ScholarshipEscrowContract
	attestorKey  *ecdsa.PrivateKey
	
	log          *zap.Logger
}

func NewScholarshipService(
	repo domain.ScholarshipRepository,
	studentRepo domain.StudentRepository,
	txMgr *blockchain.TxManager,
	escrow *bindings.ScholarshipEscrowContract,
	attestorKey *ecdsa.PrivateKey,
	log *zap.Logger,
) (*ScholarshipService, error) {
	

	return &ScholarshipService{
		repo:        repo,
		studentRepo: studentRepo,
		txMgr:       txMgr,
		escrow:      escrow,
		attestorKey: attestorKey,
		
		log:         log,
	}, nil
}

// EvaluateAndPublishRoot evaluates all students for a fund and publishes a Merkle Root.
func (s *ScholarshipService) EvaluateAndPublishRoot(ctx context.Context, fundID int64, trancheIndex int, recipient string) error {
	fund, err := s.repo.GetFundByID(ctx, fundID)
	if err != nil || fund == nil {
		return fmt.Errorf("fund not found")
	}
	if fund.Paused {
		return fmt.Errorf("fund is paused")
	}

	students, err := s.studentRepo.ListAll(ctx)
	if err != nil {
		return err
	}

	var eligibleLeaves [][]byte
	var eligibleStudents []*domain.Student

	onChainFundId := new(big.Int).SetUint64(fund.OnChainID)
	tIndex := big.NewInt(int64(trancheIndex))
	recipAddr := common.HexToAddress(recipient)
	requiredCredits := (trancheIndex + 1) * 15

	// We will track each student's dynamic amount to save to the DB later
	studentAmounts := make(map[string]*big.Int)

	for _, student := range students {
		if student.Credits < requiredCredits {
			continue
		}
		
		released, err := s.repo.HasReleased(ctx, fund.ID, student.Hash, trancheIndex)
		if err == nil && !released {
			// Calculate dynamic amount based on GPA
			amount := new(big.Int).SetUint64(fund.TrancheAmount)
			if student.GPA >= 3.5 {
				// 20% bonus
				bonus := new(big.Int).Mul(amount, big.NewInt(20))
				bonus.Div(bonus, big.NewInt(100))
				amount.Add(amount, bonus)
			}

			var sHash [32]byte
			copy(sHash[:], common.FromHex(student.Hash))
			leaf := blockchain.GenerateLeaf(onChainFundId, sHash, tIndex, recipAddr, amount)
			eligibleLeaves = append(eligibleLeaves, leaf)
			eligibleStudents = append(eligibleStudents, student)
			studentAmounts[student.Hash] = amount
		}
	}

	if len(eligibleLeaves) == 0 {
		return nil // Nobody eligible
	}

	tree := blockchain.GenerateTree(eligibleLeaves)
	root := tree[len(tree)-1][0]
	var root32 [32]byte
	copy(root32[:], root)

	opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)
	if err != nil {
		return fmt.Errorf("transact opts: %w", err)
	}

	tx, err := s.escrow.ScholarshipEscrowContractTransactor.PublishTrancheRoot(
		opts,
		onChainFundId,
		tIndex,
		root32,
	)
	if err != nil {
		rollback()
		return fmt.Errorf("publish root: %w", err)
	}

	confirm()
	fmt.Printf("\n\n🚀 SUCCESS: Generated Merkle Tree for %d students! Published Root: 0x%x\n\n", len(eligibleLeaves), root32)


	// Save all to database (as pending claims)
	for _, student := range eligibleStudents {
		tr := &domain.TrancheRelease{
			FundID:        fund.ID,
			OnChainFundID: fund.OnChainID,
			StudentHash:   student.Hash,
			TrancheIndex:  trancheIndex,
			Amount:        studentAmounts[student.Hash].Uint64(),
			Recipient:     recipient,
			TxHash:        tx.Hash().Hex(),
			ReleasedAt:    time.Now(),
		}
		_ = s.repo.CreateTrancheRelease(ctx, tr)
	}

	return nil
}

// SimulateStudentClaim simulates a student submitting their Merkle proof to the blockchain.
func (s *ScholarshipService) SimulateStudentClaim(ctx context.Context, fundID int64, studentID string, trancheIndex int, recipient string) error {
	fund, err := s.repo.GetFundByID(ctx, fundID)
	if err != nil || fund == nil {
		return fmt.Errorf("fund not found")
	}

	student, err := s.studentRepo.GetByStudentID(ctx, studentID)
	if err != nil || student == nil {
		return fmt.Errorf("student not found")
	}

	students, err := s.studentRepo.ListAll(ctx)
	if err != nil {
		return err
	}

	var eligibleLeaves [][]byte
	onChainFundId := new(big.Int).SetUint64(fund.OnChainID)
	tIndex := big.NewInt(int64(trancheIndex))
	recipAddr := common.HexToAddress(recipient)
	requiredCredits := (trancheIndex + 1) * 15

	var studentLeaf []byte
	var claimAmount *big.Int

	for _, st := range students {
		if st.Credits < requiredCredits {
			continue
		}
		
		amount := new(big.Int).SetUint64(fund.TrancheAmount)
		if st.GPA >= 3.5 {
			bonus := new(big.Int).Mul(amount, big.NewInt(20))
			bonus.Div(bonus, big.NewInt(100))
			amount.Add(amount, bonus)
		}

		var sHash [32]byte
		copy(sHash[:], common.FromHex(st.Hash))
		
		leaf := blockchain.GenerateLeaf(onChainFundId, sHash, tIndex, recipAddr, amount)
		eligibleLeaves = append(eligibleLeaves, leaf)
		
		if st.StudentID == studentID {
			studentLeaf = leaf
			claimAmount = amount
		}
	}

	if studentLeaf == nil {
		return fmt.Errorf("student not eligible for this tranche")
	}

	tree := blockchain.GenerateTree(eligibleLeaves)
	proof := blockchain.GenerateProof(tree, studentLeaf)
	
	// Convert proof to [][32]byte
	var proof32 [][32]byte
	for _, p := range proof {
		var p32 [32]byte
		copy(p32[:], p)
		proof32 = append(proof32, p32)
	}

	opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)
	if err != nil {
		return fmt.Errorf("transact opts: %w", err)
	}

	var sHash [32]byte
	copy(sHash[:], common.FromHex(student.Hash))

	_, err = s.escrow.ScholarshipEscrowContractTransactor.ClaimTranche(
		opts,
		onChainFundId,
		sHash,
		tIndex,
		recipAddr,
		claimAmount,
		proof32,
	)
	if err != nil {
		rollback()
		return fmt.Errorf("claim tranche: %w", err)
	}

	confirm()
	return nil
}

// ClawbackFund allows the admin to reclaim unused funds
func (s *ScholarshipService) ClawbackFund(ctx context.Context, fundID int64, recipient string) error {
	fund, err := s.repo.GetFundByID(ctx, fundID)
	if err != nil || fund == nil {
		return fmt.Errorf("fund not found")
	}

	onChainFundId := new(big.Int).SetUint64(fund.OnChainID)
	recipAddr := common.HexToAddress(recipient)

	opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)
	if err != nil {
		return fmt.Errorf("transact opts: %w", err)
	}

	_, err = s.escrow.ScholarshipEscrowContractTransactor.ClawbackFund(
		opts,
		onChainFundId,
		recipAddr,
	)
	if err != nil {
		rollback()
		return fmt.Errorf("clawback fund: %w", err)
	}

	confirm()
	
	// Mark fund as paused/closed in DB
	fund.Paused = true
	// Assume update fund method exists or just ignore for now since it's a PoC
	
	return nil
}
