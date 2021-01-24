# Crypto
[![Github Status](https://github.com/jpedro/crypto/workflows/tests/badge.svg)](https://github.com/jpedro/crypto/actions)
[![GoDoc](https://godoc.org/github.com/jpedro/crypto?status.svg)](https://godoc.org/github.com/jpedro/crypto)

Small go library to encrypt and decrypt text.


## Usage

```go
package main

import(
    "fmt"

    "github.com/jpedro/crypto"
)

func main() {
    original := "test"
    password := "test"
    encrypted, _ := crypto.Encrypt(original, password)
    decrypted, _ := crypto.Decrypt(encrypted, password)
    fmt.Printf("Encrypted: %s\n", encrypted)
    fmt.Printf("Decrypted: %s\n", decrypted)
}
```

You can run it with:

```bash
$ (cd example && go run main.go)
Encrypted: 5fb983fceb745148b3d714425d0de00fbe5f0ed0bd7101c01198627f5cbfecc9
Decrypted: test
```

## CLI

Check [cli/crypto](cli/crypto) for your terminal needs.
