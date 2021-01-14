package main

import(
    "fmt"

    "github.com/jpedro/crypto"
)

func main() {
    text := "test"
    pass := "test"

    encryted, _ := crypto.Encrypt(text, pass)
    fmt.Println(encryted)
}
