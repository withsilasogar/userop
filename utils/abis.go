package utils

import (
	"encoding/json"
	"fmt"
)

type ABI struct{}

// Get returns the ABI JSON string based on the contract name
func (a *ABI) Get(name string) (string, error) {
	var abi []map[string]interface{}

	switch name {
	case "ERC20":
		abi = []map[string]interface{}{
			{"constant": true, "inputs": []interface{}{}, "name": "name", "outputs": []interface{}{map[string]interface{}{"name": "", "type": "string"}}, "payable": false, "stateMutability": "view", "type": "function"},
			{"constant": false, "inputs": []interface{}{map[string]interface{}{"name": "spender", "type": "address"}, map[string]interface{}{"name": "value", "type": "uint256"}}, "name": "approve", "outputs": []interface{}{map[string]interface{}{"name": "", "type": "bool"}}, "payable": false, "stateMutability": "nonpayable", "type": "function"},
			{"constant": true, "inputs": []interface{}{}, "name": "totalSupply", "outputs": []interface{}{map[string]interface{}{"name": "", "type": "uint256"}}, "payable": false, "stateMutability": "view", "type": "function"},
			{"constant": false, "inputs": []interface{}{map[string]interface{}{"name": "sender", "type": "address"}, map[string]interface{}{"name": "recipient", "type": "address"}, map[string]interface{}{"name": "amount", "type": "uint256"}}, "name": "transferFrom", "outputs": []interface{}{map[string]interface{}{"name": "", "type": "bool"}}, "payable": false, "stateMutability": "nonpayable", "type": "function"},
			{"constant": true, "inputs": []interface{}{}, "name": "decimals", "outputs": []interface{}{map[string]interface{}{"name": "", "type": "uint8"}}, "payable": false, "stateMutability": "view", "type": "function"},
		}
	default:
		return "", fmt.Errorf("ABI does not exist for %s", name)
	}

	// Convert ABI to JSON string
	abiJSON, err := json.Marshal(abi)
	if err != nil {
		return "", err
	}

	return string(abiJSON), nil
}
