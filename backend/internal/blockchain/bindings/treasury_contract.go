// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = time.Tick
	_ = context.Background
)

// ITreasurySignerChangeProposal is an auto generated low-level Go binding around an user-defined struct.
type ITreasurySignerChangeProposal struct {
	Id            *big.Int
	Proposer      common.Address
	TargetSigner  common.Address
	NewSigner     common.Address
	ChangeType    uint8
	Status        uint8
	ApprovalCount *big.Int
	CreatedAt     *big.Int
}

// ITreasuryWithdrawalProposal is an auto generated low-level Go binding around an user-defined struct.
type ITreasuryWithdrawalProposal struct {
	Id            *big.Int
	Proposer      common.Address
	Recipient     common.Address
	Amount        *big.Int
	Purpose       string
	Status        uint8
	ApprovalCount *big.Int
	CreatedAt     *big.Int
}

// TreasuryContractMetaData contains all meta data concerning the TreasuryContract contract.
var TreasuryContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialApprovers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_requiredApprovals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_dailyLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_usdcToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"APPROVER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EXECUTOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROPOSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approveSignerChange\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approveWithdrawal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelSignerChange\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelWithdrawal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dailyLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSignerChange\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeWithdrawal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freeze\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDailyWithdrawnAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProposal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITreasury.WithdrawalProposal\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"purpose\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumITreasury.ProposalStatus\"},{\"name\":\"approvalCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSignerProposal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITreasury.SignerChangeProposal\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"targetSigner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newSigner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"changeType\",\"type\":\"uint8\",\"internalType\":\"enumITreasury.ChangeType\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumITreasury.ProposalStatus\"},{\"name\":\"approvalCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasApproved\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasApprovedSignerChange\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isFrozen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeSignerChange\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"replacement\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"changeType\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proposeWithdrawal\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"purpose\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requiredApprovals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDailyLimit\",\"inputs\":[{\"name\":\"newLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRequiredApprovals\",\"inputs\":[{\"name\":\"newRequired\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unfreeze\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"usdcToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DailyLimitUpdated\",\"inputs\":[{\"name\":\"oldLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newBalance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequiredApprovalsUpdated\",\"inputs\":[{\"name\":\"oldRequired\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newRequired\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerChangeApproved\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"approver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approvalCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerChangeExecuted\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"replacement\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"changeType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerChangeProposed\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"replacement\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"changeType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TreasuryFrozen\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TreasuryUnfrozen\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalApproved\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"approver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approvalCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalCancelled\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"canceller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalExecuted\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalProposed\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"purpose\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Treasury__AlreadyApproved\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Treasury__DailyLimitExceeded\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"remainingToday\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Treasury__Frozen\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Treasury__InsufficientApprovals\",\"inputs\":[{\"name\":\"have\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"need\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Treasury__InsufficientBalance\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Treasury__NotAuthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Treasury__ProposalNotFound\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Treasury__ProposalNotPending\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Treasury__ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Treasury__ZeroAmount\",\"inputs\":[]}]",
}

// TreasuryContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TreasuryContractMetaData.ABI instead.
var TreasuryContractABI = TreasuryContractMetaData.ABI

// TreasuryContract is an auto generated Go binding around an Ethereum contract.
type TreasuryContract struct {
	TreasuryContractCaller     // Read-only binding to the contract
	TreasuryContractTransactor // Write-only binding to the contract
	TreasuryContractFilterer   // Log filterer for contract events
}

// TreasuryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TreasuryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TreasuryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TreasuryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TreasuryContractSession struct {
	Contract     *TreasuryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TreasuryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TreasuryContractCallerSession struct {
	Contract *TreasuryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TreasuryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TreasuryContractTransactorSession struct {
	Contract     *TreasuryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TreasuryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TreasuryContractRaw struct {
	Contract *TreasuryContract // Generic contract binding to access the raw methods on
}

// TreasuryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TreasuryContractCallerRaw struct {
	Contract *TreasuryContractCaller // Generic read-only contract binding to access the raw methods on
}

// TreasuryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TreasuryContractTransactorRaw struct {
	Contract *TreasuryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTreasuryContract creates a new instance of TreasuryContract, bound to a specific deployed contract.
func NewTreasuryContract(address common.Address, backend bind.ContractBackend) (*TreasuryContract, error) {
	contract, err := bindTreasuryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TreasuryContract{TreasuryContractCaller: TreasuryContractCaller{contract: contract}, TreasuryContractTransactor: TreasuryContractTransactor{contract: contract}, TreasuryContractFilterer: TreasuryContractFilterer{contract: contract}}, nil
}

// NewTreasuryContractCaller creates a new read-only instance of TreasuryContract, bound to a specific deployed contract.
func NewTreasuryContractCaller(address common.Address, caller bind.ContractCaller) (*TreasuryContractCaller, error) {
	contract, err := bindTreasuryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractCaller{contract: contract}, nil
}

// NewTreasuryContractTransactor creates a new write-only instance of TreasuryContract, bound to a specific deployed contract.
func NewTreasuryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TreasuryContractTransactor, error) {
	contract, err := bindTreasuryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractTransactor{contract: contract}, nil
}

// NewTreasuryContractFilterer creates a new log filterer instance of TreasuryContract, bound to a specific deployed contract.
func NewTreasuryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TreasuryContractFilterer, error) {
	contract, err := bindTreasuryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractFilterer{contract: contract}, nil
}

// bindTreasuryContract binds a generic wrapper to an already deployed contract.
func bindTreasuryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TreasuryContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TreasuryContract *TreasuryContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TreasuryContract.Contract.TreasuryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TreasuryContract *TreasuryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryContract.Contract.TreasuryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TreasuryContract *TreasuryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TreasuryContract.Contract.TreasuryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TreasuryContract *TreasuryContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TreasuryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TreasuryContract *TreasuryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TreasuryContract *TreasuryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TreasuryContract.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_TreasuryContract *TreasuryContractCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_TreasuryContract *TreasuryContractSession) ADMINROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.ADMINROLE(&_TreasuryContract.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_TreasuryContract *TreasuryContractCallerSession) ADMINROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.ADMINROLE(&_TreasuryContract.CallOpts)
}

// APPROVERROLE is a free data retrieval call binding the contract method 0x4245962b.
//
// Solidity: function APPROVER_ROLE() view returns(bytes32)
func (_TreasuryContract *TreasuryContractCaller) APPROVERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "APPROVER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// APPROVERROLE is a free data retrieval call binding the contract method 0x4245962b.
//
// Solidity: function APPROVER_ROLE() view returns(bytes32)
func (_TreasuryContract *TreasuryContractSession) APPROVERROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.APPROVERROLE(&_TreasuryContract.CallOpts)
}

// APPROVERROLE is a free data retrieval call binding the contract method 0x4245962b.
//
// Solidity: function APPROVER_ROLE() view returns(bytes32)
func (_TreasuryContract *TreasuryContractCallerSession) APPROVERROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.APPROVERROLE(&_TreasuryContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TreasuryContract *TreasuryContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TreasuryContract *TreasuryContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.DEFAULTADMINROLE(&_TreasuryContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TreasuryContract *TreasuryContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.DEFAULTADMINROLE(&_TreasuryContract.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_TreasuryContract *TreasuryContractCaller) EXECUTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "EXECUTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_TreasuryContract *TreasuryContractSession) EXECUTORROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.EXECUTORROLE(&_TreasuryContract.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_TreasuryContract *TreasuryContractCallerSession) EXECUTORROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.EXECUTORROLE(&_TreasuryContract.CallOpts)
}

// PROPOSERROLE is a free data retrieval call binding the contract method 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (_TreasuryContract *TreasuryContractCaller) PROPOSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "PROPOSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROPOSERROLE is a free data retrieval call binding the contract method 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (_TreasuryContract *TreasuryContractSession) PROPOSERROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.PROPOSERROLE(&_TreasuryContract.CallOpts)
}

// PROPOSERROLE is a free data retrieval call binding the contract method 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (_TreasuryContract *TreasuryContractCallerSession) PROPOSERROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.PROPOSERROLE(&_TreasuryContract.CallOpts)
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() view returns(uint256)
func (_TreasuryContract *TreasuryContractCaller) DailyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "dailyLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() view returns(uint256)
func (_TreasuryContract *TreasuryContractSession) DailyLimit() (*big.Int, error) {
	return _TreasuryContract.Contract.DailyLimit(&_TreasuryContract.CallOpts)
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() view returns(uint256)
func (_TreasuryContract *TreasuryContractCallerSession) DailyLimit() (*big.Int, error) {
	return _TreasuryContract.Contract.DailyLimit(&_TreasuryContract.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_TreasuryContract *TreasuryContractCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_TreasuryContract *TreasuryContractSession) GetBalance() (*big.Int, error) {
	return _TreasuryContract.Contract.GetBalance(&_TreasuryContract.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_TreasuryContract *TreasuryContractCallerSession) GetBalance() (*big.Int, error) {
	return _TreasuryContract.Contract.GetBalance(&_TreasuryContract.CallOpts)
}

// GetDailyWithdrawnAmount is a free data retrieval call binding the contract method 0xb4b5ed27.
//
// Solidity: function getDailyWithdrawnAmount() view returns(uint256)
func (_TreasuryContract *TreasuryContractCaller) GetDailyWithdrawnAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "getDailyWithdrawnAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDailyWithdrawnAmount is a free data retrieval call binding the contract method 0xb4b5ed27.
//
// Solidity: function getDailyWithdrawnAmount() view returns(uint256)
func (_TreasuryContract *TreasuryContractSession) GetDailyWithdrawnAmount() (*big.Int, error) {
	return _TreasuryContract.Contract.GetDailyWithdrawnAmount(&_TreasuryContract.CallOpts)
}

// GetDailyWithdrawnAmount is a free data retrieval call binding the contract method 0xb4b5ed27.
//
// Solidity: function getDailyWithdrawnAmount() view returns(uint256)
func (_TreasuryContract *TreasuryContractCallerSession) GetDailyWithdrawnAmount() (*big.Int, error) {
	return _TreasuryContract.Contract.GetDailyWithdrawnAmount(&_TreasuryContract.CallOpts)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 proposalId) view returns((uint256,address,address,uint256,string,uint8,uint256,uint256))
func (_TreasuryContract *TreasuryContractCaller) GetProposal(opts *bind.CallOpts, proposalId *big.Int) (ITreasuryWithdrawalProposal, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "getProposal", proposalId)

	if err != nil {
		return *new(ITreasuryWithdrawalProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(ITreasuryWithdrawalProposal)).(*ITreasuryWithdrawalProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 proposalId) view returns((uint256,address,address,uint256,string,uint8,uint256,uint256))
func (_TreasuryContract *TreasuryContractSession) GetProposal(proposalId *big.Int) (ITreasuryWithdrawalProposal, error) {
	return _TreasuryContract.Contract.GetProposal(&_TreasuryContract.CallOpts, proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 proposalId) view returns((uint256,address,address,uint256,string,uint8,uint256,uint256))
func (_TreasuryContract *TreasuryContractCallerSession) GetProposal(proposalId *big.Int) (ITreasuryWithdrawalProposal, error) {
	return _TreasuryContract.Contract.GetProposal(&_TreasuryContract.CallOpts, proposalId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TreasuryContract *TreasuryContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TreasuryContract *TreasuryContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TreasuryContract.Contract.GetRoleAdmin(&_TreasuryContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TreasuryContract *TreasuryContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TreasuryContract.Contract.GetRoleAdmin(&_TreasuryContract.CallOpts, role)
}

// GetSignerProposal is a free data retrieval call binding the contract method 0xee9ac367.
//
// Solidity: function getSignerProposal(uint256 proposalId) view returns((uint256,address,address,address,uint8,uint8,uint256,uint256))
func (_TreasuryContract *TreasuryContractCaller) GetSignerProposal(opts *bind.CallOpts, proposalId *big.Int) (ITreasurySignerChangeProposal, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "getSignerProposal", proposalId)

	if err != nil {
		return *new(ITreasurySignerChangeProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(ITreasurySignerChangeProposal)).(*ITreasurySignerChangeProposal)

	return out0, err

}

// GetSignerProposal is a free data retrieval call binding the contract method 0xee9ac367.
//
// Solidity: function getSignerProposal(uint256 proposalId) view returns((uint256,address,address,address,uint8,uint8,uint256,uint256))
func (_TreasuryContract *TreasuryContractSession) GetSignerProposal(proposalId *big.Int) (ITreasurySignerChangeProposal, error) {
	return _TreasuryContract.Contract.GetSignerProposal(&_TreasuryContract.CallOpts, proposalId)
}

// GetSignerProposal is a free data retrieval call binding the contract method 0xee9ac367.
//
// Solidity: function getSignerProposal(uint256 proposalId) view returns((uint256,address,address,address,uint8,uint8,uint256,uint256))
func (_TreasuryContract *TreasuryContractCallerSession) GetSignerProposal(proposalId *big.Int) (ITreasurySignerChangeProposal, error) {
	return _TreasuryContract.Contract.GetSignerProposal(&_TreasuryContract.CallOpts, proposalId)
}

// HasApproved is a free data retrieval call binding the contract method 0x2358d5a8.
//
// Solidity: function hasApproved(uint256 proposalId, address approver) view returns(bool)
func (_TreasuryContract *TreasuryContractCaller) HasApproved(opts *bind.CallOpts, proposalId *big.Int, approver common.Address) (bool, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "hasApproved", proposalId, approver)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasApproved is a free data retrieval call binding the contract method 0x2358d5a8.
//
// Solidity: function hasApproved(uint256 proposalId, address approver) view returns(bool)
func (_TreasuryContract *TreasuryContractSession) HasApproved(proposalId *big.Int, approver common.Address) (bool, error) {
	return _TreasuryContract.Contract.HasApproved(&_TreasuryContract.CallOpts, proposalId, approver)
}

// HasApproved is a free data retrieval call binding the contract method 0x2358d5a8.
//
// Solidity: function hasApproved(uint256 proposalId, address approver) view returns(bool)
func (_TreasuryContract *TreasuryContractCallerSession) HasApproved(proposalId *big.Int, approver common.Address) (bool, error) {
	return _TreasuryContract.Contract.HasApproved(&_TreasuryContract.CallOpts, proposalId, approver)
}

// HasApprovedSignerChange is a free data retrieval call binding the contract method 0xf12ca327.
//
// Solidity: function hasApprovedSignerChange(uint256 proposalId, address approver) view returns(bool)
func (_TreasuryContract *TreasuryContractCaller) HasApprovedSignerChange(opts *bind.CallOpts, proposalId *big.Int, approver common.Address) (bool, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "hasApprovedSignerChange", proposalId, approver)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasApprovedSignerChange is a free data retrieval call binding the contract method 0xf12ca327.
//
// Solidity: function hasApprovedSignerChange(uint256 proposalId, address approver) view returns(bool)
func (_TreasuryContract *TreasuryContractSession) HasApprovedSignerChange(proposalId *big.Int, approver common.Address) (bool, error) {
	return _TreasuryContract.Contract.HasApprovedSignerChange(&_TreasuryContract.CallOpts, proposalId, approver)
}

// HasApprovedSignerChange is a free data retrieval call binding the contract method 0xf12ca327.
//
// Solidity: function hasApprovedSignerChange(uint256 proposalId, address approver) view returns(bool)
func (_TreasuryContract *TreasuryContractCallerSession) HasApprovedSignerChange(proposalId *big.Int, approver common.Address) (bool, error) {
	return _TreasuryContract.Contract.HasApprovedSignerChange(&_TreasuryContract.CallOpts, proposalId, approver)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TreasuryContract *TreasuryContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TreasuryContract *TreasuryContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TreasuryContract.Contract.HasRole(&_TreasuryContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TreasuryContract *TreasuryContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TreasuryContract.Contract.HasRole(&_TreasuryContract.CallOpts, role, account)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_TreasuryContract *TreasuryContractCaller) IsFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "isFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_TreasuryContract *TreasuryContractSession) IsFrozen() (bool, error) {
	return _TreasuryContract.Contract.IsFrozen(&_TreasuryContract.CallOpts)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_TreasuryContract *TreasuryContractCallerSession) IsFrozen() (bool, error) {
	return _TreasuryContract.Contract.IsFrozen(&_TreasuryContract.CallOpts)
}

// RequiredApprovals is a free data retrieval call binding the contract method 0x99c1aadc.
//
// Solidity: function requiredApprovals() view returns(uint256)
func (_TreasuryContract *TreasuryContractCaller) RequiredApprovals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "requiredApprovals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredApprovals is a free data retrieval call binding the contract method 0x99c1aadc.
//
// Solidity: function requiredApprovals() view returns(uint256)
func (_TreasuryContract *TreasuryContractSession) RequiredApprovals() (*big.Int, error) {
	return _TreasuryContract.Contract.RequiredApprovals(&_TreasuryContract.CallOpts)
}

// RequiredApprovals is a free data retrieval call binding the contract method 0x99c1aadc.
//
// Solidity: function requiredApprovals() view returns(uint256)
func (_TreasuryContract *TreasuryContractCallerSession) RequiredApprovals() (*big.Int, error) {
	return _TreasuryContract.Contract.RequiredApprovals(&_TreasuryContract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TreasuryContract *TreasuryContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TreasuryContract *TreasuryContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TreasuryContract.Contract.SupportsInterface(&_TreasuryContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TreasuryContract *TreasuryContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TreasuryContract.Contract.SupportsInterface(&_TreasuryContract.CallOpts, interfaceId)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_TreasuryContract *TreasuryContractCaller) UsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "usdcToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_TreasuryContract *TreasuryContractSession) UsdcToken() (common.Address, error) {
	return _TreasuryContract.Contract.UsdcToken(&_TreasuryContract.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_TreasuryContract *TreasuryContractCallerSession) UsdcToken() (common.Address, error) {
	return _TreasuryContract.Contract.UsdcToken(&_TreasuryContract.CallOpts)
}

// ApproveSignerChange is a paid mutator transaction binding the contract method 0xe5531e8a.
//
// Solidity: function approveSignerChange(uint256 proposalId) returns()
func (_TreasuryContract *TreasuryContractTransactor) ApproveSignerChange(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "approveSignerChange", proposalId)
}

// ApproveSignerChange is a paid mutator transaction binding the contract method 0xe5531e8a.
//
// Solidity: function approveSignerChange(uint256 proposalId) returns()
func (_TreasuryContract *TreasuryContractSession) ApproveSignerChange(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ApproveSignerChange(&_TreasuryContract.TransactOpts, proposalId)
}

// ApproveSignerChange is a paid mutator transaction binding the contract method 0xe5531e8a.
//
// Solidity: function approveSignerChange(uint256 proposalId) returns()
func (_TreasuryContract *TreasuryContractTransactorSession) ApproveSignerChange(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ApproveSignerChange(&_TreasuryContract.TransactOpts, proposalId)
}

// ApproveWithdrawal is a paid mutator transaction binding the contract method 0x9eceddea.
//
// Solidity: function approveWithdrawal(uint256 proposalId) returns()
func (_TreasuryContract *TreasuryContractTransactor) ApproveWithdrawal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "approveWithdrawal", proposalId)
}

// ApproveWithdrawal is a paid mutator transaction binding the contract method 0x9eceddea.
//
// Solidity: function approveWithdrawal(uint256 proposalId) returns()
func (_TreasuryContract *TreasuryContractSession) ApproveWithdrawal(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ApproveWithdrawal(&_TreasuryContract.TransactOpts, proposalId)
}

// ApproveWithdrawal is a paid mutator transaction binding the contract method 0x9eceddea.
//
// Solidity: function approveWithdrawal(uint256 proposalId) returns()
func (_TreasuryContract *TreasuryContractTransactorSession) ApproveWithdrawal(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ApproveWithdrawal(&_TreasuryContract.TransactOpts, proposalId)
}

// CancelSignerChange is a paid mutator transaction binding the contract method 0xec9b4cd1.
//
// Solidity: function cancelSignerChange(uint256 proposalId) returns()
func (_TreasuryContract *TreasuryContractTransactor) CancelSignerChange(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "cancelSignerChange", proposalId)
}

// CancelSignerChange is a paid mutator transaction binding the contract method 0xec9b4cd1.
//
// Solidity: function cancelSignerChange(uint256 proposalId) returns()
func (_TreasuryContract *TreasuryContractSession) CancelSignerChange(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.CancelSignerChange(&_TreasuryContract.TransactOpts, proposalId)
}

// CancelSignerChange is a paid mutator transaction binding the contract method 0xec9b4cd1.
//
// Solidity: function cancelSignerChange(uint256 proposalId) returns()
func (_TreasuryContract *TreasuryContractTransactorSession) CancelSignerChange(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.CancelSignerChange(&_TreasuryContract.TransactOpts, proposalId)
}

// CancelWithdrawal is a paid mutator transaction binding the contract method 0x3efcfda4.
//
// Solidity: function cancelWithdrawal(uint256 proposalId) returns()
func (_TreasuryContract *TreasuryContractTransactor) CancelWithdrawal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "cancelWithdrawal", proposalId)
}

// CancelWithdrawal is a paid mutator transaction binding the contract method 0x3efcfda4.
//
// Solidity: function cancelWithdrawal(uint256 proposalId) returns()
func (_TreasuryContract *TreasuryContractSession) CancelWithdrawal(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.CancelWithdrawal(&_TreasuryContract.TransactOpts, proposalId)
}

// CancelWithdrawal is a paid mutator transaction binding the contract method 0x3efcfda4.
//
// Solidity: function cancelWithdrawal(uint256 proposalId) returns()
func (_TreasuryContract *TreasuryContractTransactorSession) CancelWithdrawal(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.CancelWithdrawal(&_TreasuryContract.TransactOpts, proposalId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_TreasuryContract *TreasuryContractTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_TreasuryContract *TreasuryContractSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.Deposit(&_TreasuryContract.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_TreasuryContract *TreasuryContractTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.Deposit(&_TreasuryContract.TransactOpts, amount)
}

// ExecuteSignerChange is a paid mutator transaction binding the contract method 0x5efde9a4.
//
// Solidity: function executeSignerChange(uint256 proposalId) returns()
func (_TreasuryContract *TreasuryContractTransactor) ExecuteSignerChange(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "executeSignerChange", proposalId)
}

// ExecuteSignerChange is a paid mutator transaction binding the contract method 0x5efde9a4.
//
// Solidity: function executeSignerChange(uint256 proposalId) returns()
func (_TreasuryContract *TreasuryContractSession) ExecuteSignerChange(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ExecuteSignerChange(&_TreasuryContract.TransactOpts, proposalId)
}

// ExecuteSignerChange is a paid mutator transaction binding the contract method 0x5efde9a4.
//
// Solidity: function executeSignerChange(uint256 proposalId) returns()
func (_TreasuryContract *TreasuryContractTransactorSession) ExecuteSignerChange(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ExecuteSignerChange(&_TreasuryContract.TransactOpts, proposalId)
}

// ExecuteWithdrawal is a paid mutator transaction binding the contract method 0x24f13a76.
//
// Solidity: function executeWithdrawal(uint256 proposalId) returns()
func (_TreasuryContract *TreasuryContractTransactor) ExecuteWithdrawal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "executeWithdrawal", proposalId)
}

// ExecuteWithdrawal is a paid mutator transaction binding the contract method 0x24f13a76.
//
// Solidity: function executeWithdrawal(uint256 proposalId) returns()
func (_TreasuryContract *TreasuryContractSession) ExecuteWithdrawal(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ExecuteWithdrawal(&_TreasuryContract.TransactOpts, proposalId)
}

// ExecuteWithdrawal is a paid mutator transaction binding the contract method 0x24f13a76.
//
// Solidity: function executeWithdrawal(uint256 proposalId) returns()
func (_TreasuryContract *TreasuryContractTransactorSession) ExecuteWithdrawal(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ExecuteWithdrawal(&_TreasuryContract.TransactOpts, proposalId)
}

// Freeze is a paid mutator transaction binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() returns()
func (_TreasuryContract *TreasuryContractTransactor) Freeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "freeze")
}

// Freeze is a paid mutator transaction binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() returns()
func (_TreasuryContract *TreasuryContractSession) Freeze() (*types.Transaction, error) {
	return _TreasuryContract.Contract.Freeze(&_TreasuryContract.TransactOpts)
}

// Freeze is a paid mutator transaction binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() returns()
func (_TreasuryContract *TreasuryContractTransactorSession) Freeze() (*types.Transaction, error) {
	return _TreasuryContract.Contract.Freeze(&_TreasuryContract.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TreasuryContract *TreasuryContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TreasuryContract *TreasuryContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasuryContract.Contract.GrantRole(&_TreasuryContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TreasuryContract *TreasuryContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasuryContract.Contract.GrantRole(&_TreasuryContract.TransactOpts, role, account)
}

// ProposeSignerChange is a paid mutator transaction binding the contract method 0xf3816fad.
//
// Solidity: function proposeSignerChange(address target, address replacement, uint8 changeType) returns(uint256 proposalId)
func (_TreasuryContract *TreasuryContractTransactor) ProposeSignerChange(opts *bind.TransactOpts, target common.Address, replacement common.Address, changeType uint8) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "proposeSignerChange", target, replacement, changeType)
}

// ProposeSignerChange is a paid mutator transaction binding the contract method 0xf3816fad.
//
// Solidity: function proposeSignerChange(address target, address replacement, uint8 changeType) returns(uint256 proposalId)
func (_TreasuryContract *TreasuryContractSession) ProposeSignerChange(target common.Address, replacement common.Address, changeType uint8) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ProposeSignerChange(&_TreasuryContract.TransactOpts, target, replacement, changeType)
}

// ProposeSignerChange is a paid mutator transaction binding the contract method 0xf3816fad.
//
// Solidity: function proposeSignerChange(address target, address replacement, uint8 changeType) returns(uint256 proposalId)
func (_TreasuryContract *TreasuryContractTransactorSession) ProposeSignerChange(target common.Address, replacement common.Address, changeType uint8) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ProposeSignerChange(&_TreasuryContract.TransactOpts, target, replacement, changeType)
}

// ProposeWithdrawal is a paid mutator transaction binding the contract method 0x2cdfb41d.
//
// Solidity: function proposeWithdrawal(address recipient, uint256 amount, string purpose) returns(uint256 proposalId)
func (_TreasuryContract *TreasuryContractTransactor) ProposeWithdrawal(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, purpose string) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "proposeWithdrawal", recipient, amount, purpose)
}

// ProposeWithdrawal is a paid mutator transaction binding the contract method 0x2cdfb41d.
//
// Solidity: function proposeWithdrawal(address recipient, uint256 amount, string purpose) returns(uint256 proposalId)
func (_TreasuryContract *TreasuryContractSession) ProposeWithdrawal(recipient common.Address, amount *big.Int, purpose string) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ProposeWithdrawal(&_TreasuryContract.TransactOpts, recipient, amount, purpose)
}

// ProposeWithdrawal is a paid mutator transaction binding the contract method 0x2cdfb41d.
//
// Solidity: function proposeWithdrawal(address recipient, uint256 amount, string purpose) returns(uint256 proposalId)
func (_TreasuryContract *TreasuryContractTransactorSession) ProposeWithdrawal(recipient common.Address, amount *big.Int, purpose string) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ProposeWithdrawal(&_TreasuryContract.TransactOpts, recipient, amount, purpose)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TreasuryContract *TreasuryContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TreasuryContract *TreasuryContractSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TreasuryContract.Contract.RenounceRole(&_TreasuryContract.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TreasuryContract *TreasuryContractTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TreasuryContract.Contract.RenounceRole(&_TreasuryContract.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TreasuryContract *TreasuryContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TreasuryContract *TreasuryContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasuryContract.Contract.RevokeRole(&_TreasuryContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TreasuryContract *TreasuryContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasuryContract.Contract.RevokeRole(&_TreasuryContract.TransactOpts, role, account)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0xb20d30a9.
//
// Solidity: function setDailyLimit(uint256 newLimit) returns()
func (_TreasuryContract *TreasuryContractTransactor) SetDailyLimit(opts *bind.TransactOpts, newLimit *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "setDailyLimit", newLimit)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0xb20d30a9.
//
// Solidity: function setDailyLimit(uint256 newLimit) returns()
func (_TreasuryContract *TreasuryContractSession) SetDailyLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.SetDailyLimit(&_TreasuryContract.TransactOpts, newLimit)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0xb20d30a9.
//
// Solidity: function setDailyLimit(uint256 newLimit) returns()
func (_TreasuryContract *TreasuryContractTransactorSession) SetDailyLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.SetDailyLimit(&_TreasuryContract.TransactOpts, newLimit)
}

// SetRequiredApprovals is a paid mutator transaction binding the contract method 0x222a242e.
//
// Solidity: function setRequiredApprovals(uint256 newRequired) returns()
func (_TreasuryContract *TreasuryContractTransactor) SetRequiredApprovals(opts *bind.TransactOpts, newRequired *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "setRequiredApprovals", newRequired)
}

// SetRequiredApprovals is a paid mutator transaction binding the contract method 0x222a242e.
//
// Solidity: function setRequiredApprovals(uint256 newRequired) returns()
func (_TreasuryContract *TreasuryContractSession) SetRequiredApprovals(newRequired *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.SetRequiredApprovals(&_TreasuryContract.TransactOpts, newRequired)
}

// SetRequiredApprovals is a paid mutator transaction binding the contract method 0x222a242e.
//
// Solidity: function setRequiredApprovals(uint256 newRequired) returns()
func (_TreasuryContract *TreasuryContractTransactorSession) SetRequiredApprovals(newRequired *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.SetRequiredApprovals(&_TreasuryContract.TransactOpts, newRequired)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6a28f000.
//
// Solidity: function unfreeze() returns()
func (_TreasuryContract *TreasuryContractTransactor) Unfreeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "unfreeze")
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6a28f000.
//
// Solidity: function unfreeze() returns()
func (_TreasuryContract *TreasuryContractSession) Unfreeze() (*types.Transaction, error) {
	return _TreasuryContract.Contract.Unfreeze(&_TreasuryContract.TransactOpts)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6a28f000.
//
// Solidity: function unfreeze() returns()
func (_TreasuryContract *TreasuryContractTransactorSession) Unfreeze() (*types.Transaction, error) {
	return _TreasuryContract.Contract.Unfreeze(&_TreasuryContract.TransactOpts)
}

// TreasuryContractDailyLimitUpdatedIterator is returned from FilterDailyLimitUpdated and is used to iterate over the raw logs and unpacked data for DailyLimitUpdated events raised by the TreasuryContract contract.
type TreasuryContractDailyLimitUpdatedIterator struct {
	Event *TreasuryContractDailyLimitUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryContractDailyLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryContractDailyLimitUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TreasuryContractDailyLimitUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TreasuryContractDailyLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryContractDailyLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryContractDailyLimitUpdated represents a DailyLimitUpdated event raised by the TreasuryContract contract.
type TreasuryContractDailyLimitUpdated struct {
	OldLimit *big.Int
	NewLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDailyLimitUpdated is a free log retrieval operation binding the contract event 0x207c4cbdf55ec315a13f0d5e047732ec5d947da056e706593aa509909941cedf.
//
// Solidity: event DailyLimitUpdated(uint256 oldLimit, uint256 newLimit)
func (_TreasuryContract *TreasuryContractFilterer) FilterDailyLimitUpdated(opts *bind.FilterOpts) (*TreasuryContractDailyLimitUpdatedIterator, error) {

	logs, sub, err := _TreasuryContract.contract.FilterLogs(opts, "DailyLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &TreasuryContractDailyLimitUpdatedIterator{contract: _TreasuryContract.contract, event: "DailyLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchDailyLimitUpdated is a free log subscription operation binding the contract event 0x207c4cbdf55ec315a13f0d5e047732ec5d947da056e706593aa509909941cedf.
//
// Solidity: event DailyLimitUpdated(uint256 oldLimit, uint256 newLimit)
func (_TreasuryContract *TreasuryContractFilterer) WatchDailyLimitUpdated(opts *bind.WatchOpts, sink chan<- *TreasuryContractDailyLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _TreasuryContract.contract.WatchLogs(opts, "DailyLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryContractDailyLimitUpdated)
				if err := _TreasuryContract.contract.UnpackLog(event, "DailyLimitUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDailyLimitUpdated is a log parse operation binding the contract event 0x207c4cbdf55ec315a13f0d5e047732ec5d947da056e706593aa509909941cedf.
//
// Solidity: event DailyLimitUpdated(uint256 oldLimit, uint256 newLimit)
func (_TreasuryContract *TreasuryContractFilterer) ParseDailyLimitUpdated(log types.Log) (*TreasuryContractDailyLimitUpdated, error) {
	event := new(TreasuryContractDailyLimitUpdated)
	if err := _TreasuryContract.contract.UnpackLog(event, "DailyLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryContractDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the TreasuryContract contract.
type TreasuryContractDepositedIterator struct {
	Event *TreasuryContractDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryContractDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryContractDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TreasuryContractDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TreasuryContractDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryContractDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryContractDeposited represents a Deposited event raised by the TreasuryContract contract.
type TreasuryContractDeposited struct {
	Depositor  common.Address
	Amount     *big.Int
	NewBalance *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed depositor, uint256 amount, uint256 newBalance)
func (_TreasuryContract *TreasuryContractFilterer) FilterDeposited(opts *bind.FilterOpts, depositor []common.Address) (*TreasuryContractDepositedIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _TreasuryContract.contract.FilterLogs(opts, "Deposited", depositorRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractDepositedIterator{contract: _TreasuryContract.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed depositor, uint256 amount, uint256 newBalance)
func (_TreasuryContract *TreasuryContractFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *TreasuryContractDeposited, depositor []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _TreasuryContract.contract.WatchLogs(opts, "Deposited", depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryContractDeposited)
				if err := _TreasuryContract.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed depositor, uint256 amount, uint256 newBalance)
func (_TreasuryContract *TreasuryContractFilterer) ParseDeposited(log types.Log) (*TreasuryContractDeposited, error) {
	event := new(TreasuryContractDeposited)
	if err := _TreasuryContract.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryContractRequiredApprovalsUpdatedIterator is returned from FilterRequiredApprovalsUpdated and is used to iterate over the raw logs and unpacked data for RequiredApprovalsUpdated events raised by the TreasuryContract contract.
type TreasuryContractRequiredApprovalsUpdatedIterator struct {
	Event *TreasuryContractRequiredApprovalsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryContractRequiredApprovalsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryContractRequiredApprovalsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TreasuryContractRequiredApprovalsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TreasuryContractRequiredApprovalsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryContractRequiredApprovalsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryContractRequiredApprovalsUpdated represents a RequiredApprovalsUpdated event raised by the TreasuryContract contract.
type TreasuryContractRequiredApprovalsUpdated struct {
	OldRequired *big.Int
	NewRequired *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRequiredApprovalsUpdated is a free log retrieval operation binding the contract event 0xcbd3eb803c8088c0677f88231c298df176fddba3dfd2a266ebcc538c71885ef6.
//
// Solidity: event RequiredApprovalsUpdated(uint256 oldRequired, uint256 newRequired)
func (_TreasuryContract *TreasuryContractFilterer) FilterRequiredApprovalsUpdated(opts *bind.FilterOpts) (*TreasuryContractRequiredApprovalsUpdatedIterator, error) {

	logs, sub, err := _TreasuryContract.contract.FilterLogs(opts, "RequiredApprovalsUpdated")
	if err != nil {
		return nil, err
	}
	return &TreasuryContractRequiredApprovalsUpdatedIterator{contract: _TreasuryContract.contract, event: "RequiredApprovalsUpdated", logs: logs, sub: sub}, nil
}

// WatchRequiredApprovalsUpdated is a free log subscription operation binding the contract event 0xcbd3eb803c8088c0677f88231c298df176fddba3dfd2a266ebcc538c71885ef6.
//
// Solidity: event RequiredApprovalsUpdated(uint256 oldRequired, uint256 newRequired)
func (_TreasuryContract *TreasuryContractFilterer) WatchRequiredApprovalsUpdated(opts *bind.WatchOpts, sink chan<- *TreasuryContractRequiredApprovalsUpdated) (event.Subscription, error) {

	logs, sub, err := _TreasuryContract.contract.WatchLogs(opts, "RequiredApprovalsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryContractRequiredApprovalsUpdated)
				if err := _TreasuryContract.contract.UnpackLog(event, "RequiredApprovalsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequiredApprovalsUpdated is a log parse operation binding the contract event 0xcbd3eb803c8088c0677f88231c298df176fddba3dfd2a266ebcc538c71885ef6.
//
// Solidity: event RequiredApprovalsUpdated(uint256 oldRequired, uint256 newRequired)
func (_TreasuryContract *TreasuryContractFilterer) ParseRequiredApprovalsUpdated(log types.Log) (*TreasuryContractRequiredApprovalsUpdated, error) {
	event := new(TreasuryContractRequiredApprovalsUpdated)
	if err := _TreasuryContract.contract.UnpackLog(event, "RequiredApprovalsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TreasuryContract contract.
type TreasuryContractRoleAdminChangedIterator struct {
	Event *TreasuryContractRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryContractRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TreasuryContractRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TreasuryContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryContractRoleAdminChanged represents a RoleAdminChanged event raised by the TreasuryContract contract.
type TreasuryContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TreasuryContract *TreasuryContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TreasuryContractRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TreasuryContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractRoleAdminChangedIterator{contract: _TreasuryContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TreasuryContract *TreasuryContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TreasuryContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TreasuryContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryContractRoleAdminChanged)
				if err := _TreasuryContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TreasuryContract *TreasuryContractFilterer) ParseRoleAdminChanged(log types.Log) (*TreasuryContractRoleAdminChanged, error) {
	event := new(TreasuryContractRoleAdminChanged)
	if err := _TreasuryContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TreasuryContract contract.
type TreasuryContractRoleGrantedIterator struct {
	Event *TreasuryContractRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryContractRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TreasuryContractRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TreasuryContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryContractRoleGranted represents a RoleGranted event raised by the TreasuryContract contract.
type TreasuryContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasuryContract *TreasuryContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TreasuryContractRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasuryContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractRoleGrantedIterator{contract: _TreasuryContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasuryContract *TreasuryContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TreasuryContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasuryContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryContractRoleGranted)
				if err := _TreasuryContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasuryContract *TreasuryContractFilterer) ParseRoleGranted(log types.Log) (*TreasuryContractRoleGranted, error) {
	event := new(TreasuryContractRoleGranted)
	if err := _TreasuryContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TreasuryContract contract.
type TreasuryContractRoleRevokedIterator struct {
	Event *TreasuryContractRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryContractRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TreasuryContractRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TreasuryContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryContractRoleRevoked represents a RoleRevoked event raised by the TreasuryContract contract.
type TreasuryContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasuryContract *TreasuryContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TreasuryContractRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasuryContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractRoleRevokedIterator{contract: _TreasuryContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasuryContract *TreasuryContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TreasuryContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasuryContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryContractRoleRevoked)
				if err := _TreasuryContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasuryContract *TreasuryContractFilterer) ParseRoleRevoked(log types.Log) (*TreasuryContractRoleRevoked, error) {
	event := new(TreasuryContractRoleRevoked)
	if err := _TreasuryContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryContractSignerChangeApprovedIterator is returned from FilterSignerChangeApproved and is used to iterate over the raw logs and unpacked data for SignerChangeApproved events raised by the TreasuryContract contract.
type TreasuryContractSignerChangeApprovedIterator struct {
	Event *TreasuryContractSignerChangeApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryContractSignerChangeApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryContractSignerChangeApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TreasuryContractSignerChangeApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TreasuryContractSignerChangeApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryContractSignerChangeApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryContractSignerChangeApproved represents a SignerChangeApproved event raised by the TreasuryContract contract.
type TreasuryContractSignerChangeApproved struct {
	ProposalId    *big.Int
	Approver      common.Address
	ApprovalCount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSignerChangeApproved is a free log retrieval operation binding the contract event 0x83859c56d4c0c74904a558a929632ad210da2cb88276444dc62a07555c4a6789.
//
// Solidity: event SignerChangeApproved(uint256 indexed proposalId, address indexed approver, uint256 approvalCount)
func (_TreasuryContract *TreasuryContractFilterer) FilterSignerChangeApproved(opts *bind.FilterOpts, proposalId []*big.Int, approver []common.Address) (*TreasuryContractSignerChangeApprovedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _TreasuryContract.contract.FilterLogs(opts, "SignerChangeApproved", proposalIdRule, approverRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractSignerChangeApprovedIterator{contract: _TreasuryContract.contract, event: "SignerChangeApproved", logs: logs, sub: sub}, nil
}

// WatchSignerChangeApproved is a free log subscription operation binding the contract event 0x83859c56d4c0c74904a558a929632ad210da2cb88276444dc62a07555c4a6789.
//
// Solidity: event SignerChangeApproved(uint256 indexed proposalId, address indexed approver, uint256 approvalCount)
func (_TreasuryContract *TreasuryContractFilterer) WatchSignerChangeApproved(opts *bind.WatchOpts, sink chan<- *TreasuryContractSignerChangeApproved, proposalId []*big.Int, approver []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _TreasuryContract.contract.WatchLogs(opts, "SignerChangeApproved", proposalIdRule, approverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryContractSignerChangeApproved)
				if err := _TreasuryContract.contract.UnpackLog(event, "SignerChangeApproved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignerChangeApproved is a log parse operation binding the contract event 0x83859c56d4c0c74904a558a929632ad210da2cb88276444dc62a07555c4a6789.
//
// Solidity: event SignerChangeApproved(uint256 indexed proposalId, address indexed approver, uint256 approvalCount)
func (_TreasuryContract *TreasuryContractFilterer) ParseSignerChangeApproved(log types.Log) (*TreasuryContractSignerChangeApproved, error) {
	event := new(TreasuryContractSignerChangeApproved)
	if err := _TreasuryContract.contract.UnpackLog(event, "SignerChangeApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryContractSignerChangeExecutedIterator is returned from FilterSignerChangeExecuted and is used to iterate over the raw logs and unpacked data for SignerChangeExecuted events raised by the TreasuryContract contract.
type TreasuryContractSignerChangeExecutedIterator struct {
	Event *TreasuryContractSignerChangeExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryContractSignerChangeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryContractSignerChangeExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TreasuryContractSignerChangeExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TreasuryContractSignerChangeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryContractSignerChangeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryContractSignerChangeExecuted represents a SignerChangeExecuted event raised by the TreasuryContract contract.
type TreasuryContractSignerChangeExecuted struct {
	ProposalId  *big.Int
	Target      common.Address
	Replacement common.Address
	ChangeType  uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSignerChangeExecuted is a free log retrieval operation binding the contract event 0x8640ae84b71f2feb507b6c3ce24a4ade255a2c18616f614e6343d6f6e6236757.
//
// Solidity: event SignerChangeExecuted(uint256 indexed proposalId, address target, address replacement, uint8 changeType)
func (_TreasuryContract *TreasuryContractFilterer) FilterSignerChangeExecuted(opts *bind.FilterOpts, proposalId []*big.Int) (*TreasuryContractSignerChangeExecutedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _TreasuryContract.contract.FilterLogs(opts, "SignerChangeExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractSignerChangeExecutedIterator{contract: _TreasuryContract.contract, event: "SignerChangeExecuted", logs: logs, sub: sub}, nil
}

// WatchSignerChangeExecuted is a free log subscription operation binding the contract event 0x8640ae84b71f2feb507b6c3ce24a4ade255a2c18616f614e6343d6f6e6236757.
//
// Solidity: event SignerChangeExecuted(uint256 indexed proposalId, address target, address replacement, uint8 changeType)
func (_TreasuryContract *TreasuryContractFilterer) WatchSignerChangeExecuted(opts *bind.WatchOpts, sink chan<- *TreasuryContractSignerChangeExecuted, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _TreasuryContract.contract.WatchLogs(opts, "SignerChangeExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryContractSignerChangeExecuted)
				if err := _TreasuryContract.contract.UnpackLog(event, "SignerChangeExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignerChangeExecuted is a log parse operation binding the contract event 0x8640ae84b71f2feb507b6c3ce24a4ade255a2c18616f614e6343d6f6e6236757.
//
// Solidity: event SignerChangeExecuted(uint256 indexed proposalId, address target, address replacement, uint8 changeType)
func (_TreasuryContract *TreasuryContractFilterer) ParseSignerChangeExecuted(log types.Log) (*TreasuryContractSignerChangeExecuted, error) {
	event := new(TreasuryContractSignerChangeExecuted)
	if err := _TreasuryContract.contract.UnpackLog(event, "SignerChangeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryContractSignerChangeProposedIterator is returned from FilterSignerChangeProposed and is used to iterate over the raw logs and unpacked data for SignerChangeProposed events raised by the TreasuryContract contract.
type TreasuryContractSignerChangeProposedIterator struct {
	Event *TreasuryContractSignerChangeProposed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryContractSignerChangeProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryContractSignerChangeProposed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TreasuryContractSignerChangeProposed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TreasuryContractSignerChangeProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryContractSignerChangeProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryContractSignerChangeProposed represents a SignerChangeProposed event raised by the TreasuryContract contract.
type TreasuryContractSignerChangeProposed struct {
	ProposalId  *big.Int
	Proposer    common.Address
	Target      common.Address
	Replacement common.Address
	ChangeType  uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSignerChangeProposed is a free log retrieval operation binding the contract event 0xcc444a2a8b8581d87b963b87f8d992c3ce64fcce3ae94a580204902b309c490e.
//
// Solidity: event SignerChangeProposed(uint256 indexed proposalId, address indexed proposer, address target, address replacement, uint8 changeType)
func (_TreasuryContract *TreasuryContractFilterer) FilterSignerChangeProposed(opts *bind.FilterOpts, proposalId []*big.Int, proposer []common.Address) (*TreasuryContractSignerChangeProposedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _TreasuryContract.contract.FilterLogs(opts, "SignerChangeProposed", proposalIdRule, proposerRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractSignerChangeProposedIterator{contract: _TreasuryContract.contract, event: "SignerChangeProposed", logs: logs, sub: sub}, nil
}

// WatchSignerChangeProposed is a free log subscription operation binding the contract event 0xcc444a2a8b8581d87b963b87f8d992c3ce64fcce3ae94a580204902b309c490e.
//
// Solidity: event SignerChangeProposed(uint256 indexed proposalId, address indexed proposer, address target, address replacement, uint8 changeType)
func (_TreasuryContract *TreasuryContractFilterer) WatchSignerChangeProposed(opts *bind.WatchOpts, sink chan<- *TreasuryContractSignerChangeProposed, proposalId []*big.Int, proposer []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _TreasuryContract.contract.WatchLogs(opts, "SignerChangeProposed", proposalIdRule, proposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryContractSignerChangeProposed)
				if err := _TreasuryContract.contract.UnpackLog(event, "SignerChangeProposed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignerChangeProposed is a log parse operation binding the contract event 0xcc444a2a8b8581d87b963b87f8d992c3ce64fcce3ae94a580204902b309c490e.
//
// Solidity: event SignerChangeProposed(uint256 indexed proposalId, address indexed proposer, address target, address replacement, uint8 changeType)
func (_TreasuryContract *TreasuryContractFilterer) ParseSignerChangeProposed(log types.Log) (*TreasuryContractSignerChangeProposed, error) {
	event := new(TreasuryContractSignerChangeProposed)
	if err := _TreasuryContract.contract.UnpackLog(event, "SignerChangeProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryContractTreasuryFrozenIterator is returned from FilterTreasuryFrozen and is used to iterate over the raw logs and unpacked data for TreasuryFrozen events raised by the TreasuryContract contract.
type TreasuryContractTreasuryFrozenIterator struct {
	Event *TreasuryContractTreasuryFrozen // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryContractTreasuryFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryContractTreasuryFrozen)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TreasuryContractTreasuryFrozen)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TreasuryContractTreasuryFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryContractTreasuryFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryContractTreasuryFrozen represents a TreasuryFrozen event raised by the TreasuryContract contract.
type TreasuryContractTreasuryFrozen struct {
	By  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTreasuryFrozen is a free log retrieval operation binding the contract event 0xe5d3ade628a7ba68637214109f9c0b87cf65fdbf07d3f718eba657fb14c605b9.
//
// Solidity: event TreasuryFrozen(address indexed by)
func (_TreasuryContract *TreasuryContractFilterer) FilterTreasuryFrozen(opts *bind.FilterOpts, by []common.Address) (*TreasuryContractTreasuryFrozenIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _TreasuryContract.contract.FilterLogs(opts, "TreasuryFrozen", byRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractTreasuryFrozenIterator{contract: _TreasuryContract.contract, event: "TreasuryFrozen", logs: logs, sub: sub}, nil
}

// WatchTreasuryFrozen is a free log subscription operation binding the contract event 0xe5d3ade628a7ba68637214109f9c0b87cf65fdbf07d3f718eba657fb14c605b9.
//
// Solidity: event TreasuryFrozen(address indexed by)
func (_TreasuryContract *TreasuryContractFilterer) WatchTreasuryFrozen(opts *bind.WatchOpts, sink chan<- *TreasuryContractTreasuryFrozen, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _TreasuryContract.contract.WatchLogs(opts, "TreasuryFrozen", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryContractTreasuryFrozen)
				if err := _TreasuryContract.contract.UnpackLog(event, "TreasuryFrozen", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTreasuryFrozen is a log parse operation binding the contract event 0xe5d3ade628a7ba68637214109f9c0b87cf65fdbf07d3f718eba657fb14c605b9.
//
// Solidity: event TreasuryFrozen(address indexed by)
func (_TreasuryContract *TreasuryContractFilterer) ParseTreasuryFrozen(log types.Log) (*TreasuryContractTreasuryFrozen, error) {
	event := new(TreasuryContractTreasuryFrozen)
	if err := _TreasuryContract.contract.UnpackLog(event, "TreasuryFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryContractTreasuryUnfrozenIterator is returned from FilterTreasuryUnfrozen and is used to iterate over the raw logs and unpacked data for TreasuryUnfrozen events raised by the TreasuryContract contract.
type TreasuryContractTreasuryUnfrozenIterator struct {
	Event *TreasuryContractTreasuryUnfrozen // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryContractTreasuryUnfrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryContractTreasuryUnfrozen)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TreasuryContractTreasuryUnfrozen)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TreasuryContractTreasuryUnfrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryContractTreasuryUnfrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryContractTreasuryUnfrozen represents a TreasuryUnfrozen event raised by the TreasuryContract contract.
type TreasuryContractTreasuryUnfrozen struct {
	By  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTreasuryUnfrozen is a free log retrieval operation binding the contract event 0x3a3527f242847642c35d31ec7eaf163ba3149cd9899ccc19f1049c9d0d8c5ea6.
//
// Solidity: event TreasuryUnfrozen(address indexed by)
func (_TreasuryContract *TreasuryContractFilterer) FilterTreasuryUnfrozen(opts *bind.FilterOpts, by []common.Address) (*TreasuryContractTreasuryUnfrozenIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _TreasuryContract.contract.FilterLogs(opts, "TreasuryUnfrozen", byRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractTreasuryUnfrozenIterator{contract: _TreasuryContract.contract, event: "TreasuryUnfrozen", logs: logs, sub: sub}, nil
}

// WatchTreasuryUnfrozen is a free log subscription operation binding the contract event 0x3a3527f242847642c35d31ec7eaf163ba3149cd9899ccc19f1049c9d0d8c5ea6.
//
// Solidity: event TreasuryUnfrozen(address indexed by)
func (_TreasuryContract *TreasuryContractFilterer) WatchTreasuryUnfrozen(opts *bind.WatchOpts, sink chan<- *TreasuryContractTreasuryUnfrozen, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _TreasuryContract.contract.WatchLogs(opts, "TreasuryUnfrozen", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryContractTreasuryUnfrozen)
				if err := _TreasuryContract.contract.UnpackLog(event, "TreasuryUnfrozen", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTreasuryUnfrozen is a log parse operation binding the contract event 0x3a3527f242847642c35d31ec7eaf163ba3149cd9899ccc19f1049c9d0d8c5ea6.
//
// Solidity: event TreasuryUnfrozen(address indexed by)
func (_TreasuryContract *TreasuryContractFilterer) ParseTreasuryUnfrozen(log types.Log) (*TreasuryContractTreasuryUnfrozen, error) {
	event := new(TreasuryContractTreasuryUnfrozen)
	if err := _TreasuryContract.contract.UnpackLog(event, "TreasuryUnfrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryContractWithdrawalApprovedIterator is returned from FilterWithdrawalApproved and is used to iterate over the raw logs and unpacked data for WithdrawalApproved events raised by the TreasuryContract contract.
type TreasuryContractWithdrawalApprovedIterator struct {
	Event *TreasuryContractWithdrawalApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryContractWithdrawalApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryContractWithdrawalApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TreasuryContractWithdrawalApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TreasuryContractWithdrawalApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryContractWithdrawalApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryContractWithdrawalApproved represents a WithdrawalApproved event raised by the TreasuryContract contract.
type TreasuryContractWithdrawalApproved struct {
	ProposalId    *big.Int
	Approver      common.Address
	ApprovalCount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalApproved is a free log retrieval operation binding the contract event 0x26fdbb8cb013508232284b67258d179e90a315277e0687b0a9f4fd1e94d2d632.
//
// Solidity: event WithdrawalApproved(uint256 indexed proposalId, address indexed approver, uint256 approvalCount)
func (_TreasuryContract *TreasuryContractFilterer) FilterWithdrawalApproved(opts *bind.FilterOpts, proposalId []*big.Int, approver []common.Address) (*TreasuryContractWithdrawalApprovedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _TreasuryContract.contract.FilterLogs(opts, "WithdrawalApproved", proposalIdRule, approverRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractWithdrawalApprovedIterator{contract: _TreasuryContract.contract, event: "WithdrawalApproved", logs: logs, sub: sub}, nil
}

// WatchWithdrawalApproved is a free log subscription operation binding the contract event 0x26fdbb8cb013508232284b67258d179e90a315277e0687b0a9f4fd1e94d2d632.
//
// Solidity: event WithdrawalApproved(uint256 indexed proposalId, address indexed approver, uint256 approvalCount)
func (_TreasuryContract *TreasuryContractFilterer) WatchWithdrawalApproved(opts *bind.WatchOpts, sink chan<- *TreasuryContractWithdrawalApproved, proposalId []*big.Int, approver []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _TreasuryContract.contract.WatchLogs(opts, "WithdrawalApproved", proposalIdRule, approverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryContractWithdrawalApproved)
				if err := _TreasuryContract.contract.UnpackLog(event, "WithdrawalApproved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalApproved is a log parse operation binding the contract event 0x26fdbb8cb013508232284b67258d179e90a315277e0687b0a9f4fd1e94d2d632.
//
// Solidity: event WithdrawalApproved(uint256 indexed proposalId, address indexed approver, uint256 approvalCount)
func (_TreasuryContract *TreasuryContractFilterer) ParseWithdrawalApproved(log types.Log) (*TreasuryContractWithdrawalApproved, error) {
	event := new(TreasuryContractWithdrawalApproved)
	if err := _TreasuryContract.contract.UnpackLog(event, "WithdrawalApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryContractWithdrawalCancelledIterator is returned from FilterWithdrawalCancelled and is used to iterate over the raw logs and unpacked data for WithdrawalCancelled events raised by the TreasuryContract contract.
type TreasuryContractWithdrawalCancelledIterator struct {
	Event *TreasuryContractWithdrawalCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryContractWithdrawalCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryContractWithdrawalCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TreasuryContractWithdrawalCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TreasuryContractWithdrawalCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryContractWithdrawalCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryContractWithdrawalCancelled represents a WithdrawalCancelled event raised by the TreasuryContract contract.
type TreasuryContractWithdrawalCancelled struct {
	ProposalId *big.Int
	Canceller  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalCancelled is a free log retrieval operation binding the contract event 0x13c630dd37c22cad8d02f5d249d447222f855d1ca8986dd17b248ca4b9d8c29a.
//
// Solidity: event WithdrawalCancelled(uint256 indexed proposalId, address indexed canceller)
func (_TreasuryContract *TreasuryContractFilterer) FilterWithdrawalCancelled(opts *bind.FilterOpts, proposalId []*big.Int, canceller []common.Address) (*TreasuryContractWithdrawalCancelledIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var cancellerRule []interface{}
	for _, cancellerItem := range canceller {
		cancellerRule = append(cancellerRule, cancellerItem)
	}

	logs, sub, err := _TreasuryContract.contract.FilterLogs(opts, "WithdrawalCancelled", proposalIdRule, cancellerRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractWithdrawalCancelledIterator{contract: _TreasuryContract.contract, event: "WithdrawalCancelled", logs: logs, sub: sub}, nil
}

// WatchWithdrawalCancelled is a free log subscription operation binding the contract event 0x13c630dd37c22cad8d02f5d249d447222f855d1ca8986dd17b248ca4b9d8c29a.
//
// Solidity: event WithdrawalCancelled(uint256 indexed proposalId, address indexed canceller)
func (_TreasuryContract *TreasuryContractFilterer) WatchWithdrawalCancelled(opts *bind.WatchOpts, sink chan<- *TreasuryContractWithdrawalCancelled, proposalId []*big.Int, canceller []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var cancellerRule []interface{}
	for _, cancellerItem := range canceller {
		cancellerRule = append(cancellerRule, cancellerItem)
	}

	logs, sub, err := _TreasuryContract.contract.WatchLogs(opts, "WithdrawalCancelled", proposalIdRule, cancellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryContractWithdrawalCancelled)
				if err := _TreasuryContract.contract.UnpackLog(event, "WithdrawalCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalCancelled is a log parse operation binding the contract event 0x13c630dd37c22cad8d02f5d249d447222f855d1ca8986dd17b248ca4b9d8c29a.
//
// Solidity: event WithdrawalCancelled(uint256 indexed proposalId, address indexed canceller)
func (_TreasuryContract *TreasuryContractFilterer) ParseWithdrawalCancelled(log types.Log) (*TreasuryContractWithdrawalCancelled, error) {
	event := new(TreasuryContractWithdrawalCancelled)
	if err := _TreasuryContract.contract.UnpackLog(event, "WithdrawalCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryContractWithdrawalExecutedIterator is returned from FilterWithdrawalExecuted and is used to iterate over the raw logs and unpacked data for WithdrawalExecuted events raised by the TreasuryContract contract.
type TreasuryContractWithdrawalExecutedIterator struct {
	Event *TreasuryContractWithdrawalExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryContractWithdrawalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryContractWithdrawalExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TreasuryContractWithdrawalExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TreasuryContractWithdrawalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryContractWithdrawalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryContractWithdrawalExecuted represents a WithdrawalExecuted event raised by the TreasuryContract contract.
type TreasuryContractWithdrawalExecuted struct {
	ProposalId *big.Int
	Recipient  common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalExecuted is a free log retrieval operation binding the contract event 0xd6cddb3d69146e96ebc2c87b1b3dd0b20ee2d3b0eadf134e011afb434a3e56e6.
//
// Solidity: event WithdrawalExecuted(uint256 indexed proposalId, address indexed recipient, uint256 amount)
func (_TreasuryContract *TreasuryContractFilterer) FilterWithdrawalExecuted(opts *bind.FilterOpts, proposalId []*big.Int, recipient []common.Address) (*TreasuryContractWithdrawalExecutedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TreasuryContract.contract.FilterLogs(opts, "WithdrawalExecuted", proposalIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractWithdrawalExecutedIterator{contract: _TreasuryContract.contract, event: "WithdrawalExecuted", logs: logs, sub: sub}, nil
}

// WatchWithdrawalExecuted is a free log subscription operation binding the contract event 0xd6cddb3d69146e96ebc2c87b1b3dd0b20ee2d3b0eadf134e011afb434a3e56e6.
//
// Solidity: event WithdrawalExecuted(uint256 indexed proposalId, address indexed recipient, uint256 amount)
func (_TreasuryContract *TreasuryContractFilterer) WatchWithdrawalExecuted(opts *bind.WatchOpts, sink chan<- *TreasuryContractWithdrawalExecuted, proposalId []*big.Int, recipient []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TreasuryContract.contract.WatchLogs(opts, "WithdrawalExecuted", proposalIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryContractWithdrawalExecuted)
				if err := _TreasuryContract.contract.UnpackLog(event, "WithdrawalExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalExecuted is a log parse operation binding the contract event 0xd6cddb3d69146e96ebc2c87b1b3dd0b20ee2d3b0eadf134e011afb434a3e56e6.
//
// Solidity: event WithdrawalExecuted(uint256 indexed proposalId, address indexed recipient, uint256 amount)
func (_TreasuryContract *TreasuryContractFilterer) ParseWithdrawalExecuted(log types.Log) (*TreasuryContractWithdrawalExecuted, error) {
	event := new(TreasuryContractWithdrawalExecuted)
	if err := _TreasuryContract.contract.UnpackLog(event, "WithdrawalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryContractWithdrawalProposedIterator is returned from FilterWithdrawalProposed and is used to iterate over the raw logs and unpacked data for WithdrawalProposed events raised by the TreasuryContract contract.
type TreasuryContractWithdrawalProposedIterator struct {
	Event *TreasuryContractWithdrawalProposed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryContractWithdrawalProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryContractWithdrawalProposed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TreasuryContractWithdrawalProposed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TreasuryContractWithdrawalProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryContractWithdrawalProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryContractWithdrawalProposed represents a WithdrawalProposed event raised by the TreasuryContract contract.
type TreasuryContractWithdrawalProposed struct {
	ProposalId *big.Int
	Proposer   common.Address
	Recipient  common.Address
	Amount     *big.Int
	Purpose    string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalProposed is a free log retrieval operation binding the contract event 0x4014eec3cea3dc25ac3d039e650db43e87bc0156ee996c08de64199c2ce35f8c.
//
// Solidity: event WithdrawalProposed(uint256 indexed proposalId, address indexed proposer, address indexed recipient, uint256 amount, string purpose)
func (_TreasuryContract *TreasuryContractFilterer) FilterWithdrawalProposed(opts *bind.FilterOpts, proposalId []*big.Int, proposer []common.Address, recipient []common.Address) (*TreasuryContractWithdrawalProposedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TreasuryContract.contract.FilterLogs(opts, "WithdrawalProposed", proposalIdRule, proposerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractWithdrawalProposedIterator{contract: _TreasuryContract.contract, event: "WithdrawalProposed", logs: logs, sub: sub}, nil
}

// WatchWithdrawalProposed is a free log subscription operation binding the contract event 0x4014eec3cea3dc25ac3d039e650db43e87bc0156ee996c08de64199c2ce35f8c.
//
// Solidity: event WithdrawalProposed(uint256 indexed proposalId, address indexed proposer, address indexed recipient, uint256 amount, string purpose)
func (_TreasuryContract *TreasuryContractFilterer) WatchWithdrawalProposed(opts *bind.WatchOpts, sink chan<- *TreasuryContractWithdrawalProposed, proposalId []*big.Int, proposer []common.Address, recipient []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TreasuryContract.contract.WatchLogs(opts, "WithdrawalProposed", proposalIdRule, proposerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryContractWithdrawalProposed)
				if err := _TreasuryContract.contract.UnpackLog(event, "WithdrawalProposed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalProposed is a log parse operation binding the contract event 0x4014eec3cea3dc25ac3d039e650db43e87bc0156ee996c08de64199c2ce35f8c.
//
// Solidity: event WithdrawalProposed(uint256 indexed proposalId, address indexed proposer, address indexed recipient, uint256 amount, string purpose)
func (_TreasuryContract *TreasuryContractFilterer) ParseWithdrawalProposed(log types.Log) (*TreasuryContractWithdrawalProposed, error) {
	event := new(TreasuryContractWithdrawalProposed)
	if err := _TreasuryContract.contract.UnpackLog(event, "WithdrawalProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
