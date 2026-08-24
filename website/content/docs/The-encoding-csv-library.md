## Overview 

The `csv` library reads and writes comma-separated values (CSV) files. There are many kinds of CSV files; this package supports the format described in RFC 4180, except that it uses LF instead of CRLF as the newline character by default. 

White space is considered part of a field. 

Carriage returns before newline characters are silently removed. 

Blank lines are ignored. A line with only whitespace characters (excluding the ending newline character) is not considered a blank line. 

Fields which start and stop with the quote character " are called quoted-fields. The beginning and ending quote are not part of the field. 

The source `normal string,"quoted-field"` results in the fields `normal string`, `quoted-field`. 

Within a quoted-field a quote character followed by a second quote character is considered a single quote: `"the ""word"" is true","a ""quoted-field"""` results in `the "word" is true`, `a "quoted-field"`. Newlines and commas may be included in a quoted-field  

## Modules 

### ` ("encoding/csv")`

### ` ("errors")`

### ` ("strings")`
## Functions 

### `encode(L list) -> string / error`

Encodes a list into a CSV string, returning an error if the elements of the list aren't all of type `string`.  

### `decode(s string) -> list / error`

Decodes a CSV string into a `list`, returning an `error` if the input is not in fact a well-formed CSV string.  

