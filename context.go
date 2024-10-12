package userop

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// UserOperationMiddlewareCtx implements IUserOperationMiddlewareCtx
type UserOperationMiddlewareCtx struct {
	Op         *IUserOperation
	EntryPoint common.Address
	ChainId    *big.Int
}

// NewUserOperationMiddlewareCtx initializes a new UserOperationMiddlewareCtx
func NewUserOperationMiddlewareCtx(op *IUserOperation, entryPoint common.Address, chainId *big.Int) *UserOperationMiddlewareCtx {
	return &UserOperationMiddlewareCtx{
		Op:         op,
		EntryPoint: entryPoint,
		ChainId:    chainId,
	}
}
