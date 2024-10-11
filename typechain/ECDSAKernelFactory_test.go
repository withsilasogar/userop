package typechain

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	// Adjust the import path to your project
)

// MockEthClient is a mock implementation of ethclient.Client
type MockEthClient struct {
	mock.Mock
	ethclient.Client
}

func (m *MockEthClient) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	args := m.Called(ctx, call, blockNumber)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockEthClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	args := m.Called(ctx, tx)
	return args.Error(0)
}

func (m *MockEthClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return nil, nil
}

func TestNewECDSAKernelFactory(t *testing.T) {
	mockClient := new(MockEthClient)
	address := "0xYourContractAddress"

	factory, err := NewECDSAKernelFactory(&mockClient.Client, address)
	assert.NoError(t, err)
	assert.NotNil(t, factory)
	assert.Equal(t, common.HexToAddress(address), factory.address)
}

func TestECDSAKernelFactory_CreateAccount(t *testing.T) {
	mockClient := new(MockEthClient)
	factory, err := NewECDSAKernelFactory(&mockClient.Client, "0xYourContractAddress")
	assert.NoError(t, err)

	owner := common.HexToAddress("0xOwnerAddress")
	index := big.NewInt(1)
	opts := &bind.TransactOpts{
		From:     common.HexToAddress("0xFromAddress"),
		GasPrice: big.NewInt(1000),
		GasLimit: uint64(3000000),
		Nonce:    big.NewInt(0),
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
	}

	mockClient.On("SendTransaction", mock.Anything, mock.Anything).Return(nil)

	address, err := factory.CreateAccount(owner, index, opts)
	assert.NoError(t, err)
	assert.NotEmpty(t, address)
	mockClient.AssertExpectations(t)
}

func TestECDSAKernelFactory_EntryPoint(t *testing.T) {
	mockClient := new(MockEthClient)
	factory, err := NewECDSAKernelFactory(&mockClient.Client, "0xYourContractAddress")
	assert.NoError(t, err)

	expectedEntryPoint := common.HexToAddress("0xExpectedEntryPoint")
	mockClient.On("CallContract", mock.Anything, mock.Anything, mock.Anything).
		Return([]byte(expectedEntryPoint.Bytes()), nil)

	entryPoint, err := factory.EntryPoint()
	assert.NoError(t, err)
	assert.Equal(t, expectedEntryPoint, entryPoint)
	mockClient.AssertExpectations(t)
}

func TestECDSAKernelFactory_GetAccountAddress(t *testing.T) {
	mockClient := new(MockEthClient)
	factory, err := NewECDSAKernelFactory(&mockClient.Client, "0xYourContractAddress")
	assert.NoError(t, err)

	owner := common.HexToAddress("0xOwnerAddress")
	index := big.NewInt(1)
	expectedAccountAddress := common.HexToAddress("0xExpectedAccountAddress")
	mockClient.On("CallContract", mock.Anything, mock.Anything, mock.Anything).
		Return([]byte(expectedAccountAddress.Bytes()), nil)

	accountAddress, err := factory.GetAccountAddress(owner, index)
	assert.NoError(t, err)
	assert.Equal(t, expectedAccountAddress, accountAddress)
	mockClient.AssertExpectations(t)
}

func TestECDSAKernelFactory_SingletonFactory(t *testing.T) {
	mockClient := new(MockEthClient)
	factory, err := NewECDSAKernelFactory(&mockClient.Client, "0xYourContractAddress")
	assert.NoError(t, err)

	expectedSingletonFactory := common.HexToAddress("0xExpectedSingletonFactory")
	mockClient.On("CallContract", mock.Anything, mock.Anything, mock.Anything).
		Return([]byte(expectedSingletonFactory.Bytes()), nil)

	singletonFactory, err := factory.SingletonFactory()
	assert.NoError(t, err)
	assert.Equal(t, expectedSingletonFactory, singletonFactory)
	mockClient.AssertExpectations(t)
}

func TestECDSAKernelFactory_Validator(t *testing.T) {
	mockClient := new(MockEthClient)
	factory, err := NewECDSAKernelFactory(&mockClient.Client, "0xYourContractAddress")
	assert.NoError(t, err)

	expectedValidator := common.HexToAddress("0xExpectedValidator")
	mockClient.On("CallContract", mock.Anything, mock.Anything, mock.Anything).
		Return([]byte(expectedValidator.Bytes()), nil)

	validator, err := factory.Validator()
	assert.NoError(t, err)
	assert.Equal(t, expectedValidator, validator)
	mockClient.AssertExpectations(t)
}
