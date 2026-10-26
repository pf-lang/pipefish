## Introduction

The type system that I invented for Pipefish turned out to resemble that of Julia down to some of the terminology. As "type-system-like-Julia" is awkward to say, and as there doesn't seem to be an existing term, I have coined "latticial type system" as a natural parallel to the term "algebraic type system".

In this document I shall look at what a latticial type system is definitionally, and the other features that seem to naturally go with it. It will not be an overview of everything the Julia and/or Pipefish type systems can do: e.g. I don't mention enums, since there's no particularly latticial way to do enums.

My reinvention of the Julia authors' idea is particularly interesting because Julia is in a different paradigm oriented to a different domain, being an imperative math-oriented language, and so this suggests that this system is a method, independent of the other details of the language, for squeezing as much juice as we can out of a dynamic type system.

"Dynamic type system" is a term which I should now define, or it will cause confusion, as it usually does.

## What is a dynamic type system?

The term "dynamic type system" makes people think of interpreted languages, and of languages that aren't typechecked, and of languages that have *weak* typing where one can (e.g.) promiscuously add a string to an integer, or compare it with a boolean.

None of these is necessarily true. A **dynamic type system** is one in which the values of the language are tagged at runtime with their types, so that the runtime behavior can depend either on explicit inspection of the type by the code (**reflection**) or on decisions made automatically by the runtime (**dispatch**).

You will see from this description that many "static" languages are partly dynamic: for example all Java `Object`s are dynamic; so are all Go interfaces.

As an important practical constraint, it is important that the tags should be simple integers, that they should have no *structure*: we can represent a type as `list{set{int/float}}` at compile time, but by runtime we want to have reduced it to `57`, so that the reflection/dispatch can work efficiently on the tags.

## Latticial types versus algebraic types

We assume the reader is familiar with algebraic type systems, with their sum and product types: if not, hopefully the exposition below and a couple of minutes with Google will make the concept clear. Latticial type systems use product types (i.e. structs) in just the same way, but use union types instead of sum types.

To explain the distinction, suppose we want the idea "A circle is a shape with one parameter, a rectangle is a shape with two parameters". In a language with sum types, such as Rust, we would express it like this:

```
enum Shape {
    Circle(f64),
    Rectangle(f64, f64),
}
```

When we use either of the constructors, `Circle` or `Rectangle`, we get a value of type `Shape`, which contains a tag saying whether it's a `Circle` or a `Rectangle`. These are *not types*: they are varieties of the `Shape` type: you can't construct something of type `Circle`, or of type `Rectangle`. Rather, we have one type, `Shape`, with two constructors, `Circle` and `Rectangle`.

You might do something like this (again in Rust):

```
struct Circle {
    radius: f64,
}

enum Shape {
    Circle(Circle),
    Rectangle(f64, f64),
}
```

Now a `Circle` type *exists*, but note that it still isn't a subtype of `Shape`, and the typechecker would reject any attempt to pass a `Circle` to something expecting a `Shape`.

By contrast, let's express the same idea ("a circle is a shape with one parameter and a rectangle is a shape with two parameters") in Pipefish's latticial type system.

```
newtype

Circle = struct(radius float)
Rectangle = struct(width, height float)

Shape = abstract Circle/Square
```

Now we have two constructors for two types, `Circle` and `Square`. `Shape` is also a type, but it is an **abstract type**: you can never have a value of type `Shape`. Rather, abstract types like `Shape` are used as filters on the parameters of a function, or the fields of a struct. So in a latticial type system, you definitely *can* pass things of type `Circle` and type `Rectangle` to functions expecting a `Shape`: that's just what `Shape` exists for.

And so in a latticial type system, we have **concrete types** such as `Circle` and `Square`, which can be instantiated but not subtyped; and **abstract types**, which can be subtyped but not instantiated: plus a bottom `empty` type which can neither be instantiated nor subtyped.

The resulting type system forms a lattice of sets of values of which the concrete types are the minimal non-empty elements.

Strictly speaking that should just be the definition of a latticial type system, but there are some features that go along with being latticial in the same way that for example pattern matching goes along with being algebraic.

## Dispatch

The definition of an abstract type gives us a very efficient way to test for membership at runtime: just as we can test whether a value belongs to a concrete type by comparing its tag to the appropriate integer, so we can test its membership of an abstract type by indexing an array of booleans by the type tag.

This is in fact what motivates the construction of a latticial type system in the first place: we can overload anything any way we like, duck-type on it, and have the compiler/runtime dispatch on it. For example, let's make a `double` function which will work on things that can be multiplied by a scalar, by multiplying them by `2.0`; on lists, by doubling every element of the list; and on structs, by doubling every field of the struct:

```
newtype

Scalable = interface :
    (scalar float) * (x self) -> self

def

double(x Scalable) :
    2 * x

double(L list) :
    L >> double

double(S struct) :
    cast type(S), from L = [] for _::field = range S :
        L & double field
```

In fact, both Pipefish and Julia practice *multiple dispatch*: unlike languages such as Java or Go which dispatch only on the type of the value on which a method is called, Julia and Pipefish will dispatch on all the type parameters of a function, so that for example if we define a `Vec` type we could e.g. define `(x float) * (v Vec)` as scalar multiplication and `(v Vec) * (w Vec)` as the dot product, and have them both work alongside the built-in definition of `(x float) * (y float)` and `(x int) * (y int)`.

What we get from this is what the authors of Julia call "a dance between abstraction and specification" and which I more prosaically call "having your cake and eating it". When we want to, we can be strict about types and make sure you can only add apples to apples; and when we want to, for example when performing serialization, we can write functions which just treat all types as data.

In principle you *could* have a lattical type system (in the sense of having union types) that didn't use these facilities for dispatch, but why?

## Nominal types

Similarly it would be a missed opportunity *not* to have nominal types, so that we can distinguish an integer which actually represents a dimensionless quantity from one that represents a UID, or a quantity of apples, or a shoe-size.

In Pipefish, these "clone types" don't automatically get the full suite of operations of the parent type, so that by default you cannot add UIDs and UIDs, or multiply apples by apples.

## Consequences for generics

It follows from the rules of a latticial type system that while we can have for example concrete types `list{int}` and `list{any}`, the former cannot be a subtype of the latter.

## Conseqences for interfaces

In a latticial type system, an interface, as in our example above, is an abstract type which is defined according to some criterion rather than by explicitly listing its members. Julia doesn't have them, for which it has been criticized; their extra semantic significance in Pipefish is discussed in the main documentation.

## Consequences for nullability

In such a system, it is natural to make all the concrete types non-nullable, except for one null type containing a null element; the union of this null type with any other type gives a nullable type.
