package models

import (
	"encoding/json"
)

// VerifyingPaymasterResult represents the result of a verifying paymaster operation
type VerifyingPaymasterResult struct {
	PaymasterAndData     string `json:"paymasterAndData"`
	PreVerificationGas   string `json:"preVerificationGas"`
	VerificationGasLimit string `json:"verificationGasLimit"`
	CallGasLimit         string `json:"callGasLimit"`
}

// NewVerifyingPaymasterResult creates a new instance of VerifyingPaymasterResult
func NewVerifyingPaymasterResult(paymasterAndData, preVerificationGas, verificationGasLimit, callGasLimit string) *VerifyingPaymasterResult {
	return &VerifyingPaymasterResult{
		PaymasterAndData:     paymasterAndData,
		PreVerificationGas:   preVerificationGas,
		VerificationGasLimit: verificationGasLimit,
		CallGasLimit:         callGasLimit,
	}
}

// FromJSON unmarshals a JSON object into a VerifyingPaymasterResult instance
func (v *VerifyingPaymasterResult) FromJSON(data []byte) (*VerifyingPaymasterResult, error) {
	var result VerifyingPaymasterResult
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ToJSON marshals a VerifyingPaymasterResult instance into JSON
func (v *VerifyingPaymasterResult) ToJSON() ([]byte, error) {
	return json.Marshal(v)
}
