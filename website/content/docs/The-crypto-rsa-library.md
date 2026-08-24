## Overview 

The `rsa` library implements RSA encryption as specified in PKCS #1 and RFC 8017. 

The original specification for encryption and signatures with RSA is PKCS #1 and the terms "RSA encryption" and "RSA signatures" by default refer to PKCS #1 version 1.5. However, that specification has flaws and new designs should use version 2, usually called by just OAEP and PSS, where possible.  

## Modules 

### ` ("crypto/rsa")`

### ` ("math/big")`

### ` ("crypto/rand")`

### ` ("encoding/base64")`

### ` ("crypto/sha256")`

### `big ("/home/tim/Pipefish/source/initializer/libraries/math/big.pf")`
## Types 

### `GenerateKey = struct (bits int)`

The data to pass to `get ... from` to generate a key.  

### `PublicKey = struct (modulus Int, exponent int)`

A public key/  

### `PrivateKey = struct (publicKey PublicKey, exponent Int, primeA Int, primeB Int)`

A private key.  

### `EncryptOaep = struct (publicKey PublicKey, plaintext string, oaepLabel string)`

The data to pass to `get ... from` to get OAEP-encrypted text.  

### `EncryptPkcs = struct (publicKey PublicKey, plaintext string)`

The data to pass to `get ... from` to get PKCS-encrypted text.  
## Commands 

### `get(x ref) from (gk GenerateKey)`

Populates the reference variable with a private key.  

### `get(x ref) from (e EncryptOaep)`

Gets the ciphertext from the plaintext using OAEP.  

### `get(x ref) from (e EncryptPkcs)`

Gets the ciphertext from the plaintext using PKCS.  
## Functions 

### `EncryptOaep(key PublicKey, plaintext string)`

Constructor for `EncryptOaep` defaulting to `""` for the label.  

### `decryptOaep(k PrivateKey, ciphertext string) -> string / error`

Decrypts an OAEP-encoded cyphertext into plaintext, with the OEAP label taken to be `""`.  

### `decryptOaep(k PrivateKey, ciphertext string, oaepLabel string) -> string / error`

Decrypts an OAEP-encoded cyphertext into plaintext.  

### `decryptPkcs(k PrivateKey, ciphertext string) -> string / error`

Decrypts a PKCS-encoded cyphertext into plaintext.  

