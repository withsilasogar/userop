package utils

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ContractsHelper provides methods to interact with Ethereum smart contracts.
type ContractsHelper struct {
	client *ethclient.Client
}

// NewContractsHelper creates a new ContractsHelper instance.
func NewContractsHelper(client *ethclient.Client) *ContractsHelper {
	return &ContractsHelper{client: client}
}

// ReadFromContract reads data from a deployed contract using the specified function and parameters.
func (ch *ContractsHelper) ReadFromContract(contractName, contractAddress, functionName string, params []interface{}, abiJSON string) ([]interface{}, error) {
	contractABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %v", err)
	}

	contractAddr := common.HexToAddress(contractAddress)

	// Prepare the call data
	data, err := contractABI.Pack(functionName, params...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack function call: %v", err)
	}

	// Execute the call
	callMsg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: data,
	}

	output, err := ch.client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %v", err)
	}

	// Unpack the returned data
	results, err := contractABI.Unpack(functionName, output)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack output: %v", err)
	}

	return results, nil
}

// EncodedDataForContractCall encodes data for a contract call using the specified function and parameters.
func (ch *ContractsHelper) EncodedDataForContractCall(contractName, contractAddress, functionName string, params []interface{}, abiJSON string) ([]byte, error) {
	contractABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %v", err)
	}

	// Encode the function call with the parameters
	data, err := contractABI.Pack(functionName, params...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack function call: %v", err)
	}

	return data, nil
}
