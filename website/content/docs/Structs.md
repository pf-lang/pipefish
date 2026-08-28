In this page we will discuss the `struct` type, its constructors, and its operations.

## Defining and creating structs

As an example, consider the file `examples/docs/structs.pf`.

```
newtype // We define some struct types.

Person = struct(name string, age int)
Cat = struct(name string, nobelPrizes int, pink bool)

var // And make some struct values.

doug = Person("Douglas", 42)
joe = Person with name::"Joseph", age::22
tom = Cat with pink::false, name::"Tom", nobelPrizes::2
```

The struct type is defined in the `newtype` section of the script, as for example `Person = struct(name string, age int)`. This creates the `Person` type and specifies the types of its fields. It also supplies us with a "short-form constructor" e.g. `Person("Douglas", 42)`, which is just an ordinary function, and a "long-form constructor" e.g `Person with name::"Joseph", age::22`, where the fields come in any order. As we will see later on, functions can be overloaded, and this includes the short-form constructor.

Struct types and indeed all user-defined types should be named in `PascalCase`.

## Labels

As you can see, structs are indexed with square brackets, like everything else in Pipefish.

The things you index them by (`name`, `age`, `pink`) etc are first-class values, of type `label`.

The expressions that follow `with` in the long-form constructors, e.g. `Person with name::"Joseph", age::22` are therefore also first-class values: `name::"Joseph", age::22` is a tuple of pairs, where the `[0]`th element of each pair is a label.