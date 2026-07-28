package crypto

import "golang.org/x/crypto/argon2"

const (
	Memory     = 64 * 1024
	Iterations = 3
	Threads    = 4
	KeyLen     = 32
)

func DeriveKey(password string, salt [16]byte) [32]byte {
	passwordBytes := []byte(password)

	cryptoKey := argon2.IDKey(passwordBytes, salt[:], Iterations, Memory, Threads, KeyLen)
	return [KeyLen]byte(cryptoKey)
}
