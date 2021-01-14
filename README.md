# crypto

The smallest go library to encrypt and decrypt text with a password.

## Usage

```go
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
```

## Cli

```bash
$ go get github.com/jpedro/crypto/cli/crypto

$ export CRYPTO_PASSWORD="test"

$ echo test | crypto encrypt
039df0154dbce96f8302d30e9263314f0d88a2538e040640a051cd359076f5b4

$ echo '039df0154dbce96f8302d30e9263314f0d88a2538e040640a051cd359076f5b4' | crypto decrypt
test
```
