package service

import (
	"io"
	"os"
	"path/filepath"

	"github.com/kyotomin/epstein/internal/container"
	"github.com/kyotomin/epstein/internal/crypto"
)

// v1
func EncryptFile(srcPath string, password string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// fi, err := srcFile.Stat()
	// if err != nil {
	// 	return err
	// }

	// totalSize := fi.Size()

	ext := filepath.Ext(srcPath)
	dstPath := srcPath[:len(srcPath)-len(ext)] + ".epst"
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	plaintext, err := io.ReadAll(srcFile)
	if err != nil {
		return err
	}

	salt, err := generateSalt()
	if err != nil {
		return err
	}

	nonce, err := generateNonce()
	if err != nil {
		return err
	}
	key := crypto.DeriveKey(password, salt)

	header := container.Header{
		Version:  1,
		Salt:     salt,
		Nonce:    nonce,
		Filename: filepath.Base(srcPath),
	}

	if err := container.WriteHeader(dstFile, header); err != nil {
		return err
	}

	cipherText, err := crypto.Encrypt(plaintext, key, nonce)
	if err != nil {
		return err
	}

	if _, err := dstFile.Write(cipherText); err != nil {
		return err
	}

	return nil
}
