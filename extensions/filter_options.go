package extensions

import (
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// UserOperationEventFilter defines the filter options for the UserOperation event
type UserOperationEventFilter struct {
	Contract  common.Address  // The contract address to filter
	Event     string          // The event signature/topic
	FromBlock *big.Int        // Optional start block to filter from
	ToBlock   *big.Int        // Optional end block to filter to
	Topics    [][]common.Hash // Filter by topics, i.e., event arguments
}

// NewUserOperationEventFilter creates a new filter for UserOperation events
func NewUserOperationEventFilter(contract common.Address, event string, fromBlock, toBlock *big.Int, userOpHash string) *UserOperationEventFilter {
	// Initialize the filter with the provided parameters
	filter := &UserOperationEventFilter{
		Contract:  contract,
		Event:     event,
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Topics:    [][]common.Hash{},
	}

	// Add the user operation hash to the topics if it's not empty
	if userOpHash != "" {
		filter.Topics = append(filter.Topics, []common.Hash{common.HexToHash(userOpHash)})
	}

	return filter
}

// ToFilterQuery converts UserOperationEventFilter to the ethereum FilterQuery object
func (f *UserOperationEventFilter) ToFilterQuery() ethereum.FilterQuery {
	return ethereum.FilterQuery{
		Addresses: []common.Address{f.Contract},
		Topics:    f.Topics,
		FromBlock: f.FromBlock,
		ToBlock:   f.ToBlock,
	}
}
