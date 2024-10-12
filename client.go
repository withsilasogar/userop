package userop

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/withsilasogar/userop/constants"
)

// Client for interacting with an ERC-4337 bundler.
type Client struct {
	web3Client   *rpc.Client
	chainId      *big.Int
	entryPoint   common.Address
	waitTimeout  time.Duration
	waitInterval time.Duration
}

// NewClient initializes a new Client.
func NewClient(rpcUrl string, opts *IClientOpts) (*Client, error) {
	client, err := rpc.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return nil, err
	}

	entryPoint := common.HexToAddress(constants.ENTRY_POINT)

	return &Client{
		web3Client:   client,
		entryPoint:   entryPoint,
		waitTimeout:  30 * time.Second,
		waitInterval: 5 * time.Second,
	}, nil
}

// Init initializes the client and fetches the chain ID.
func Init(rpcUrl string, opts *IClientOpts) (*Client, error) {
	client, err := NewClient(rpcUrl, opts)
	if err != nil {
		return nil, err
	}

	var chainId *big.Int
	err = client.web3Client.CallContext(context.Background(), &chainId, "eth_chainId")
	if err != nil {
		return nil, err
	}
	client.chainId = chainId

	return client, nil
}

// BuildUserOperation builds a user operation using the provided builder.
func (c *Client) BuildUserOperation(builder IUserOperationBuilder) (*IUserOperation, error) {
	return builder.BuildOp(c.entryPoint, c.chainId)
}
