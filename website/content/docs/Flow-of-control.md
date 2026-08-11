Pipefish commands have flow of control in a way that Pipefish _functions_ don't and can't. This is demonstrated in the file `examples/flow.pf`.

Let's look at the commands in it one at a time.

## Commands are sequences of instructions

Unlike a Pipefish function, which just evaluates things, Pipefish commands *do* things, and so can and indeed must do them one after another.

```
seq :
    post "It's just one thing ..." to Output()
    post "... after another." to Output()
```

## Conditionals

And so this means that in Imperative Pipefish one conditional can follow on another:

```
check :
    5 % 2 == 0 :
        post "5 is even" to Output()
    else :
        post "5 is odd" to Output()
    6 % 2 == 0 :
        post "6 is even" to Output()
    else :
        post "6 is odd" to Output()
```

It is recommended that every conditional block should end with an `else` clause. First, for consistency and second because when one conditional block follows on another like this, Pipefish identifies the end of the first conditional block with the `else` clause.

## `OK`

If you want one branch of a conditional to do nothing at all, then you should tell it to return `OK`: e.g. the following rewrite of the `check` command above will do nothing unless 5 becomes even or 6 becomes odd:

```
check :
    5 % 2 == 0 :
        post "5 is even" to Output()
    else :
        OK
    6 % 2 == 0 :
        OK
    else :
        post "6 is odd" to Output()
```

## `for` loops

In Imperative Pipefish a `for` loop can contain statements rather than expressions.

```
cmd

hello(n) : 
    for i = 0; i < n; i + 1 : 
        post "Hello!"
```

Note that this particular `for` loop neither has nor needs bound variables, although they can if you wish, when needed. As the loop body is imperative, it will either return `OK`, which you needn't concern yourself with; or an error, which you can deal with as explained in the page on [imperative error-handling](https://github.com/tim-hardcastle/Pipefish/wiki/Imperative-error-handling).