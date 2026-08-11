**Note:** we have interpreted the term "builtin" very broadly in this appendix and included many things that are properly speaking special forms; but we have excluded [`for` loops](https://github.com/tim-hardcastle/Pipefish/wiki/For-loops), [`with` and `without`](https://github.com/tim-hardcastle/Pipefish/wiki/With-and-without), and the various [piping](https://github.com/tim-hardcastle/Pipefish/wiki/Piping) operators, except by giving links to extended articles here and in the appropriate places (if any) below rather than attempting a precis.

## Builtins that act on all or most types

`(x any) == (y any)`
Returns `true` if the values are the same, `false` if they are different but of the same concrete type, and will return an error if the types are different, or if they are the same but are either `func` or `error`.

`(x any) != (y any)`
The opposite of the previous operator.

`(x any) -> <expression>`
A piping operator. For further information we direct the reader to [the entire article on piping](https://github.com/tim-hardcastle/Pipefish/wiki/Piping).

`literal(x any)`
Returns a Pipefish string literal representing the given value, e.g. `literal 2 * 2` returns `"4"`.

`string(x any)`
Converts a value to a string representation which is prettier than a literal but therefore is not in fact a literal. Can be recursively overloaded to make it still prettier.

`type(x any)`
Returns the type of the given value.

## Bool

`(x bool) and (y bool)`
Returns `true` if both `x` and `y` are true, otherwise `false`. If `x` is `false`, `y` is not computed.

`not(x bool)`
Returns the opposite of `x`.

`(x bool) or (y bool)`
Returns `true` if either `x` or `y` (or both) is true, otherwise `false`. If `x` is `true`, `y` is not computed.

For the behavior of the conditional  `:` and `;`/`newline` operators, see the main article on [Functions and commands](https://github.com/tim-hardcastle/Pipefish/wiki/Functions-and-commands).

## Clone types

`cast(x single, y type)`
Returns a cast of a value to a given type; or an error if `y` is not a clone of the type of `x`. Note that this can be overloaded, like most builtins, to allow other casts where they are desirable.

`<name of clone type>(x <base type>)`
A constructor for each clone type.

## Enum types

`(x <any enum type>)[(y int)]`
Returns the `y`-indexed element of the type, counting as usual from `0`, or an error if this element does not exist.

## Errors

`error(x string)`
Returns an error with the given error text.

`unwrap(x error)`
Returns a value of type `Error` having fields `errorCode` and `errorMessage` if it is passed an error, and returns an error otherwise.

`valid(x single?/tuple/error)`
Returns `true` if the value is not a error; `false` if it is.

## Floats

`(x float) + (y float)`
Returns the sum of two floating-point numbers.

`(x float) - (y float)`
Returns the difference of two floating-point numbers.

`(x float) * (y float)`
Returns the product of two floating-point numbers.

`(x float) / (y float)`
Returns the quotient of two floating-point numbers, or an error if the divisor is zero.

`(x floatlike) > (y floatlike)`
Returns `true` if `x` is strictly greater than `y`, or `false` otherwise.

`(x floatlike) >= (y floatlike)`
Returns `true` if `x` is greater than or equal to `y`, or `false` otherwise.

`(x floatlike) < (y floatlike)`
Returns `true` if `x` is strictly less than `y`, or `false` otherwise.

`(x floatlike) <= (y floatlike)`
Returns `true` if `x` is less than or equal to `y`, or `false` otherwise.

`float(x floatlike)`
Casts any floatlike value to a float, but returns an error if passed any other type.

`float(x int)`
Converts any integer to a float.

`float(x string)`
Converts a string to a float if the string is indeed a valid representation of a float, and returns an error otherwise.

## Ints

`(x int) + (y int)`
Returns the sum of two integers.

`(x int) - (y int)`
Returns the difference between two integers.

`(x int) * (y int)`
Returns the product of two integers.

`(x int) / (y int)`
Returns the quotient of two integers, or an error if the divisor is zero.

`(x intlike) > (y intlike)`
Returns `true` if `x` is strictly greater than `y`, or `false` otherwise.

`(x intlike) >= (y intlike)`
Returns `true` if `x` is greater than or equal to `y`, or `false` otherwise.

`(x intlike) < (y intlike)`
Returns `true` if `x` is strictly less than `y`, or `false` otherwise.

`(x intlike) <= (y intlike)`
Returns `true` if `x` is less than or equal to `y`, or `false` otherwise.

`int(x intlike)`
Casts any intlike value to an integer, but returns an error if passed a value of any other type.

`int(x string)`
Converts a string to a float if the string is indeed a valid representation of a float, and returns an error otherwise.

`int(x float)`
Converts a float to an integer

## Labels

`label(x string)`
Converts the given string to the label, if the label exists.

## Lists

`(x list) + (y list)`
Returns the concatenation of two lists.

`(x listlike)[(y int)]`
Returns the `y`-indexed element of the list, counting as usual from `0`, or an error if this element does not exist.

`(x list)[(y::z pair)]`
Returns the elements of the list from and including `y` up to and excluding `z`, or an error if some or all of those elements do not exist.

`(L list) >> <expression>`
A mapping operator. For further information we direct the reader to [the entire article on piping](https://github.com/tim-hardcastle/Pipefish/wiki/Piping).

`(x single?) in (y listlike)`
Returns `true` if the list contains the given element and `false` otherwise.

`len(x listlike)`
Returns the length of a list or a clone of `list`.

`list(x listlike)`
Casts any listlike value to an list, but returns an error if passed a value of any other type.

`(x list) ?> <expression>`
A filter operator. For further information we direct the reader to [the entire article on piping](https://github.com/tim-hardcastle/Pipefish/wiki/Piping).

`(x list) with (y ... pair)`
For this we refer the reader to [the entire page on the `with` and `without` operators](https://github.com/tim-hardcastle/Pipefish/wiki/With-and-without), as they cannot be briefly summarized.

## Maps

`(x maplike)[(y single?)]`
Returns the value in the map indexed by the key `y`, or an error if the key does not exist.

`keys(x maplike)`
Returns the keys of a map.

`len(x maplike)`
Returns the length of a map, where the length is the number of key-value pairs.

`map(x ... pair)`
Converts a tuple of pairs into a map, so long as the keys are hashable, and returns an error otherwise.

`(x map) with (y ... single?)`
For this we refer the reader to [the entire page on the `with` and `without` operators](https://github.com/tim-hardcastle/Pipefish/wiki/With-and-without), as they cannot be briefly summarized.

## Pairs

`(x pair)[(y int)]`
Returns the `y`-indexed element of the pair, where `y` must be `0` or `1`, and an error if it isn't.

`(x single?) :: (y single?)`
Returns a pair of which `x` is the 0-indexed element and `y` is the 1-indexed element.

`pair(x pairlike)`
Casts any pairlke value to a pair, but returns an error if passed a value of any other type.

## Runes

`(x rune) + (y rune)`
Returns the concatenation of two runes as a string.

`(x rune) + (y string)`
Returns the result of prepending a rune to a string.

`(x string) + (y rune)`
Returns the result of appending the rune `y` to a string `x`.

`codepoint(x rune)`
Returns the Unicode value of the rune.

`rune(x int)`
Returns a rune corresponding to the given codepoint.

## Sets

`(x single?) in (y set)`
Returns `true` if the list contains the given element and `false` otherwise.

`(x set) + (y set)`
Returns the union of the given sets.

`len(x setlike)`
Returns the length of a set or a clone of `set`.

`set(x setlike)`
Casts any value in a clone of `set` (or `set` itself) to a pair, but returns an error if passed a value of any other type.

`set(x ... single?)`
Converts a tuple to a set containing the same elements.

## Strings

`(x rune) + (y string)`
Returns the result of prepending a rune to a string.

`(x string) + (y rune)`
Returns the result of appending the rune `y` to a string `x`.

`(x string) + (y string)`
Returns the concatenation of two strings.

`(x stringlike)[(y int)]`
Returns the `y`-indexed rune of the string, counting as usual from `0`, or an error if this rune does not exist.

`(x string)[(y::z pair)]`
Returns the substring of the list from and including `y` up to and excluding `z`, or an error if some or all of those elements do not exist.

`len(x stringlike)`
Returns the length of a string or a clone of `string`.

`string(x any)`
Converts a value to a string representation which is prettier than a literal but therefore is not in fact a literal. Can be recursively overloaded.

`literal(x any)`
Returns a Pipefish string literal representing the given value, e.g. `literal 2 * 2` returns `"4"`

## Struct types

`(x <any struct type>)[(y label)]`
Returns the value in the struct indexed by the label `y`, or an error if the key does not exist.

`keys(x struct)`
Returns the keys of a struct, which will belong to the type `label`. 

`(x map) without (y ... single?)`
For this we refer the reader to [the entire page on the `with` and `without` operators](https://github.com/tim-hardcastle/Pipefish/wiki/With-and-without), as they cannot be briefly summarized.

`(x struct) with (y ... single?)`
For this we refer the reader to [the entire page on the `with` and `without` operators](https://github.com/tim-hardcastle/Pipefish/wiki/With-and-without), as they cannot be briefly summarized.

See the [the page on structs](https://github.com/tim-hardcastle/Pipefish/wiki/Structs) for details of their constructors.

## Tuples

`(x tuple)[(y int)]`
Returns the `y`-indexed element of the tuple, counting as usual from `0`, or an error if this element does not exist.

`(x tuple)[(y::z pair)]`
Returns the elements of the list from and including `y` up to and excluding `z`, or an error if some or all of those elements do not exist.

`(x single?) in (y tuple)`
Returns `true` if the tuple contains the given element and `false` otherwise.

`first(x tuple)`
Returns the first element of a tuple.

`last(x tuple)`
Returns the last element of a tuple.

`len(x tuple)`
Returns the length of a tuple.

`tuple(x ... single?)`
Coverts a sequence of values to a tuple. Useful for constructing tuples of one element.

`tuple(x tuple)`
Converts a tuple to the same tuple.

## The `type` type

`(x type) / (y type)`
Returns the union of the two types.

`(x single?) in (y type)`
Returns `true` if the type contains the given value and `false` otherwise.

`type(x any)`
Returns the type of the given value.

`varchar(x int)`
Returns the type `varchar(x)`.