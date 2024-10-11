package constants

type KernelModes struct{}

const (
	SUDO   = "0x00000000"
	PLUGIN = "0x00000001"
	ENABLE = "0x00000002"
)

func NewKernelModes() *KernelModes {
	return &KernelModes{}
}
