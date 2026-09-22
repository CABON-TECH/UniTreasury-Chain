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
var ZKEnrollmentRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_verifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRAR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"enrollmentMerkleRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isEnrollmentVerified\",\"inputs\":[{\"name\":\"nullifierHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"student\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proveEnrollment\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nullifierHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"claimedRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateMerkleRoot\",\"inputs\":[{\"name\":\"newRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifiedNullifier\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractMockZKVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EnrollmentProofVerified\",\"inputs\":[{\"name\":\"nullifierHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"student\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MerkleRootUpdated\",\"inputs\":[{\"name\":\"newRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"updatedBy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZKRegistry__InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZKRegistry__NullifierAlreadyUsed\",\"inputs\":[]}]",
}
var ZKEnrollmentRegistryABI = ZKEnrollmentRegistryMetaData.ABI
type ZKEnrollmentRegistry struct {
	ZKEnrollmentRegistryCaller     
	ZKEnrollmentRegistryTransactor 
	ZKEnrollmentRegistryFilterer   
}
type ZKEnrollmentRegistryCaller struct {
	contract *bind.BoundContract 
}
type ZKEnrollmentRegistryTransactor struct {
	contract *bind.BoundContract 
}
type ZKEnrollmentRegistryFilterer struct {
	contract *bind.BoundContract 
}
type ZKEnrollmentRegistrySession struct {
	Contract     *ZKEnrollmentRegistry 
	CallOpts     bind.CallOpts         
	TransactOpts bind.TransactOpts     
}
type ZKEnrollmentRegistryCallerSession struct {
	Contract *ZKEnrollmentRegistryCaller 
	CallOpts bind.CallOpts               
}
type ZKEnrollmentRegistryTransactorSession struct {
	Contract     *ZKEnrollmentRegistryTransactor 
	TransactOpts bind.TransactOpts               
}
type ZKEnrollmentRegistryRaw struct {
	Contract *ZKEnrollmentRegistry 
}
type ZKEnrollmentRegistryCallerRaw struct {
	Contract *ZKEnrollmentRegistryCaller 
}
type ZKEnrollmentRegistryTransactorRaw struct {
	Contract *ZKEnrollmentRegistryTransactor 
}
func NewZKEnrollmentRegistry(address common.Address, backend bind.ContractBackend) (*ZKEnrollmentRegistry, error) {
	contract, err := bindZKEnrollmentRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZKEnrollmentRegistry{ZKEnrollmentRegistryCaller: ZKEnrollmentRegistryCaller{contract: contract}, ZKEnrollmentRegistryTransactor: ZKEnrollmentRegistryTransactor{contract: contract}, ZKEnrollmentRegistryFilterer: ZKEnrollmentRegistryFilterer{contract: contract}}, nil
}
func NewZKEnrollmentRegistryCaller(address common.Address, caller bind.ContractCaller) (*ZKEnrollmentRegistryCaller, error) {
	contract, err := bindZKEnrollmentRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZKEnrollmentRegistryCaller{contract: contract}, nil
}
func NewZKEnrollmentRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ZKEnrollmentRegistryTransactor, error) {
	contract, err := bindZKEnrollmentRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZKEnrollmentRegistryTransactor{contract: contract}, nil
}
func NewZKEnrollmentRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ZKEnrollmentRegistryFilterer, error) {
	contract, err := bindZKEnrollmentRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZKEnrollmentRegistryFilterer{contract: contract}, nil
}
func bindZKEnrollmentRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZKEnrollmentRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKEnrollmentRegistry.Contract.ZKEnrollmentRegistryCaller.contract.Call(opts, result, method, params...)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.ZKEnrollmentRegistryTransactor.contract.Transfer(opts)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.ZKEnrollmentRegistryTransactor.contract.Transact(opts, method, params...)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKEnrollmentRegistry.Contract.contract.Call(opts, result, method, params...)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.contract.Transfer(opts)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.contract.Transact(opts, method, params...)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "ADMIN_ROLE")
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) ADMINROLE() ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.ADMINROLE(&_ZKEnrollmentRegistry.CallOpts)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) ADMINROLE() ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.ADMINROLE(&_ZKEnrollmentRegistry.CallOpts)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.DEFAULTADMINROLE(&_ZKEnrollmentRegistry.CallOpts)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.DEFAULTADMINROLE(&_ZKEnrollmentRegistry.CallOpts)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) REGISTRARROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "REGISTRAR_ROLE")
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) REGISTRARROLE() ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.REGISTRARROLE(&_ZKEnrollmentRegistry.CallOpts)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) REGISTRARROLE() ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.REGISTRARROLE(&_ZKEnrollmentRegistry.CallOpts)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) EnrollmentMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "enrollmentMerkleRoot")
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) EnrollmentMerkleRoot() ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.EnrollmentMerkleRoot(&_ZKEnrollmentRegistry.CallOpts)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) EnrollmentMerkleRoot() ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.EnrollmentMerkleRoot(&_ZKEnrollmentRegistry.CallOpts)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "getRoleAdmin", role)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.GetRoleAdmin(&_ZKEnrollmentRegistry.CallOpts, role)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.GetRoleAdmin(&_ZKEnrollmentRegistry.CallOpts, role)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "hasRole", role, account)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZKEnrollmentRegistry.Contract.HasRole(&_ZKEnrollmentRegistry.CallOpts, role, account)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZKEnrollmentRegistry.Contract.HasRole(&_ZKEnrollmentRegistry.CallOpts, role, account)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) IsEnrollmentVerified(opts *bind.CallOpts, nullifierHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "isEnrollmentVerified", nullifierHash)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) IsEnrollmentVerified(nullifierHash [32]byte) (common.Address, error) {
	return _ZKEnrollmentRegistry.Contract.IsEnrollmentVerified(&_ZKEnrollmentRegistry.CallOpts, nullifierHash)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) IsEnrollmentVerified(nullifierHash [32]byte) (common.Address, error) {
	return _ZKEnrollmentRegistry.Contract.IsEnrollmentVerified(&_ZKEnrollmentRegistry.CallOpts, nullifierHash)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZKEnrollmentRegistry.Contract.SupportsInterface(&_ZKEnrollmentRegistry.CallOpts, interfaceId)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZKEnrollmentRegistry.Contract.SupportsInterface(&_ZKEnrollmentRegistry.CallOpts, interfaceId)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) VerifiedNullifier(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "verifiedNullifier", arg0)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) VerifiedNullifier(arg0 [32]byte) (common.Address, error) {
	return _ZKEnrollmentRegistry.Contract.VerifiedNullifier(&_ZKEnrollmentRegistry.CallOpts, arg0)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) VerifiedNullifier(arg0 [32]byte) (common.Address, error) {
	return _ZKEnrollmentRegistry.Contract.VerifiedNullifier(&_ZKEnrollmentRegistry.CallOpts, arg0)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "verifier")
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) Verifier() (common.Address, error) {
	return _ZKEnrollmentRegistry.Contract.Verifier(&_ZKEnrollmentRegistry.CallOpts)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) Verifier() (common.Address, error) {
	return _ZKEnrollmentRegistry.Contract.Verifier(&_ZKEnrollmentRegistry.CallOpts)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.contract.Transact(opts, "grantRole", role, account)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.GrantRole(&_ZKEnrollmentRegistry.TransactOpts, role, account)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.GrantRole(&_ZKEnrollmentRegistry.TransactOpts, role, account)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactor) ProveEnrollment(opts *bind.TransactOpts, proof []byte, nullifierHash [32]byte, claimedRoot [32]byte) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.contract.Transact(opts, "proveEnrollment", proof, nullifierHash, claimedRoot)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) ProveEnrollment(proof []byte, nullifierHash [32]byte, claimedRoot [32]byte) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.ProveEnrollment(&_ZKEnrollmentRegistry.TransactOpts, proof, nullifierHash, claimedRoot)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactorSession) ProveEnrollment(proof []byte, nullifierHash [32]byte, claimedRoot [32]byte) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.ProveEnrollment(&_ZKEnrollmentRegistry.TransactOpts, proof, nullifierHash, claimedRoot)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.RenounceRole(&_ZKEnrollmentRegistry.TransactOpts, role, callerConfirmation)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.RenounceRole(&_ZKEnrollmentRegistry.TransactOpts, role, callerConfirmation)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.contract.Transact(opts, "revokeRole", role, account)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.RevokeRole(&_ZKEnrollmentRegistry.TransactOpts, role, account)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.RevokeRole(&_ZKEnrollmentRegistry.TransactOpts, role, account)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactor) UpdateMerkleRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.contract.Transact(opts, "updateMerkleRoot", newRoot)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) UpdateMerkleRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.UpdateMerkleRoot(&_ZKEnrollmentRegistry.TransactOpts, newRoot)
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactorSession) UpdateMerkleRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.UpdateMerkleRoot(&_ZKEnrollmentRegistry.TransactOpts, newRoot)
}
type ZKEnrollmentRegistryEnrollmentProofVerifiedIterator struct {
	Event *ZKEnrollmentRegistryEnrollmentProofVerified 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *ZKEnrollmentRegistryEnrollmentProofVerifiedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEnrollmentRegistryEnrollmentProofVerified)
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
		it.Event = new(ZKEnrollmentRegistryEnrollmentProofVerified)
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
func (it *ZKEnrollmentRegistryEnrollmentProofVerifiedIterator) Error() error {
	return it.fail
}
func (it *ZKEnrollmentRegistryEnrollmentProofVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type ZKEnrollmentRegistryEnrollmentProofVerified struct {
	NullifierHash [32]byte
	Student       common.Address
	Raw           types.Log 
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) FilterEnrollmentProofVerified(opts *bind.FilterOpts, nullifierHash [][32]byte, student []common.Address) (*ZKEnrollmentRegistryEnrollmentProofVerifiedIterator, error) {
	var nullifierHashRule []interface{}
	for _, nullifierHashItem := range nullifierHash {
		nullifierHashRule = append(nullifierHashRule, nullifierHashItem)
	}
	var studentRule []interface{}
	for _, studentItem := range student {
		studentRule = append(studentRule, studentItem)
	}
	logs, sub, err := _ZKEnrollmentRegistry.contract.FilterLogs(opts, "EnrollmentProofVerified", nullifierHashRule, studentRule)
	if err != nil {
		return nil, err
	}
	return &ZKEnrollmentRegistryEnrollmentProofVerifiedIterator{contract: _ZKEnrollmentRegistry.contract, event: "EnrollmentProofVerified", logs: logs, sub: sub}, nil
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) WatchEnrollmentProofVerified(opts *bind.WatchOpts, sink chan<- *ZKEnrollmentRegistryEnrollmentProofVerified, nullifierHash [][32]byte, student []common.Address) (event.Subscription, error) {
	var nullifierHashRule []interface{}
	for _, nullifierHashItem := range nullifierHash {
		nullifierHashRule = append(nullifierHashRule, nullifierHashItem)
	}
	var studentRule []interface{}
	for _, studentItem := range student {
		studentRule = append(studentRule, studentItem)
	}
	logs, sub, err := _ZKEnrollmentRegistry.contract.WatchLogs(opts, "EnrollmentProofVerified", nullifierHashRule, studentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(ZKEnrollmentRegistryEnrollmentProofVerified)
				if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "EnrollmentProofVerified", log); err != nil {
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
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) ParseEnrollmentProofVerified(log types.Log) (*ZKEnrollmentRegistryEnrollmentProofVerified, error) {
	event := new(ZKEnrollmentRegistryEnrollmentProofVerified)
	if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "EnrollmentProofVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type ZKEnrollmentRegistryMerkleRootUpdatedIterator struct {
	Event *ZKEnrollmentRegistryMerkleRootUpdated 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *ZKEnrollmentRegistryMerkleRootUpdatedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEnrollmentRegistryMerkleRootUpdated)
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
		it.Event = new(ZKEnrollmentRegistryMerkleRootUpdated)
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
func (it *ZKEnrollmentRegistryMerkleRootUpdatedIterator) Error() error {
	return it.fail
}
func (it *ZKEnrollmentRegistryMerkleRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type ZKEnrollmentRegistryMerkleRootUpdated struct {
	NewRoot   [32]byte
	UpdatedBy common.Address
	Raw       types.Log 
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) FilterMerkleRootUpdated(opts *bind.FilterOpts, newRoot [][32]byte, updatedBy []common.Address) (*ZKEnrollmentRegistryMerkleRootUpdatedIterator, error) {
	var newRootRule []interface{}
	for _, newRootItem := range newRoot {
		newRootRule = append(newRootRule, newRootItem)
	}
	var updatedByRule []interface{}
	for _, updatedByItem := range updatedBy {
		updatedByRule = append(updatedByRule, updatedByItem)
	}
	logs, sub, err := _ZKEnrollmentRegistry.contract.FilterLogs(opts, "MerkleRootUpdated", newRootRule, updatedByRule)
	if err != nil {
		return nil, err
	}
	return &ZKEnrollmentRegistryMerkleRootUpdatedIterator{contract: _ZKEnrollmentRegistry.contract, event: "MerkleRootUpdated", logs: logs, sub: sub}, nil
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) WatchMerkleRootUpdated(opts *bind.WatchOpts, sink chan<- *ZKEnrollmentRegistryMerkleRootUpdated, newRoot [][32]byte, updatedBy []common.Address) (event.Subscription, error) {
	var newRootRule []interface{}
	for _, newRootItem := range newRoot {
		newRootRule = append(newRootRule, newRootItem)
	}
	var updatedByRule []interface{}
	for _, updatedByItem := range updatedBy {
		updatedByRule = append(updatedByRule, updatedByItem)
	}
	logs, sub, err := _ZKEnrollmentRegistry.contract.WatchLogs(opts, "MerkleRootUpdated", newRootRule, updatedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(ZKEnrollmentRegistryMerkleRootUpdated)
				if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "MerkleRootUpdated", log); err != nil {
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
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) ParseMerkleRootUpdated(log types.Log) (*ZKEnrollmentRegistryMerkleRootUpdated, error) {
	event := new(ZKEnrollmentRegistryMerkleRootUpdated)
	if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "MerkleRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type ZKEnrollmentRegistryRoleAdminChangedIterator struct {
	Event *ZKEnrollmentRegistryRoleAdminChanged 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *ZKEnrollmentRegistryRoleAdminChangedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEnrollmentRegistryRoleAdminChanged)
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
		it.Event = new(ZKEnrollmentRegistryRoleAdminChanged)
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
func (it *ZKEnrollmentRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}
func (it *ZKEnrollmentRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type ZKEnrollmentRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log 
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ZKEnrollmentRegistryRoleAdminChangedIterator, error) {
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
	logs, sub, err := _ZKEnrollmentRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ZKEnrollmentRegistryRoleAdminChangedIterator{contract: _ZKEnrollmentRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ZKEnrollmentRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {
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
	logs, sub, err := _ZKEnrollmentRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(ZKEnrollmentRegistryRoleAdminChanged)
				if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*ZKEnrollmentRegistryRoleAdminChanged, error) {
	event := new(ZKEnrollmentRegistryRoleAdminChanged)
	if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type ZKEnrollmentRegistryRoleGrantedIterator struct {
	Event *ZKEnrollmentRegistryRoleGranted 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *ZKEnrollmentRegistryRoleGrantedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEnrollmentRegistryRoleGranted)
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
		it.Event = new(ZKEnrollmentRegistryRoleGranted)
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
func (it *ZKEnrollmentRegistryRoleGrantedIterator) Error() error {
	return it.fail
}
func (it *ZKEnrollmentRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type ZKEnrollmentRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log 
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZKEnrollmentRegistryRoleGrantedIterator, error) {
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
	logs, sub, err := _ZKEnrollmentRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZKEnrollmentRegistryRoleGrantedIterator{contract: _ZKEnrollmentRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ZKEnrollmentRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {
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
	logs, sub, err := _ZKEnrollmentRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(ZKEnrollmentRegistryRoleGranted)
				if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) ParseRoleGranted(log types.Log) (*ZKEnrollmentRegistryRoleGranted, error) {
	event := new(ZKEnrollmentRegistryRoleGranted)
	if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type ZKEnrollmentRegistryRoleRevokedIterator struct {
	Event *ZKEnrollmentRegistryRoleRevoked 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *ZKEnrollmentRegistryRoleRevokedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEnrollmentRegistryRoleRevoked)
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
		it.Event = new(ZKEnrollmentRegistryRoleRevoked)
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
func (it *ZKEnrollmentRegistryRoleRevokedIterator) Error() error {
	return it.fail
}
func (it *ZKEnrollmentRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type ZKEnrollmentRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log 
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZKEnrollmentRegistryRoleRevokedIterator, error) {
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
	logs, sub, err := _ZKEnrollmentRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZKEnrollmentRegistryRoleRevokedIterator{contract: _ZKEnrollmentRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ZKEnrollmentRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {
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
	logs, sub, err := _ZKEnrollmentRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(ZKEnrollmentRegistryRoleRevoked)
				if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) ParseRoleRevoked(log types.Log) (*ZKEnrollmentRegistryRoleRevoked, error) {
	event := new(ZKEnrollmentRegistryRoleRevoked)
	if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
