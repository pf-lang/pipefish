## Overview 

The`base_32` library implements base32 encoding as specified by RFC 4648.  

## Modules 

### ` ("encoding/base32")`
## Types 

### `Encoding = clone string`

A 32-character string giving the encoding of the numbers 0 ... 31.  
## Constants 

### `HEX_ENCODING`

HexEncoding is the “Extended Hex Alphabet” defined in RFC 4648. It is typically used in DNS.  

### `STD_ENCODING`

StdEncoding is the standard base32 encoding, as defined in RFC 4648.  
## Functions 

### `encode(encoding Encoding, plaintext string) -> string`

Encodes the plaintext using the given encoding.  

### `decode(encoding Encoding, codedText string) -> string / error`

Decodes the coded text using the given encoding.  

