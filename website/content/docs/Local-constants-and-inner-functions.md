Pipefish functions allow and encourage the declaration of local variables and local functions (usually called "inner functions").

These are defined in a block at the bottom of the function. This is, at least superficially, one of Pipefish's most distinctive features. The idea behind it is that if you give your local constants and inner functions sufficiently verbose and descriptive names, then people will largely be able to read your code without looking at the definitions, which can serve as footnotes to your code from a human perspective as well as from the point of view of the interpreter.

## Local variables

An example is found in `examples/given.pf` : 

```
def

g(x) : 
    a * b * x
given : 
    a = x + 1
    b = x + 2
```

The meaning is, we hope, obvious: it's an imitation of how mathematicians talk, but with syntactic whitespace.

**NOTE 1** : local variables aren't calculated unless they turn out to be needed. So if your function has a lot of conditional branches which might require the calculation of different local variables, you can put them all into the `given` block secure in the knowledge that you're not wasting time computing the ones that aren't needed.

**NOTE 2** : despite the name, local variables don't actually vary, at least not during the course of one function call.

## Inner functions

In the same way, we can define inner functions. These can do everything normal functions can do, except that (a) they can't be overloaded (b) they must have plain old "vanilla" function syntax.

```
def

happyAddition(a, b) : 
    makeHappyString a + b
given : 
    makeHappyString(i) : "*!* " + (string i) + " *!*"
```

Inner functions can have their own `given` blocks. This will often be a bad idea, as illustrated by this example:

```
def

f(x) : a * b(x)
given : 
    a = x + 1
    b(x) : x + c(x)
    given :
        c(x): m * x
        given m = 5
```

Just because you can, doesn't mean you should. The ability to define inner functions is a convenience to help you organize your code: when it becomes an inconvenience, you should stop.