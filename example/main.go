package main

import(
    "fmt"

    "github.com/jpedro/crypto"
)

func main() {
    original := "test"
    password := "test"
    encrypted, _ := crypto.Encrypt(original, password)
    decryptaed, _ := crypto.Decrypt(encrypted, password)
    fmt.Printf("Encrypted: %s\n", encrypted)
    fmt.Printf("Decrypted: %s\n", decrypted)
}
