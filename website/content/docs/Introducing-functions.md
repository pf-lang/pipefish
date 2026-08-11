In this page we will look at the basics of functions.

## Examples

Here are some examples, in `examples/functions.pf`.

```
def

square(x) : x * x

parity(x) : 
    x mod 2 == 0 : 
        "even"
    else : 
        "odd"

swap(x, y) : y, x
```

## Basic syntax: defining functions

Functions are defined in the `def` section of the script.

The body of a Pipefish function consists of a single expression to be evaluated. There is no `return` keyword because all a Pipefish function ever does is return something, so it would be superfluous to have a keyword to point this out.

In Pipefish, a `:` precedes the body of a code block, which may be on the same line as the `:` so long as it fits into one line. Otherwise the block must also be marked off by indentation, as in the `parity` function.

Parentheses must be used to identify the parameters of a function.

## Basic syntax: calling functions

Parentheses are not required when calling functions, except as necessary to establish precedence. Let's demonstrate in the REPL:

```
→ hub run "examples/functions.pf" as "Fun"                             
Starting script 'examples/functions.pf' as service 'Fun'.
Fun → parity 5
odd
Fun → swap "walrus", 42 
42, walrus  
Fun → square 3 + 4
49
Fun → (square 3) + 4 
13
Fun → square(3) + 4                                                    
13
Fun →
``` 

## Conditionals

We have briefly met Pipefish's conditional operator, the example given being `2 + 2 == 3 : "Oops" ; 2 + 2 == 4 : "Whew!"; else : "Who even knows?"`, which evaluates to `"Whew!"`.

In the body of a function, the `;` can be replaced by a newline. So the following two functions are equivalent:

```
parity (x) : 
    x % 0 == 0 : 
        "even"
    else : 
        "odd"
```

```
parity (x) : 
    x % 0 == 0 : "even" ; else : "odd"
```

But the first is easier to read.

This becomes still more useful when we wish to nest conditionals, as demonstrated in `examples/conditionals.pf`:

```
classify (i int) :
    i > 0 :
        i < 10 : "small number"
        else : "big number"
    i == 0 :
        "zero"
    else :
        i > -10 : "small negative number"
        else : "big negative number"
```

Usually a conditional should have an `else` as the final condition, but there is one case in which it is useful to omit it, which we will discuss [later](https://github.com/tim-hardcastle/Pipefish/wiki/Unsatisfied-conditionals).