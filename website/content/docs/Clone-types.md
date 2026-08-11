Besides the other user-defined types, it can be useful to have types which resemble, but are distinct from, the built-in types. If, for example, we use an integer to represent a UID, then it is convenient to declare a UID type which can be distinguished from other integers. Or it may be convenient to have an `Apples` type which can't be added to `Oranges`, or a vector type which can't be mistaken for a list:

<img width="710" height="139" alt="image" src="https://github.com/user-attachments/assets/4e99c42a-a35c-443a-90b8-a9b28606e041" />

We should explain about the `using` clause. Some operations, such as `<` for clones of integers or `len` for clones of lists, are supplied automatically when you declare the type, and work just the same as for the parent type. Others, such as `+` for integers or slicing for lists, must be explicitly requested in a `using` clause.

The underlying rule is that an operation must be requested if it would be expected to return a value in the clone type. Hence the operations that need requesting are `+`, `-`, `*`, `/`, `div`, `mod`,`with`, `without`, the operators `>>` and `?>` for lists, and slicing clones of strings and lists. (Request the slicing operation with the word `slice`, the other operations by their names or symbols.)

An operation which you don't request in the `using` clause may still be overloaded by hand, if you don't want the operation to work like it does on the parent type. Here for example is how we'd implement addition for the `Vec` type. Don't worry if you don't quite understand it: it uses language features you haven't met yet, and is here just to show that you can do so when you need to.

<img width="710" height="174" alt="image" src="https://github.com/user-attachments/assets/3904fd7d-e408-4bdc-ac5c-dc3ba8a38b01" />

As with structs and enums, each clone type has a constructor function with the same name as the type.

Let's demonstrate all this in the TUI.

<img width="710" height="179" alt="image" src="https://github.com/user-attachments/assets/12e2d412-1839-4611-8a67-a218c09f5a5e" />