package container

const Magic = "EPST"

type Header struct {
	Version byte

	Salt  [16]byte
	Nonce [24]byte

	Filename string
}
