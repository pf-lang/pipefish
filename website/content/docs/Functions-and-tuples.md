## Introduction

The tuple holds a special place in the type system. As we have seen, it is constructed simply by concatenating values with commas: `1, true, "walrus"` is a tuple. We can surround it with parentheses: `(1, true, "walrus")`, or we can emphasize the fact that it is a tuple with the `tuple` keyword and write `tuple(1, true "walrus")`, but these all refer to the same thing.

The purpose of the tuple type is to be completely invisible and something that you can completely take for granted. Nonetheless, we should still look at why it exists and how it works.

## Tuples as function arguments

When passed to a function, tuples split up into their component parts.

If this wasn't true, it wouldn't be possible to pass multiple values to a function at all! Consider the following function, one of several example functions defined in `examples/tuples.pf`

```
swap (x, y) : y, x
```

When we call (e.g.) `swap 1, 2`, then from the definition of a tuple above, we are in fact passing it the tuple `1, 2`. But tuples break up into their component parts when you pass them to functions, and so the `1` gets assigned to `x` and the `2` gets assigned to `y`.

## Tuples as function parameters

You can prevent a tuple from exploding itself by using `tuple` as a parameter:

```
rotateLeft(t tuple) : 
    t == () : 
        ()
    else : 
        t[1::len t] , t[0]
```

## Tuples as function return values

Again, the very definition of a tuple means that if you return more than one value from a function, you are in fact *still* returning one value, namely one tuple! But when you put that together with the way tuples deconstruct themselves when passed to a function, you find that 99% of the time the net result is that you don't have to think about tuples or know that they exist.

The one case where you *do* have to remember that tuples are a type is when you've written a function to return multiple values and then you find a circumstance where you want to call the function but then consume just one or two of those values rather than all of them. At this point it's useful to remember that your multiple return values are in fact a tuple, and to start indexing and slicing it.

## Varargs

We can also collect single values into tuples (including single values made by tuples exploding themselves):

```
max(args ... int) : 
    from result = args[0] for _::v = range args[1::len args] : 
        v > result : 
            v 
        else : 
            continue 
```