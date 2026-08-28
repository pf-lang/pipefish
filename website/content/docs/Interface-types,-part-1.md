## What is an interface?

An interface type is an abstract type defined not just by arbitrarily joining types together with `/`, but by saying that it is the union of all the concrete types with a given function or collection of functions defined on them.

For example, there is a built-in `Addable` type:

```
newtype

Addable = interface :
    (x self) + (y self) -> self
```

This type includes every type with an operation which adds a value of that type to another value of that type and returns a values of the same type. So it contains at least `int`, `float`, `list`, `string` and `set`, and then whatever other types you decide to define addition on. 

Note that if we had e.g. a type `Thing` and we defined a function `(x Thing) + (y Thing)` that would *not* qualify it as an `Addable`, because we left the return type off; this will be true even if it's possible to infer that the result of the addition must in fact be a `Thing`.

You can also define your own interface types as you please:

```
Fooable = inteface :
    foo(x self) -> int
    bar(x rune, y self)
```

Besides merely defining abstract types they play another role, but we will discuss this on a later page, after we have introduced modules.

## The built-in interfaces

The built-in interfaces are as follows:

```
newtype

Addable = interface :
    (x self) + (y self) -> self

Ltable = interface :
    (x self) < (y self) -> bool

Lteable = interface :
    (x self) <= (y self) -> bool

Gtable = interface :
    (x self) > (y self) -> bool

Gteable = interface :
    (x self) >= (y self) -> bool

Comparable = interface :
    (x self) < (y self) -> bool
    (x self) <= (y self) -> bool
    (x self) > (y self) -> bool
    (x self) >= (y self) -> bool

Lenable = interface :
    len(x self) -> int 

Multipliable = interface :
    (x self) * (y self) -> self

Negatable = interface :
    -(x self) -> self 

Stringable = interface :
    string(x self) -> string

Subtractable = interface :
    (x self) - (y self) -> self
```

## Naming conventions

As usual, user-defined types should be defined in `PascalCase`.

The `-able` suffix is very natural: we are defining types by saying what can be done to them. However, the style guide doesn't make it mandatory, and if there is some more natural name for the type, then you should feel free to use it.