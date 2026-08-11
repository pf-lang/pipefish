## Enums

A Pipefish enum is a very simple thing. The script `examples/wiki/enums.pf` shows how to declare enums:

<img width="710" height="159" alt="image" src="https://github.com/user-attachments/assets/377535e5-0a24-444e-8576-5e69e8c53a03" />

These have only the most basic built-in functions, plus each enum type has a constructor which will make an element of the type out of an integer. We demonstrate this as usual in the TUI:

<img width="710" height="236" alt="image" src="https://github.com/user-attachments/assets/3d40525b-fff8-4f8d-b85c-331606ef5874" />

## General note on type declaration

So far we have seen how to declare structs and enums. Let's look at what they have in common.

* They both appear under the `newtype` headword.
* The declarations are parallel in form: struct declarations begin with `TypeName = struct`; enum declarations begin with `TypeName = enum`.
* The types are conventionally named in `PascalCase`.
* Each struct or enum type declared automatically has a constructor with the same name as the type. This is an ordinary function and can be overloaded like other functions.

This pattern will be repeated with other type declarations.