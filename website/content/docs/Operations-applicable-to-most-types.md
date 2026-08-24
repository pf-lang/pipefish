The following operators and functions all come as standard for all types except `error` and `func`.

* Comparison using `==` and `!=`.
* The `type` function (works with `func`).
* The `in` operator (works with `func`).
* The `literal` function.
* The `string` function.

Let's demonstrate them in the TUI:

```tui
→ 2 + 2 == 4
true
→ "foo" != "bar"
true
→ type true
bool
→ type bool
type
→ 99 in bool
false
→ string 5
5
→ string "foo"
foo
→ literal "foo"
"foo"
→
```

The `in` operator also works for membership of lists, sets, and tuples, as we will discuss when we introduce those types. It is an error to compare elements of different types. `"walrus" == 42` produces a compile-time error and not `false`.

All comparison in Pipefish is by-value: structs with identical fields are equal, lists with the same elements in the same order are equal, etc.