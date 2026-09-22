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
var FeeRegistryContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recorder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdcToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RECORDER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"studentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"termId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recordFee\",\"inputs\":[{\"name\":\"studentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"termId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdcToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FeeRecorded\",\"inputs\":[{\"name\":\"studentHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"termId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}
var FeeRegistryContractABI = FeeRegistryContractMetaData.ABI
type FeeRegistryContract struct {
	FeeRegistryContractCaller     
	FeeRegistryContractTransactor 
	FeeRegistryContractFilterer   
}
type FeeRegistryContractCaller struct {
	contract *bind.BoundContract 
}
type FeeRegistryContractTransactor struct {
	contract *bind.BoundContract 
}
type FeeRegistryContractFilterer struct {
	contract *bind.BoundContract 
}
type FeeRegistryContractSession struct {
	Contract     *FeeRegistryContract 
	CallOpts     bind.CallOpts        
	TransactOpts bind.TransactOpts    
}
type FeeRegistryContractCallerSession struct {
	Contract *FeeRegistryContractCaller 
	CallOpts bind.CallOpts              
}
type FeeRegistryContractTransactorSession struct {
	Contract     *FeeRegistryContractTransactor 
	TransactOpts bind.TransactOpts              
}
type FeeRegistryContractRaw struct {
	Contract *FeeRegistryContract 
}
type FeeRegistryContractCallerRaw struct {
	Contract *FeeRegistryContractCaller 
}
type FeeRegistryContractTransactorRaw struct {
	Contract *FeeRegistryContractTransactor 
}
func NewFeeRegistryContract(address common.Address, backend bind.ContractBackend) (*FeeRegistryContract, error) {
	contract, err := bindFeeRegistryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeRegistryContract{FeeRegistryContractCaller: FeeRegistryContractCaller{contract: contract}, FeeRegistryContractTransactor: FeeRegistryContractTransactor{contract: contract}, FeeRegistryContractFilterer: FeeRegistryContractFilterer{contract: contract}}, nil
}
func NewFeeRegistryContractCaller(address common.Address, caller bind.ContractCaller) (*FeeRegistryContractCaller, error) {
	contract, err := bindFeeRegistryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeRegistryContractCaller{contract: contract}, nil
}
func NewFeeRegistryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeRegistryContractTransactor, error) {
	contract, err := bindFeeRegistryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeRegistryContractTransactor{contract: contract}, nil
}
func NewFeeRegistryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeRegistryContractFilterer, error) {
	contract, err := bindFeeRegistryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeRegistryContractFilterer{contract: contract}, nil
}
func bindFeeRegistryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeeRegistryContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}
func (_FeeRegistryContract *FeeRegistryContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeRegistryContract.Contract.FeeRegistryContractCaller.contract.Call(opts, result, method, params...)
}
func (_FeeRegistryContract *FeeRegistryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.FeeRegistryContractTransactor.contract.Transfer(opts)
}
func (_FeeRegistryContract *FeeRegistryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.FeeRegistryContractTransactor.contract.Transact(opts, method, params...)
}
func (_FeeRegistryContract *FeeRegistryContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeRegistryContract.Contract.contract.Call(opts, result, method, params...)
}
func (_FeeRegistryContract *FeeRegistryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.contract.Transfer(opts)
}
func (_FeeRegistryContract *FeeRegistryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.contract.Transact(opts, method, params...)
}
func (_FeeRegistryContract *FeeRegistryContractCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FeeRegistryContract.contract.Call(opts, &out, "ADMIN_ROLE")
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_FeeRegistryContract *FeeRegistryContractSession) ADMINROLE() ([32]byte, error) {
	return _FeeRegistryContract.Contract.ADMINROLE(&_FeeRegistryContract.CallOpts)
}
func (_FeeRegistryContract *FeeRegistryContractCallerSession) ADMINROLE() ([32]byte, error) {
	return _FeeRegistryContract.Contract.ADMINROLE(&_FeeRegistryContract.CallOpts)
}
func (_FeeRegistryContract *FeeRegistryContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FeeRegistryContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_FeeRegistryContract *FeeRegistryContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FeeRegistryContract.Contract.DEFAULTADMINROLE(&_FeeRegistryContract.CallOpts)
}
func (_FeeRegistryContract *FeeRegistryContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FeeRegistryContract.Contract.DEFAULTADMINROLE(&_FeeRegistryContract.CallOpts)
}
func (_FeeRegistryContract *FeeRegistryContractCaller) RECORDERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FeeRegistryContract.contract.Call(opts, &out, "RECORDER_ROLE")
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_FeeRegistryContract *FeeRegistryContractSession) RECORDERROLE() ([32]byte, error) {
	return _FeeRegistryContract.Contract.RECORDERROLE(&_FeeRegistryContract.CallOpts)
}
func (_FeeRegistryContract *FeeRegistryContractCallerSession) RECORDERROLE() ([32]byte, error) {
	return _FeeRegistryContract.Contract.RECORDERROLE(&_FeeRegistryContract.CallOpts)
}
func (_FeeRegistryContract *FeeRegistryContractCaller) GetFee(opts *bind.CallOpts, studentHash [32]byte, termId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeeRegistryContract.contract.Call(opts, &out, "getFee", studentHash, termId)
	if err != nil {
		return *new(*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return out0, err
}
func (_FeeRegistryContract *FeeRegistryContractSession) GetFee(studentHash [32]byte, termId *big.Int) (*big.Int, error) {
	return _FeeRegistryContract.Contract.GetFee(&_FeeRegistryContract.CallOpts, studentHash, termId)
}
func (_FeeRegistryContract *FeeRegistryContractCallerSession) GetFee(studentHash [32]byte, termId *big.Int) (*big.Int, error) {
	return _FeeRegistryContract.Contract.GetFee(&_FeeRegistryContract.CallOpts, studentHash, termId)
}
func (_FeeRegistryContract *FeeRegistryContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _FeeRegistryContract.contract.Call(opts, &out, "getRoleAdmin", role)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_FeeRegistryContract *FeeRegistryContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FeeRegistryContract.Contract.GetRoleAdmin(&_FeeRegistryContract.CallOpts, role)
}
func (_FeeRegistryContract *FeeRegistryContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FeeRegistryContract.Contract.GetRoleAdmin(&_FeeRegistryContract.CallOpts, role)
}
func (_FeeRegistryContract *FeeRegistryContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _FeeRegistryContract.contract.Call(opts, &out, "hasRole", role, account)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}
func (_FeeRegistryContract *FeeRegistryContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FeeRegistryContract.Contract.HasRole(&_FeeRegistryContract.CallOpts, role, account)
}
func (_FeeRegistryContract *FeeRegistryContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FeeRegistryContract.Contract.HasRole(&_FeeRegistryContract.CallOpts, role, account)
}
func (_FeeRegistryContract *FeeRegistryContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FeeRegistryContract.contract.Call(opts, &out, "supportsInterface", interfaceId)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}
func (_FeeRegistryContract *FeeRegistryContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeeRegistryContract.Contract.SupportsInterface(&_FeeRegistryContract.CallOpts, interfaceId)
}
func (_FeeRegistryContract *FeeRegistryContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeeRegistryContract.Contract.SupportsInterface(&_FeeRegistryContract.CallOpts, interfaceId)
}
func (_FeeRegistryContract *FeeRegistryContractCaller) UsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeRegistryContract.contract.Call(opts, &out, "usdcToken")
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}
func (_FeeRegistryContract *FeeRegistryContractSession) UsdcToken() (common.Address, error) {
	return _FeeRegistryContract.Contract.UsdcToken(&_FeeRegistryContract.CallOpts)
}
func (_FeeRegistryContract *FeeRegistryContractCallerSession) UsdcToken() (common.Address, error) {
	return _FeeRegistryContract.Contract.UsdcToken(&_FeeRegistryContract.CallOpts)
}
func (_FeeRegistryContract *FeeRegistryContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeRegistryContract.contract.Transact(opts, "grantRole", role, account)
}
func (_FeeRegistryContract *FeeRegistryContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.GrantRole(&_FeeRegistryContract.TransactOpts, role, account)
}
func (_FeeRegistryContract *FeeRegistryContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.GrantRole(&_FeeRegistryContract.TransactOpts, role, account)
}
func (_FeeRegistryContract *FeeRegistryContractTransactor) RecordFee(opts *bind.TransactOpts, studentHash [32]byte, termId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _FeeRegistryContract.contract.Transact(opts, "recordFee", studentHash, termId, amount)
}
func (_FeeRegistryContract *FeeRegistryContractSession) RecordFee(studentHash [32]byte, termId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.RecordFee(&_FeeRegistryContract.TransactOpts, studentHash, termId, amount)
}
func (_FeeRegistryContract *FeeRegistryContractTransactorSession) RecordFee(studentHash [32]byte, termId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.RecordFee(&_FeeRegistryContract.TransactOpts, studentHash, termId, amount)
}
func (_FeeRegistryContract *FeeRegistryContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _FeeRegistryContract.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}
func (_FeeRegistryContract *FeeRegistryContractSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.RenounceRole(&_FeeRegistryContract.TransactOpts, role, callerConfirmation)
}
func (_FeeRegistryContract *FeeRegistryContractTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.RenounceRole(&_FeeRegistryContract.TransactOpts, role, callerConfirmation)
}
func (_FeeRegistryContract *FeeRegistryContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeRegistryContract.contract.Transact(opts, "revokeRole", role, account)
}
func (_FeeRegistryContract *FeeRegistryContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.RevokeRole(&_FeeRegistryContract.TransactOpts, role, account)
}
func (_FeeRegistryContract *FeeRegistryContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.RevokeRole(&_FeeRegistryContract.TransactOpts, role, account)
}
type FeeRegistryContractFeeRecordedIterator struct {
	Event *FeeRegistryContractFeeRecorded 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *FeeRegistryContractFeeRecordedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeRegistryContractFeeRecorded)
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
		it.Event = new(FeeRegistryContractFeeRecorded)
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
func (it *FeeRegistryContractFeeRecordedIterator) Error() error {
	return it.fail
}
func (it *FeeRegistryContractFeeRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type FeeRegistryContractFeeRecorded struct {
	StudentHash [32]byte
	TermId      *big.Int
	Amount      *big.Int
	Raw         types.Log 
}
func (_FeeRegistryContract *FeeRegistryContractFilterer) FilterFeeRecorded(opts *bind.FilterOpts, studentHash [][32]byte, termId []*big.Int) (*FeeRegistryContractFeeRecordedIterator, error) {
	var studentHashRule []interface{}
	for _, studentHashItem := range studentHash {
		studentHashRule = append(studentHashRule, studentHashItem)
	}
	var termIdRule []interface{}
	for _, termIdItem := range termId {
		termIdRule = append(termIdRule, termIdItem)
	}
	logs, sub, err := _FeeRegistryContract.contract.FilterLogs(opts, "FeeRecorded", studentHashRule, termIdRule)
	if err != nil {
		return nil, err
	}
	return &FeeRegistryContractFeeRecordedIterator{contract: _FeeRegistryContract.contract, event: "FeeRecorded", logs: logs, sub: sub}, nil
}
func (_FeeRegistryContract *FeeRegistryContractFilterer) WatchFeeRecorded(opts *bind.WatchOpts, sink chan<- *FeeRegistryContractFeeRecorded, studentHash [][32]byte, termId []*big.Int) (event.Subscription, error) {
	var studentHashRule []interface{}
	for _, studentHashItem := range studentHash {
		studentHashRule = append(studentHashRule, studentHashItem)
	}
	var termIdRule []interface{}
	for _, termIdItem := range termId {
		termIdRule = append(termIdRule, termIdItem)
	}
	logs, sub, err := _FeeRegistryContract.contract.WatchLogs(opts, "FeeRecorded", studentHashRule, termIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(FeeRegistryContractFeeRecorded)
				if err := _FeeRegistryContract.contract.UnpackLog(event, "FeeRecorded", log); err != nil {
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
func (_FeeRegistryContract *FeeRegistryContractFilterer) ParseFeeRecorded(log types.Log) (*FeeRegistryContractFeeRecorded, error) {
	event := new(FeeRegistryContractFeeRecorded)
	if err := _FeeRegistryContract.contract.UnpackLog(event, "FeeRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type FeeRegistryContractRoleAdminChangedIterator struct {
	Event *FeeRegistryContractRoleAdminChanged 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *FeeRegistryContractRoleAdminChangedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeRegistryContractRoleAdminChanged)
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
		it.Event = new(FeeRegistryContractRoleAdminChanged)
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
func (it *FeeRegistryContractRoleAdminChangedIterator) Error() error {
	return it.fail
}
func (it *FeeRegistryContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type FeeRegistryContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log 
}
func (_FeeRegistryContract *FeeRegistryContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FeeRegistryContractRoleAdminChangedIterator, error) {
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
	logs, sub, err := _FeeRegistryContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FeeRegistryContractRoleAdminChangedIterator{contract: _FeeRegistryContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}
func (_FeeRegistryContract *FeeRegistryContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FeeRegistryContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {
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
	logs, sub, err := _FeeRegistryContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(FeeRegistryContractRoleAdminChanged)
				if err := _FeeRegistryContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_FeeRegistryContract *FeeRegistryContractFilterer) ParseRoleAdminChanged(log types.Log) (*FeeRegistryContractRoleAdminChanged, error) {
	event := new(FeeRegistryContractRoleAdminChanged)
	if err := _FeeRegistryContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type FeeRegistryContractRoleGrantedIterator struct {
	Event *FeeRegistryContractRoleGranted 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *FeeRegistryContractRoleGrantedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeRegistryContractRoleGranted)
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
		it.Event = new(FeeRegistryContractRoleGranted)
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
func (it *FeeRegistryContractRoleGrantedIterator) Error() error {
	return it.fail
}
func (it *FeeRegistryContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type FeeRegistryContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log 
}
func (_FeeRegistryContract *FeeRegistryContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FeeRegistryContractRoleGrantedIterator, error) {
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
	logs, sub, err := _FeeRegistryContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FeeRegistryContractRoleGrantedIterator{contract: _FeeRegistryContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}
func (_FeeRegistryContract *FeeRegistryContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FeeRegistryContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {
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
	logs, sub, err := _FeeRegistryContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(FeeRegistryContractRoleGranted)
				if err := _FeeRegistryContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_FeeRegistryContract *FeeRegistryContractFilterer) ParseRoleGranted(log types.Log) (*FeeRegistryContractRoleGranted, error) {
	event := new(FeeRegistryContractRoleGranted)
	if err := _FeeRegistryContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type FeeRegistryContractRoleRevokedIterator struct {
	Event *FeeRegistryContractRoleRevoked 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *FeeRegistryContractRoleRevokedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeRegistryContractRoleRevoked)
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
		it.Event = new(FeeRegistryContractRoleRevoked)
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
func (it *FeeRegistryContractRoleRevokedIterator) Error() error {
	return it.fail
}
func (it *FeeRegistryContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type FeeRegistryContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log 
}
func (_FeeRegistryContract *FeeRegistryContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FeeRegistryContractRoleRevokedIterator, error) {
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
	logs, sub, err := _FeeRegistryContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FeeRegistryContractRoleRevokedIterator{contract: _FeeRegistryContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}
func (_FeeRegistryContract *FeeRegistryContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FeeRegistryContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {
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
	logs, sub, err := _FeeRegistryContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(FeeRegistryContractRoleRevoked)
				if err := _FeeRegistryContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_FeeRegistryContract *FeeRegistryContractFilterer) ParseRoleRevoked(log types.Log) (*FeeRegistryContractRoleRevoked, error) {
	event := new(FeeRegistryContractRoleRevoked)
	if err := _FeeRegistryContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
