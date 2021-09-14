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

	"github.com/jpedro/crypto"
)

type From int

const (
	FROM_ARGS From = iota
	FROM_STDIN
)

var (
	VERSION = "v0.1.6"
	USAGE   = `USAGE:
    crypto encrypt [TEXT]   # Encrypts TEXT or uses the STDIN
    crypto decrypt [TEXT]   # Decrypts TEXT or uses the STDIN
    crypto help             # Shows this help
    crypto version          # Shows the current version

ENVIRONMENT VARIABLES:
    CRYPTO_PASSWORD        # The password to use (avoids the prompt, required if you use stdin)
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
		from, payload := getPayload(command)
		password := getPassword(from)
		encrypted, err := crypto.Encrypt(payload, password)
		if err != nil {
			bail("Error: Failed to encrypt (%s).", err)
			os.Exit(1)
		}
		fmt.Println(encrypted)

	} else if command == "decrypt" {
		from, payload := getPayload(command)
		password := getPassword(from)
		payload = strings.Replace(payload, "\n", "", -1)
		decrypted, err := crypto.Decrypt(payload, password)
		if err != nil {
			bail("Error: Failed to decrypt (%s)\n", err)
			os.Exit(1)
		}
		fmt.Println(decrypted)

	} else {
		bail("Error: Command '%s' not found.\n", command)
		bail("Run 'crypto help' to check available commands.\n")
		os.Exit(1)
	}
}

func getPayload(command string) (From, string) {
	if len(os.Args) < 3 {
		return FROM_STDIN, readStdin(command)
	} else {
		return FROM_ARGS, os.Args[2]
	}
}

func getPassword(from From) string {
	password := os.Getenv("CRYPTO_PASSWORD")
	if password != "" {
		return password
	}

	if from == FROM_STDIN {
		bail("Error: The 'CRYPTO_PASSWORD' env var is required if you use stdin.")
		os.Exit(2)
	}

	password = askPassword("Enter the password: ")
	if password == "" {
		bail("Error: The password can't be empty.\n")
		bail("You can use the 'CRYPTO_PASSWORD' env var instead.\n")
		os.Exit(2)
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

func bail(message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, message, args...)
}
