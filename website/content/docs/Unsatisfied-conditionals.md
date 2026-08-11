In this page we will discuss why most conditional statements should end with an `else` clause, and the circumstances under which we might not do that.

This allows some useful idioms which are, so far as I know, unique to Pipefish.

## Why conditionals should usually end with `else`

There are two possibilities. Either you know for certain that the various possibilities in your conditional are exhaustive, in which case you might as well exploit that fact:

```
def

isEven(x int) :
    x % 2 == 0 :
        "even"
    else : 
        "odd"
```

... or you do *not* know that your conditional is exhaustive, in which case you should write an `else` statement returning a specialized error explaining where your particular function has gone off the rails. Leaving the `else` statement off and relying on Pipefish to return a generic error is usually bad Pipefish.

## Unsatisfied conditionals

We should now think about what exactly happens if you leave the `else` block off.

An example script has been provided in `examples/unsat.pf`. Here's the first function in the script.

```
def

isEven(x int) :
    x % 2 == 0 :
        "even"
```

Let's play with this function in the REPL and see what it does.

```
→ hub run "examples/unsat.pf" as "Unsat"                       
Starting script 'examples/unsat.pf' as service 'Unsat'.
Unsat → isEven 2                                               
even
Unsat → isEven 3 

[0] Error: unsatisfied conditional at line 1:0-6 of REPL input

Unsat → 
```

This error makes perfect sense. We have passed the value `3` to a function which doesn't say what to do if it gets an odd number.

But we can compose the function with something which _does_ say what to do:

```
Unsat → isEven 3 ; else : "odd"
odd
Unsat →
```

The `;`/`newline` operator _means_: "If the conditional on the left is unsatisfied, try the one on the right." And so when a function has an unsatisfied conditional, what it returns to the calling context is _flow of control_.

## Flow of control

This supplies us with some nice idioms, as illustrated by the last two functions in `examples/unsat.pf`:

```
safeNull(t tuple) :
    ([t] ?> that in null -> len) > 0 :
        NULL

safeAdd(a int?, b int?) :
    safeNull(a, b)
    else :
        a + b
```

The `safeNull` function will return `NULL` if one of its parameters is `NULL`. Otherwise its conditional will be unsatisfied, and so when called from `safeAdd` with two integer parameters, it will simply pass flow of control back to `safeAdd`, which will perform the addition.

``` 
Unsat → safeAdd(NULL, 7) 
NULL
Unsat → safeAdd(8, NULL) 
NULL
Unsat → safeAdd(8, 7)    
15
Unsat → 
```

The practical advantage of this is that the gatekeeping logic in `safeNull` can now be reused by any other function.