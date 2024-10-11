package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// decodeABI decodes ABI-encoded data into a list of Go types.
func DecodeABI(types []string, value []byte) ([]interface{}, error) {
	arguments := make(abi.Arguments, len(types))
	for i, t := range types {
		parsedType, err := abi.NewType(t, "", nil)
		if err != nil {
			return nil, err
		}
		arguments[i] = abi.Argument{Type: parsedType}
	}

	// Decode the value into a slice of interfaces
	decodedData, err := arguments.Unpack(value)
	if err != nil {
		return nil, err
	}
	return decodedData, nil
}

// encodeABI encodes a list of Go types into ABI-encoded data.
func EncodeABI(types []string, values []interface{}) ([]byte, error) {
	arguments := make(abi.Arguments, len(types))
	for i, t := range types {
		parsedType, err := abi.NewType(t, "", nil)
		if err != nil {
			return nil, err
		}
		arguments[i] = abi.Argument{Type: parsedType}
	}

	// Pack the values into ABI-encoded bytes
	encodedData, err := arguments.Pack(values...)
	if err != nil {
		return nil, err
	}
	return encodedData, nil
}
