## Overview 

The `strings` library supplies functions for handling the built-in `string` type.  

## Modules 

### ` ("strings")`

### ` ("errors")`
## Functions 

### `compare(a string, b string) -> int`

`compare(a, b)` returns an integer comparing two strings lexicographically. The result is `0` if `a == b`, `-1` if `a < b`, and `+1` if `a > b`.  

### `contains(s string, substr string) -> bool`

`contains(s, substr)` reports whether `substr` is contained within `s`.  

### `containsAny(s string, chars string) -> bool`

`containsAny(s, chars)` reports whether any Unicode code points in `chars` are within `s`.  

### `count(s string, substr string) -> int`

`count(s, substr)` counts the number of non-overlapping instances of `substr` in `s`.  

### `cut(s string, sep string) -> string, string, bool`

`cut(s, sep)` slices `s` around the first instance of `sep`, returning the text before and after `sep`.  

### `equalFold(s string, t string) -> bool`

`equalFold(s, t)` reports whether `s` and `t`, interpreted as UTF-8 strings, are equal under simple Unicode case-folding.  

### `fields(s string)`

`fields(s)` splits `s` around one or more consecutive white-space characters.  

### `hasPrefix(s string, prefix string) -> bool`

`hasPrefix(s, prefix)` reports whether `s` begins with `prefix`.  

### `hasSuffix(s string, suffix string) -> bool`

`hasSuffix(s, suffix)` reports whether `s` ends with `suffix`.  

### `index(s string, substr string) -> int`

`index(s, substr)` returns the index of the first instance of `substr` in `s`, or -1 if `substr` is not present in `s`.  

### `indexAny(s string, chars string) -> int`

`indexAny(s, chars)` returns the index of the first instance in `s` of any Unicode code point from `chars`, or -1 if none is present.  

### `join(elems list, sep string) -> string / error`

`join(elems, sep)` concatenates the elements of `elems` to create a single string separated by `sep`. It returns an error if `elems` contains a non-string value.  

### `lastIndex(s string, substr string) -> int`

`lastIndex(s, substr)` returns the index of the last instance of `substr` in `s`, or -1 if `substr` is not present in `s`.  

### `lastIndexAny(s string, chars string) -> int`

`lastIndexAny(s, chars)` returns the index of the last instance in `s` of any Unicode code point from `chars`, or -1 if none is present.  

### `repeat(s string, c int) -> string`

`repeat(s, c)` returns a new string consisting of `c` copies of `s`.  

### `replace(s string, old string, new string, n int) -> string`

`replace(s, old, new, n)` returns a copy of `s` with the first `n` non-overlapping instances of `old` replaced by `new`.  

### `replaceAll(s string, old string, new string) -> string`

`replaceAll(s, old, new)` returns a copy of `s` with all non-overlapping instances of `old` replaced by `new`.  

### `split(s string, sep string)`

`split(s, sep)` slices `s` into all substrings separated by `sep`.  

### `splitAfter(s string, sep string)`

`splitAfter(s, sep)` slices `s` into all substrings after each instance of `sep` and includes `sep` in the returned substrings.  

### `splitAfterN(s string, sep string, n int)`

`splitAfterN(s, sep, n)` slices `s` into substrings after each instance of `sep` and includes `sep` in the substrings returned.  

### `splitN(s string, sep string, n int)`

`splitN(s, sep, n)` slices `s` into substrings separated by `sep`.  

### `toLower(s string) -> string`

`toLower(s)` returns a copy of `s` with all Unicode letters mapped to their lower case.  

### `toTitle(s string) -> string`

`toTitle(s)` returns a copy of `s` with all Unicode letters mapped to their title case.  

### `toUpper(s string) -> string`

`toUpper(s)` returns a copy of `s` with all Unicode letters mapped to their upper case.  

### `toValidUTF_8(s string, replacementString string) -> string`

`toValidUTF_8(s, replacementString)` returns a copy of `s` with each run of invalid UTF-8 byte sequences replaced by `replacementString`.  

### `trim(s string, cutset string) -> string`

`trim(s, cutset)` returns a slice of `s` with all leading and trailing Unicode code points contained in `cutset` removed.  

### `trimLeft(s string, cutset string) -> string`

`trimLeft(s, cutset)` returns a slice of `s` with all leading Unicode code points contained in `cutset` removed.  

### `trimPrefix(s string, prefix string) -> string`

`trimPrefix(s, prefix)` returns `s` without the provided leading prefix. If `s` does not start with `prefix`, `s` is returned unchanged.  

### `trimRight(s string, cutset string) -> string`

`trimRight(s, cutset)` returns a slice of `s` with all trailing Unicode code points contained in `cutset` removed.  

### `trimSuffix(s string, prefix string) -> string`

`trimSuffix(s, suffix)` returns `s` without the provided trailing suffix. If `s` does not end with the suffix, `s` is returned unchanged.  

