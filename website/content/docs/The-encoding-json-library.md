## Overview 

The `json` library implements semantic processing of JSON as specified in RFC 8259.  

## Modules 

### `strings ("/home/tim/Pipefish/source/initializer/libraries/strings.pf")`
## Types 

### `Json = clone snippet`

A type for containing fragments of Json together with values to be injected in it.  
## Functions 

### `string(J Json) -> string`

Converts the `Json` type to a `string` by injecting the values, performing sanitation as it does so.  

### `decode(j string)`

Default decoding of Json.  

### `decode(j string) as (T type)`

Decodes Json to type `T`.  

### `decode(j string) like (T type)`

Decodes Json like type `T`: i.e. if `T` is list{MyStructType}, then the result will be of type `list` with elements of type `MyStructType`.  

### `encode(x any?) -> string`

Encodes the given value as Json.  

### `prettyEncode(x any)`

Encodes the given value as Json with indentation.  

