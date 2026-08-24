## Overview 

The `bcrypt` library implements Provos and Mazières's bcrypt adaptive hashing algorithm. See (http://www.usenix.org/event/usenix99/provos/provos.pdf.)  

## Modules 

### ` ("golang.org/x/crypto/bcrypt")`
## Constants 

### `DEFAULT_COST`

The default cost used by the hashing function.  
## Functions 

### `compare(hashedPassword string, plaintext string) -> bool`

Compares a bcrypt hashed password with its possible plaintext equivalent. on success, or an error on failure.  

### `hash(s string) -> string / error`

Returns the bcrypt hash of the password using the default cost defined in this package. 

`hash` does not accept passwords longer than 72 characters, which is the longest password bcrypt will operate on.  

### `hash(s string, cost int) -> string / error`

Returns the bcrypt hash of the password at the given cost. If the cost given is less than 4, the cost will be set to DEFAULT_COST,  

