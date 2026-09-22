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
type UserOperation struct {
	Sender               common.Address
	Nonce                *big.Int
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     []byte
	Signature            []byte
}
var MockEntryPointMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"handleOps\",\"inputs\":[{\"name\":\"ops\",\"type\":\"tuple[]\",\"internalType\":\"structUserOperation[]\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verificationGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}
var MockEntryPointABI = MockEntryPointMetaData.ABI
type MockEntryPoint struct {
	MockEntryPointCaller     
	MockEntryPointTransactor 
	MockEntryPointFilterer   
}
type MockEntryPointCaller struct {
	contract *bind.BoundContract 
}
type MockEntryPointTransactor struct {
	contract *bind.BoundContract 
}
type MockEntryPointFilterer struct {
	contract *bind.BoundContract 
}
type MockEntryPointSession struct {
	Contract     *MockEntryPoint   
	CallOpts     bind.CallOpts     
	TransactOpts bind.TransactOpts 
}
type MockEntryPointCallerSession struct {
	Contract *MockEntryPointCaller 
	CallOpts bind.CallOpts         
}
type MockEntryPointTransactorSession struct {
	Contract     *MockEntryPointTransactor 
	TransactOpts bind.TransactOpts         
}
type MockEntryPointRaw struct {
	Contract *MockEntryPoint 
}
type MockEntryPointCallerRaw struct {
	Contract *MockEntryPointCaller 
}
type MockEntryPointTransactorRaw struct {
	Contract *MockEntryPointTransactor 
}
func NewMockEntryPoint(address common.Address, backend bind.ContractBackend) (*MockEntryPoint, error) {
	contract, err := bindMockEntryPoint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockEntryPoint{MockEntryPointCaller: MockEntryPointCaller{contract: contract}, MockEntryPointTransactor: MockEntryPointTransactor{contract: contract}, MockEntryPointFilterer: MockEntryPointFilterer{contract: contract}}, nil
}
func NewMockEntryPointCaller(address common.Address, caller bind.ContractCaller) (*MockEntryPointCaller, error) {
	contract, err := bindMockEntryPoint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockEntryPointCaller{contract: contract}, nil
}
func NewMockEntryPointTransactor(address common.Address, transactor bind.ContractTransactor) (*MockEntryPointTransactor, error) {
	contract, err := bindMockEntryPoint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockEntryPointTransactor{contract: contract}, nil
}
func NewMockEntryPointFilterer(address common.Address, filterer bind.ContractFilterer) (*MockEntryPointFilterer, error) {
	contract, err := bindMockEntryPoint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockEntryPointFilterer{contract: contract}, nil
}
func bindMockEntryPoint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockEntryPointMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}
func (_MockEntryPoint *MockEntryPointRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockEntryPoint.Contract.MockEntryPointCaller.contract.Call(opts, result, method, params...)
}
func (_MockEntryPoint *MockEntryPointRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockEntryPoint.Contract.MockEntryPointTransactor.contract.Transfer(opts)
}
func (_MockEntryPoint *MockEntryPointRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockEntryPoint.Contract.MockEntryPointTransactor.contract.Transact(opts, method, params...)
}
func (_MockEntryPoint *MockEntryPointCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockEntryPoint.Contract.contract.Call(opts, result, method, params...)
}
func (_MockEntryPoint *MockEntryPointTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockEntryPoint.Contract.contract.Transfer(opts)
}
func (_MockEntryPoint *MockEntryPointTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockEntryPoint.Contract.contract.Transact(opts, method, params...)
}
func (_MockEntryPoint *MockEntryPointTransactor) HandleOps(opts *bind.TransactOpts, ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _MockEntryPoint.contract.Transact(opts, "handleOps", ops, beneficiary)
}
func (_MockEntryPoint *MockEntryPointSession) HandleOps(ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _MockEntryPoint.Contract.HandleOps(&_MockEntryPoint.TransactOpts, ops, beneficiary)
}
func (_MockEntryPoint *MockEntryPointTransactorSession) HandleOps(ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _MockEntryPoint.Contract.HandleOps(&_MockEntryPoint.TransactOpts, ops, beneficiary)
}
