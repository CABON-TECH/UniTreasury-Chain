package blockchain
import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)
type Client struct {
	clients []*ethclient.Client
	chainID *big.Int
	log     *zap.Logger
}
func NewClient(ctx context.Context, rpcURLs string, chainID int64, log *zap.Logger) (*Client, error) {
	urls := strings.Split(rpcURLs, ",")
	var clients []*ethclient.Client
	var firstChainID *big.Int
	for _, u := range urls {
		u = strings.TrimSpace(u)
		if u == "" {
			continue
		}
		ec, err := ethclient.DialContext(ctx, u)
		if err != nil {
			log.Warn("failed to dial node, skipping", zap.String("url", u), zap.Error(err))
			continue
		}
		gotChainID, err := ec.ChainID(ctx)
		if err != nil {
			log.Warn("failed to get chain ID, skipping", zap.String("url", u), zap.Error(err))
			continue
		}
		if gotChainID.Int64() != chainID {
			log.Warn("chain ID mismatch, skipping", zap.String("url", u), zap.Int64("expected", chainID), zap.Int64("got", gotChainID.Int64()))
			continue
		}
		if len(clients) == 0 {
			firstChainID = gotChainID
		}
		clients = append(clients, ec)
		log.Info("connected to Ethereum node", zap.String("rpc_url", u), zap.Int64("chain_id", gotChainID.Int64()))
	}
	if len(clients) == 0 {
		return nil, fmt.Errorf("blockchain: could not connect to any RPC nodes for chain %d", chainID)
	}
	return &Client{
		clients: clients,
		chainID: firstChainID,
		log:     log,
	}, nil
}
func (c *Client) ChainID() *big.Int {
	return c.chainID
}
func (c *Client) Close() {
	for _, client := range c.clients {
		client.Close()
	}
}
func (c *Client) TransactOpts(ctx context.Context, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, c.chainID)
	if err != nil {
		return nil, fmt.Errorf("blockchain: create transactor: %w", err)
	}
	opts.Context = ctx
	return opts, nil
}
func ParsePrivateKey(hexKey string) (*ecdsa.PrivateKey, error) {
	if len(hexKey) > 2 && hexKey[:2] == "0x" {
		hexKey = hexKey[2:]
	}
	pk, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return nil, fmt.Errorf("blockchain: parse private key: %w", err)
	}
	return pk, nil
}
func AddressFromKey(pk *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(pk.PublicKey)
}
func (c *Client) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	var err error
	for i, client := range c.clients {
		var res []byte
		res, err = client.CodeAt(ctx, contract, blockNumber)
		if err == nil {
			return res, nil
		}
		c.log.Warn("CodeAt failed, falling back", zap.Int("node_idx", i), zap.Error(err))
	}
	return nil, err
}
func (c *Client) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var err error
	for i, client := range c.clients {
		var res []byte
		res, err = client.CallContract(ctx, call, blockNumber)
		if err == nil {
			return res, nil
		}
		c.log.Warn("CallContract failed, falling back", zap.Int("node_idx", i), zap.Error(err))
	}
	return nil, err
}
func (c *Client) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	var err error
	for i, client := range c.clients {
		var res []byte
		res, err = client.PendingCodeAt(ctx, account)
		if err == nil {
			return res, nil
		}
		c.log.Warn("PendingCodeAt failed, falling back", zap.Int("node_idx", i), zap.Error(err))
	}
	return nil, err
}
func (c *Client) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	var err error
	for i, client := range c.clients {
		var res uint64
		res, err = client.PendingNonceAt(ctx, account)
		if err == nil {
			return res, nil
		}
		c.log.Warn("PendingNonceAt failed, falling back", zap.Int("node_idx", i), zap.Error(err))
	}
	return 0, err
}
func (c *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	var err error
	for i, client := range c.clients {
		var res *big.Int
		res, err = client.SuggestGasPrice(ctx)
		if err == nil {
			return res, nil
		}
		c.log.Warn("SuggestGasPrice failed, falling back", zap.Int("node_idx", i), zap.Error(err))
	}
	return nil, err
}
func (c *Client) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	var err error
	for i, client := range c.clients {
		var res *big.Int
		res, err = client.SuggestGasTipCap(ctx)
		if err == nil {
			return res, nil
		}
		c.log.Warn("SuggestGasTipCap failed, falling back", zap.Int("node_idx", i), zap.Error(err))
	}
	return nil, err
}
func (c *Client) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	var err error
	for i, client := range c.clients {
		var res uint64
		res, err = client.EstimateGas(ctx, call)
		if err == nil {
			return res, nil
		}
		c.log.Warn("EstimateGas failed, falling back", zap.Int("node_idx", i), zap.Error(err))
	}
	return 0, err
}
func (c *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	var err error
	for i, client := range c.clients {
		err = client.SendTransaction(ctx, tx)
		if err == nil {
			return nil
		}
		c.log.Warn("SendTransaction failed, falling back", zap.Int("node_idx", i), zap.Error(err))
	}
	return err
}
func (c *Client) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	var err error
	for i, client := range c.clients {
		var res []types.Log
		res, err = client.FilterLogs(ctx, query)
		if err == nil {
			return res, nil
		}
		c.log.Warn("FilterLogs failed, falling back", zap.Int("node_idx", i), zap.Error(err))
	}
	return nil, err
}
func (c *Client) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	var err error
	for i, client := range c.clients {
		var res ethereum.Subscription
		res, err = client.SubscribeFilterLogs(ctx, query, ch)
		if err == nil {
			return res, nil
		}
		c.log.Warn("SubscribeFilterLogs failed, falling back", zap.Int("node_idx", i), zap.Error(err))
	}
	return nil, err
}
func (c *Client) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	var err error
	for i, client := range c.clients {
		var res *types.Header
		res, err = client.HeaderByNumber(ctx, number)
		if err == nil {
			return res, nil
		}
		c.log.Warn("HeaderByNumber failed, falling back", zap.Int("node_idx", i), zap.Error(err))
	}
	return nil, err
}
func (c *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	var err error
	for i, client := range c.clients {
		var res *types.Receipt
		res, err = client.TransactionReceipt(ctx, txHash)
		if err == nil {
			return res, nil
		}
		if err == ethereum.NotFound {
			return nil, err
		}
		c.log.Warn("TransactionReceipt failed, falling back", zap.Int("node_idx", i), zap.Error(err))
	}
	return nil, err
}
func (c *Client) BlockNumber(ctx context.Context) (uint64, error) {
	var err error
	for i, client := range c.clients {
		var res uint64
		res, err = client.BlockNumber(ctx)
		if err == nil {
			return res, nil
		}
		c.log.Warn("BlockNumber failed, falling back", zap.Int("node_idx", i), zap.Error(err))
	}
	return 0, err
}
