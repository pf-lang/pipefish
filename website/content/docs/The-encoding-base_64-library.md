## Overview 

The `base_64` library implements base64 encoding as specified by RFC 4648.  

## Modules 

### ` ("encoding/base64")`
## Types 

### `Encoding = struct (padding bool, key string)`

Contains the data specifying how encoding and decoding should be done, i.e. whether the encoding is padded and which characters correspond to the numbers 0 ... 63.  
## Constants 

### `STD_ENCODING`

`STD_ENCODING` is the standard base64 encoding, as defined in RFC 4648.  

### `URL_ENCODING`

`URL_ENCODING` is the alternate base64 encoding defined in RFC 4648. It is typically used in URLs and file names.  

### `RAW_STD_ENCODING`

`RAW_STD_ENCODING` is the standard raw, unpadded base64 encoding, as defined in RFC 4648 section 3.2. This is the same as `STD_ENCODING` but omits padding characters.  

### `RAW_URL_ENCODING`

`RAW_URL_ENCODING` is the unpadded alternate base64 encoding defined in RFC 4648. It is typically used in URLs and file names. This is the same as URLEncoding but omits padding characters.  
## Functions 

### `encode(encoding Encoding, plaintext string) -> string`

Encodes the plaintext using the given encoding.  

### `decode(encoding Encoding, codedText string) -> string / error`

Decodes the coded text using the given encoding.  

