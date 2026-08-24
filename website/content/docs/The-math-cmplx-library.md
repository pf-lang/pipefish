## Overview 

The `cmplx` library provides basic constants and mathematical functions for complex numbers. Special case handling conforms to the C99 standard Annex G IEC 60559- compatible complex arithmetic.  

## Modules 

### ` ("math/cmplx")`

### ` ("math")`

### ` ("strconv")`
## Types 

### `Complex = wrapper complex128`

The complex number type.  
## Functions 

### `Complex(a float, b float) -> Complex`

Constructs a `Complex` value from two `float` values representing its real and imaginary parts.  

### `(x Complex) + (y Complex) -> Complex`

Adds two `Complex` values.  

### `(x Complex) - (y Complex) -> Complex`

Subtracts one `Complex` value from another.  

### `(x Complex) * (y Complex) -> Complex`

Multiplies two `Complex` values.  

### `(x Complex) / (y Complex) -> Complex`

Divides one `Complex` value by another.  

### `(x Complex) ^ (y Complex) -> Complex`

Exponentiates one `Complex` value by another.  

### `(x Complex) ^ (y float) -> Complex`

Exponentiates a `Complex` value by a `float`.  

### `abs(x Complex) -> float`

Returns the modulus of `x`.  

### `acos(x Complex) -> Complex`

Returns the inverse cosine of `x`.  

### `acosh(x Complex) -> Complex`

Returns the inverse hyperbolic cosine of `x`.  

### `asin(x Complex) -> Complex`

Returns the inverse sine of `x`.  

### `asinh(x Complex) -> Complex`

Returns the inverse hyperbolic sine of `x`.  

### `atan(x Complex) -> Complex`

Returns the inverse tangent of `x`.  

### `atanh(x Complex) -> Complex`

Returns the inverse hyperbolic tangent of `x`.  

### `conj(x Complex) -> Complex`

Returns the complex conjugate of `x`.  

### `cos(x Complex) -> Complex`

Returns the cosine of `x`.  

### `cosh(x Complex) -> Complex`

Returns the hyberbolic cosine of `x`.  

### `cot(x Complex) -> Complex`

Returns the cotangent of `x`.  

### `exp(x Complex) -> Complex`

Returns the exponent of `x`.  

### `imag(x Complex) -> float`

Returns the imaginary part of `x`.  

### `inf -> Complex`

Returns infinity.  

### `isInf(x Complex) -> bool`

Tests whether `x` is infinity.  

### `isNaN(x Complex) -> bool`

Tests whether `x` is `naN`.  

### `log(x Complex) -> Complex`

Returns the natural logarithm of `x`.  

### `log_10(x Complex) -> Complex`

Returns the logarithm to the base 10 of `x`.  

### `naN -> Complex`

Returns `naN`  

### `phase(x Complex) -> float`

Returns the phase of `x`.  

### `polar(x Complex) -> float, float`

Returns the polar co-ordinates of `x`.  

### `real(x Complex) -> float`

Returns the real part of `x`.  

### `rect(r float, θ float) -> Complex`

Converts polar co-ordinates into a `Complex` value.  

### `sin(x Complex) -> Complex`

Returns the sine of `x`.  

### `sinh(x Complex) -> Complex`

Returns the hyberbolic since of `x`.  

### `sqrt(x Complex) -> Complex`

Returns a square root of `x`.  

### `string(z Complex) -> string`

Represents `z` as a string.  

### `tan(x Complex) -> Complex`

Returns the tangent of `x`.  

### `tanh(x Complex) -> Complex`

Returns the hyberbolic tangent of `x`.  

