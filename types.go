package userop

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// IUserOperation represents an ERC-4337 User Operation.
type IUserOperation struct {
	Sender               common.Address
	Nonce                *big.Int
	InitCode             string
	CallData             string
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     string
	Signature            string
}

// NewDefaultUserOperation creates a default IUserOperation.
func NewDefaultUserOperation() *IUserOperation {
	return &IUserOperation{
		Sender:               common.Address{}, // AddressZero
		Nonce:                big.NewInt(0),
		InitCode:             "0x",
		CallData:             "0x",
		CallGasLimit:         big.NewInt(35000),
		VerificationGasLimit: big.NewInt(70000),
		PreVerificationGas:   big.NewInt(21000),
		MaxFeePerGas:         big.NewInt(0),
		MaxPriorityFeePerGas: big.NewInt(0),
		PaymasterAndData:     "0x",
		Signature:            "0x",
	}
}

// ToJSON converts the IUserOperation to a JSON-like map.
func (op *IUserOperation) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"sender":               op.Sender.Hex(),
		"nonce":                "0x" + op.Nonce.Text(16),
		"initCode":             op.InitCode,
		"callData":             op.CallData,
		"callGasLimit":         "0x" + op.CallGasLimit.Text(16),
		"verificationGasLimit": "0x" + op.VerificationGasLimit.Text(16),
		"preVerificationGas":   "0x" + op.PreVerificationGas.Text(16),
		"maxFeePerGas":         "0x" + op.MaxFeePerGas.Text(16),
		"maxPriorityFeePerGas": "0x" + op.MaxPriorityFeePerGas.Text(16),
		"paymasterAndData":     op.PaymasterAndData,
		"signature":            op.Signature,
	}
}

// IUserOperationBuilder provides a flexible way to construct an IUserOperation.
type IUserOperationBuilder interface {
	GetSender() common.Address
	GetNonce() *big.Int
	GetInitCode() string
	GetCallData() string
	GetCallGasLimit() *big.Int
	GetVerificationGasLimit() *big.Int
	GetPreVerificationGas() *big.Int
	GetMaxFeePerGas() *big.Int
	GetMaxPriorityFeePerGas() *big.Int
	GetPaymasterAndData() string
	GetSignature() string
	GetOp() *IUserOperation
	SetSender(address common.Address) IUserOperationBuilder
	SetNonce(nonce *big.Int) IUserOperationBuilder
	SetInitCode(code string) IUserOperationBuilder
	SetCallData(data string) IUserOperationBuilder
	SetCallGasLimit(gas *big.Int) IUserOperationBuilder
	SetVerificationGasLimit(gas *big.Int) IUserOperationBuilder
	SetPreVerificationGas(gas *big.Int) IUserOperationBuilder
	SetMaxFeePerGas(fee *big.Int) IUserOperationBuilder
	SetMaxPriorityFeePerGas(fee *big.Int) IUserOperationBuilder
	SetPaymasterAndData(data string) IUserOperationBuilder
	SetSignature(bytes string) IUserOperationBuilder
	SetPartial(partialOp map[string]interface{}) IUserOperationBuilder
	UseDefaults(partialOp map[string]interface{}) IUserOperationBuilder
	ResetDefaults() IUserOperationBuilder
	UseMiddleware(fn UserOperationMiddlewareFn) IUserOperationBuilder
	ResetMiddleware() IUserOperationBuilder
	BuildOp(entryPoint common.Address, chainID *big.Int) (*IUserOperation, error)
	ResetOp() IUserOperationBuilder
}

// UserOperationMiddlewareFn is a function type for middleware in user operations.
type UserOperationMiddlewareFn func(ctx *IUserOperationMiddlewareCtx) error

// IUserOperationMiddlewareCtx provides context for middleware functions.
type IUserOperationMiddlewareCtx struct {
	Op         *IUserOperation
	EntryPoint common.Address
	ChainID    *big.Int
}

// GetUserOpHash returns the hash of the user operation.
func (ctx *IUserOperationMiddlewareCtx) GetUserOpHash() []byte {
	// Implement hashing logic based on the user operation structure
	// Example: Keccak256 of the serialized operation
	hashInput := ctx.Op.ToJSON() // Serialize this correctly for hashing
	hashInputBytes, err := json.Marshal(hashInput)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	hash := crypto.Keccak256Hash(hashInputBytes)
	return hash.Bytes()
}

// IClient represents an interface for the client class.
type IClient interface {
	SendUserOperation(builder IUserOperationBuilder, opts *ISendUserOperationOpts) (*ISendUserOperationResponse, error)
	BuildUserOperation(builder IUserOperationBuilder) (*IUserOperation, error)
}

// IClientOpts contains options for the client.
type IClientOpts struct {
	EntryPoint         common.Address
	OverrideBundlerRpc string
	SocketConnector    func() StreamChannel
}

// ISendUserOperationOpts contains options for sending user operations.
type ISendUserOperationOpts struct {
	DryRun  bool
	OnBuild func(op *IUserOperation)
}

// ISendUserOperationResponse represents the response for sendUserOperation.
type ISendUserOperationResponse struct {
	UserOpHash string
	Wait       func() (*FilterEvent, error)
}

// IPresetBuilderOpts contains options for the preset builder.
type IPresetBuilderOpts struct {
	EntryPoint          common.Address
	Salt                *big.Int
	FactoryAddress      common.Address
	PaymasterMiddleware UserOperationMiddlewareFn
	NonceKey            *big.Int
	OverrideBundlerRpc  string
}

// Call represents a call operation.
type Call struct {
	To    common.Address
	Value *big.Int
	Data  []byte
}

// Example of a StreamChannel interface (define as needed)
type StreamChannel interface {
	// Define necessary methods for the StreamChannel
	Send(msg string)
	Receive() string
}

// FilterEvent represents the structure for filter events (define as needed)
type FilterEvent struct {
	// Add fields as necessary
}
