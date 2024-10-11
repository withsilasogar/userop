package constants

type Addresses struct{}

const AddressZero = "0x0000000000000000000000000000000000000000"

func NewAddresses() *Addresses {
	return &Addresses{}
}
