**Note:** we have interpreted the term "builtin" very broadly in this appendix and included many things that are properly speaking special forms; but we have excluded [`for` loops](https://github.com/tim-hardcastle/Pipefish/wiki/For-loops), [`with` and `without`](https://github.com/tim-hardcastle/Pipefish/wiki/With-and-without), and the various [piping](https://github.com/tim-hardcastle/Pipefish/wiki/Piping) operators, except by giving links to extended articles here and in the appropriate places below rather than attempting a precis. For the conditional `:` and `;` operators so the article on [Functions and commands](https://github.com/tim-hardcastle/Pipefish/wiki/Functions-and-commands), which also introduces conditionals. 

## Builtins with symbols

**Note:** these are arranged in alphabetic order of how someone might describe them in words, e.g. addition comes before division because "a" comes before "d".

`(x float) + (y float)`
Returns the sum of two floating-point numbers.

`(x int) + (y int)`
Returns the sum of two integers.

`(x list) + (y list)`
Returns the concatenation of two lists.

`(x rune) + (y rune)`
Returns the concatenation of two runes as a string.

`(x rune) + (y string)`
Returns the result of prepending a rune to a string.

`(x set) + (y set)`
Returns the union of the given sets.

`(x string) + (y rune)`
Returns the result of appending the rune `y` to a string `x`.

`(x string) + (y string)`
Returns the concatenation of two strings.

`(x float) / (y float)`
Returns the quotient of two floating-point numbers, or an error if the divisor is zero.

`(x int) / (y int)`
Returns the quotient of two integers, or an error if the divisor is zero.

`(x list) ?> <expression>`
A filter operator. For further information we direct the reader to [the entire article on piping](https://github.com/tim-hardcastle/Pipefish/wiki/Piping).

`(x float) > (y float)`
Returns `true` if `x` is strictly greater than `y`, or `false` otherwise.

`(x float) >= (y float)`
Returns `true` if `x` is greater than or equal to `y`, or `false` otherwise.

`(x int) > (y int)`
Returns `true` if `x` is strictly greater than `y`, or `false` otherwise.

`(x int) >= (y int)`
Returns `true` if `x` is greater than or equal to `y`, or `false` otherwise.

`(x <any enum type>)[(y int)]`
Returns the `y`-indexed element of the type, counting as usual from `0`, or an error if this element does not exist.

`(x list)[(y int)]`
Returns the `y`-indexed element of the list, counting as usual from `0`, or an error if this element does not exist.

`(x list)[(y::z pair)]`
Returns the elements of the list from and including `y` up to and excluding `z`, or an error if some or all of those elements do not exist.

`(x map)[(y single?)]`
Returns the value in the map indexed by the key `y`, or an error if the key does not exist.

`(x pair)[(y int)]`
Returns the `y`-indexed element of the pair, where `y` must be `0` or `1`, and an error if it isn't.

`(x <any struct type>)[(y label)]`
Returns the value in the struct indexed by the label `y`, or an error if the key does not exist.

`(x string)[(y int)]`
Returns the `y`-indexed rune of the string, counting as usual from `0`, or an error if this rune does not exist.

`(x string)[(y::z pair)]`
Returns the substring of the list from and including `y` up to and excluding `z`, or an error if some or all of those elements do not exist.

`(x tuple)[(y int)]`
Returns the `y`-indexed element of the tuple, counting as usual from `0`, or an error if this element does not exist.

`(x tuple)[(y::z pair)]`
Returns the elements of the list from and including `y` up to and excluding `z`, or an error if some or all of those elements do not exist.

`(x float) < (y float)`
Returns `true` if `x` is strictly less than `y`, or `false` otherwise.

`(x float) <= (y float)`
Returns `true` if `x` is less than or equal to `y`, or `false` otherwise.

`(x int) < (y int)`
Returns `true` if `x` is strictly less than `y`, or `false` otherwise.

`(x int) <= (y int)`
Returns `true` if `x` is less than or equal to `y`, or `false` otherwise.

`(x float) * (y float)`
Returns the product of two floating-point numbers.

`(x int) * (y int)`
Returns the product of two integers.

`(x list) >> <expression>`
A mapping operator. For further information we direct the reader to [the entire article on piping](https://github.com/tim-hardcastle/Pipefish/wiki/Piping).

`(x int) % (y int)`
Returns the modulus of two integers, or an error if the divisor is zero. The result is negative if the dividend is negative.

`(x any) -> <expression>`
A piping operator. For further information we direct the reader to [the entire article on piping](https://github.com/tim-hardcastle/Pipefish/wiki/Piping).

`(x float) - (y float)`
Returns the difference between two floating-point numbers.

`(x int) - (y int)`
Returns the difference between two integers.

`(x single?) :: (y single?)`
Returns a pair of which `x` is the 0-indexed element and `y` is the 1-indexed element.

`(x type) / (y type)`
Returns the union of the two types.

## Builtins with names

(x bool) and (y bool)
Returns `true` if both `x and `y` are true, otherwise `false`. It doesn't bother to compute `y` if x is false.

`cast(x single, y type)`
Returns a cast of a value to a given type; or an error if `y` is not a clone of the type of `x`. Note that this can be overloaded, like most builtins, to allow other casts where they are desirable.

`codepoint(x rune)`
Returns the Unicode value of the rune.

`error(x string)`
Returns an error with the given error text.

`for <and various forms of expression following it>`
For this we direct the reader to [the entire article on for loops](https://github.com/tim-hardcastle/Pipefish/wiki/For-loops).

`float(x floatlike)`
Casts any value in a clone of `float` to a float, but returns an error if passed any other type.

`float(x int)`
Converts any integer to a float.

`float(x string)`
Converts a string to a float if the string is indeed a valid representation of a float, and returns an error otherwise.

`(x single?) in (y list)`
Returns `true` if the list contains the given element and `false` otherwise.

`(x single?) in (y set)`
Returns `true` if the list contains the given element and `false` otherwise.

`(x single?) in (y tuple)`
Returns `true` if the tuple contains the given element and `false` otherwise.

`(x single?) in (y type)`
Returns `true` if the type contains the given value and `false` otherwise.

`first(x tuple)`
Returns the first element of a tuple.

`int(x intlike)`
Casts any value in a clone of `int` (or `int` itself) to an integer, but returns an error if passed a value of any other type.

`int(x string)`
Converts a string to a float if the string is indeed a valid representation of a float, and returns an error otherwise.

`int(x float)`
Converts a float to an integer

`keys(x maplike)`
Returns the keys of a map or a clone of a map.

`keys(x struct)`
Returns the keys of a struct, which will belong to the type `label`. 

`label(x string)`
Converts the given string to the label, if the label exists.

`last(x tuple)`
Returns the last element of a tuple.

`len(x listlike)`
Returns the length of a list or a clone of `list`.

`len(x maplike)`
Returns the length of a map or a clone of `map`, where the length is the number of key-value pairs.

`len(x setlike)`
Returns the length of a set or a clone of `set`.

`len(x stringlike)`
Returns the length of a string or a clone of `string`.

`len(x tuple)`
Returns the length of a tuple.

`list(x listlike)`
Casts any value in a clone of `list` (or `list` itself) to an list, but returns an error if passed a value of any other type.

`literal(x single?)`
Returns a Pipefish string literal representing the given value, e.g. `literal 2 * 2` returns `"4"`

`literal(x tuple)`
Returns a Pipefish string literal representing the given value.

`map(x maplike)`
Casts any value in a clone of `map` (or `map` itself) to an map, but returns an error if passed a value of any other type.

`(x list) ?> <expression>`
A filter operator. For further information we direct the reader to [the entire article on piping](https://github.com/tim-hardcastle/Pipefish/wiki/Piping).

`map(x ... pair)`
Converts a tuple of pairs into a map, so long as the keys are hashable, and returns an error otherwise.

not(x bool)
Returns the opposite of `x`.

(x bool) or (y bool)
Returns `true` if either `x or y` are true, otherwise `true`. It doesn't bother to compute `y` if x is true.

`pair(x pairlike)`
Casts any value in a clone of `pair` (or `pair` itself) to a pair, but returns an error if passed a value of any other type.

`rune(x int)`
Returns a rune corresponding to the given codepoint.

`set(x setlike)`
Casts any value in a clone of `set` (or `set` itself) to a pair, but returns an error if passed a value of any other type.

`set(x ... single?)`
Converts a tuple to a set containing the same elements.
  
`string`(x single?)`
Converts a value to a string representation which is prettier than a literal but therefore is not in fact a literal. Can be recursively overloaded.

`tuple(x ... single?)`
Coverts a sequence of values to a tuple. Useful for constructing tuples of one element.

`tuple(x tuple)`
Converts a tuple to the same tuple.

`type(x single?)`
Returns the type of the given value.

`type(x tuple)`
Returns `tuple`.

`unwrap(x error)`
Returns a value of type `Error` having fields `errorCode` and `errorMessage` if it is passed an error, and returns an error otherwise.

`valid(x single?/tuple/error)`
Returns `true` if the value is not a error; `false` if it is.

`varchar(x int)`
Returns the type `varchar(x)`.

- `(x list) with (y ... pair)`
- `(x map) with (y ... single?)`
- `(x type) with (y ... single?)`
- `(x map) without (y ... single?)`
For these we refer the reader to [the entire page on these operators](https://github.com/tim-hardcastle/Pipefish/wiki/With-and-without), as they cannot be briefly summarized.