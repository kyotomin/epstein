package service

import "crypto/rand"

func generateSalt() ([16]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return [16]byte{}, err
	}

	return [16]byte(salt), nil
}
func generateNonce() ([24]byte, error) {
	nonce := make([]byte, 24)
	if _, err := rand.Read(nonce); err != nil {
		return [24]byte{}, err
	}

	return [24]byte(nonce), nil
}
