## Overview 

The `math` package implements basic math functions for the `float` type.  

## Modules 

### ` ("math")`
## Functions 

### `abs(x float) -> float`

Returns the absolute value of `x`.  

### `acos(x float) -> float`

Returns the inverse cosine of `x`.  

### `acosh(x float) -> float`

Returns the inverse hyperbolic cosine of `x`.  

### `asin(x float) -> float`

Returns the inverse sine of `x`.  

### `asinh(x float) -> float`

Returns the inverse hyperbolic sine of `x`.  

### `atan(x float) -> float`

Returns the inverse tangent of `x`.  

### `atan_2(y float, x float) -> float`

Returns the inverse tangent of `y`/`x`, using the signs of the two to determine the quadrant of the return value.  

### `atanh(x float) -> float`

Returns the inverse hyperbolic tangent of `x`.  

### `cbrt(x float) -> float`

Returns the cube root of `x`.  

### `ceil(x float) -> float`

Returns the ceiling of `x`.  

### `copysign(f float, sign float) -> float`

Returns the sign of `x`.  

### `cos(x float) -> float`

Returns the cosine of `x`.  

### `cosh(x float) -> float`

Returns the hyperbolic cosine of `x`.  

### `dim(x float, y float) -> float`

Returns the the maximum of `x - y` or `0`.  

### `erf(x float) -> float`

Returns the error function of `x`.  

### `erfc(x float) -> float`

Returns the complementary error function of `x`.  

### `erfcinv(x float) -> float`

Returns the inverse of `erfc(x)`.  

### `erfinv(x float) -> float`

Returns the inverse of `erv(x)`.  

### `exp(x float) -> float`

Returns the base-e exponential of `x`.  

### `exp_2(x float) -> float`

Returns the base-2 exponential of `x`.  

### `expm_1(x float) -> float`

Returns the base-e exponential of `x`, minus 1. It is more accurate than `exp(x) - 1` when x is near zero.  

### `fMA(x float, y float, z float) -> float`

Returns `x * y + z`, computed with only one rounding. (That is, FMA returns the fused multiply-add of `x`, `y`, and `z`.)  

### `floor(x float) -> float`

Returns the floor of `x`.  

### `frexp(f float) -> float, int`

`frexp` breaks `x` into a normalized fraction and an integral power of two, such that if `frac, exp = frexp(x)` then `x == frac × 2^exp`, with the absolute value of `frac` in the interval [½, 1).  

### `gamma(x float) -> float`

Returns the Gamma function of `x`.  

### `hypot(p float, q float) -> float`

Returns `sqrt(p*p + q*q)`, taking care to avoid unnecessary overflow and underflow.  

### `inf(sign int) -> float`

Returns positive infinity if `sign >= 0`, negative infinity if `sign < 0`.  

### `isInf(f float, sign int) -> bool`

Tests whether `f` is an infinity, according to `sign`. If `sign > 0`, `isInf` reports whether `f` is positive infinity. If `sign < 0`, `isInf` reports whether `f` is negative infinity. If `sign == 0`, `isInf` reports whether `f` is either infinity.  

### `isNaN(f float) -> bool`

Tests whether `x` is `naN`.  

### `j_0(x float) -> float`

Returns the order-zero Bessel function of the first kind.  

### `j_1(x float) -> float`

Returns the order-one Bessel function of the first kind.  

### `j_n(n int, x float) -> float`

Returns the order-n Bessel function of the first kind.  

### `ldexp(frac float, ex int) -> float`

The inverse of `frexp`: it returns `frac * 2^exp`.  

### `lgamma(x float) -> float, int`

Returns the natural logarithm and sign (`-1` or `+1`) of `gamma(x)`.  

### `log(x float) -> float`

Returns the natural logarithm of `x`.  

### `log_10(x float) -> float`

Returns the base-10 logarithm of `x`.  

### `log_1_p(x float) -> float`

Returns the natural logarithm of `1 + x`. It is more accurate than `log(1 + x)` when `x` is near zero.  

### `log_2(x float) -> float`

Returns the base-2 logarithm of `x`.  

### `log_b(x float) -> float`

Returns the base-2 exponent of `x`.  

### `max(x float, y float) -> float`

Returns the larger of `x` and `y`.  

### `min(x float, y float) -> float`

Returns the smaller of `x` and `y`.  

### `mod(x float, y float) -> float`

Returns the floating-point remainder of `x/y`. The magnitude of the result is less than `y` and its sign agrees with that of `x`.  

### `modf(x float) -> float, float`

Returns integer and fractional floating-point numbers that sum to `x`. Both values have the same sign as `x`.  

### `nextafter(x float, y float) -> float`

Returns the next representable float value after `x` towards `y`.  

### `pow(x float, y float) -> float`

Returns `x`^`y`.  

### `pow_10(n int) -> float`

Returns 10^`x`.  

### `remainder(x float, y float) -> float`

Returns the IEEE 754 floating-point remainder of `x/y`.  

### `round(x float) -> float`

Returns the nearest integer to `x` as a `float`, rounding o.5 away from zero.  

### `roundToEven(x float) -> float`

Returns the nearest integer to `x`, rounding ties to the even integer.  

### `signbit(x float) -> bool`

Tests whether `x` is negative, or negative zero  

### `sin(x float) -> float`

Returns the sin of `x`.  

### `sincos(x float) -> float, float`

Returns the sine and the cosine of `x`.  

### `sinh(x float) -> float`

Returns the hyperbolic sine of `x`.  

### `sqrt(x float) -> float`

Returns the square root of `x`.  

### `tan(x float) -> float`

Returns the tangent of `x`.  

### `tanh(x float) -> float`

Returns the hyperbolic tangent of `x`.  

### `trunc(x float) -> float`

Returns the integer part of `x`, as a float.  

### `y_0(x float) -> float`

Returns the order-zero Bessel function of the second kind.  

### `y_1(x float) -> float`

Returns the order-one Bessel function of the second kind.  

### `y_n(n int, x float) -> float`

Returns the order-`n` Bessel function of the second kind.  

