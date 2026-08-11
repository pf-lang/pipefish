## Introduction

In this page we will discuss how the `for` loops in Pipefish work and what you can do with them.

## Functional `for` loops

The `for` loops in Pipefish are based syntactically on its parent language Go, which is in turn based on C. For a variety of reasons, some good and some bad, most functional languages don't have C-like `for` loops. To make them work, we need to make some slight changes to the paradigm. Here is an example, a `for` loop which sums the elements of a list:

```
sum(L list) :
    from a = L[0] for i = 1; i < len L; i + 1 :
        a + L[i]
```

In an imperative language the equivalent loop would look like this.

```
sum(L list) :
    a := L[0]
    for i := 1; i < len L; i = i + 1 :
        a = a + L[i]
    return a
```

That is, we would start off by assigning values to mutable variables `a` and `i`. We would then reassign them every time we go around the loop (with the imperative statements `i = i + 1` and `a = a + L[i]`, and return the final value of `a`.

In the functional version, we can't and don't mutate anything, and there is no "final value of `a`". Instead, the `for` loop is an expression in which the `a` and `i` are *bound variables*, just like the `i` in the mathematical expression $` \sum_{i=0} L_i `$ (which means much the same thing). And the result is simply the final value *of the* `for` *expression* — `i` and `a` don't exist or have any meaning outside of the `for` loop.

What difference does this make? It means that we write our `for` loops in pure expressions rather than in terms of mutating variables. Let's look at the actual, functional version again:

```
sum(L list) :
    from a = L[0] for i = 1; i < len L; i + 1 :
        a + L[i]
```

The third part of the "header" of the `for` loop, the `i + 1`, *is an expression* that says what happens to the index variable `i` each time we go round the loop, and the body of the `for` loop *is an expression* that says what happens to the bound variable `a` each time we go round.

## Multiple bound variables

We can bind more than one variable. Here's an example of a Fibonacci function:

```
fib(n int) :
    from a, b = 0, 1 for i = 0; i < n; i + 1 :
        b, a + b
```

However, if you try this you will find that it returns a 2-tuple of numbers of which we are interested only in the first, e.g. `fib 6` will return `8, 13`. The ergonomic way to fix this is by using the built-in `first` function on the tuple returned by the `for` loop:

```
fib(n int) :
    first from a, b = 0, 1 for i = 0; i < n; i + 1 :
        b, a + b
```

It is also possible to have more than one index variable by the same sort of multiple assignment, although it's not clear why you would want to. More usefully, you can constrain the types of bound and index variables just like any other variables:

```
fib(n int) :
    first from a, b int = 0, 1 for i int = 0; i < n; i + 1 :
        b, a + b
```

## `break` and `continue`

Pipefish supplies you with `break` and `continue` statements. This function will search through a list `L` for a given element `x`, returning the index of `x` if it's present or `-1` if it isn't. 

```
find(x single?, L list) :
    from result = -1 for i = 0; i < len L; i + 1 :
        L[i] == x :
            break i
        else :
            continue
```

When the `break` statement takes an argument, as in the example above, this is what the loop returns; if not, it returns whatever the bound variable is when the `break` is encountered.

As with Go, we can use `for` with just the condition as a `while` loop, as in this implementation of the Collatz function, which will return `1` if (as we hope) the function terminates.

```
collatz(n int) :
    from x = n for x != 1 :
        x % 2 == 0 :
            x / 2
        else :
            3 * x + 1
```

... or with no condition at all as an infinite loop:

```
collatz(n int) :
    from x = n for :
        x == 1 :
            break
        x % 2 == 0 :
            x / 2
        else :
            3 * x + 1
```

## Using `range`

And we can likewise imitate the `range` form of Go's `for` loop, though we will use Pipefish's pair operator `::` to do so.
```
selectEvenIndexedElements(L list):
    from a = [] for i::x = range L :
        i % 2 == 0 :
            a + [x]
        else :
            continue
```

Just as in Go, we can use the data-eater symbol `_` to indicate that we don't want either the index or the value of the container. Let's rewrite the `sum` function from the top of the page:

```
sum(L list) :
    from a = L[0] for _::v = range L[1::len L] :
        a + v
```

You can range over lists, maps, sets, and strings. In the case of lists and strings, the index is an integer from 0 to one less than the length of the string, for maps it's the key of the map, and for sets the index and the value are the same thing, both ranging over the elements of the set, to save you having to remember which is which.

Finally, you can use a numerical range given as usual with the pair operator `::`. This will sum the numbers from and including `a` to and excluding `b`.

```
sumBetween(a, b) :
    from a = 0 for _::v = range a::b :
        a + v
```

The index in such a case is the numbers from and including `0` to and excluding `b`-`a`. If the first number in the given range is higher than the second, then the value counts down from and excluding the higher number to and including the lower number, while the index still counts up from `0`. So for example this will find if the given string is a palindrome:

```
palindrome(s string) :
    from result = true for i::j = range len(s)::0 :
        s[i] != s[j] : 
            break false 
        else : 
            continue
```

## The `given` block

Like a function or a lambda, a `for` loop can have a `given` block of local variables. For example, this converts integers to Roman numerals. The variables `textToUse` and `numberToUse` are local to the loop.

```
const

ROMAN_NUMERALS = ["M"::1000, "D"::500, "C"::100, "L"::50, "X"::10, "IX"::9, "V"::5, "IV"::4, "I"::1]

def 

string(i int) :
    first from result, number = "", i for number > 0 :
        result + textToUse, number - numberToUse
    given :
        textToUse, numberToUse = from t, n = "", -1 for _::p = range ROMAN_NUMERALS :
            p[1] <= number :
                break p[0], p[1]
            else :
                continue
```