## Overview 

The `rand` library supplies a random number generator. It is not cryptographically secure: for that, you should use the `crypto/rand` package instead.  

## Modules 

### `reflect ("/home/tim/Pipefish/source/initializer/libraries/reflect.pf")`

### ` ("math/rand")`
## Types 

### `Exponential = struct `

Data structure to pass to `get ... from` to get a random number with an exponential distribution.  

### `Normal = struct `

Data structure to pass to `get ... from` to get a random number with an exponential distribution.  

### `Random = struct (elements any?)`

Data structure to pass to `get ... from` to get a random value depending on the `elements` field; see the documentation for `get (x ref) from (r Random)` for further details.  

### `Seed = struct `

Data structure to pass to `put ... into` when setting the random seed.  

### `Shuffle = struct (elements any?)`

### `Zipf = struct (s float, v float, imax int)`
## Commands 

### `put(i int) into (seed Seed)`

Sets the random seed to `i`. If unset, the RNG is given a strong random seed on initialization.  

### `get(x ref) from (r Random)`

This populates the random value from the range given by the `elements` field of the `Random` value, according to the type of `elements`: 


- If `elements` is a list, it selects a random element of the list. E.g: `get x from Random ["one", "two", "three]`.
- If it's an integer `i`, it selects a random integer in `0::i`.
- If it's `int`, it selects a random integer from the whole 64-bit range of integers.
- If it's a float `x`, it selects a random float between 0 and `x`.
- If it's `float`, it selects a random float between 0 and 1.
- If it's `bool`, it selects randomly from `true` and `false`.
- If it's the name of an enum type, it selects a random element of the enum.
- If it's a pair of integers, it selects a random integer from the range.
- If it's a pair of floats, it selects a random float from the range.
 

### `get(x ref) from (s Shuffle)`

This populates the reference variable with a random permutation of the range given by the `elements` field of the `Shuffle` value, according to the type of `elements`: 


- If `elements` is a list, it returns a permutation of the list.
- If it's the name of an enum type, it returns a permutation of the elements of the enum.
- If it's a pair of integers, if returns a permutation of the integers in that range, e.g. `get x from Shuffle 0::10` will get a permutation of the numbers `0` ... `9`.
 

### `get(x ref) from (n Normal)`

This populates the reference variable with a random `float` with a normal distribution having mean 0 and standard deviation 1.  

### `get(x ref) from (e Exponential)`

This populates the reference variable with a random `float` greater than or equal to 0, with an exponential distribution and mean both equal to 1.  

### `get(x ref) from (z Zipf)`

This populates the reference variable with a `float` chosen from values k ∈ [0, `z[imax]`] such that P(k) is proportional to (`z[v]`+k)^(-`z[s]`). Note that in this one case, the top of the range (`z[imax]`) is among the possibilities.  

