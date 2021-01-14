# crypto

A small go library to encrypt and decrypt text.


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

    encryted, err := crypto.Encrypt(text, pass)
    fmt.Println(encryted, err)
}
```

Check also [the cli README.MD](cli/crypto/README.md) for the cli tool docs.
