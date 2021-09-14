// Simple go library and cli to encrypt and decrypt data
//
// This packages uses the Galois/Counter_Mode
// https://en.wikipedia.org/wiki/Galois/Counter_Mode and the package
// https://pkg.go.dev/crypto/cipher#NewGCM
package crypto

import (
	"fmt"
	"io"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"

	"encoding/hex"
)

// Encrypts a payload with a password
func Encrypt(payload string, password string) (string, error) {
	plain := []byte(payload)
	data := []byte(password)
	key := sha256.Sum256(data)

	// Create a new cipher block from the key
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt the data using aesGCM.Seal. Since we don't want to save the nonce
	// somewhere else in this case, we add it as a prefix to the encrypted data.
	// The first nonce argument in Seal is the prefix.
	ciphered := gcm.Seal(nonce, nonce, plain, nil)
	return fmt.Sprintf("%x", ciphered), nil
}

// Decrypts an encrypted payload with password
func Decrypt(payload string, password string) (string, error) {
	data := []byte(password)
	key := sha256.Sum256(data)
	enc, _ := hex.DecodeString(payload)

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphered := enc[:nonceSize], enc[nonceSize:]

	plain, err := gcm.Open(nil, nonce, ciphered, nil)
	if err != nil {
		return "", err
	}

	return string(plain), nil
}
