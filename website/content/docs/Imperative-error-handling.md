Pipefish commands don't return values. But everything can return an error! The way Imperative Pipefish works with this is demonstrated in `examples/try.pf`.

## `try`

The `try` keyword acts as a condition that succeeds if an error is not returned:

```
tryWithoutCapture :
    try :
        get x from File("non-existent file")
    else :
        post "Well, that didn't work."
```

## `try` with capture

If `try` is followed by the name of a variable, and an error value is returned, then the error will be assigned to that variable.

```
tryWithCapture :
    try e :
        get x from File("non-existent file")
    else :
        post "Well, that didn't work. Error was: " + e[errorMessage]
```

**Note** that the error can only be assigned to a local variable: it is never possible to assign an error value to a global variable by this or any other means.

## `try` as a conditional

In other respects `try` works like an ordinary condition. For example, it doesn't have to be the first condition in its block:

```
dividePotentialNulls(a, b int?) :
    a in null or b in null :
        post NULL to Output()
    try :
        post a / b to Output()
    else :
        error "tried to divide by zero"
```