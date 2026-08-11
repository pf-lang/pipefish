Constants can be defined under the heading `const`. The Pipefish style guide recommends that you use `SCREAMING_SNAKE_CASE` to name them.

Variables can be declared under the heading `var`.

An example is given in `examples/wiki/variables.pf`

<img width="710" height="459" alt="image" src="https://github.com/user-attachments/assets/42738cba-4abc-488a-ac37-6f09a225e64a" />

Constants and the initial values of variables can be computed at compile time, as `x` is in the example above.

You can change the values of variables via the REPL, but you can't create them: they must be declared in the script.

Constants are in scope for all commands and functions. Variables can only be seen by commands, and even then you *must* bring them into the scope of the command using the `global` keyword. E.g if we add this to the script above, it will do what you think it would do.

<img width="710" height="135" alt="image" src="https://github.com/user-attachments/assets/2ea1ea9f-115d-443c-80c0-ecf6b4067259" />

In Pipefish, variables by default take on the type of the thing assigned to them on declaration. So `x` is of type `int`, and `h` is of type `string`, and trying to store anything else in them causes an error. This is only the default behavior — Pipefish is after all a dynamic language! We'll come back to this when we look at the type system.