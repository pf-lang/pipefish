Here is a diagram of the type system. It contains a few types we haven't met yet, such as `func`.

It does *not* show the [nullable twin](https://github.com/tim-hardcastle/Pipefish/wiki/Abstract-types#nullable-twins) of each standard type, with the exception of `any?` because that would complicate the diagram to little purpose.

The `error` type is shown off to one side because it doesn't belong anywhere in the type hierarchy: not belonging is what it's for.

```mermaid
graph TD;
    any?-->any
    any?-->null
    any-->intlike;
    intlike-->int
    intlike-->I[clones of int];
    any-->floatlike;
    floatlike-->float
    floatlike-->F[clones of float];
    any-->stringlike;
    stringlike-->string
    string-->V[varchar];
    stringlike-->S[clones of string]
    any-->bool;
    any-->struct;
    any-->listlike;
    listlike-->list
    listlike-->L[clones of list]
    any-->pairlike;
    pairlike-->pair
    pairlike-->P[clones of pair]
    any-->label;
    any-->setlike;
    setlike-->set
    setlike-->Z[clones of set]
    any-->maplike
    maplike-->M[map];
    maplike-->N[clones of map]
    any-->func;
    any-->type;
    struct-->X[user-defined structs];
    any-->enum
    enum-->Y[user-defined enums];
    error
```