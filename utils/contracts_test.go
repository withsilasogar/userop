package utils

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockClient is a mock implementation of the Ethereum client.
type MockClient struct {
	mock.Mock
	ethclient.Client
}

// CallContract is a mocked method to simulate Ethereum contract calls.
func (m *MockClient) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	args := m.Called(ctx, call, blockNumber)
	return args.Get(0).([]byte), args.Error(1)
}

// TestContractsHelper tests the ContractsHelper methods.
func TestContractsHelper(t *testing.T) {
	// Create a mock Ethereum client
	mockClient := new(MockClient)

	// Create an instance of ContractsHelper with the mock client
	contractsHelper := &ContractsHelper{client: &mockClient.Client}

	// Define the contract ABI
	abiJSON := `[{"constant":true,"inputs":[{"name":"account","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`
	functionName := "balanceOf"
	contractAddress := "0xYourContractAddress"

	// Define parameters and expected results
	address := common.HexToAddress("0xYourAddress")
	expectedResult := big.NewInt(100) // Example balance
	encodedResult, _ := json.Marshal(expectedResult)

	// Mock the CallContract method to return the expected result
	mockClient.On("CallContract", mock.Anything, mock.Anything, mock.Anything).Return(encodedResult, nil)

	// Call the ReadFromContract method
	result, err := contractsHelper.ReadFromContract("ERC20", contractAddress, functionName, []interface{}{address}, abiJSON)

	// Assert no error occurred
	assert.NoError(t, err)

	// Assert the result matches the expected result
	assert.Equal(t, expectedResult.String(), result[0].(*big.Int).String())

	// Test encoded data for contract call
	encodedData, err := contractsHelper.EncodedDataForContractCall("ERC20", contractAddress, functionName, []interface{}{address}, abiJSON)

	// Assert no error occurred
	assert.NoError(t, err)

	// Check if encoded data is not empty
	assert.NotEmpty(t, encodedData)

	// Check if the first 4 bytes match the function signature hash
	functionSignature := "0x70a08231" // Keccak256("balanceOf(address)")[0:4]
	assert.Equal(t, functionSignature, "0x"+hex.EncodeToString(encodedData[:4]))
}
