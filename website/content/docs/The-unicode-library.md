## Overview 

The `unicode` library provides data and functions to test some properties of Unicode code points.  

## Modules 

### ` ("unicode")`
## Types 

### `Case = enum UPPER_CASE, LOWER_CASE, TITLE_CASE, MAX_CASE`

An enum representing the various sorts of case recognized by the Unicode standard.  
## Functions 

### `isControl(r rune) -> bool`

`isControl(r)` reports whether `r` is a control character.  

### `isDigit(r rune) -> bool`

`isDigit(r)` reports whether `r` is a decimal digit.  

### `isGraphic(r rune) -> bool`

`isGraphic(r)` reports whether `r` is defined as a Graphic character by Unicode.  

### `isIn(r rune, bounds ... pair)`

`isIn` reports whether the rune lies within any of the bounds in `bounds`.  

### `isLetter(r rune) -> bool`

`isLetter(r)` reports whether `r` is a Unicode letter.  

### `isLower(r rune) -> bool`

`isLower(r)` reports whether `r` is a lower-case letter.  

### `isMark(r rune) -> bool`

`isMark(r)` reports whether `r` is a Unicode mark character.  

### `isNumber(r rune) -> bool`

`isNumber(r)` reports whether `r` is a Unicode number.  

### `isPrint(r rune) -> bool`

`isPrint(r)` reports whether `r` is defined as printable by Pipefish.  

### `isPunct(r rune) -> bool`

`isPunct(r)` reports whether `r` is a Unicode punctuation character.  

### `isSpace(r rune) -> bool`

`isSpace(r)` reports whether `r` is a Unicode white space character.  

### `isSymbol(r rune) -> bool`

`isSymbol(r)` reports whether `r` is a Unicode symbolic character.  

### `isTitle(r rune) -> bool`

`isTitle(r)` reports whether `r` is a title-case letter.  

### `isUpper(r rune) -> bool`

`isUpper(r)` reports whether `r` is an upper-case letter.  

### `simpleFold(r rune) -> rune`

`simpleFold(r)` iterates over Unicode code points equivalent to `r` under Unicode simple case folding.  

### `toCase(c Case, r rune) -> rune`

Converts `r` to the given case.  

### `toLower(r rune) -> rune`

`toLower(r)` maps the Unicode letter `r` to lower case.  

### `toTitle(r rune) -> rune`

`toTitle(r)` maps the Unicode letter `r` to title case.  

### `toUpper(r rune) -> rune`

`toUpper(r)` maps the Unicode letter `r` to upper case.  

