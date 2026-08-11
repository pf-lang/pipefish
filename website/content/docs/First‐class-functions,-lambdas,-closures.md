## First-class functions

A function is called "first-class" if it can be treated like a value: that is if it can be assigned to a variable, passed to a function, made the value in a key-value pair, etc. First-class functions are demonstrated by the file `examples/fcf.pf`. Here we have a simple function `applyTwice` which takes a function and applies it twice to a string.

```
def

applyTwice(f func, arg string) : 
    f f arg

makeFancy(s string) : "*!* " + s + " *!*"

decorator(decoration string) :
    func(x) : decoration + " " + x + " " + decoration
```

If we try it out in the REPL:

```
→ hub run "examples/fcf.pf" as "Fcf"                                   
Starting script 'examples/fcf.pf' as service 'Fcf'.
Fcf → applyTwice makeFancy, "hello"
*!* *!* hello *!* *!*       
```

This is straightforward enough. We pass a function and a string to the `applyTwice` function, and it applies the function twice to the string.

## Lambdas and closures

But what about the `decorator` function? Here we are taking a string, and returning an "anonymous function" or "lambda function". It has no name (we write the keyword `func` where the name would otherwise go): but it has a perfectly good definition. And so if we continue to toy with the REPL:

```
Fcf → applyTwice decorator("~~~"), "hello" 
~~~ ~~~ hello ~~~ ~~~
Fcf → applyTwice decorator("<*>"), "hello"
<*> <*> hello <*> <*>
Fcf →  
```

You will notice that the anonymous functions we're returning are also "closures". That is, without needing to be passed the value of `decoration`, each anonymous function on being created stores and retains that information, so that in effect calling `decorator("~~~")` returns a function `func(x) : "~~~" + " " + x + " " + "~~~"`.

Inner functions are also closures. That is, if we had written the `decorator` function as follows, then it would do exactly the same thing.

```
decorator(decoration string) :
    innerDecorator
given :
    innerDecorator(x) : decoration + " " + x + " " + decoration
```