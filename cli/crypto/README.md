# crypto

Go cli to encrypt and decrypt payloads.


## Usage

```bash
# Install the cli
$ go install github.com/jpedro/crypto/cli/crypto

# Set the password as env var
$ export CRYPTO_PASSWORD="test"

# Encrypt
$ echo test | crypto encrypt
039df0154dbce96f8302d30e9263314f0d88a2538e040640a051cd359076f5b4

# Decrypt
$ echo 039df0154dbce96f8302d30e9263314f0d88a2538e040640a051cd359076f5b4 | crypto decrypt
test

# Chain commands
$ echo test | crypto encrypt | crypto decrypt
test
```
