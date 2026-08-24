## Overview 

The `markdown` library implements functions to render markdown in the terminal using the terminal control codes. 

The markdown is "Pipefish-flavored", i.e. the colors red, yellow, green, cyan, blue, and purple can be produced by e.g. `<R>foo</>` for red.  

## Modules 

### ` ("github.com/tim-hardcastle/pipefish/source/text")`
## Types 

### `MdFormat = struct (leftMargin string, width int)`

Specifies a left margin as a string and a width as an int that includes the width of the margin.  
## Functions 

### `render(fmt MdFormat, s string) -> string`

Renders the string in the given format.  

### `render(fmt MdFormat, S snippet) -> string`

Converts a snippet to a string by applying the `string` function to each of its elements and concatenates them, then renders the resulting string in the given format.  

### `render(s string) -> string`

Renders the string in a default format with no left margin and a width of 92.  

### `render(S snippet) -> string`

Converts a snippet to a string by applying the `string` function to each of its elements and concatenates them, then renders the resulting string in the given format.  

