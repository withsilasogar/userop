package userop

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

// BundlerJsonRpcProvider is a wrapper over JsonRPC, specifically for the Bundler RPC.
type BundlerJsonRpcProvider struct {
	*rpc.Client
	bundlerRpc     *rpc.Client
	bundlerMethods map[string]struct{}
}

// NewBundlerJsonRpcProvider creates a new BundlerJsonRpcProvider with the given URL and HTTP client.
func NewBundlerJsonRpcProvider(url string) (*BundlerJsonRpcProvider, error) {
	rpcClient, err := rpc.DialHTTP(url)
	if err != nil {
		return nil, fmt.Errorf("failed to create RPC client: %w", err)
	}

	return &BundlerJsonRpcProvider{
		Client: rpcClient,
		bundlerMethods: map[string]struct{}{
			"eth_sendUserOperation":        {},
			"eth_estimateUserOperationGas": {},
			"eth_getUserOperationByHash":   {},
			"eth_getUserOperationReceipt":  {},
			"eth_supportedEntryPoints":     {},
		},
	}, nil
}

// SetBundlerRpc sets a new RPC client for the bundler.
func (p *BundlerJsonRpcProvider) SetBundlerRpc(bundlerRpcURL string) error {
	if bundlerRpcURL != "" {
		bundlerRpcClient, err := rpc.DialHTTP(bundlerRpcURL)
		if err != nil {
			return fmt.Errorf("failed to create bundler RPC client: %w", err)
		}
		p.bundlerRpc = bundlerRpcClient
	}
	return nil
}

// Call overrides the call method to handle bundler-specific methods.
func (p *BundlerJsonRpcProvider) Call(ctx context.Context, method string, args interface{}, result interface{}) error {
	if _, exists := p.bundlerMethods[method]; exists && p.bundlerRpc != nil {
		return p.bundlerRpc.CallContext(ctx, result, method, args)
	}
	return p.Client.CallContext(ctx, result, method, args)
}
