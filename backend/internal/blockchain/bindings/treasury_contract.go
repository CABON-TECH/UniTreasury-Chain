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
var TreasuryContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialApprovers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_requiredApprovals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_dailyLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_usdcToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aavePool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aUsdcToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_timelockDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"APPROVER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EXECUTOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROPOSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aUsdcToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aavePool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approveSignerChange\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approveWithdrawal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelSignerChange\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelWithdrawal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dailyLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSignerChange\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeWithdrawal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freeze\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDailyWithdrawnAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProposal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITreasury.WithdrawalProposal\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"purpose\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumITreasury.ProposalStatus\"},{\"name\":\"approvalCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSignerProposal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITreasury.SignerChangeProposal\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"targetSigner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newSigner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"changeType\",\"type\":\"uint8\",\"internalType\":\"enumITreasury.ChangeType\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumITreasury.ProposalStatus\"},{\"name\":\"approvalCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasApproved\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasApprovedSignerChange\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isFrozen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeSignerChange\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"replacement\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"changeType\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proposeWithdrawal\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"purpose\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requiredApprovals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDailyLimit\",\"inputs\":[{\"name\":\"newLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRequiredApprovals\",\"inputs\":[{\"name\":\"newRequired\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTimelockDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"timelockDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unfreeze\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"usdcToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DailyLimitUpdated\",\"inputs\":[{\"name\":\"oldLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newBalance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequiredApprovalsUpdated\",\"inputs\":[{\"name\":\"oldRequired\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newRequired\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerChangeApproved\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"approver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approvalCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerChangeExecuted\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"replacement\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"changeType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerChangeProposed\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"replacement\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"changeType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TimelockDelayUpdated\",\"inputs\":[{\"name\":\"oldDelay\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newDelay\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TreasuryFrozen\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TreasuryUnfrozen\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalApproved\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"approver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approvalCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalCancelled\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"canceller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalExecuted\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalProposed\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"purpose\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Treasury__AlreadyApproved\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Treasury__DailyLimitExceeded\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"remainingToday\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Treasury__Frozen\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Treasury__InsufficientApprovals\",\"inputs\":[{\"name\":\"have\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"need\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Treasury__InsufficientBalance\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Treasury__NotAuthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Treasury__ProposalNotFound\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Treasury__ProposalNotPending\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Treasury__TimelockNotExpired\",\"inputs\":[{\"name\":\"readyAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Treasury__ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Treasury__ZeroAmount\",\"inputs\":[]}]",
}
var TreasuryContractABI = TreasuryContractMetaData.ABI
type TreasuryContract struct {
	TreasuryContractCaller     
	TreasuryContractTransactor 
	TreasuryContractFilterer   
}
type TreasuryContractCaller struct {
	contract *bind.BoundContract 
}
type TreasuryContractTransactor struct {
	contract *bind.BoundContract 
}
type TreasuryContractFilterer struct {
	contract *bind.BoundContract 
}
type TreasuryContractSession struct {
	Contract     *TreasuryContract 
	CallOpts     bind.CallOpts     
	TransactOpts bind.TransactOpts 
}
type TreasuryContractCallerSession struct {
	Contract *TreasuryContractCaller 
	CallOpts bind.CallOpts           
}
type TreasuryContractTransactorSession struct {
	Contract     *TreasuryContractTransactor 
	TransactOpts bind.TransactOpts           
}
type TreasuryContractRaw struct {
	Contract *TreasuryContract 
}
type TreasuryContractCallerRaw struct {
	Contract *TreasuryContractCaller 
}
type TreasuryContractTransactorRaw struct {
	Contract *TreasuryContractTransactor 
}
func NewTreasuryContract(address common.Address, backend bind.ContractBackend) (*TreasuryContract, error) {
	contract, err := bindTreasuryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TreasuryContract{TreasuryContractCaller: TreasuryContractCaller{contract: contract}, TreasuryContractTransactor: TreasuryContractTransactor{contract: contract}, TreasuryContractFilterer: TreasuryContractFilterer{contract: contract}}, nil
}
func NewTreasuryContractCaller(address common.Address, caller bind.ContractCaller) (*TreasuryContractCaller, error) {
	contract, err := bindTreasuryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractCaller{contract: contract}, nil
}
func NewTreasuryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TreasuryContractTransactor, error) {
	contract, err := bindTreasuryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractTransactor{contract: contract}, nil
}
func NewTreasuryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TreasuryContractFilterer, error) {
	contract, err := bindTreasuryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TreasuryContractFilterer{contract: contract}, nil
}
func bindTreasuryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TreasuryContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}
func (_TreasuryContract *TreasuryContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TreasuryContract.Contract.TreasuryContractCaller.contract.Call(opts, result, method, params...)
}
func (_TreasuryContract *TreasuryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryContract.Contract.TreasuryContractTransactor.contract.Transfer(opts)
}
func (_TreasuryContract *TreasuryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TreasuryContract.Contract.TreasuryContractTransactor.contract.Transact(opts, method, params...)
}
func (_TreasuryContract *TreasuryContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TreasuryContract.Contract.contract.Call(opts, result, method, params...)
}
func (_TreasuryContract *TreasuryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryContract.Contract.contract.Transfer(opts)
}
func (_TreasuryContract *TreasuryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TreasuryContract.Contract.contract.Transact(opts, method, params...)
}
func (_TreasuryContract *TreasuryContractCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "ADMIN_ROLE")
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) ADMINROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.ADMINROLE(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCallerSession) ADMINROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.ADMINROLE(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCaller) APPROVERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "APPROVER_ROLE")
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) APPROVERROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.APPROVERROLE(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCallerSession) APPROVERROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.APPROVERROLE(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.DEFAULTADMINROLE(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.DEFAULTADMINROLE(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCaller) EXECUTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "EXECUTOR_ROLE")
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) EXECUTORROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.EXECUTORROLE(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCallerSession) EXECUTORROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.EXECUTORROLE(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCaller) PROPOSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "PROPOSER_ROLE")
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) PROPOSERROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.PROPOSERROLE(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCallerSession) PROPOSERROLE() ([32]byte, error) {
	return _TreasuryContract.Contract.PROPOSERROLE(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCaller) AUsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "aUsdcToken")
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) AUsdcToken() (common.Address, error) {
	return _TreasuryContract.Contract.AUsdcToken(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCallerSession) AUsdcToken() (common.Address, error) {
	return _TreasuryContract.Contract.AUsdcToken(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCaller) AavePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "aavePool")
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) AavePool() (common.Address, error) {
	return _TreasuryContract.Contract.AavePool(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCallerSession) AavePool() (common.Address, error) {
	return _TreasuryContract.Contract.AavePool(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCaller) DailyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "dailyLimit")
	if err != nil {
		return *new(*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) DailyLimit() (*big.Int, error) {
	return _TreasuryContract.Contract.DailyLimit(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCallerSession) DailyLimit() (*big.Int, error) {
	return _TreasuryContract.Contract.DailyLimit(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "getBalance")
	if err != nil {
		return *new(*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) GetBalance() (*big.Int, error) {
	return _TreasuryContract.Contract.GetBalance(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCallerSession) GetBalance() (*big.Int, error) {
	return _TreasuryContract.Contract.GetBalance(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCaller) GetDailyWithdrawnAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "getDailyWithdrawnAmount")
	if err != nil {
		return *new(*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) GetDailyWithdrawnAmount() (*big.Int, error) {
	return _TreasuryContract.Contract.GetDailyWithdrawnAmount(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCallerSession) GetDailyWithdrawnAmount() (*big.Int, error) {
	return _TreasuryContract.Contract.GetDailyWithdrawnAmount(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCaller) GetProposal(opts *bind.CallOpts, proposalId *big.Int) (ITreasuryWithdrawalProposal, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "getProposal", proposalId)
	if err != nil {
		return *new(ITreasuryWithdrawalProposal), err
	}
	out0 := *abi.ConvertType(out[0], new(ITreasuryWithdrawalProposal)).(*ITreasuryWithdrawalProposal)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) GetProposal(proposalId *big.Int) (ITreasuryWithdrawalProposal, error) {
	return _TreasuryContract.Contract.GetProposal(&_TreasuryContract.CallOpts, proposalId)
}
func (_TreasuryContract *TreasuryContractCallerSession) GetProposal(proposalId *big.Int) (ITreasuryWithdrawalProposal, error) {
	return _TreasuryContract.Contract.GetProposal(&_TreasuryContract.CallOpts, proposalId)
}
func (_TreasuryContract *TreasuryContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "getRoleAdmin", role)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TreasuryContract.Contract.GetRoleAdmin(&_TreasuryContract.CallOpts, role)
}
func (_TreasuryContract *TreasuryContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TreasuryContract.Contract.GetRoleAdmin(&_TreasuryContract.CallOpts, role)
}
func (_TreasuryContract *TreasuryContractCaller) GetSignerProposal(opts *bind.CallOpts, proposalId *big.Int) (ITreasurySignerChangeProposal, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "getSignerProposal", proposalId)
	if err != nil {
		return *new(ITreasurySignerChangeProposal), err
	}
	out0 := *abi.ConvertType(out[0], new(ITreasurySignerChangeProposal)).(*ITreasurySignerChangeProposal)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) GetSignerProposal(proposalId *big.Int) (ITreasurySignerChangeProposal, error) {
	return _TreasuryContract.Contract.GetSignerProposal(&_TreasuryContract.CallOpts, proposalId)
}
func (_TreasuryContract *TreasuryContractCallerSession) GetSignerProposal(proposalId *big.Int) (ITreasurySignerChangeProposal, error) {
	return _TreasuryContract.Contract.GetSignerProposal(&_TreasuryContract.CallOpts, proposalId)
}
func (_TreasuryContract *TreasuryContractCaller) HasApproved(opts *bind.CallOpts, proposalId *big.Int, approver common.Address) (bool, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "hasApproved", proposalId, approver)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) HasApproved(proposalId *big.Int, approver common.Address) (bool, error) {
	return _TreasuryContract.Contract.HasApproved(&_TreasuryContract.CallOpts, proposalId, approver)
}
func (_TreasuryContract *TreasuryContractCallerSession) HasApproved(proposalId *big.Int, approver common.Address) (bool, error) {
	return _TreasuryContract.Contract.HasApproved(&_TreasuryContract.CallOpts, proposalId, approver)
}
func (_TreasuryContract *TreasuryContractCaller) HasApprovedSignerChange(opts *bind.CallOpts, proposalId *big.Int, approver common.Address) (bool, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "hasApprovedSignerChange", proposalId, approver)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) HasApprovedSignerChange(proposalId *big.Int, approver common.Address) (bool, error) {
	return _TreasuryContract.Contract.HasApprovedSignerChange(&_TreasuryContract.CallOpts, proposalId, approver)
}
func (_TreasuryContract *TreasuryContractCallerSession) HasApprovedSignerChange(proposalId *big.Int, approver common.Address) (bool, error) {
	return _TreasuryContract.Contract.HasApprovedSignerChange(&_TreasuryContract.CallOpts, proposalId, approver)
}
func (_TreasuryContract *TreasuryContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "hasRole", role, account)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TreasuryContract.Contract.HasRole(&_TreasuryContract.CallOpts, role, account)
}
func (_TreasuryContract *TreasuryContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TreasuryContract.Contract.HasRole(&_TreasuryContract.CallOpts, role, account)
}
func (_TreasuryContract *TreasuryContractCaller) IsFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "isFrozen")
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) IsFrozen() (bool, error) {
	return _TreasuryContract.Contract.IsFrozen(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCallerSession) IsFrozen() (bool, error) {
	return _TreasuryContract.Contract.IsFrozen(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCaller) RequiredApprovals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "requiredApprovals")
	if err != nil {
		return *new(*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) RequiredApprovals() (*big.Int, error) {
	return _TreasuryContract.Contract.RequiredApprovals(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCallerSession) RequiredApprovals() (*big.Int, error) {
	return _TreasuryContract.Contract.RequiredApprovals(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "supportsInterface", interfaceId)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TreasuryContract.Contract.SupportsInterface(&_TreasuryContract.CallOpts, interfaceId)
}
func (_TreasuryContract *TreasuryContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TreasuryContract.Contract.SupportsInterface(&_TreasuryContract.CallOpts, interfaceId)
}
func (_TreasuryContract *TreasuryContractCaller) TimelockDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "timelockDelay")
	if err != nil {
		return *new(*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) TimelockDelay() (*big.Int, error) {
	return _TreasuryContract.Contract.TimelockDelay(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCallerSession) TimelockDelay() (*big.Int, error) {
	return _TreasuryContract.Contract.TimelockDelay(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCaller) UsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TreasuryContract.contract.Call(opts, &out, "usdcToken")
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}
func (_TreasuryContract *TreasuryContractSession) UsdcToken() (common.Address, error) {
	return _TreasuryContract.Contract.UsdcToken(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractCallerSession) UsdcToken() (common.Address, error) {
	return _TreasuryContract.Contract.UsdcToken(&_TreasuryContract.CallOpts)
}
func (_TreasuryContract *TreasuryContractTransactor) ApproveSignerChange(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "approveSignerChange", proposalId)
}
func (_TreasuryContract *TreasuryContractSession) ApproveSignerChange(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ApproveSignerChange(&_TreasuryContract.TransactOpts, proposalId)
}
func (_TreasuryContract *TreasuryContractTransactorSession) ApproveSignerChange(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ApproveSignerChange(&_TreasuryContract.TransactOpts, proposalId)
}
func (_TreasuryContract *TreasuryContractTransactor) ApproveWithdrawal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "approveWithdrawal", proposalId)
}
func (_TreasuryContract *TreasuryContractSession) ApproveWithdrawal(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ApproveWithdrawal(&_TreasuryContract.TransactOpts, proposalId)
}
func (_TreasuryContract *TreasuryContractTransactorSession) ApproveWithdrawal(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ApproveWithdrawal(&_TreasuryContract.TransactOpts, proposalId)
}
func (_TreasuryContract *TreasuryContractTransactor) CancelSignerChange(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "cancelSignerChange", proposalId)
}
func (_TreasuryContract *TreasuryContractSession) CancelSignerChange(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.CancelSignerChange(&_TreasuryContract.TransactOpts, proposalId)
}
func (_TreasuryContract *TreasuryContractTransactorSession) CancelSignerChange(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.CancelSignerChange(&_TreasuryContract.TransactOpts, proposalId)
}
func (_TreasuryContract *TreasuryContractTransactor) CancelWithdrawal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "cancelWithdrawal", proposalId)
}
func (_TreasuryContract *TreasuryContractSession) CancelWithdrawal(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.CancelWithdrawal(&_TreasuryContract.TransactOpts, proposalId)
}
func (_TreasuryContract *TreasuryContractTransactorSession) CancelWithdrawal(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.CancelWithdrawal(&_TreasuryContract.TransactOpts, proposalId)
}
func (_TreasuryContract *TreasuryContractTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "deposit", amount)
}
func (_TreasuryContract *TreasuryContractSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.Deposit(&_TreasuryContract.TransactOpts, amount)
}
func (_TreasuryContract *TreasuryContractTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.Deposit(&_TreasuryContract.TransactOpts, amount)
}
func (_TreasuryContract *TreasuryContractTransactor) ExecuteSignerChange(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "executeSignerChange", proposalId)
}
func (_TreasuryContract *TreasuryContractSession) ExecuteSignerChange(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ExecuteSignerChange(&_TreasuryContract.TransactOpts, proposalId)
}
func (_TreasuryContract *TreasuryContractTransactorSession) ExecuteSignerChange(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ExecuteSignerChange(&_TreasuryContract.TransactOpts, proposalId)
}
func (_TreasuryContract *TreasuryContractTransactor) ExecuteWithdrawal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "executeWithdrawal", proposalId)
}
func (_TreasuryContract *TreasuryContractSession) ExecuteWithdrawal(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ExecuteWithdrawal(&_TreasuryContract.TransactOpts, proposalId)
}
func (_TreasuryContract *TreasuryContractTransactorSession) ExecuteWithdrawal(proposalId *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ExecuteWithdrawal(&_TreasuryContract.TransactOpts, proposalId)
}
func (_TreasuryContract *TreasuryContractTransactor) Freeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "freeze")
}
func (_TreasuryContract *TreasuryContractSession) Freeze() (*types.Transaction, error) {
	return _TreasuryContract.Contract.Freeze(&_TreasuryContract.TransactOpts)
}
func (_TreasuryContract *TreasuryContractTransactorSession) Freeze() (*types.Transaction, error) {
	return _TreasuryContract.Contract.Freeze(&_TreasuryContract.TransactOpts)
}
func (_TreasuryContract *TreasuryContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "grantRole", role, account)
}
func (_TreasuryContract *TreasuryContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasuryContract.Contract.GrantRole(&_TreasuryContract.TransactOpts, role, account)
}
func (_TreasuryContract *TreasuryContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasuryContract.Contract.GrantRole(&_TreasuryContract.TransactOpts, role, account)
}
func (_TreasuryContract *TreasuryContractTransactor) ProposeSignerChange(opts *bind.TransactOpts, target common.Address, replacement common.Address, changeType uint8) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "proposeSignerChange", target, replacement, changeType)
}
func (_TreasuryContract *TreasuryContractSession) ProposeSignerChange(target common.Address, replacement common.Address, changeType uint8) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ProposeSignerChange(&_TreasuryContract.TransactOpts, target, replacement, changeType)
}
func (_TreasuryContract *TreasuryContractTransactorSession) ProposeSignerChange(target common.Address, replacement common.Address, changeType uint8) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ProposeSignerChange(&_TreasuryContract.TransactOpts, target, replacement, changeType)
}
func (_TreasuryContract *TreasuryContractTransactor) ProposeWithdrawal(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, purpose string) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "proposeWithdrawal", recipient, amount, purpose)
}
func (_TreasuryContract *TreasuryContractSession) ProposeWithdrawal(recipient common.Address, amount *big.Int, purpose string) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ProposeWithdrawal(&_TreasuryContract.TransactOpts, recipient, amount, purpose)
}
func (_TreasuryContract *TreasuryContractTransactorSession) ProposeWithdrawal(recipient common.Address, amount *big.Int, purpose string) (*types.Transaction, error) {
	return _TreasuryContract.Contract.ProposeWithdrawal(&_TreasuryContract.TransactOpts, recipient, amount, purpose)
}
func (_TreasuryContract *TreasuryContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}
func (_TreasuryContract *TreasuryContractSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TreasuryContract.Contract.RenounceRole(&_TreasuryContract.TransactOpts, role, callerConfirmation)
}
func (_TreasuryContract *TreasuryContractTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TreasuryContract.Contract.RenounceRole(&_TreasuryContract.TransactOpts, role, callerConfirmation)
}
func (_TreasuryContract *TreasuryContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "revokeRole", role, account)
}
func (_TreasuryContract *TreasuryContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasuryContract.Contract.RevokeRole(&_TreasuryContract.TransactOpts, role, account)
}
func (_TreasuryContract *TreasuryContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasuryContract.Contract.RevokeRole(&_TreasuryContract.TransactOpts, role, account)
}
func (_TreasuryContract *TreasuryContractTransactor) SetDailyLimit(opts *bind.TransactOpts, newLimit *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "setDailyLimit", newLimit)
}
func (_TreasuryContract *TreasuryContractSession) SetDailyLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.SetDailyLimit(&_TreasuryContract.TransactOpts, newLimit)
}
func (_TreasuryContract *TreasuryContractTransactorSession) SetDailyLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.SetDailyLimit(&_TreasuryContract.TransactOpts, newLimit)
}
func (_TreasuryContract *TreasuryContractTransactor) SetRequiredApprovals(opts *bind.TransactOpts, newRequired *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "setRequiredApprovals", newRequired)
}
func (_TreasuryContract *TreasuryContractSession) SetRequiredApprovals(newRequired *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.SetRequiredApprovals(&_TreasuryContract.TransactOpts, newRequired)
}
func (_TreasuryContract *TreasuryContractTransactorSession) SetRequiredApprovals(newRequired *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.SetRequiredApprovals(&_TreasuryContract.TransactOpts, newRequired)
}
func (_TreasuryContract *TreasuryContractTransactor) SetTimelockDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "setTimelockDelay", newDelay)
}
func (_TreasuryContract *TreasuryContractSession) SetTimelockDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.SetTimelockDelay(&_TreasuryContract.TransactOpts, newDelay)
}
func (_TreasuryContract *TreasuryContractTransactorSession) SetTimelockDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _TreasuryContract.Contract.SetTimelockDelay(&_TreasuryContract.TransactOpts, newDelay)
}
func (_TreasuryContract *TreasuryContractTransactor) Unfreeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryContract.contract.Transact(opts, "unfreeze")
}
func (_TreasuryContract *TreasuryContractSession) Unfreeze() (*types.Transaction, error) {
	return _TreasuryContract.Contract.Unfreeze(&_TreasuryContract.TransactOpts)
}
func (_TreasuryContract *TreasuryContractTransactorSession) Unfreeze() (*types.Transaction, error) {
	return _TreasuryContract.Contract.Unfreeze(&_TreasuryContract.TransactOpts)
}
type TreasuryContractDailyLimitUpdatedIterator struct {
	Event *TreasuryContractDailyLimitUpdated 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *TreasuryContractDailyLimitUpdatedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *TreasuryContractDailyLimitUpdatedIterator) Error() error {
	return it.fail
}
func (it *TreasuryContractDailyLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type TreasuryContractDailyLimitUpdated struct {
	OldLimit *big.Int
	NewLimit *big.Int
	Raw      types.Log 
}
func (_TreasuryContract *TreasuryContractFilterer) FilterDailyLimitUpdated(opts *bind.FilterOpts) (*TreasuryContractDailyLimitUpdatedIterator, error) {
	logs, sub, err := _TreasuryContract.contract.FilterLogs(opts, "DailyLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &TreasuryContractDailyLimitUpdatedIterator{contract: _TreasuryContract.contract, event: "DailyLimitUpdated", logs: logs, sub: sub}, nil
}
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
func (_TreasuryContract *TreasuryContractFilterer) ParseDailyLimitUpdated(log types.Log) (*TreasuryContractDailyLimitUpdated, error) {
	event := new(TreasuryContractDailyLimitUpdated)
	if err := _TreasuryContract.contract.UnpackLog(event, "DailyLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type TreasuryContractDepositedIterator struct {
	Event *TreasuryContractDeposited 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *TreasuryContractDepositedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *TreasuryContractDepositedIterator) Error() error {
	return it.fail
}
func (it *TreasuryContractDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type TreasuryContractDeposited struct {
	Depositor  common.Address
	Amount     *big.Int
	NewBalance *big.Int
	Raw        types.Log 
}
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
func (_TreasuryContract *TreasuryContractFilterer) ParseDeposited(log types.Log) (*TreasuryContractDeposited, error) {
	event := new(TreasuryContractDeposited)
	if err := _TreasuryContract.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type TreasuryContractRequiredApprovalsUpdatedIterator struct {
	Event *TreasuryContractRequiredApprovalsUpdated 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *TreasuryContractRequiredApprovalsUpdatedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *TreasuryContractRequiredApprovalsUpdatedIterator) Error() error {
	return it.fail
}
func (it *TreasuryContractRequiredApprovalsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type TreasuryContractRequiredApprovalsUpdated struct {
	OldRequired *big.Int
	NewRequired *big.Int
	Raw         types.Log 
}
func (_TreasuryContract *TreasuryContractFilterer) FilterRequiredApprovalsUpdated(opts *bind.FilterOpts) (*TreasuryContractRequiredApprovalsUpdatedIterator, error) {
	logs, sub, err := _TreasuryContract.contract.FilterLogs(opts, "RequiredApprovalsUpdated")
	if err != nil {
		return nil, err
	}
	return &TreasuryContractRequiredApprovalsUpdatedIterator{contract: _TreasuryContract.contract, event: "RequiredApprovalsUpdated", logs: logs, sub: sub}, nil
}
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
func (_TreasuryContract *TreasuryContractFilterer) ParseRequiredApprovalsUpdated(log types.Log) (*TreasuryContractRequiredApprovalsUpdated, error) {
	event := new(TreasuryContractRequiredApprovalsUpdated)
	if err := _TreasuryContract.contract.UnpackLog(event, "RequiredApprovalsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type TreasuryContractRoleAdminChangedIterator struct {
	Event *TreasuryContractRoleAdminChanged 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *TreasuryContractRoleAdminChangedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *TreasuryContractRoleAdminChangedIterator) Error() error {
	return it.fail
}
func (it *TreasuryContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type TreasuryContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log 
}
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
func (_TreasuryContract *TreasuryContractFilterer) ParseRoleAdminChanged(log types.Log) (*TreasuryContractRoleAdminChanged, error) {
	event := new(TreasuryContractRoleAdminChanged)
	if err := _TreasuryContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type TreasuryContractRoleGrantedIterator struct {
	Event *TreasuryContractRoleGranted 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *TreasuryContractRoleGrantedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *TreasuryContractRoleGrantedIterator) Error() error {
	return it.fail
}
func (it *TreasuryContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type TreasuryContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log 
}
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
func (_TreasuryContract *TreasuryContractFilterer) ParseRoleGranted(log types.Log) (*TreasuryContractRoleGranted, error) {
	event := new(TreasuryContractRoleGranted)
	if err := _TreasuryContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type TreasuryContractRoleRevokedIterator struct {
	Event *TreasuryContractRoleRevoked 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *TreasuryContractRoleRevokedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *TreasuryContractRoleRevokedIterator) Error() error {
	return it.fail
}
func (it *TreasuryContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type TreasuryContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log 
}
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
func (_TreasuryContract *TreasuryContractFilterer) ParseRoleRevoked(log types.Log) (*TreasuryContractRoleRevoked, error) {
	event := new(TreasuryContractRoleRevoked)
	if err := _TreasuryContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type TreasuryContractSignerChangeApprovedIterator struct {
	Event *TreasuryContractSignerChangeApproved 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *TreasuryContractSignerChangeApprovedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *TreasuryContractSignerChangeApprovedIterator) Error() error {
	return it.fail
}
func (it *TreasuryContractSignerChangeApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type TreasuryContractSignerChangeApproved struct {
	ProposalId    *big.Int
	Approver      common.Address
	ApprovalCount *big.Int
	Raw           types.Log 
}
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
func (_TreasuryContract *TreasuryContractFilterer) ParseSignerChangeApproved(log types.Log) (*TreasuryContractSignerChangeApproved, error) {
	event := new(TreasuryContractSignerChangeApproved)
	if err := _TreasuryContract.contract.UnpackLog(event, "SignerChangeApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type TreasuryContractSignerChangeExecutedIterator struct {
	Event *TreasuryContractSignerChangeExecuted 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *TreasuryContractSignerChangeExecutedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *TreasuryContractSignerChangeExecutedIterator) Error() error {
	return it.fail
}
func (it *TreasuryContractSignerChangeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type TreasuryContractSignerChangeExecuted struct {
	ProposalId  *big.Int
	Target      common.Address
	Replacement common.Address
	ChangeType  uint8
	Raw         types.Log 
}
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
func (_TreasuryContract *TreasuryContractFilterer) ParseSignerChangeExecuted(log types.Log) (*TreasuryContractSignerChangeExecuted, error) {
	event := new(TreasuryContractSignerChangeExecuted)
	if err := _TreasuryContract.contract.UnpackLog(event, "SignerChangeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type TreasuryContractSignerChangeProposedIterator struct {
	Event *TreasuryContractSignerChangeProposed 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *TreasuryContractSignerChangeProposedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *TreasuryContractSignerChangeProposedIterator) Error() error {
	return it.fail
}
func (it *TreasuryContractSignerChangeProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type TreasuryContractSignerChangeProposed struct {
	ProposalId  *big.Int
	Proposer    common.Address
	Target      common.Address
	Replacement common.Address
	ChangeType  uint8
	Raw         types.Log 
}
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
func (_TreasuryContract *TreasuryContractFilterer) ParseSignerChangeProposed(log types.Log) (*TreasuryContractSignerChangeProposed, error) {
	event := new(TreasuryContractSignerChangeProposed)
	if err := _TreasuryContract.contract.UnpackLog(event, "SignerChangeProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type TreasuryContractTimelockDelayUpdatedIterator struct {
	Event *TreasuryContractTimelockDelayUpdated 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *TreasuryContractTimelockDelayUpdatedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryContractTimelockDelayUpdated)
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
	select {
	case log := <-it.logs:
		it.Event = new(TreasuryContractTimelockDelayUpdated)
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
func (it *TreasuryContractTimelockDelayUpdatedIterator) Error() error {
	return it.fail
}
func (it *TreasuryContractTimelockDelayUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type TreasuryContractTimelockDelayUpdated struct {
	OldDelay *big.Int
	NewDelay *big.Int
	Raw      types.Log 
}
func (_TreasuryContract *TreasuryContractFilterer) FilterTimelockDelayUpdated(opts *bind.FilterOpts) (*TreasuryContractTimelockDelayUpdatedIterator, error) {
	logs, sub, err := _TreasuryContract.contract.FilterLogs(opts, "TimelockDelayUpdated")
	if err != nil {
		return nil, err
	}
	return &TreasuryContractTimelockDelayUpdatedIterator{contract: _TreasuryContract.contract, event: "TimelockDelayUpdated", logs: logs, sub: sub}, nil
}
func (_TreasuryContract *TreasuryContractFilterer) WatchTimelockDelayUpdated(opts *bind.WatchOpts, sink chan<- *TreasuryContractTimelockDelayUpdated) (event.Subscription, error) {
	logs, sub, err := _TreasuryContract.contract.WatchLogs(opts, "TimelockDelayUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(TreasuryContractTimelockDelayUpdated)
				if err := _TreasuryContract.contract.UnpackLog(event, "TimelockDelayUpdated", log); err != nil {
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
func (_TreasuryContract *TreasuryContractFilterer) ParseTimelockDelayUpdated(log types.Log) (*TreasuryContractTimelockDelayUpdated, error) {
	event := new(TreasuryContractTimelockDelayUpdated)
	if err := _TreasuryContract.contract.UnpackLog(event, "TimelockDelayUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type TreasuryContractTreasuryFrozenIterator struct {
	Event *TreasuryContractTreasuryFrozen 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *TreasuryContractTreasuryFrozenIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *TreasuryContractTreasuryFrozenIterator) Error() error {
	return it.fail
}
func (it *TreasuryContractTreasuryFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type TreasuryContractTreasuryFrozen struct {
	By  common.Address
	Raw types.Log 
}
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
func (_TreasuryContract *TreasuryContractFilterer) ParseTreasuryFrozen(log types.Log) (*TreasuryContractTreasuryFrozen, error) {
	event := new(TreasuryContractTreasuryFrozen)
	if err := _TreasuryContract.contract.UnpackLog(event, "TreasuryFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type TreasuryContractTreasuryUnfrozenIterator struct {
	Event *TreasuryContractTreasuryUnfrozen 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *TreasuryContractTreasuryUnfrozenIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *TreasuryContractTreasuryUnfrozenIterator) Error() error {
	return it.fail
}
func (it *TreasuryContractTreasuryUnfrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type TreasuryContractTreasuryUnfrozen struct {
	By  common.Address
	Raw types.Log 
}
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
func (_TreasuryContract *TreasuryContractFilterer) ParseTreasuryUnfrozen(log types.Log) (*TreasuryContractTreasuryUnfrozen, error) {
	event := new(TreasuryContractTreasuryUnfrozen)
	if err := _TreasuryContract.contract.UnpackLog(event, "TreasuryUnfrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type TreasuryContractWithdrawalApprovedIterator struct {
	Event *TreasuryContractWithdrawalApproved 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *TreasuryContractWithdrawalApprovedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *TreasuryContractWithdrawalApprovedIterator) Error() error {
	return it.fail
}
func (it *TreasuryContractWithdrawalApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type TreasuryContractWithdrawalApproved struct {
	ProposalId    *big.Int
	Approver      common.Address
	ApprovalCount *big.Int
	Raw           types.Log 
}
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
func (_TreasuryContract *TreasuryContractFilterer) ParseWithdrawalApproved(log types.Log) (*TreasuryContractWithdrawalApproved, error) {
	event := new(TreasuryContractWithdrawalApproved)
	if err := _TreasuryContract.contract.UnpackLog(event, "WithdrawalApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type TreasuryContractWithdrawalCancelledIterator struct {
	Event *TreasuryContractWithdrawalCancelled 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *TreasuryContractWithdrawalCancelledIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *TreasuryContractWithdrawalCancelledIterator) Error() error {
	return it.fail
}
func (it *TreasuryContractWithdrawalCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type TreasuryContractWithdrawalCancelled struct {
	ProposalId *big.Int
	Canceller  common.Address
	Raw        types.Log 
}
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
func (_TreasuryContract *TreasuryContractFilterer) ParseWithdrawalCancelled(log types.Log) (*TreasuryContractWithdrawalCancelled, error) {
	event := new(TreasuryContractWithdrawalCancelled)
	if err := _TreasuryContract.contract.UnpackLog(event, "WithdrawalCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type TreasuryContractWithdrawalExecutedIterator struct {
	Event *TreasuryContractWithdrawalExecuted 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *TreasuryContractWithdrawalExecutedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *TreasuryContractWithdrawalExecutedIterator) Error() error {
	return it.fail
}
func (it *TreasuryContractWithdrawalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type TreasuryContractWithdrawalExecuted struct {
	ProposalId *big.Int
	Recipient  common.Address
	Amount     *big.Int
	Raw        types.Log 
}
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
func (_TreasuryContract *TreasuryContractFilterer) ParseWithdrawalExecuted(log types.Log) (*TreasuryContractWithdrawalExecuted, error) {
	event := new(TreasuryContractWithdrawalExecuted)
	if err := _TreasuryContract.contract.UnpackLog(event, "WithdrawalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type TreasuryContractWithdrawalProposedIterator struct {
	Event *TreasuryContractWithdrawalProposed 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *TreasuryContractWithdrawalProposedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *TreasuryContractWithdrawalProposedIterator) Error() error {
	return it.fail
}
func (it *TreasuryContractWithdrawalProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type TreasuryContractWithdrawalProposed struct {
	ProposalId *big.Int
	Proposer   common.Address
	Recipient  common.Address
	Amount     *big.Int
	Purpose    string
	Raw        types.Log 
}
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
func (_TreasuryContract *TreasuryContractFilterer) ParseWithdrawalProposed(log types.Log) (*TreasuryContractWithdrawalProposed, error) {
	event := new(TreasuryContractWithdrawalProposed)
	if err := _TreasuryContract.contract.UnpackLog(event, "WithdrawalProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
