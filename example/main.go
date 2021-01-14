package main

import(
    "fmt"

    "github.com/jpedro/crypto"
)

func main() {
    encryted, _ := crypto.Encrypt("test", "test")
    fmt.Println(encryted)
}
