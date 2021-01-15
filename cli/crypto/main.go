// Swine leberkas venison
//
// Burgdoggen sirloin biltong chuck drumstick shank capicola porchetta. Turkey pork loin
// chuck fatback jowl. T-bone short ribs turducken cupim, brisket cow pork belly leberkas.
// Landjaeger ham hock fatback pig corned beef bresaola beef ribs. Pork pork chop boudin
// strip steak landjaeger, pork belly kevin pork loin capicola ham. Pastrami spare ribs
// porchetta, drumstick leberkas t-bone short loin doner filet mignon hamburger corned
// beef. Venison short loin flank, cupim fatback spare ribs pork loin buffalo turducken
// tail.package main
package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"

    "github.com/jpedro/crypto"
    "github.com/jpedro/color"
)

var USAGE = `USAGE:
    crypto encrypt [TEXT]  # Encrypts TEXT or uses the STDIN
    crypto decrypt [TEXT]  # Decrypts TEXT or uses the STDIN

ENVIRONMENT VARIABLES:
    CRYPTO_PASSWORD        # The password to use (to avoid the prompt)
`

func main() {
    if len(os.Args) < 2 {
        fmt.Println(USAGE)
        os.Exit(0)
    }

    command  := os.Args[1]
    password := os.Getenv("CRYPTO_PASSWORD")
    if password == "" {
        fmt.Printf("Enter the password: ")
        reader := bufio.NewReader(os.Stdin)
        pass, _ := reader.ReadString('\n')
        pass = strings.TrimSpace(pass)
        if pass == "" {
            // fmt.Printf("Warning: Environment variable %s is not set.",
            //     color.Paint("green", "CRYPTO_PASSWORD"))
            fmt.Println("Error: Password can't be empty.")
            os.Exit(1)
        }
        password = pass
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
        encrypted, err := crypto.Encrypt(text, password)
        if err != nil {
            fmt.Println("Error: Failed to encrypt.")
            os.Exit(1)
        }
        fmt.Println(encrypted)

    } else if command == "decrypt" {
        text = strings.Replace(text, "\n", "", -1)
        decrypted, err := crypto.Decrypt(text, password)
        if err != nil {
            fmt.Println("Error: Failed to decrypt.")
            os.Exit(1)
        }
        fmt.Println(decrypted)

    } else {
        fmt.Printf("Error: Command %s not found.\n",
            color.Paint("green", command))
        os.Exit(1)
    }
}
