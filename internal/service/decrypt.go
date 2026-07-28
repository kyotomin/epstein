package service

import (
	"io"
	"os"

	"github.com/kyotomin/epstein/internal/container"
	"github.com/kyotomin/epstein/internal/crypto"
)

// v1
func DecryptFile(srcPath string, password string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	header, err := container.ReadHeader(srcFile)
	if err != nil {
		return err
	}

	salt, nonce, filename := header.Salt, header.Nonce, header.Filename

	key := crypto.DeriveKey(password, salt)

	ciphertext, err := io.ReadAll(srcFile)
	if err != nil {
		return err
	}

	plaintext, err := crypto.Decrypt(ciphertext, key, nonce)
	if err != nil {
		return err
	}

	dstPath := filename
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err := dstFile.Write(plaintext); err != nil {
		return err
	}

	return nil
}
