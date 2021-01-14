# crypto

The smallest go library to encrypt and decrypt text.


## Usage

### Library

```go
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
```

### CLI

```bash
# Install the cli
$ go get github.com/jpedro/crypto/cli/crypto

# Set the password as env var
$ export CRYPTO_PASSWORD="test"

# Encrypt some text
$ echo test | crypto encrypt
039df0154dbce96f8302d30e9263314f0d88a2538e040640a051cd359076f5b4

# Decrypt it back
$ echo '039df0154dbce96f8302d30e9263314f0d88a2538e040640a051cd359076f5b4' | crypto decrypt
test

# Chain the commands
$ echo test | crypto encrypt | crypto decrypt
test
```
