## Overview 

The `rand` library implements a cryptographically secure random number generator.  

## Modules 

### ` ("crypto/rand")`

### ` ("math/big")`

### `big ("/home/tim/Pipefish/source/initializer/libraries/math/big.pf")`
## Types 

### `RandomInt = struct (max string / Int)`

A data structure to pass to `get ... from` to get a random integer.  

### `RandomPrime = struct (bits int)`

A data structure to pass to `get ... from` to get a random prime.  

### `RandomText = struct `

A data structure to pass to `get ... from` to get a random integer.  
## Commands 

### `get(x ref) from (r RandomInt)`

Returns a uniform random value in [0, max).  

### `get(x ref) from (r RandomPrime)`

Returns a number of the given bit length that is prime with high probability.  

### `get(x ref) from (r RandomText)`

Returns a cryptographically random string using the standard RFC 4648 base32 alphabet for use when a secret string, token, password, or other text is needed. The result contains at least 128 bits of randomness, enough to prevent brute force guessing attacks and to make the likelihood of collisions vanishingly small. A future version may return longer texts as needed to maintain those properties.  

