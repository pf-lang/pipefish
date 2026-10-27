## Introduction

In this document we will look at the rationale behind Pipefish having `for` loops, something which is of course commonplace in imperative languages but seems (at first) to be at odds with the functional paradigm.

We will begin with the only piece of sarcastic mathematics that I know of.

## Sarcasm

<sarcasm>


When John Backus invented the concept of functional programming in [his seminal paper](https://worrydream.com/refs/Backus_1978_-_Can_Programming_Be_Liberated_from_the_von_Neumann_Style.pdf) "Can Programming Be Liberated from the von Neumann Style?" he explained that a major part of the definition of FP, and its chief merit, was the ability to write point-free code. With a prescient insight, he realized that the real problem with programming is that we keep giving names to concepts instead of floating in a realm of pure abstraction. And so as he explains, functional programmers "eschew lambda functions", and even to this day there is nothing so antithetical to functional programming as lambdas, with their nasty, nasty, *nasty* named parameters ... except, maybe, the `for` loop, coincidentally invented by one John Backus, and absolutely dripping with named variables. We are better off without such things. Let's look at the rich rewards of getting rid of them.

Von Neumann languages, complains Backus, are "fat and flabby". Functional programming will replace their plethora of special cases with powerful combinators. We can see this in action by looking at e.g. Haskell's prelude and some of its powerful combinators. We have `map`, `filter`, `reverse`, `any`, `all`, `scanl`, `scanl1`, `scanr`, `scanr1`, `replicate`, `zip`, `zip3`, `zipWith`, `zipWith3`, `unzip`, `unzip3` ... damn, a number of those do look like they might be special cases after all! If only there was just *one* wonderful magical Omega Function by which we could express all of these things, something that subsumes anamorphisms, catamorphisms, hylomorphisms and paramorphisms into one glorious panmorphism. But of course such a thing is just an impossible dream, isn't it?

Not at all. After literally minutes of unremitting toil, I give you … the Omega Function. For types `π` and `τ`, let us define a function `Ω[π, τ]` having type  `π → ℕ → τ → τ → π → ℕ → τ → τ` and defined by :

    Ω[π, τ](F, p, n) = F(p, n - 1) ∘ F(p, n - 2) ∘ … ∘ F(p, 0) ∘ ι

for `F : π → ℕ → τ → τ` , `p : π`, `n : ℕ`.

Going forward we will omit the `[π, τ]` in  `Ω[π, τ]` for convenience and just write `Ω`, since it will always be clear from context what the types are.

Examples:

* For any type `σ` on which addition is defined and having zero element `zero(σ)`, let `π = list[σ]` and let `τ = σ`. Then if we define `F(L, i, t) = t + Lᵢ`, then `Ω(F, L)(|L|, zero(σ))` is the sum of the elements of the list.
* For any type `σ` let `π = list[σ] × σ -> Bool` and let `τ = Bool`. Then if we define `F(L, ϕ, i, t) = t ∧ ϕ(Lᵢ)` (where `L : list[σ]`,  `ϕ : σ → Bool`), then `Ω(F, L, ϕ)(|L|, true)` returns `true` if `ϕ` holds for every element of `L`.
* For any types `σ`, `ρ`, let `π = list[σ] × σ → ρ` and let `τ = list[ρ]`. Then if we define `F(L, f, i, t) = cons(t, f(Lᵢ))` (where `L : list[σ]`,  `f : σ → ρ`, then `Ω(F, L, f)(|L|, [])` returns `[f(L₀) , f(L₁), … ]`
* Let `π = Unit` (where `Unit` is a type having a sole element `u`) and let `τ = ℕ × ℕ`. Then if we define `F(u, i, t) = (t₁, t₀ + t₁)`, then `(Ω(F, u)(n, 0, 1))₀` returns the nth Fibonacci number.

So the Omega Function is clearly very powerful. But can it be presented in an ergonomic way that the average coder will understand? At first it seemed unlikely, but after a period of intense thought lasting many seconds, I invented something I call ... the `for` loop.

That's not the most extended piece of sarcasm I've ever written but it's the only one with sarcastic math in it.

My point is, the `for` loop was a good idea, it's *the* abstraction of what people do with computers: you iterate over a thing doing a thing. Making stuff like `foldr` and `zip` and so on core-adjacent makes the number of forms you need to learn much larger than if you just used a `for` loop but without even managing to cover the same area. Add to that the fact that people like `for` loops and understand them, and there's a pretty good case for a functional language having them, and having them in the convenient form that people are used to, so far as is possible.

## Introducing `for` loops in Pipefish

Pipefish is written in Go and in some ways sticks to it closely, so syntacticcally I have copied the form of Go's `for` loops, systematically changed to reflect the facts that (1) Pipefish doesn't distinguish between `=` and `:=`; (2) Pipefish has a `::` operator specifically for key-value pairs (3) Pipefish has syntactic whitespace (4) functional `for` loops need one more clause.

Here by way of example is a `sum` function:

```
def

sum(L list) :
    from a = L[0] for i = 1; i < len L; i + 1 :
        a + L[i]
```

## Three interpretations of the `for` loop

The *procedural* interpretation of the `for` loop given above is that the `from` clause tells us what to assign to `a` when we initialize it; the body of the `for` loop tells us what we're substituting for the variable `a` each time we go around the loop; and the header of the `for` loop tells us that we initialize `i` as `0` and increment `i` each time we go around the loop.

The *mathematical* interpretation of the `for` loop is that `a` and `i` don't change at all *because they are bound variables*. They don't change any more than the `n` in `{n ∈ ℕ : n > 5}` changes, or the `i` in mathematicians' `Σ` or `Π` notation (which are indeed just specialized versions of a `for` loop, or in the Omega Function). And we can maintain that point of view so long as the body of the loop is a pure expression.

The `from` is there to bind the `a` to the `for` loop, to make it clear that it is not in fact an assignment.

And the *functional programming* interpretation of the `for` loop is that we are passing some initial values an anonymous function where all the recursion is in tail-call form to a polymorphic functor which will recursively apply the function to the initial values until it reaches a fixpoint.

## The rest of the specification

As usual in functional languages if we need to iterate over a tuple of values we're going to get a tuple of values back. E.g. if we implement a Fibonacci function like this:

```
def

fib(n int) :
    from a, b = 0, 1 for i = 0; i < n; i + 1 :
        b, a + b
```

... then this will return a 2-tuple of numbers of which we are interested only in the first, e.g. `fib(6)` will return `8, 13`. As in other functional languages, the ergonomic fix is to supply a function `first` which we can use to select the relevant member of the tuple. This goes nicely with our choice of `from` as syntax:

```
def

fib(n int) :
    first from a, b = 0, 1 for i = 0; i < n; i + 1 :
        b, a + b
```

There is no difficulty in providing `break` and `continue` expressions, as these are just syntactic sugar. E.g. we can write:

```
def

find(x single?, L list) :
    from result = -1 for i = 0; i < len L; i + 1 :            
        L[i] == x :
            break i
        else :
            continue
```

Although the `for` loop looks very Go-like or C-like with its `for <initialization>; <condition>; <expression>` format, the initialization must consist of giving an initial value to the loop variable and the expression must say what it is we want to do with that variable each time we go round the loop. The imperative construct lacks these constraints, but that's no big loss because people rarely use the additional freedom of the imperative version and when they do it is often confusing.

As with Go, we can use `for` with just the condition as a `while` loop:

```
def

collatz(n int) :
    from x = n for x != 1 :
        x % 2 == 0 :
            x / 2
        else :
            3 * x + 1
```

... or with no condition at all as an infinite loop:

```
def

collatz(n int) :
    from x = n for :
        x == 1 :
            break
        x % 2 == 0 :
            x / 2
        else :
            3 * x + 1
```

And we can likewise imitate the `range` form of Go's `for` loop, though we will use Pipefish's pair operator `::` to do so:

```
selectEvenIndexedElements(L list):
    from a = [] for i::x = range L :
        i % 2 == 0 :
            a + [x]
        else :
            continue
```

## Prior art

The semantics of pure functional `for` loops was previously implemented in the Futhark programming language, designed for parallelization for high-performance math. [In the case of Futhark](https://futhark-lang.org/blog/2026-01-20-why-not-tail-recursion.html), the authors were motivated by the difficulty of doing recursion in the hardware they were targeting: my own motivations were [more to do with ergonomics and DX](https://youtu.be/ADqLBc1vFwI?si=UsAQb2kGg1kFJBhn&t=79).
