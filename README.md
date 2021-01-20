# crypto
![Build Status](https://action-badges.now.sh/jpedro/crypto)

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

## CLI

Check [cli/crypto](cli/crypto) for your terminal needs.
