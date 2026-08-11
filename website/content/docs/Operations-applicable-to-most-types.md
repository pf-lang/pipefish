The following operators and functions all come as standard for all types except `error` and `func`.

* Comparison using `==` and `!=`.
* The `type` function (works with `func`).
* The `in` operator (works with `func`).
* The `literal` function.
* The `string` function.

Let's demonstrate them in the TUI:

<img width="710" height="262" alt="image" src="https://github.com/user-attachments/assets/eb070835-e903-497d-8af2-f229e2dfc647" />

The `in` operator also works for membership of lists, sets, and tuples, as we will discuss when we introduce those types. It is an error to compare elements of different types. `"walrus" == 42` gives an error and not `false`.

All comparison in Pipefish is by-value: structs with identical fields are equal, lists with the same elements in the same order are equal, etc.