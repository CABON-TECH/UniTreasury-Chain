// Package blockchain provides the Ethereum client, transaction manager,
// and contract binding wrappers used by service layer.
package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

// Client wraps an ethclient.Client with convenience helpers.
type Client struct {
	inner   *ethclient.Client
	chainID *big.Int
	log     *zap.Logger
}

// NewClient dials an Ethereum node and returns a Client.
func NewClient(ctx context.Context, rpcURL string, chainID int64, log *zap.Logger) (*Client, error) {
	ec, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, fmt.Errorf("blockchain: dial %s: %w", rpcURL, err)
	}

	// Verify chain ID matches expectation
	gotChainID, err := ec.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("blockchain: get chain ID: %w", err)
	}
	if gotChainID.Int64() != chainID {
		return nil, fmt.Errorf("blockchain: chain ID mismatch: expected %d, got %d", chainID, gotChainID)
	}

	log.Info("connected to Ethereum node",
		zap.String("rpc_url", rpcURL),
		zap.Int64("chain_id", gotChainID.Int64()),
	)

	return &Client{
		inner:   ec,
		chainID: gotChainID,
		log:     log,
	}, nil
}

// Inner returns the underlying ethclient for packages that need direct access.
func (c *Client) Inner() *ethclient.Client {
	return c.inner
}

// ChainID returns the connected chain ID.
func (c *Client) ChainID() *big.Int {
	return c.chainID
}

// Close terminates the underlying connection.
func (c *Client) Close() {
	c.inner.Close()
}

// TransactOpts builds a bind.TransactOpts for signing transactions with a given
// private key. Gas is estimated automatically by go-ethereum.
func (c *Client) TransactOpts(ctx context.Context, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, c.chainID)
	if err != nil {
		return nil, fmt.Errorf("blockchain: create transactor: %w", err)
	}
	opts.Context = ctx
	return opts, nil
}

// CurrentBlock returns the latest block number.
func (c *Client) CurrentBlock(ctx context.Context) (uint64, error) {
	return c.inner.BlockNumber(ctx)
}

// ParsePrivateKey decodes a hex-encoded private key (with or without 0x prefix).
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

// AddressFromKey derives the Ethereum address from a private key.
func AddressFromKey(pk *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(pk.PublicKey)
}
