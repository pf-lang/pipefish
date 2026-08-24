## Overview 

The `aes` library implements AES encryption as defined in U.S. Federal Information Processing Standards Publication 197. 

The AES operations in this package are not implemented using constant-time algorithms. An exception is when running on systems with enabled hardware support for AES that makes these operations constant-time. Examples include amd64 systems using AES-NI extensions and s390x systems using Message-Security-Assist extensions. On such systems, when the result of NewCipher is passed to cipher.NewGCM, the GHASH operation used by GCM is also constant-time.  

## Modules 

### ` ("crypto/aes")`

### ` ("crypto/cipher")`

### ` ("crypto/rand")`

### ` ("encoding/base64")`

### ` ("io")`
## Types 

### `AesEncode = struct (key string, plaintext string)`

Contains the data to pass to `get ... from`. 

The `key` field must be the AES key, either 16, 24, or 32 characters to select AES-128, AES-192, or AES-256.  
## Commands 

### `get(ciphertext ref) from (a AesEncode)`

Sets the reference variable to the result of encoding the `plaintext` field of `a`.  
## Functions 

### `decode(key string, ciphertext string) -> string / error`

Decodes the ciphertext using the key supplied, returning the recovered plaintext if the decoding works and an error otherwise.  

