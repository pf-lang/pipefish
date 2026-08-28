## Abstract versus concrete types

Every specific Pipefish value has a *concrete type*: `int`, `bool`, `string`, a user-defined struct type, etc. We can instantiate any concrete type with literals (`42`, `"foo"`, `false`) or with constructors as with structs, enums, and clone types. Concrete types can never have subtypes.

An *abstract type* is just a union of concrete types, e.g `int/float`. Such types clearly do have subtypes, in this case `int` and `float`, but they can never be instantiated: no value can ever be of type `int/float`, it has to be one or the other.

## Declaring abstract types

While we can create and use abstract types  “on the fly” by using the `/` operator, we can also give them names in the `newtype` section, which is to be preferred if you're going to use the same abstract type more than once.

```
newtype

Number = abstract int/float

// We can then use this to define other types:
Widget = struct(height, width Number)

var // Or when declaring variables:

myNumber Number = 42

def // Or the parameters of functions:

double(n Number) :
    2 * n
```

Note that by default, variables are given the type of the value they're initialized with, so if we just wrote `myNumber = 42`, then myNumber could only ever contain an `int`.

## The `?` operator and `NULL`

As syntactic sugar, you can write e.g. `int?` for `int/null`, etc.

**Note** that you can't compare a non-null value with `NULL`, because you can't compare values of different types. So if you want to know if a variable `foo` of type `int?` is equal to `NULL`, you can't write `foo == NULL` (because that would cause an error at runtime if `foo` *wasn't* `NULL`) but instead must use the expression `foo in null` to test its type membership.

## Built-in abstract types

The abstract type `any` contains all types that aren't `null` or `tuple`. `struct` contains all struct values, `enum` contains all enum values, and `label` contains all field labels of structs.

For any clonable type, e.g. `rune`, the abstract type `clones{rune}` contains the parent type and all its clones.

There are abstract types defined my interfaces, such as `Addable`, which will be discussed in the next section of this documentation.

## Note on semantics

While concrete types are *nominal*, abstract types are *structural*: two abstract types are equal just if they are the union of the same concrete types. So in our example above, `Number` is equal to `int/float` and `float/int`. Or for example if you haven't defined any clones of `rune`, then `clones{rune}` is equal to `rune`.