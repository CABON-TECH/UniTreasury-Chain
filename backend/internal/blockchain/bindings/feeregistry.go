// Package bindings contains Go bindings for UniTreasury Chain smart contracts.
//
// These are hand-written stubs that match the Solidity ABI exactly.
// Run `make generate-bindings` (requires abigen) to regenerate from Foundry output.
// The generated file will be API-compatible with these stubs.
package bindings

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"strings"
)

// FeeRegistryContractABI is the ABI JSON for FeeRegistryContract.
// Generated from: contracts/out/FeeRegistryContract.sol/FeeRegistryContract.json
const FeeRegistryContractABI = `[
  {"type":"constructor","inputs":[{"name":"admin","type":"address"},{"name":"recorder","type":"address"}],"stateMutability":"nonpayable"},
  {"type":"function","name":"recordPayment","inputs":[{"name":"studentHash","type":"bytes32"},{"name":"receiptHash","type":"bytes32"},{"name":"amount","type":"uint256"},{"name":"semester","type":"uint256"},{"name":"feeType","type":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},
  {"type":"function","name":"setFeeStructure","inputs":[{"name":"semester","type":"uint256"},{"name":"tuitionFee","type":"uint256"},{"name":"hostingFee","type":"uint256"},{"name":"examFee","type":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},
  {"type":"function","name":"deactivateFeeStructure","inputs":[{"name":"semester","type":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},
  {"type":"function","name":"checkpoint","inputs":[],"outputs":[],"stateMutability":"nonpayable"},
  {"type":"function","name":"getPaymentRecord","inputs":[{"name":"receiptHash","type":"bytes32"}],"outputs":[{"name":"","type":"tuple","components":[{"name":"studentHash","type":"bytes32"},{"name":"receiptHash","type":"bytes32"},{"name":"amount","type":"uint256"},{"name":"semester","type":"uint256"},{"name":"feeType","type":"uint256"},{"name":"recordedAt","type":"uint256"}]}],"stateMutability":"view"},
  {"type":"function","name":"hasStudentPaidSemester","inputs":[{"name":"studentHash","type":"bytes32"},{"name":"semester","type":"uint256"}],"outputs":[{"name":"","type":"bool"}],"stateMutability":"view"},
  {"type":"function","name":"getFeeStructure","inputs":[{"name":"semester","type":"uint256"}],"outputs":[{"name":"","type":"tuple","components":[{"name":"tuitionFee","type":"uint256"},{"name":"hostingFee","type":"uint256"},{"name":"examFee","type":"uint256"},{"name":"active","type":"bool"}]}],"stateMutability":"view"},
  {"type":"function","name":"getTotalVolume","inputs":[],"outputs":[{"name":"","type":"uint256"}],"stateMutability":"view"},
  {"type":"function","name":"getTotalRecords","inputs":[],"outputs":[{"name":"","type":"uint256"}],"stateMutability":"view"},
  {"type":"event","name":"PaymentRecorded","inputs":[{"name":"studentHash","type":"bytes32","indexed":true},{"name":"receiptHash","type":"bytes32","indexed":true},{"name":"amount","type":"uint256","indexed":false},{"name":"semester","type":"uint256","indexed":false},{"name":"timestamp","type":"uint256","indexed":false}],"anonymous":false},
  {"type":"event","name":"ReconciliationCheckpoint","inputs":[{"name":"totalRecords","type":"uint256","indexed":false},{"name":"totalVolume","type":"uint256","indexed":false},{"name":"blockNumber","type":"uint256","indexed":false}],"anonymous":false}
]`

// FeeRegistryContract is a Go binding for the FeeRegistryContract Solidity contract.
type FeeRegistryContract struct {
	FeeRegistryContractCaller
	FeeRegistryContractTransactor
	FeeRegistryContractFilterer
}

// FeeRegistryContractCaller contains read-only contract methods.
type FeeRegistryContractCaller struct {
	contract *bind.BoundContract
}

// FeeRegistryContractTransactor contains state-changing contract methods.
type FeeRegistryContractTransactor struct {
	contract *bind.BoundContract
}

// FeeRegistryContractFilterer contains log filtering methods.
type FeeRegistryContractFilterer struct {
	contract *bind.BoundContract
}

// NewFeeRegistryContract creates a new instance bound to the given address.
func NewFeeRegistryContract(address common.Address, backend bind.ContractBackend) (*FeeRegistryContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeRegistryContractABI))
	if err != nil {
		return nil, err
	}
	contract := bind.NewBoundContract(address, parsed, backend, backend, backend)
	return &FeeRegistryContract{
		FeeRegistryContractCaller:     FeeRegistryContractCaller{contract: contract},
		FeeRegistryContractTransactor: FeeRegistryContractTransactor{contract: contract},
		FeeRegistryContractFilterer:   FeeRegistryContractFilterer{contract: contract},
	}, nil
}

// ── Transactor methods ────────────────────────────────────────────────────────

// RecordPayment submits a fee payment to the on-chain registry.
func (t *FeeRegistryContractTransactor) RecordPayment(
	opts *bind.TransactOpts,
	studentHash [32]byte,
	receiptHash [32]byte,
	amount *big.Int,
	semester *big.Int,
	feeType *big.Int,
) (*types.Transaction, error) {
	return t.contract.Transact(opts, "recordPayment", studentHash, receiptHash, amount, semester, feeType)
}

// SetFeeStructure configures the fee schedule for a semester.
func (t *FeeRegistryContractTransactor) SetFeeStructure(
	opts *bind.TransactOpts,
	semester *big.Int,
	tuitionFee *big.Int,
	hostingFee *big.Int,
	examFee *big.Int,
) (*types.Transaction, error) {
	return t.contract.Transact(opts, "setFeeStructure", semester, tuitionFee, hostingFee, examFee)
}

// Checkpoint emits a reconciliation anchor event.
func (t *FeeRegistryContractTransactor) Checkpoint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return t.contract.Transact(opts, "checkpoint")
}

// ── Caller methods ────────────────────────────────────────────────────────────

// PaymentRecord mirrors the Solidity struct.
type PaymentRecord struct {
	StudentHash [32]byte
	ReceiptHash [32]byte
	Amount      *big.Int
	Semester    *big.Int
	FeeType     *big.Int
	RecordedAt  *big.Int
}

// GetPaymentRecord fetches a payment record by receipt hash.
func (c *FeeRegistryContractCaller) GetPaymentRecord(opts *bind.CallOpts, receiptHash [32]byte) (PaymentRecord, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "getPaymentRecord", receiptHash)
	if err != nil {
		return PaymentRecord{}, err
	}
	return *abi.ConvertType(out[0], new(PaymentRecord)).(*PaymentRecord), nil
}

// HasStudentPaidSemester checks whether a student has a confirmed payment for a semester.
func (c *FeeRegistryContractCaller) HasStudentPaidSemester(opts *bind.CallOpts, studentHash [32]byte, semester *big.Int) (bool, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "hasStudentPaidSemester", studentHash, semester)
	if err != nil {
		return false, err
	}
	return *abi.ConvertType(out[0], new(bool)).(*bool), nil
}

// GetTotalVolume returns the cumulative confirmed payment volume in wei.
func (c *FeeRegistryContractCaller) GetTotalVolume(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "getTotalVolume")
	if err != nil {
		return nil, err
	}
	return *abi.ConvertType(out[0], new(*big.Int)).(**big.Int), nil
}

// GetTotalRecords returns the total count of recorded payments.
func (c *FeeRegistryContractCaller) GetTotalRecords(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "getTotalRecords")
	if err != nil {
		return nil, err
	}
	return *abi.ConvertType(out[0], new(*big.Int)).(**big.Int), nil
}

// ── Event types ───────────────────────────────────────────────────────────────

// FeeRegistryPaymentRecorded represents a PaymentRecorded log event.
type FeeRegistryPaymentRecorded struct {
	StudentHash [32]byte
	ReceiptHash [32]byte
	Amount      *big.Int
	Semester    *big.Int
	Timestamp   *big.Int
	Raw         types.Log
}

// ParsePaymentRecorded decodes a PaymentRecorded event from a raw log.
func (f *FeeRegistryContractFilterer) ParsePaymentRecorded(log types.Log) (*FeeRegistryPaymentRecorded, error) {
	event := new(FeeRegistryPaymentRecorded)
	err := f.contract.UnpackLog(event, "PaymentRecorded", log)
	if err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WatchPaymentRecorded subscribes to PaymentRecorded events.
func (f *FeeRegistryContractFilterer) WatchPaymentRecorded(
	opts *bind.WatchOpts,
	sink chan<- *FeeRegistryPaymentRecorded,
	studentHash [][32]byte,
	receiptHash [][32]byte,
) (event.Subscription, error) {
	_, sub, err := f.contract.WatchLogs(opts, "PaymentRecorded")
	return sub, err
}
