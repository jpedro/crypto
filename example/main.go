package main

import(
    "fmt"

    "github.com/jpedro/crypto"
)

func main() {
    text := "test"
    pass := "test"

    encryted, err := crypto.Encrypt(text, pass)
    fmt.Println(encryted, err)
}
