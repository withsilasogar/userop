package typechain

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Define the ABI for the ECDSAKernelFactory contract
const ecdsaKernelFactoryABI = `[{"inputs":[{"internalType":"contract KernelFactory","name":"_singletonFactory","type":"address"},{"internalType":"contract ECDSAValidator","name":"_validator","type":"address"},{"internalType":"contract IEntryPoint","name":"_entryPoint","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"address","name":"_owner","type":"address"},{"internalType":"uint256","name":"_index","type":"uint256"}],"name":"createAccount","outputs":[{"internalType":"contract EIP1967Proxy","name":"proxy","type":"address"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"entryPoint","outputs":[{"internalType":"contract IEntryPoint","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"_owner","type":"address"},{"internalType":"uint256","name":"_index","type":"uint256"}],"name":"getAccountAddress","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"singletonFactory","outputs":[{"internalType":"contract KernelFactory","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"validator","outputs":[{"internalType":"contract ECDSAValidator","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`

// ECDSAKernelFactory represents the contract
type ECDSAKernelFactory struct {
	client   *ethclient.Client
	contract *abi.ABI
	address  common.Address
}

// NewECDSAKernelFactory creates a new instance of ECDSAKernelFactory
func NewECDSAKernelFactory(client *ethclient.Client, address string) (*ECDSAKernelFactory, error) {
	parsedABI, err := abi.JSON(strings.NewReader(ecdsaKernelFactoryABI))
	if err != nil {
		return nil, err
	}
	return &ECDSAKernelFactory{
		client:   client,
		contract: &parsedABI,
		address:  common.HexToAddress(address),
	}, nil
}

// CreateAccount calls the createAccount function on the contract
func (f *ECDSAKernelFactory) CreateAccount(owner common.Address, index *big.Int, opts *bind.TransactOpts) (common.Address, error) {
	callData, err := f.contract.Pack("createAccount", owner, index)
	if err != nil {
		return common.Address{}, err
	}

	tx := types.NewTransaction(opts.Nonce.Uint64(), f.address, big.NewInt(0), opts.GasLimit, opts.GasPrice, callData)
	signedTx, err := opts.Signer(opts.From, tx)
	if err != nil {
		return common.Address{}, err
	}

	if err := f.client.SendTransaction(context.Background(), signedTx); err != nil {
		return common.Address{}, err
	}

	receipt, err := bind.WaitMined(context.Background(), f.client, tx)
	if err != nil {
		return common.Address{}, err
	}

	if receipt.Status != 1 {
		return common.Address{}, fmt.Errorf("transaction failed")
	}

	var result struct {
		Proxy common.Address
	}
	err = f.contract.UnpackIntoInterface(&result, "createAccount", receipt.Logs[0].Data)
	if err != nil {
		return common.Address{}, err
	}

	return result.Proxy, nil
}

// EntryPoint returns the entry point address
func (f *ECDSAKernelFactory) EntryPoint() (common.Address, error) {
	var result []interface{}
	msg := ethereum.CallMsg{
		To:   &f.address,
		Data: f.contract.Methods["entryPoint"].ID,
	}
	output, err := f.client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return common.Address{}, err
	}
	err = f.contract.UnpackIntoInterface(&result, "entryPoint", output)
	if err != nil {
		return common.Address{}, err
	}
	return result[0].(common.Address), nil
}

// GetAccountAddress returns the account address based on owner and index
func (f *ECDSAKernelFactory) GetAccountAddress(owner common.Address, index *big.Int) (common.Address, error) {
	var result []interface{}
	msg := ethereum.CallMsg{
		To:   &f.address,
		Data: f.contract.Methods["getAccountAddress"].ID,
	}
	output, err := f.client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return common.Address{}, err
	}
	err = f.contract.UnpackIntoInterface(&result, "getAccountAddress", output)
	if err != nil {
		return common.Address{}, err
	}
	return result[0].(common.Address), nil
}

// SingletonFactory returns the singleton factory address
func (f *ECDSAKernelFactory) SingletonFactory() (common.Address, error) {
	var result []interface{}
	msg := ethereum.CallMsg{
		To:   &f.address,
		Data: f.contract.Methods["singletonFactory"].ID,
	}
	output, err := f.client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return common.Address{}, err
	}
	err = f.contract.UnpackIntoInterface(&result, "singletonFactory", output)
	if err != nil {
		return common.Address{}, err
	}
	return result[0].(common.Address), nil
}

// Validator returns the validator address
func (f *ECDSAKernelFactory) Validator() (common.Address, error) {
	var result []interface{}
	msg := ethereum.CallMsg{
		To:   &f.address,
		Data: f.contract.Methods["validator"].ID,
	}
	output, err := f.client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return common.Address{}, err
	}
	err = f.contract.UnpackIntoInterface(&result, "validator", output)
	if err != nil {
		return common.Address{}, err
	}
	return result[0].(common.Address), nil
}
