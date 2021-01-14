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
    fmt.Println(crypto.Encrypt(text, pass))
}
```
