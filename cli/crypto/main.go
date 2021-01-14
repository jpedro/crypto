package main

import (
    "os"
    "fmt"
    "bufio"

    "github.com/jpedro/crypto"
    "github.com/jpedro/color"
)

var USAGE = `USAGE:
    crypto encrypt [TEXT]  # Encrypts TEXT or uses the STDIN
    crypto decrypt [TEXT]  # Decrypts TEXT or uses the STDIN

ENVIRONMENT VARIABLES:
    CRYPTO_PASSWORD        # The password to use
`

func main() {
    if len(os.Args) < 2 {
        fmt.Println(USAGE)
        os.Exit(0)
    }

    command  := os.Args[1]
    password := os.Getenv("CRYPTO_PASSWORD")
    if password == "" {
		fmt.Printf("Error: Environment variable %s is not set.", color.Paint("green", "CRYPTO_PASSWORD"))
		return
    }

    text := ""
    if len(os.Args) < 3 {
        // fmt.Printf("==> Enter the text to %s (stop with Ctrl+D or cancel with Ctrl+C):\n", command)
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
            text = scanner.Text()
        }
    } else {
        text = os.Args[2]
    }

    if command == "encrypt" {
        encrypted, _ := crypto.Encrypt(text, password)
        fmt.Println(encrypted)

	} else if command == "decrypt" {
        decrypted, _ := crypto.Decrypt(text, password)
        fmt.Println(decrypted)

	} else {
		fmt.Printf("Error: Command %s not found.\n", color.Paint("green", command))
		os.Exit(1)
    }
}
