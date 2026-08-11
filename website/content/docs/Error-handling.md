## Introduction

In this page we will discuss what an error is and how to handle them.

## What is an error?

A runtime error is a first-class object. When a function/operation returns an error, it literally returns an error:

```
→ 1 / 0           
                                                    
[0] Error: division by zero at line 1 of REPL input.

→ 
```

With a few exceptions (pun unavoidable) when a function/operator is passed an error, it automatically returns that error:

```
→ 1 / 0                                                               

[0] Error: division by zero at line 1 of REPL input.

→ 1 + 1 / 0                                                           

[0] Error: division by zero at line 1 of REPL input.

string 1 + 1 / 0                                                    

[0] Error: division by zero at line 1 of REPL input.

→ 
```

An uncaught error will therefore propagate upwards through the calling functions until it is returned to the REPL.

## Creating errors

You can create your own errors by type conversion from `string` to `error`, and they work the same way:

```
error "this is my error"       
                                     
[0] Error: this is my error at line 1 of REPL input.

1 + error "this is my error"             
                           
[0] Error: this is my error at line 1 of REPL input.

→    
```

Note that the rule that a function or operator applied to an error yields that same error applies also to the `,` operator, the comma. This means that even in a function which normally returns multiple return values, you can’t return an error and another value.

## Handling errors

But there are things you can do to stop an error from propagating.

- You can use the `valid` function, which returns `false` if the value passed is an error and `true` otherwise.
- You can use the `unwrap` function on something you know is an `error`, turning it into an `Error`, a perfectly normal struct with field labels `errorCode` and `errorText`.
- You can assign it to a *local* variable.

Let’s have a look at `examples/error.pf`. This shows four possible implementations of a true `mod` function (where the result is never negative) to supplement the `%` operator.

```
def

// We could propagate the error thrown by %.

(x int) modA (y int) :
    remainder >= 0 : remainder
    else : remainder + y
given :
    remainder = x % y

// We could anticipate the error and throw our own.

(x int) modB (y int) :
    y == 0 : error "taking the modulus of a number by zero"
    remainder >= 0 : remainder
    else : remainder + y
given :
    remainder = x % y

// We could catch the error and throw our own.

(x int) modC (y int) :
    valid remainder : 
        remainder >= 0 : remainder
        else : remainder + y
    else : 
        error "taking the modulus of a number by zero"
given :
    remainder = x % y

// We could catch the error and return something other than an error.

(x int) modD (y int) :
    valid remainder : 
        remainder >= 0 : remainder
        else : remainder + y
    else : 
        "this isn't an error, just a friendly warning"
given :
    remainder = x % y

// ... etc, etc.
```