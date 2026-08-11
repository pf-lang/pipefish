## Function signatures

So far we haven't given types to the parameters of our functions. But we can if we want to! For example this function is constrained only to accept strings:

```
double(x string) : 
    x + x
```

Whereas this will only accept integers:

```
mul3(x, y, z int) : 
    x * y * z
```

And this will accept a float and a bool:

```
conditionallyInvert(x float64, c bool):
    c : - x
    else : x
```

It is also possible to specify the return type of the function, if you want to make absolutely sure it's doing what it's supposed to. So we could write the previous function as:

```
conditionallyInvert(x float, c bool) -> float :
    c : - x
    else : x
```

## Overloading

We can overload functions on the types of their calling parameters. An example is given in `examples/overloading.pf`:

```
def

twice(x string) : x + ", " + x

twice(x int) : 2 * x

twice(b bool) :
    b : "That's as true as things get!"
    else : "That's as false as things get!"

twice(t single) :
    error "I don't know how to double that!"

twice(t tuple) :
	t, t
```

Functions with more specific type signatures (lower down in the [type system diagram](https://github.com/tim-hardcastle/Pipefish/wiki/Type-system-diagram)) take precedence over less specific signatures. So if we pass a `bool`, we expect this to be handled by the third of the functions above, not the fourth or the fifth. We can demonstrate this, as usual, in the REPL:

```
→ hub run "examples/overloading.pf" as "OverL"
Starting script 'examples/overloading.pf' as service 'OverL'.
OverL → twice "hello"
hello, hello
OverL → twice 3 
6
OverL → twice true 
That's as true as things get!
OverL → twice NULL           

[0] Runtime error: I don't know how to double that!

OverL → twice 3, 4
3, 4, 3, 4
OverL →   
```

## Overloading built-in functions and operators

We can do this too! In the script `examples/money.pf` we overload the built-in `string` function and the `+` operator.

```
def

Money = struct(dollars, cents int)

ONE_DOLLAR = Money(1, 0)
TREE_FIDDY = Money(3, 50)

(m Money) + (n Money) :
    m[cents] + n[cents] >= 100 :
        Money(m[dollars] + n[dollars] + 1, m[cents] + n[cents] - 100)
    else :
        Money(m[dollars] + n[dollars], m[cents] + n[cents])

string(m Money) :
    "$" + string(m[dollars]) + "." .. 
     .. + (m[cents] < 10 : "0"; else : "") ..
     .. + string(m[cents])
```

If we run this in the REPL it works as you would hope.

```
→ hub run "examples/money.pf" as "Money"
Starting script 'examples/money.pf' as service 'Money'.
Money → string ONE_DOLLAR
$1.00
Money → string TREE_FIDDY 
$3.50
Money → string (ONE_DOLLAR + TREE_FIDDY) 
$4.50
Money → 
```

As with all the flexibility in Pipefish, this should be used with caution. You could, just the same way, overload the `+` operator so that it adds integers to strings like it was JavaScript. But there is a reason why this was left out of the language in the first place.