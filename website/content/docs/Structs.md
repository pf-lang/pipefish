In this page we will discuss the `struct` type, its constructors, and its operations.

## Defining and creating structs

As an example, consider the first few lines of the file `examples/wiki/structs.pf`.

<img width="710" height="450" alt="image" src="https://github.com/user-attachments/assets/23f18150-ee5d-4120-8b88-2666eee29680" />

The struct type is defined in the `newtype` section of the script, as for example `Person = struct(name string, age int)`. This creates the `Person` type and specifies the types of its fields. It also supplies us with a "short-form constructor" e.g. `Person("Douglas", 42)`, which is just an ordinary function, and a "long-form constructor" e.g `Person with name::"Joseph", age::22`, where the fields come in any order. As we will see later on, functions can be overloaded, and this includes the short-form constructor.

Struct types should be named in `PascalCase`.

## Labels

As you can see, structs are indexed with square brackets, like everything else in Pipefish. This means that often you can treat them syntactically and semantically as though they were container types.

The things you index them by (`name`, `age`, `pink`) etc are first-class values, of type `label`.

The expressions that follow `with` in the long-form constructors, e.g. `Person with name::"Joseph", age::22` are therefore also first-class values: `name::"Joseph", age::22` is a tuple of pairs, where the `[0]`th element of each pair is a label.

We will use the remainder of the `examples/wiki/structs.pf` file, and the TUI, to demonstrate this.

<img width="710" height="479" alt="image" src="https://github.com/user-attachments/assets/197d1d73-5a0f-40ba-bd0a-62738566c2a5" />
