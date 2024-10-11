package utils

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestEncodeABI(t *testing.T) {
	// Example types and values for encoding
	types := []string{"address", "uint256"}
	values := []interface{}{common.HexToAddress("0x000000000000000000000000000000000000dead"), big.NewInt(1000)}

	encoded, err := EncodeABI(types, values)
	assert.NoError(t, err)
	assert.NotNil(t, encoded)

	// Expected encoding result length should be greater than 0
	assert.True(t, len(encoded) > 0, "Encoded ABI data should not be empty")
}

func TestDecodeABI(t *testing.T) {
	// Example types and values for decoding
	types := []string{"address", "uint256"}
	values := []interface{}{common.HexToAddress("0x000000000000000000000000000000000000dead"), big.NewInt(1000)}

	// First, we encode the values using EncodeABI
	encoded, err := EncodeABI(types, values)
	assert.NoError(t, err)
	assert.NotNil(t, encoded)

	// Now, we attempt to decode it
	decoded, err := DecodeABI(types, encoded)
	assert.NoError(t, err)
	assert.NotNil(t, decoded)

	// Check that decoded values match original input values
	assert.Equal(t, values[0].(common.Address).Hex(), decoded[0].(common.Address).Hex(), "Decoded address should match original")
	assert.Equal(t, values[1].(*big.Int).Cmp(decoded[1].(*big.Int)), 0, "Decoded uint256 should match original")
}

func TestEncodeDecodeABI(t *testing.T) {
	// Example types and values for encoding and decoding
	types := []string{"address", "uint256"}
	values := []interface{}{common.HexToAddress("0x000000000000000000000000000000000000dead"), big.NewInt(1000)}

	// Encode the values
	encoded, err := EncodeABI(types, values)
	assert.NoError(t, err)
	assert.NotNil(t, encoded)

	// Decode the encoded values
	decoded, err := DecodeABI(types, encoded)
	assert.NoError(t, err)
	assert.NotNil(t, decoded)

	// Verify that the decoded values match the original ones
	assert.Equal(t, values[0].(common.Address), decoded[0].(common.Address))
	assert.Equal(t, values[1].(*big.Int).Cmp(decoded[1].(*big.Int)), 0)
}
