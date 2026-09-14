package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	domainSep    [32]byte
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
	// Fetch DOMAIN_SEPARATOR from contract
	domainSep, err := escrow.ScholarshipEscrowContractCaller.DomainSeparator(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("fetch domain separator: %w", err)
	}

	return &ScholarshipService{
		repo:        repo,
		studentRepo: studentRepo,
		txMgr:       txMgr,
		escrow:      escrow,
		attestorKey: attestorKey,
		domainSep:   domainSep,
		log:         log,
	}, nil
}

// EvaluateAndRelease is called by the orchestrator job.
func (s *ScholarshipService) EvaluateAndRelease(ctx context.Context, fundID int64, studentID string, trancheIndex int, recipient string) error {
	fund, err := s.repo.GetFundByID(ctx, fundID)
	if err != nil || fund == nil {
		return fmt.Errorf("fund not found")
	}
	if fund.Paused {
		return fmt.Errorf("fund is paused")
	}

	student, err := s.studentRepo.GetByStudentID(ctx, studentID)
	if err != nil || student == nil {
		return fmt.Errorf("student not found")
	}

	// 1. Evaluate credits (business rule: e.g. 15 credits per tranche)
	requiredCredits := (trancheIndex + 1) * 15
	if student.Credits < requiredCredits {
		return fmt.Errorf("student does not meet credit threshold (has %d, needs %d)", student.Credits, requiredCredits)
	}

	// 2. Check if already released
	released, err := s.repo.HasReleased(ctx, fund.ID, student.Hash, trancheIndex)
	if err != nil {
		return fmt.Errorf("check has_released: %w", err)
	}
	if released {
		return nil // already processed
	}

	var sHash [32]byte
	copy(sHash[:], common.FromHex(student.Hash))
	recipAddr := common.HexToAddress(recipient)
	onChainFundId := new(big.Int).SetUint64(fund.OnChainID)

	// 3. Compute EIP-712 Signature
	// uint256 nonce = uint256(keccak256(abi.encode(fundId, studentHash, trancheIndex)));
	nonceHash := crypto.Keccak256(
		common.LeftPadBytes(onChainFundId.Bytes(), 32),
		sHash[:],
		common.LeftPadBytes(big.NewInt(int64(trancheIndex)).Bytes(), 32),
	)
	nonce := new(big.Int).SetBytes(nonceHash)

	// TRANCHE_RELEASE_TYPEHASH = keccak256("TrancheRelease(uint256 fundId,bytes32 studentHash,uint256 trancheIndex,address recipient,uint256 nonce)")
	typeHash := crypto.Keccak256([]byte("TrancheRelease(uint256 fundId,bytes32 studentHash,uint256 trancheIndex,address recipient,uint256 nonce)"))

	// structHash = keccak256(abi.encode(typeHash, fundId, studentHash, trancheIndex, recipient, nonce))
	structHash := crypto.Keccak256(
		typeHash,
		common.LeftPadBytes(onChainFundId.Bytes(), 32),
		sHash[:],
		common.LeftPadBytes(big.NewInt(int64(trancheIndex)).Bytes(), 32),
		common.LeftPadBytes(recipAddr.Bytes(), 32),
		common.LeftPadBytes(nonce.Bytes(), 32),
	)

	// digest = keccak256(abi.encodePacked("\x19\x01", DOMAIN_SEPARATOR, structHash))
	digest := crypto.Keccak256(
		[]byte{0x19, 0x01},
		s.domainSep[:],
		structHash,
	)

	sig, err := crypto.Sign(digest, s.attestorKey)
	if err != nil {
		return fmt.Errorf("sign EIP-712 digest: %w", err)
	}
	// Ethereum v is 27 or 28, but crypto.Sign outputs 0 or 1.
	sig[64] += 27

	// 4. Record pending release in DB
	tr := &domain.TrancheRelease{
		FundID:        fund.ID,
		OnChainFundID: fund.OnChainID,
		StudentHash:   student.Hash,
		TrancheIndex:  trancheIndex,
		Amount:        fund.TrancheAmount,
		Recipient:     recipient,
	}
	if err := s.repo.CreateTrancheRelease(ctx, tr); err != nil {
		return fmt.Errorf("create tranche release record: %w", err)
	}

	// 5. Submit to blockchain
	opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)
	if err != nil {
		return fmt.Errorf("transact opts: %w", err)
	}

	tx, err := s.escrow.ScholarshipEscrowContractTransactor.ReleaseTranche(
		opts,
		onChainFundId,
		sHash,
		big.NewInt(int64(trancheIndex)),
		recipAddr,
		sig,
	)
	if err != nil {
		rollback()
		return fmt.Errorf("releaseTranche tx: %w", err)
	}
	confirm()
	
	s.log.Info("tranche released", 
		zap.Int64("fund_id", fund.ID), 
		zap.String("student", student.StudentID),
		zap.Int("tranche", trancheIndex),
		zap.String("tx_hash", tx.Hash().Hex()),
	)
	return nil
}
