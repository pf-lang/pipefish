In this page we will introduce some basic data types: `int`, `float`, `bool`, `rune`, `string`, and `null`.

## `int`

The `int` type contains whole numbers which may be either positive or negative.

Integer literals are as you'd expect: `36`, `0`, `-17`, etc. Besides these there are representations of integers in binary (e.g `0b10011` represents the number 19) octal (`0o23` represents 19) and hex (`0x13` represents 19.)

Arithmetic operations on integers are `+`, `-`, `*`, `/`, `div`, `mod` and `%`. The first three of these require no explanation. 

The `div` operation when applied to integers returns a whole number, e.g. `5 div 2` evaluates to `2`.

The `mod` operator is the mathematical modulus, e.g. `5 mod 2` is `1`, and `-3 mod 7` is `4`.

The `/` operator returns a float (see below).

The `%` operator is a suffix returning a percentage, e.g. `50%` evaluates to the floating-point value `0.5`.

There are also the comparison operators `<`, `>`, `<=` and `>=`, which work as you would expect them to.

And of course there are the `==` and `!=` operators, as discussed [here](https://github.com/tim-hardcastle/Pipefish/wiki/The-type-system#note-on-equality-and-comparison), but as these are defined for nearly every type we will take their existence for granted from now on.

## `float`

The `float` type contains decimal numbers: `6.5`, `-9.33`, `3.141592`, etc.

Its operations are `+`, `-`, `*`, `/`, `%`, `<`, `>`, `<=` and `>=`, the `%` operator being a suffix meaning "divide by 100" as with the integers, and the rest of which work as you'd expect.

Mixing the arithmetic of integers and floating-point numbers usually results in an error, with the exception of multiplication and division operations, which will yield floats as results.

## `bool`

The `bool` type consists of the boolean constants `true` and `false`. Its operations are `and`, `or`, and `not`:

`and` and `or` are lazily evaluated, as in most languages.

There is also a conditional operator. Let's demonstrate it in the REPL:

<img width="710" height="157" alt="image" src="https://github.com/user-attachments/assets/3ae10b2e-fb4a-4d0d-b5d7-d93f80ca018e" />

We will have a lot more to say about this in later pages.

## `rune`

A rune is a single Unicode codepoint. Rune literals are enclosed in single quotes, e.g. `'f'`, `'魚'`.

Besides the usual operations of equality and inequality, runes can be prepended or appended to strings using the `&` operator.

## `string`

There are two kinds of string literals. Double quotes, `"like this"`, represent escaped strings with the following escape codes:

| Sequence | Description |
| --- | --- |
| `\\` | Backslash |
| `\t` | Horizontal tab |
| `\n` | Line feed or newline |
| `\r` | Carriage return |
| `\"` | Double quote |

But if we use `` `strings with backticks` `` then these sequences are unescaped.

Strings can be added together with the `+` operator: `"mac" + "hine"` evaluates to `"machine"`.

The `len` function returns the length of a string. 

A string can be indexed by a number from and including 0 up to and excluding the length of the string, returning a rune: for example, `"Pipefish"[1]` is `'i'`.

A "slice" can be taken out of a string, returning another string: `"Pipefish"[1::4]` evaluates to `"ipe"`.

Note that the slice is from-including up-to-and-excluding, as is everything else in Pipefish.

## `null`

The `null` type contains one element, `NULL`, and has no operations defined on it except the default operations `==` and `!=`. (Which are not particular useful, because `NULL` is always equal to itself.) We'll see what it's good for later: in the meantime it definitely belongs on this page because it is indeed a simple type.