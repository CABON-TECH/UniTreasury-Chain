package bindings

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const ScholarshipEscrowContractABI = `[
  {"type":"constructor","inputs":[{"name":"admin","type":"address"},{"name":"attestor","type":"address"}],"stateMutability":"nonpayable"},
  {"type":"function","name":"createFund","inputs":[{"name":"sponsor","type":"address"},{"name":"trancheCount","type":"uint256"},{"name":"trancheAmount","type":"uint256"}],"outputs":[{"name":"","type":"uint256"}],"stateMutability":"payable"},
  {"type":"function","name":"releaseTranche","inputs":[{"name":"fundId","type":"uint256"},{"name":"studentHash","type":"bytes32"},{"name":"trancheIndex","type":"uint256"},{"name":"recipient","type":"address"},{"name":"signature","type":"bytes"}],"outputs":[],"stateMutability":"nonpayable"},
  {"type":"function","name":"pauseFund","inputs":[{"name":"fundId","type":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},
  {"type":"function","name":"unpauseFund","inputs":[{"name":"fundId","type":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},
  {"type":"function","name":"setTrustedAttestor","inputs":[{"name":"newAttestor","type":"address"}],"outputs":[],"stateMutability":"nonpayable"},
  {"type":"function","name":"funds","inputs":[{"name":"fundId","type":"uint256"}],"outputs":[{"name":"sponsor","type":"address"},{"name":"totalAmount","type":"uint256"},{"name":"releasedAmount","type":"uint256"},{"name":"trancheCount","type":"uint256"},{"name":"trancheAmount","type":"uint256"},{"name":"paused","type":"bool"}],"stateMutability":"view"},
  {"type":"function","name":"hasReleased","inputs":[{"name":"fundId","type":"uint256"},{"name":"studentHash","type":"bytes32"},{"name":"trancheIndex","type":"uint256"}],"outputs":[{"name":"","type":"bool"}],"stateMutability":"view"},
  {"type":"function","name":"trustedAttestor","inputs":[],"outputs":[{"name":"","type":"address"}],"stateMutability":"view"},
  {"type":"function","name":"DOMAIN_SEPARATOR","inputs":[],"outputs":[{"name":"","type":"bytes32"}],"stateMutability":"view"},
  {"type":"event","name":"FundCreated","inputs":[{"name":"fundId","type":"uint256","indexed":true},{"name":"sponsor","type":"address","indexed":true},{"name":"totalAmount","type":"uint256","indexed":false},{"name":"trancheCount","type":"uint256","indexed":false},{"name":"trancheAmount","type":"uint256","indexed":false}],"anonymous":false},
  {"type":"event","name":"TrancheReleased","inputs":[{"name":"fundId","type":"uint256","indexed":true},{"name":"studentHash","type":"bytes32","indexed":true},{"name":"trancheIndex","type":"uint256","indexed":false},{"name":"amount","type":"uint256","indexed":false},{"name":"recipient","type":"address","indexed":false}],"anonymous":false}
]`

type ScholarshipEscrowContract struct {
	ScholarshipEscrowContractCaller
	ScholarshipEscrowContractTransactor
}

type ScholarshipEscrowContractCaller struct {
	contract *bind.BoundContract
}

type ScholarshipEscrowContractTransactor struct {
	contract *bind.BoundContract
}

func NewScholarshipEscrowContract(address common.Address, backend bind.ContractBackend) (*ScholarshipEscrowContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ScholarshipEscrowContractABI))
	if err != nil {
		return nil, err
	}
	contract := bind.NewBoundContract(address, parsed, backend, backend, backend)
	return &ScholarshipEscrowContract{
		ScholarshipEscrowContractCaller:     ScholarshipEscrowContractCaller{contract: contract},
		ScholarshipEscrowContractTransactor: ScholarshipEscrowContractTransactor{contract: contract},
	}, nil
}

func (t *ScholarshipEscrowContractTransactor) CreateFund(opts *bind.TransactOpts, sponsor common.Address, trancheCount *big.Int, trancheAmount *big.Int) (*types.Transaction, error) {
	return t.contract.Transact(opts, "createFund", sponsor, trancheCount, trancheAmount)
}

func (t *ScholarshipEscrowContractTransactor) ReleaseTranche(opts *bind.TransactOpts, fundId *big.Int, studentHash [32]byte, trancheIndex *big.Int, recipient common.Address, signature []byte) (*types.Transaction, error) {
	return t.contract.Transact(opts, "releaseTranche", fundId, studentHash, trancheIndex, recipient, signature)
}

func (c *ScholarshipEscrowContractCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "DOMAIN_SEPARATOR")
	if err != nil {
		return [32]byte{}, err
	}
	return *abi.ConvertType(out[0], new([32]byte)).(*[32]byte), nil
}
