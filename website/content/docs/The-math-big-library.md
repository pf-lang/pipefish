## Overview 

The `big` library implements arbitrary-precision arithmetic (big numbers).  

## Modules 

### ` ("math/big")`

### ` ("errors")`

### ` ("strconv")`
## Types 

### `Float = wrapper *Float`

The big float type.  

### `Int = wrapper *Int`

### `Rat = wrapper *Rat`

The big rational type.  
## Functions 

### `Float(x float) -> Float`

Constructs a `Float` from a `float`.  

### `Float(s string) -> Float / error`

Constructs a `Float` from a `string`.  

### `Float(r Rat) -> Float`

Constructs a `Float` from a `Rat`.  

### `(x Float) + (y Float) -> Float`

Adds two `Float` values.  

### `(x Float) - (y Float) -> Float`

Subtracts one `Float` value from another.  

### `(x Float) * (y Float) -> Float`

Multiplies two `Float` values.  

### `(x Float) / (y Float) -> Float`

Divides one `Float` value by another.  

### `-(x Float) -> Float`

Negates a `Float` value.  

### `(x Float) < (y Float) -> bool`

Compares two `Float` values, returning `true` if `x < y`.  

### `(x Float) <= (y Float) -> bool`

Compares two `Float` values, returning `true` if `x <= y`.  

### `(x Float) > (y Float) -> bool`

Compares two `Float` values, returning `true` if `x > y`.  

### `(x Float) >= (y Float) -> bool`

Compares two `Float` values, returning `true` if `x >= y`.  

### `abs(x Float) -> Float`

Returns the absolute value of a `Float`.  

### `float(z Float) -> float`

Converts a `Float` to the built-in `float` type.  

### `int(z Float) -> int`

Converts a `Float` to the built-in `int` type.  

### `isInf(z Float) -> bool`

Tests whether the given `Float` is an infinity.  

### `isInt(z Float) -> bool`

Tests whether the given `Float` is an integer.  

### `sign(z Float) -> int`

`sign(z)` returns: 


- `-1` if `z < 0`
- `0` if `z == 0`
- `1` if `z > 0`
 

### `sqrt(z Float) -> Float`

Returns the square root of a `Float`.  

### `string(z Float)`

Overloads `string`.  

### `text(z Float, format int, prec int) -> string`

Converts the `Float` to a `string` of the given format and precision.  

### `Int(i int) -> Int`

Constructs an `Int` from an `int`.  

### `Int(s string) -> Int / error`

Constructs an `Int` from a `string`.  

### `Int(z Float) -> Int`

Constructs an `Int` from a `Float`.  

### `(x Int) + (y Int) -> Int`

Adds two `Int` values.  

### `(x Int) - (y Int) -> Int`

Subtracts one `Int` value from another.  

### `(x Int) * (y Int) -> Int`

Multiplies two `Int` values.  

### `-(x Int) -> Int`

Negates an `Int`.  

### `(x Int) / (y Int) -> Rat`

`x / y` returns the `Rat` obtained by dividing `x` by `y`.  

### `(x Int) div (y Int) -> Int`

`x div y` returns the `Int` obtained by dividing `x` by `y`.  

### `(x Int) mod (y Int) -> Int`

Returns the modulus of one `Int` by another.  

### `(x Int) ^ (y Int) -> Int`

Returns the exponent of on `Int` by another.  

### `(x Int) < (y Int) -> bool`

Compares two `Int` values and returns `true` if `x < y`.  

### `(x Int) <= (y Int) -> bool`

Compares two `Int` values and returns `true` if `x <= y`.  

### `(x Int) > (y Int) -> bool`

Compares two `Int` values and returns `true` if `x > y`.  

### `(x Int) >= (y Int) -> bool`

Compares two `Int` values and returns `true` if `x >= y`.  

### `abs(x Int) -> Int`

Returns the absolute value of an `Int`.  

### `andInt(x Int, y Int) -> Int`

Returns the result of a bitwise `and` operation on two `Int`s.  

### `andNotInt(x Int, y Int) -> Int`

Returns the result of a bitwise `and not` operation on two `Int`s.  

### `binomial(n int, k int) -> Int`

Returns the binomial coefficient C(n, k).  

### `bit(z Int, i int) -> int`

Returns the value of the `i`th bit of the `Int`, as `0` or `1`.  

### `bitLen(z Int) -> int`

Returns the length of the absolute value of x in bits. The bit length of 0 is 0.  

### `divMod(x any?, y any?) -> Int, Int`

`divMod(x, y)` returns the quotient and modulus of `x` by `y`.  

### `exp(x Int, y Int, m Int) -> Int`

Returns `x` to the power `y` mod `abs(m)`. If m == 0, z = x**y unless y <= 0, when z = 1. If m != 0, y < 0, and x and m are not relatively prime, z is unchanged and nil is returned. Modular exponentiation of inputs of a particular size is not a cryptographically constant-time operation.  

### `float(z Int) -> float`

Converts an `Int` to the built-in `float` type.  

### `fromString(s string, base int) -> Int / error`

Converts a string in the given base to an `Int`, returning an error if the string isn't well-formed.  

### `gcd(a Int, b Int) -> Int`

Returns the greatest common denominator of two `Int`s.  

### `int(z Int) -> int`

Converts an `Int` to the built-in `int` type.  

### `isSmallInt(z Int) -> bool`

Returns `true` if the given `Int` will fit into a built-in (64-bit) `int`.  

### `lsh(x Int, n int) -> Int`

Returns the result of a bitwise left shift of `x` by `n` bits.  

### `modInverse(g Int, n Int) -> Int`

Returns the inverse of `g` mod `n`.  

### `modSqrt(x Int, p Int) -> Int`

Returns a square root of `x` mod `p`.  

### `mulRange(a int, b int) -> Int`

Returns the result of multiplying the numbers in the range `a::b`.  

### `notInt(x Int) -> Int`

Resturns the result of performing a bitwise `not` on an `Int`.  

### `orInt(x Int, y Int) -> Int`

Returns the result of a bitwise `or` operation on two `Int`s.  

### `probablyPrime(z Int, n int) -> bool`

`probablyPrime` reports whether `x` is probably prime, applying the Miller-Rabin test with `n` pseudorandomly chosen bases as well as a Baillie-PSW test. 

If `x` is prime, ProbablyPrime returns true. If `x` is chosen randomly and not prime, `probablyPrime` probably returns false. The probability of returning true for a randomly chosen non-prime is at most ¼ⁿ. 

ProbablyPrime is 100% accurate for inputs less than 2⁶⁴. See Menezes et al., Handbook of Applied Cryptography, 1997, pp. 145-149, and FIPS 186-4 Appendix F for further discussion of the error probabilities. 

ProbablyPrime is not suitable for judging primes that an adversary may have crafted to fool the test.  

### `quo(x Int, y Int) -> Int`

Returns the quotient of `x` by `y` using truncated division.  

### `quoRem(x Int, y Int) -> Int, Int`

Returns the quotient and remainder of `x` by `y` using truncated division.  

### `rem(x Int, y Int) -> Int`

Returns the remainder when dividing `x` by `y` using truncated division.  

### `rsh(x Int, n int) -> Int`

Returns the result of a bitwise left shift of `x` by `n` bits.  

### `sign(z Int) -> int`

`sign(z)` returns: 


- `-1` if `z < 0`
- `0` if `z == 0`
- `1` if `z > 0`
 

### `sqrt(x Int) -> Int`

Returns ⌊√x⌋, the largest integer such that z² ≤ x.  

### `string(z Int) -> string`

Overloads `string`.  

### `text(z Int, base int) -> string`

`text` returns the string representation of x in the given base. Base must be between 2 and 62, inclusive. The result uses the lower-case letters 'a' to 'z' for digit values 10 to 35, and the upper-case letters 'A' to 'Z' for digit values 36 to 61. No prefix (such as "0x") is added to the string.  

### `trailingZeroBits(z Int) -> int`

Returns the number of consecutive least significant zero bits of |x|.  

### `xorInt(x Int, y Int) -> Int`

Returns the result of a bitwise `xor` operation on two `Int`s.  

### `Rat(i int, j int) -> Rat`

Constructs a `Rat` from two `int`s.  

### `Rat(z Float) -> Rat`

Constructs a `Rat` from a `Float`.  

### `Rat(s string) -> Rat / error`

Constructs a `Rat` from a `string`.  

### `Rat(x float) -> Rat`

Converts the built-in `float` type to a `Rat`.  

### `Rat(x Int, y Int)`

Converts two `Int`s to the `Rat` which is their quotient.  

### `Rat(x Int)`

Converts an `Int` to its representation as a `Rat`.  

### `Rat(x int)`

Converts an `int` to its representation as a `Rat`.  

### `(x Rat) + (y Rat) -> Rat`

Adds two `Rat` values.  

### `(x Rat) - (y Rat) -> Rat`

Subtracts one `Rat` value from another.  

### `(x Rat) * (y Rat) -> Rat`

Multiplies two `Rat` values.  

### `(x Rat) / (y Rat) -> Rat`

Divides one `Rat` value by another.  

### `-(x Rat) -> Rat`

Negates a `Rat` value.  

### `(x Rat) < (y Rat) -> bool`

Compares two `Rat` values and returns `true` if `x < y`.  

### `(x Rat) <= (y Rat) -> bool`

Compares two `Rat` values and returns `true` if `x <= y`.  

### `(x Rat) > (y Rat) -> bool`

Compares two `Rat` values and returns `true` if `x > y`.  

### `(x Rat) >= (y Rat) -> bool`

Compares two `Rat` values and returns `true` if `x >= y`.  

### `abs(z Rat)`

Returns the abolute value of `z`.  

### `denom(z Rat) -> Int`

Returns the denominator of `z`.  

### `float(z Rat)`

Converts a `Rat` to a built-in `float` value.  

### `floatString(z Rat, prec int) -> string`

Returns a string representation of `x` in decimal form with `prec` digits of precision after the radix point. The last digit is rounded to nearest, with halves rounded away from zero.  

### `inv(x Rat) -> Rat`

Returns the inverse of a `Rat`.  

### `isInt(z Rat) -> bool`

Tests whether `z` represents an integer.  

### `num(z Rat) -> Int`

Returns the numerator of a `Rat`.  

### `ratString(z Rat) -> string`

Returns a string representation of `z` in the form `"a/b"` if `b != 1`, and in the form `"a"` if `b == 1`.  

### `sign(z Rat) -> int`

`sign(z)` returns: 


- `-1` if `z < 0`
- `0` if `z == 0`
- `1` if `z > 0`
 

### `string(z Rat)`

Returns a string representation of `z` in the form `"a/b"` (even if `b == 1`).  

