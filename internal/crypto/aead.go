package crypto

import (
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
)

func Encrypt(plaintext []byte, key [32]byte, nonce [24]byte) ([]byte, error) {
	aead, err := chacha20poly1305.NewX(key[:])
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher text: %w", err)
	}

	ciphertext := aead.Seal(nil, nonce[:], plaintext, nil)
	return ciphertext, nil
}

func Decrypt(ciphertext []byte, key [32]byte, nonce [24]byte) ([]byte, error) {
	aead, err := chacha20poly1305.NewX(key[:])
	if err != nil {
		return nil, fmt.Errorf("failed to create aead: %w", err)
	}

	plaintext, err := aead.Open(nil, nonce[:], ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt ciphertext: %w", err)
	}

	return plaintext, nil
}
