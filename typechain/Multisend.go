package typechain

import (
	"context"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var contractAbiJSON = `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"bytes","name":"transactions","type":"bytes"}],"name":"multiSend","outputs":[],"stateMutability":"payable","type":"function"}]`

type Multisend struct {
	Address      common.Address
	Abi          abi.ABI
	Client       *ethclient.Client
	ChainID      *big.Int
	TransactOpts *bind.TransactOpts
}

func NewMultisend(address common.Address, client *ethclient.Client, chainID *big.Int) (*Multisend, error) {
	parsedAbi, err := abi.JSON(strings.NewReader(contractAbiJSON))
	if err != nil {
		return nil, err
	}

	return &Multisend{
		Address: address,
		Abi:     parsedAbi,
		Client:  client,
		ChainID: chainID,
	}, nil
}

func (m *Multisend) MultiSend(transactions []byte, privateKey string) (string, error) {
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}

	fromAddress := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)
	nonce, err := m.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := m.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, m.ChainID)
	if err != nil {
		return "", err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // set the value in wei
	auth.GasLimit = uint64(300000) // set the gas limit
	auth.GasPrice = gasPrice

	// Call the multiSend function
	tx, err := m.Abi.Pack("multiSend", transactions)
	if err != nil {
		return "", err
	}

	txObj, err := bind.NewBoundContract(m.Address, m.Abi, m.Client, m.Client, m.Client).Transact(auth, "multiSend", transactions)
	if err != nil {
		return "", err
	}

	txHash := txObj.Hash().Hex()
	log.Printf("Transaction hash: %s", txHash)

	return txHash, nil
}
