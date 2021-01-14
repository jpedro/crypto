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
    encryted, _ := crypto.Encrypt("test", "test")
    fmt.Println(encryted)
}
```

Check also the [cli/crypto](cli/crypto) cli tool usage.
