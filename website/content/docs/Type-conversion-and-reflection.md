## Type conversion

Pipefish has a number of built-in functions for type conversion.

Usually they have the same name as the type they convert to:

`float` converts ints to float, or converts strings to floats if they have the right format.

`int` converts floats to integers by truncation, or converts strings to ints if they have the right format.

`set` will turn a tuple into a set.

`string` will convert anything to a string. This has been given some sensible defaults but can be overloaded for any type except `tuple`.

`literal` converts a value into a string which is the Pipefish literal of that value.

`tuple` will leave a tuple unchanged and turn anything else in to a tuple _of length 1_.

`L ...` will convert a list `L` into a tuple with the same elements.

To turn a tuple `t` into a list it is only necessary to wrap it in square brackets: `[t]`.

`type` converts a value to its type.

## Reflection

In the list above we briefly noted that `type` converts a value to its type. That is, `type "walrus"` evaluates to `string`, `type 42` evaluates to `int`, etc.

`string` and `int` are themselves members of the type `type`. That is, `type string` returns `type`.

`type` is a perfectly normal Pipefish type itself, and belongs, of course, to type `type`. Try not to let this bother you.

We can also use the `in` keyword to check for type membership: so for example `"walrus" in string` and `string in type` are true.