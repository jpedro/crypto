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
    "crypto/sha256"
    "crypto/rand"

    "encoding/hex"
)

// Encrypts `text` with `password`
func Encrypt(text string, password string) (string, error) {
    plain := []byte(text)
    data  := []byte(password)
    key   := sha256.Sum256(data)

    // Create a new cipher block from the key
    block, err := aes.NewCipher(key[:])
    if err != nil {
        return "", err
    }

    // Create a new GCM
    // https://en.wikipedia.org/wiki/Galois/Counter_Mode
    // https://golang.org/pkg/crypto/cipher/#NewGCM
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    // Create a nonce. Nonce should be from GCM
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

// Decrypts `text` with `password`
func Decrypt(text string, password string) (string, error) {
    data   := []byte(password)
    key    := sha256.Sum256(data)
    enc, _ := hex.DecodeString(text)

    // Create a new cipher block from the key
    block, err := aes.NewCipher(key[:])
    if err != nil {
        return "", err
    }

    // Create a new GCM
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    // Get the nonce size
    nonceSize := gcm.NonceSize()

    // Extract the nonce from the encrypted data
    nonce, ciphered := enc[:nonceSize], enc[nonceSize:]

    // Decrypt the data
    plain, err := gcm.Open(nil, nonce, ciphered, nil)
    if err != nil {
        return "", err
    }

    return fmt.Sprintf("%s", plain), nil
}
