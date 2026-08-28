## Introduction

Someone proposing a new language for use in production any time these last few decades had better have a very good reason for doing so. I do. It was a good reason when it tumbleweeded across my brain five years ago; it is a good idea now. And now that Pipefish *does*, after all, [more or less exist](Roadmap.md), I can say that it ought to exist with more conviction, not just because it "works" in the trivial sense that it rarely goes *spoing* any more, and that you can code in it, and it [goes quite fast](https://github.com/lac-dcc/BenchGen/wiki/Pipefish:-A-functional-and-indentation-sensitive-programming-language-on-BenchGen), but in the more important sense that it is as ergonomic to code in as I thought it would be.

This article is my good reason for writing Pipefish.

## The world's most popular language paradigm

The most popular language paradigm in the world (perhaps by an order of magnitude) is declarative, interactive, functional, and stateless. It is garbage-collected (of course) it has no pointers or references, it has immutable values, it is referentially transparent, it never panics or segfaults, but simply returns an error value when it gets bad data. And it's so easy to use that a majority of the people who use it probaby don't have the word "programming" anywhere in their job descriptions.

This paradigm is best represented by two giants: Excel and SQL, with their formulas and queries, respectively. Four times as many people write SQL queries as write Python; Excel formulas, when you do the math, may be used by around 10% of the human race. Not only is this paradigm easy, it is also everywhere, because what it gives its users is a generalized CRUD app, where we wish to store and query data and where we don't want it to change except as a result of a command from a user/client saying: "Change this data"; and such apps are maybe 90% of code that gets written, to support the administrative work of pretty much every business everywhere.

But 90% of the code can't be written in Excel or SQL, because they aren't general-purpose languages.

Pipefish is a general-purpose language which follows the same paradigm, and is written with the expectation that it has the same primary use-case; and so with all the choices about syntax, semantics, tooling, deployment, etc, made with that expectation. It is a language with a purpose.

It would not make a good systems language, but it makes a very good CRUD language, in a world filled with very bad CRUD languages, a number of which I have used. This is the language I wish I'd used instead: a product of someone who cares about these things perhaps even more than the people working on systems languages care about their genre, because I am driven not only by aspiration but by annoyance.

## A simple functional language

In order for all this to work, it is necessary to write a *functional* general-purpose language. For historical reasons the people interested in functional programming languages have taken them in the direction of power and abstraction, and done some rather wonderful things, but not Pipefish. In the course of this, they stretched the definition of "functional language" so thin that I should say what I in particular mean by it. I mean in particular three things:

* Purity of functions
* Immutability of values
* Referential transparancy

These are the things in functional programming that reduce friction for everyone who can get them (and the things that e.g. Java deprives us of so systematically that at this point one must attribute it to malice). The other affordances of Haskell or ML or Lisp can be considered on their merits as they fit the use case.

However, before we can get onto those choices and the things that follow from them, we need to address the elephant in the room of every functional programming language: what do we do about state? This is a particularly urgent question because three of the four letters in CRUD are the initials of stateful operatations.

## Functional core/imperative shell

The FC/IS pattern is based on the observation that impurity is a function color, and that we can therefore minimize the footprint of impurity on our code by pushing impurity as far up the call tree as possible: by *excapsulating* it. This has gone by a number of names, possibly with differences of nuance: onion architecture, hexagonal architecture, ports and adapters, functional-core/imperative-shell. The last term introduced by Gary Berhardt in his talk [Boundaries](https://www.youtube.com/watch?v=yTkzNHF6rMs) at Ruby Conf 12 and appears to have stuck.

Whatever you call it, it is particularly suited to CRUD apps, because as we noted above, in CRUD apps we wish mutation of state to be tightly bound to requests coming in from a user/client saying "Change this state".

This leaves us able to implement our actual business logic in pure, referentially transparent code where *nothing ever happens*; the Nirvana of coding.

FC/IS itself, as a design pattern, can more or less be done in the language of your choosing, but without the language, or the standard libraries, or the third-party libraries, being suited to this task; and with the FC/IS pattern being maintained even within a single project only by programmer discipline, which notoriously doesn't scale.

In Pipefish, the semantics enforce a distincion between *commands*, which do things (specifically move data from place to place) and return nothing but `OK` or an error; and *functions*, which return values but don't do anything. Commands can call functions; functions can't call commands.

A fuller description can be found here; and here you can see an explanation of why Pipefish is in fact two short steps and a ton of sugar away from being the lambda calculus.

## Suitability for the paradigm and use-case

The idea that Pipefish is above all a CRUD language has guided every aspect of its development, as of course has adherence to the paradigm it was written in, so this section can't be exhaustive, but will be confined to major features. We needn't dwell on trivia: the reason `/` means the same as in Python 3 is the same as the reason Guido van Rossum changed it from Python 2, and the reason why `%` is a suffix operator meaning "divide by 100" *is completely obvious*.

Nor do I need to describe features which are important but standard, because the case for Pipefish existing doesn't rest on its achievements in being completely normal: errors, for example, work like exceptions; you know what exceptions are; and if you're interested in the semantics of doing this in a functional type system, you can read about that <>here</span>. Similarly with `for` loops, yes, they are almost exactly the same as imperative `for` loops and yet stateless and referentially transparent, but this is only amazing if you're [interested in how it's done](**********).  Etc.

So let's look at the big unusual stuff, the rhinoceroses and giraffes.

### The type system

Pipefish has a latticial type system: that is, it uses union types instead of the sum types of an algebraic type system; and it performs mutiple dispatch. This system is used in production by the imperative math-oriented language Julia: they describe it as "a dance between abstraction and specification". I re-invented it down to some of the terminology, but prefer the phrase "having your cake and eating it".

Dealing with data, we would sometimes like to be very hard-line about types and only be able to add apples to apples, and UIDs to nothing. And sometimes we'd like to be able to write a function that treats different data types all the same way and duck-types on them, for example when we want to serialize structs, and we want to write one function for each target format, not one for each struct; or when we want to write a `sum` function for lists that will add up whatever addition is defined on.

I should mention two differences from Julia: the addition of `de facto` interfaces, as a sanity check; and the fact that the labels of fields of structs are first-class values, with structs being indexed with square brackets the same as maps, making it more convenient to abstract over struct types and their fields.

## `main`-less programming

In Pipefish, since it's a GPL, it is *permitted* to write a program with a `main()` command which does batch processing, prints the result, and stops; or even to write a `main` command with a hand-written REPL inside it.

But the more natural way for a human to interact with it is to declare functions and commands which can be used from the TUI, and to talk to Pipefish in Pipefish in the same way that one talks to SQL in SQL. To make it easy to develop this way, the TUI is full of dev tools.

### One API To Rule Them All

(I am presently unable to think of a less silly name for this concept.) The idea is that the things you can do in the TUI when you run a Pipefish service in your terminal are the things you can do in your code if you import it as a library, and also the things you can do with it in your code if if's someone else's webfacing service and you give your Pipefish compiler your username and password.

We may note, smugly, in passing, that this sort of thing is only really practicable in a language where all values are immutable.

### DSLs and glue

Pipefish has a very nice flexible syntax for implementing and overloading math functions, and for making the front-end of your service more DSL-like when this is useful.

It also provides facilities for embedding existing DSLs. It would be possible to write an entire article on why there is only one good way to do SQL interop. I did: the article is [here](********).

This is also the best way to do interop with Prolog and HTML and JSON and SMTP and HTTP.

### A close relationship with Go

This may seem like the odd one out of this list, a mere implementation detail. However, it would be frivolous to write a new language for use in production, in 2026, that couldn't co-opt an existing established one. Go is very unlike Pipefish in that it's imperative and everything is a big fuzzy ball of mutable state. And it's very like Pipefish in that it's backend-oriented and places a lot of value on simplicity and predictability and the virtues of a small boring language.

Golang programs can embed Pipefish services as a struct of type `Service`, and, more importantly, Go can be embedded syntactically into Pipefish by writing a Pipefish function with a Golang body. Turning a Golang standard or third-party library into a Pipefish library is often trivial and sometimes automatable. (Also I'm sure an AI could do it very easily but I haven't tried.)

### A shallow learning curve, a small surface area

(... and other good geometric metaphors, such as the Oblong of Happiness.)

Pipefish has been designed and used for years, unlike the median new language in 2026 which hasn't been used by the person who told Claude to write it. It has achieved simplicity by being written by someone who prizes simplicity (again, because of my intellectual affinity to Winnie-the-Pooh) and who was allowed to develop for years without a committee or a manager or any users to lock in any design choices.

The result is a language which is no bigger than it needs to be, and which never requires you to do something difficult before you can do something easy.

It can be used as readable pseudocode, to illustrate DSA to people who haven't actually studied Pipefish; and in the business environment in which it's intended, someone's first use of the language, as with SQL, might well be typing something into the TUI that they don't think of as "a line of code", and which tells them exactly what they want to know.

### The hub and deployment

A hub is a Pipefish service which lives inside the TUI and which acts as a housekeeper for services running on that instance of Pipefish. It can start other services and stop them, supplies tools for debugging them, supplies them with environment variables, collects data from them, and does everything you would want to do to a service from the outside.

Making services web-facing is, of course, something you want to do from the outside, rather than boilerplating HTTP connections and salting and hashing passwords and role-based access management onto every service by hand, and so this is configured and managed from the hub.

Or to put it another way, the simplest Pipefish script from `"hello world"` onward is implicitly embedded in a web framework written in Pipefish which was designed concurrently with the Pipefish language and its tooling by the same person.

