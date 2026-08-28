## Introducing clone types

Besides the other user-defined types, it can be useful to have types which resemble, but are distinct from, the built-in types. If, for example, we use an integer to represent a UID, then it is convenient to declare a UID type which can be distinguished from other integers. Or it may be convenient to have an `Apples` type which can't be added to `Oranges`, or a vector type which can't be mistaken for a list:

```
newtype 

Uid = clone int 

Apples = clone int using +, -
Oranges = clone int using +, -

Vec = clone list 
```

Note that a clone type *is not a subtype of its parent*, but an entirely separate type.

## Using `using`

We should explain about the `using` clause. Some operations, such as `<` for clones of integers or `len` for clones of lists, are supplied automatically when you declare the type, and work just the same as for the parent type. Others, such as `+` for integers or slicing for lists, must be explicitly requested in a `using` clause.

The underlying rule is that an operation must be requested if it would be expected to return a value in the clone type. Hence the operations that need requesting are `+`, `-`, `*`, `/`, `div`, `mod`,`with`, `without`, the operators `>>` and `?>` for lists, and slicing clones of strings and lists. (Request the slicing operation with the word `slice`, the other operations by their names or symbols.)

An operation which you don't request in the `using` clause may still be overloaded by hand, if you don't want the operation to work like it does on the parent type. Here for example is how we'd implement addition for the `Vec` type. Don't worry if you don't quite understand it: it uses language features you haven't met yet, and is here just to show you what you can do if you want to.

```
def 

(v Vec) + (w Vec) -> Vec :
    len(v) != len(w) :
        error "adding vectors of different lengths"
    else :
        Vec from r = [] for i::el = range v :
            r + [el + w[i]]
```

As with structs and enums, each clone type has a constructor function with the same name as the type.

## Examples in the TUI

Let's demonstrate all this in the TUI.

```tui
→ hub run "examples/docs/clones.pf"
Starting script "clones.pf" as service "clones". 
clones → Apples(3) + Apples(5)
Apples(8)
clones → Apples(3) + Oranges(5)
[0] Error: function + cannot accept arguments of type Apples on the left of it and Oranges on the right at line 1:10-11 of REPL input.                                                                                     
clones → Uid(3) + Uid(5)
[0] Error: function + cannot accept arguments of type Uid on the left of it and Uid on the right at line 1:7-8 of REPL input.                                                                                              
clones → Vec[1, 2] + Vec[3, 4]
Vec[4, 6]
clones →  
```

## Postfix constructors

Cloning the `int` and `float` types supplies you with postfix construtors as well as prefix constructors, so that you can if you wish write `5 Apples` instead of `Apples 5`.