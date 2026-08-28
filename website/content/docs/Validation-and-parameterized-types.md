## Validation

Struct and clone types can have runtime validation attached to their type constructors.

```
newtype
    
Person = struct(name string, age int) :
    that[name] != ""
    that[age] >= 0
        
EvenNumber = clone int using +, -, * :
    that mod 2 == 0
```

## Parameterized types

We can parameterize this validation. For example, let's make some math-style vectors and some operators to go with them.

```
newtype
    
Vec{i int} = clone list : 
    len(that) == i
    
def
    
(v Vec{i int}) + (w Vec{i int}) -> Vec{i} :
    Vec{i} from a = [] for j::el = range v :
        a + [el + w[j]] 
    
(v Vec{i int}) ⋅ (w Vec{i int}) :
    from a = 0 for j::el = range v :
        a + el * w[j] 
    
(v Vec{3}) × (w Vec{3}) -> Vec{3} :
    Vec{3}[v[1]*w[2] - v[2]*w[1],
        .. v[2]*w[0] - v[0]*w[2],
        .. v[0]*w[1] - v[1]*w[0]]
```

Some notes on what's going on. The `Vec{i int}` types in the definition of `+` and `⋅` allow us to capture their type parameter under the identifier`i`, and, if it is used more than once, as in this case, to check that the parameters match.

Note that the presence of the return type ensures that the compiler recognizes that the output must be of the same concrete type as the inputs, so that for example it recognizes that each of the vector types fulfills the built-in `Addable` interface:

```
Addable = interface :
    (x self) + (y self) -> self
```

While the parameters of the types in the call signature may be captured, the parameters in the return signature are computed. This example should clarify the distinction. Suppose that sometimes we wanted to concatenate values in the `Vec` type as though they were ordinary lists. Then we can write a return type like this:

```
concat (v Vec{i int}, w Vec{j int}) -> Vec{i + j} :
    Vec{i + j}(list(v) + list(w))
```

(And indeed in the previous example of vector addition the return type was technically being computed, it's just that the computation was "evaluate `i`".)

The types we can use as parameters are `bool`, `float`, `int`, `string`, `rune`, `type`, and any enum type, or to put it another way, the types of which the values have literals rather than needing to be produced by a constructor function.

As mentioned, we can also parameterize struct types. In this particular example, we don't use the parameters for the runtime validation, but just to ensure that it treats different currencies as different types and doesn't try to add them together. An enum type is used as a parameter.

```
newtype 
    
Currency = enum USD, EURO, GBP 
    
Money{c Currency} = struct(large, small int):
    0 <= that[large]
    0 <= that[small]
    that[small] < 100
    
def
    
(x Money{c Currency}) + (y Money{c Currency}) -> Money{c} :
    x[small] + y[small] < 100 :
        Money{c}(x[large] + y[large], x[small] + y[small])
    else :
        Money{c}(x[large] + y[large] + 1, x[small] + y[small] - 100)
```

## Generics

We can of course implement generics merely as a special case where a parameterized type takes a type as a parameter. This has been done for you in the case of the built-in `list`, `set`, `map`, and `pair` types, with `list` and `set` taking one type parameter, and `pair` and `map` taking two. Here for example is the implementation of the generic `list` type:

```
newtype 
    
list{T type} = clone list using +, slice :
    from a = true for _::el = range that :
        el in T :
            continue 
        else :
            break false
```

Note that as in this example we can overload built-in types such as `list`. We can also overload parameterized types, e.g. we could have the `Vec{i int}` constructor defined above and also have a `Vec{i int, T type}` constructor which both checks the length of the vector and typechecks its elements.

## Monomorphization

Each instance of a parameterized type, e.g. `Vec{3}`, `Money{USA}`, is a distinct concrete type, and concrete types cannot be declared at runtime. Instead, they must be produced at compile-time, and the compiler declares all the instances that you actually mention in the script. Hence, if you run the following script ...

```
newtype

Z = clone{i int} int :
    0 <= that
    that < i

make Z{9}, Z{12} // See note below.

Qux = struct(zort Z{7})

def

foo(z Z{5}) :
    z

bar(i int) :
    Z{3}(i)
```

... then the types `Z{3}`, `Z{5}`, `Z{7}`, `Z{9}`, and `Z{12}` will exist at runtime, but nothing else of the form `Z{i}`.

The `make` statement is for when you want to have a type (or as in this case, types) to use at runtime, but there's no other reason to mention the type in the script: the `make` declaration ckmentions it, thus summoning it into existence, and has no other effect.