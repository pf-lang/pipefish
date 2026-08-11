We have [already seen](https://github.com/tim-hardcastle/Pipefish/wiki/Interface-types,-part-1) how you can use an interface declaration to define an abstract type by the functions that can be called on it, or that can return it.

But the very existence of an interface also sprinkles a little magic on the multiple dispatch. Readers who have used languages with "typeclasses" should find themselves n familiar ground here: for those who haven't, let's demonstrate with a small example.

First, let's make a little file `complex.pf` supplying us with a complex integer type which we equip with a couple of functions, `+` and `rotate`.

```
newtype

C = struct(real, imaginary int)

def

(w C) + (z C) -> C :
    C(w[real] + z[real], w[imaginary] + z[imaginary])

rotate(z C) -> C :
    C(-z[imaginary], z[real])
```

Then the `C` type will be in `Addable`. Now let's add to the script an import of a library `summer.pf` which among other things contains the following function to sum a list:

```
sum(L list) :
    from a = L[0] for _::v = range L[1::len(L)} :
        a + v
```

Try out our modified program in the REPL:

```
→ hub run "examples/complex.pf"
Starting script 'examples/complex.pf' as service '#0'.

#0 → summer.sum [1, 2, 3]
6
#0 → summer.sum [C(1, 2), C(3, 4), C(5, 6)]
C with (real::9, imaginary::12)
#0 →
```

Note that the summer.sum function can't refer to the `C` type, nor construct an instance of it. How could it? It can't *see* it --- `complex` imports `summer` and not vice versa.

But it _can_ correctly dispatch on it, because the `summer` module _does_ know that `C` belongs to the `Addable` type.

So if we were to add to the `summer` library a function like this one ...

```
rotAll(L list) :
    from a = [] for _::v = range L :
        a + [rotate v]
```

... then this would fail to compile, complaining that `rotate` is an unknown identifier. We would also need to add an interface to `summer` like this:

```
Rotatable = interface :
    rotate(x self) -> self
```

... after which `rotAll` will work just fine.

## The small print

This works on one condition that I haven't yet mentioned. The `+` operation is defined _in the same namespace_ as the `C` type, if not for which while the operation would work as such it would not mean that `C` was `Addable`, and functions like `summer.sum` wouldn't know how to dispatch `+`.

By using a `NULL` import namespace, you can wrap a type you don't own in a function it lacks, e.g. if you couldn't change the code in `complex.pf` but wanted it to have subtraction as well, this would work:

```
import

NULL::"namespace/complex.pf"

def

(w C) - (z C) -> C :
    C(w[real] - z[real], w[imaginary] - z[imaginary])
```