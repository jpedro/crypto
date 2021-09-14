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
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"

	"github.com/jpedro/color"
	"github.com/jpedro/crypto"
)

var (
	VERSION = "v0.1.0"
	USAGE   = `USAGE:
    crypto encrypt [TEXT]  # Encrypts TEXT or uses the STDIN
    crypto decrypt [TEXT]  # Decrypts TEXT or uses the STDIN

ENVIRONMENT VARIABLES:
    CRYPTO_PASSWORD        # The password to use (to avoid the prompt)
`
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(USAGE)
		os.Exit(0)
	}

	if os.Args[1] == "help" {
		fmt.Println(USAGE)
		os.Exit(0)
	}

	if os.Args[1] == "version" {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	command := os.Args[1]

	if command == "encrypt" {
		payload := getPayload(command)
		password := getPassword()
		encrypted, err := crypto.Encrypt(payload, password)
		if err != nil {
			fmt.Println("Error: Failed to encrypt.")
			os.Exit(1)
		}
		fmt.Println(encrypted)

	} else if command == "decrypt" {
		payload := getPayload(command)
		password := getPassword()
		payload = strings.Replace(payload, "\n", "", -1)
		decrypted, err := crypto.Decrypt(payload, password)
		if err != nil {
			fmt.Println("Error: Failed to decrypt.")
			os.Exit(1)
		}
		fmt.Println(decrypted)

	} else {
		fmt.Printf("Error: Command %s not found. Run 'crypt help' to check available options.\n",
			color.Paint("green", command))
		os.Exit(1)
	}
}

func getPayload(command string) string {
	if len(os.Args) < 3 {
		return readStdin(command)
	} else {
		return os.Args[2]
	}
}

func getPassword() string {
	password := os.Getenv("CRYPTO_PASSWORD")
	if password != "" {
		return password
	}

	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
	// password = strings.TrimSpace(input)
	password = askPassword("Enter the password: ")
	if password == "" {
		fmt.Println("Error: The password can't be empty. You can use the 'CRYPTO_PASSWORD' env var instead.")
		os.Exit(1)
	}

	return password
}

func askPassword(prompt string) string {
	fmt.Print(prompt)
	bytes, err := term.ReadPassword(syscall.Stdin)
	fmt.Println()
	if err != nil {
		return ""
	}

	password := string(bytes)
	return password
}

func readStdin(command string) string {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeNamedPipe == 0 {
		fmt.Printf("Enter the text to %s and finish with Ctrl+D:\n", command)
	}

	payload := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		payload = scanner.Text()
	}

	return payload
}
