package main

import(
    "fmt"

    "github.com/jpedro/crypto"
)

func main() {
    password := "test"
    encrypted, _ := crypto.Encrypt("test", password)
    decrypted, _ := crypto.Decrypt(encrypted, password)
    fmt.Printf("Encrypted: %s\n", encrypted)
    fmt.Printf("Decrypted: %s\n", decrypted)
}
