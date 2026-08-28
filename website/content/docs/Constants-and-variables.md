Constants can be defined under the heading `const`. The Pipefish style guide recommends that you use `SCREAMING_SNAKE_CASE` to name them.

Variables can be declared under the heading `var`.

An example is given in `examples/docs/variables.pf`:

```
var

h = "Hello world!"
x = MONTHS_IN_A_YEAR * 2

const

MONTHS_IN_A_YEAR = 12
```

Constants and the initial values of variables can be computed at compile time, as `x` is in the example above.

You can change the values of variables via the REPL, but you can't create them: they must be declared in the script.

Constants are in scope for all commands and functions. Variables can only be seen by commands.

In Pipefish, variables by default take on the type of the thing assigned to them on declaration. So `x` is of type `int`, and `h` is of type `string`, and trying to store anything else in them causes an error. This is only the default behavior — Pipefish is after all a dynamic language! We'll come back to this when we look at the type system.