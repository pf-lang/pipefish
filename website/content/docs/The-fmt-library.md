## Overview 

The `fmt` library implements formatting that wraps around Go's string formatting functions.  

## Modules 

### ` ("fmt")`
## Functions 

### `errorf(format string, a ... any)`

Returns a formatted `error` value.  

### `sprint(a ... any) -> string`

`sprint` formats using the default formats for its operands and returns the resulting string. Spaces are added between operands when neither is a string.  

### `sprintf(format string, a ... any) -> string`

`sprintf` formats according to a format specifier and returns the resulting string.  

### `sprintln(a ... any) -> string`

`sprintln` formats using the default formats for its operands and returns the resulting string. Spaces are always added between operands and a newline is appended.  

