package models

import (
	"encoding/json"
)

// GasEstimate represents the gas estimate for transactions
type GasEstimate struct {
	VerificationGasLimit *string `json:"verificationGasLimit"`
	PreVerificationGas   string  `json:"preVerificationGas"`
	CallGasLimit         string  `json:"callGasLimit"`
	VerificationGas      string  `json:"verificationGas"`
}

// NewGasEstimate creates a new instance of GasEstimate
func NewGasEstimate(verificationGasLimit *string, preVerificationGas, callGasLimit, verificationGas string) *GasEstimate {
	return &GasEstimate{
		VerificationGasLimit: verificationGasLimit,
		PreVerificationGas:   preVerificationGas,
		CallGasLimit:         callGasLimit,
		VerificationGas:      verificationGas,
	}
}

// FromJSON unmarshals a JSON object into a GasEstimate instance
func (g *GasEstimate) FromJSON(data []byte) (*GasEstimate, error) {
	var gasEstimate GasEstimate
	err := json.Unmarshal(data, &gasEstimate)
	if err != nil {
		return nil, err
	}
	return &gasEstimate, nil
}

// ToJSON marshals a GasEstimate instance into JSON
func (g *GasEstimate) ToJSON() ([]byte, error) {
	return json.Marshal(g)
}
