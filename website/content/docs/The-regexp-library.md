## Overview 

The `regexp` library implements regular expression search. 

The syntax of the regular expressions accepted is the same general syntax used by Perl, Python, and other languages. More precisely, it is the syntax accepted by RE2 and described at https://golang.org/s/re2syntax, except for \C.  

## Modules 

### ` ("regexp")`

### ` ("errors")`
## Functions 

### `match(pattern string, text string) -> bool / error`

Returns true if a match is found.  

### `find(pattern string, text string) -> string / error`

Returns the first matching string.  

### `findAllString(pattern string, text string, start int) -> list / error`

Returns a list of all the matching strings.  

### `findAllIndex(pattern string, text string, start int)`

Returns a list of the indices of all the matching strings.  

### `replaceAll(pattern pair, text string) -> string / error`

Finds `pattern[0]` in `text` and replaces it with `pattern[1]`.  

