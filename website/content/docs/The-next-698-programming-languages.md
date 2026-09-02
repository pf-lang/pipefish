## Introduction

This is a historical approach to what "functional programming" actually means. It barely mentions Pipefish, but it does in passing answer a question which may have bothered you while reading some of the other documents, and which bothered me for a long time.

## "If this is such a good idea ... ?"

All my life, I have been subject to occasional glimpses of how there must be something *better* than the sled, or the travois, or the mule, some *exciting new way* of reducing friction while transporting heavy loads.

And so I have set about to reinvent the wheel. Sometimes my wheels have been hexagonal, for ease of storage. Sometimes they have been spherical, for greater generality. Sometimes they have been made of styrofoam, to save weight.

And then most usually when I've finally hit on something circular and got my axle working, someone points out that my bright idea was done in the Late Neolithic, or in 1962, and I read about how they did it, and move on to something else. Just wait 'til you see what I'm "cooking up" (that's a little hint!) with two sticks and friction!

And so obviously I expected my ideas for Pipefish to have largely been pre-empted by someone else; and when that didn't happen, and years went by with no-one saying: "oh, this is X but with Y", the hypothesis that [despite all appearances](******) Pipefish is just a *really bad idea* must become proportionally more plausible in your eyes and mine, because of [Bayes' Theorem]().

The history I am going to recount, rapidly covering a period from the [*********]() to the [********](), does seem among other things to point to a set of historical and social reasons why Pipefish and a number of other potentially useful languages failed to be designed.

***** is chosen as the end of the story for dramatic effect. We should celebrate loudly that it *wasn't* the end of the FPL story, and that a new wave of user-friendly and problem-orented languages (Gleam, Elixir, Elm) have been coming out over the past 15 years or so. But for my purposes, we can and will stop there.

## What is a functional programming language?

Some years ago, but after I'd begun work on Pipefish, I sat quietly observing a very informal discussion among functional programmers of what they enjoyed about functional programming languages. I was struck by the fact that almost every answer *must* be answering the wrong question, because each of them was naming his or her favorite features of his or her own particular favorite FPL, and these features aren't common to them all; some of them are what we might classify as "bells and whistles". The Lisp enjoyers spoke of macros, the Haskell fans of laziness, the ML sodality of pattern-matching, the Lispers of dynamic freedom and the ML-Haskell faction of algebraic type systems ... and so on.

(This is not really suprising. What would a group of imperative programmers say if you asked them what they liked about imperative programming? How many of them would touch on the essence of imperativity, and what *is* that?)

This left me somewhat puzzled as to what a functional programming language actually is, and whether if I claimed to have implemented one, half of the FP community would politely say "no not really". (Pipefish, for example, has no pattern-matching and its type system is [not algebraic](*******************).)

This puzzlement drove me to look more carefully not only at the history of individual early functional programming languages, but in particular at how the term "functional programming" was first used, on the *apparently* reasonable grounds that just as Adolphe Sax can't really be wrong about what a saxophone is, so the high sage or wizard who first named "functional programming" must have some claim to be right about what it is.

## The purposes of programming languages

Before we go through the key points of our history, I should say what we're looking out for.

The most important question about any programming language is *"What is it for?"*; and whenever anyone proposes a new language, or tries to sell you on the virtues of one, that should be your first question --- and if they don't have an answer, you needn't bother with a second question.

(My own answer, for Pipefish, is found [here](Why-Pipefish-needs-to-exist.md), but this essay is not about that.)

And as we survey the history of programming languages, we see that the most successful and/or loveable languages almost always arose from an attempt to solve a very narrow use-case for a particular kind of user, or were even designed for essentially for a single use and for users as specific as Alice and Bob --- and ended up hitting a much bigger target than they were aimed at. But however big the target ends up being, the whole character of the language is set indelibly by the sort of use-case and user the designer had in mind.

Some examples (in alphabetical order):

* C was designed solely to write UNIX in.
* The C++ origin story is that Bjarne Stroustrup was trying to model computers interacting in a network, and wrote a beautiful version in Simula which would take a hundred years to yield enough data; followed by an ugly version in C which was performant.
* Erlang (the suprise sleeper hit of languages, now the basis for Elixir and Gleam) was designed to make Ericsen's telephone exchanges more stable.
* We all know what JavaScript was designed for. It works for that, and has now been used for everything else. JavaScript is something of a special case of course, because of its association with web-browsers, but we shuld consider that for example Java applets failed at the same task, along with ActionScript, VBScript, JScript, Silverlight  and NaCl, whereas JavaScript, despite its faults, *is* fit for its primary use-case.
* Go has not unkindly been called "a DSL for writing servers".
* Lua was designed to solve the needs of scientists in Brazil's nuclear power industry. (It became wedded to the games industry when someone at LucasFilms Games read about it in Dr Dobbs Journal and realized it was better than what they'd done in-house.)
* PHP, despite being objectively terrible, and coming from (and being aimed at) a dysfunctional "worse is better" culture, managed to spread like wildfire or a plague of poorly-designed locusts because its initials stood for Personal Home Page: it knew exactly what it was for.
* Python was designed because ********.
* Rust was famously inspired by a stuck elevator reminding ****** how terribly often things break, a rather more "meta" inspiration than for example C's origin story but still a moment that gave Rust a fixed purpose.


(It's worth noting in passing, for the moral of it, that *none* of the "easy" languages on out list were designed in an attempt to dumb programming down for a hypothetical userbase of grateful simple folks, like BASIC was, or Pascal. Rather, they were designed to make things (*specific* things) ergonomic for clever people, which, by a non-coincidence, makes them ergonomic for everyone. SQL is a DSL example of the same thing.)

The glaring exception to the general rule about the genesis of successful languages is of course Java. Its original use-case was embedding into consumer electronics, at which it failed, so to rescue the project it was pivoted into powering "applets" (embeddable client-side applications to be embedded in HTML), at which it failed, and then it became the giant among languages that it is today because Sun Microsystems spent a billion dollars in advertising to persuade Business Idiots that this was How To Do Business Properly. There should be a law against marketing technical solutions directly to management just like there are laws against marketing cigarettes to children *but what do I know*?

## Fortran

And now to our narrative. In 1953, John Backus began work on Fortran. Here is [his account of its invention, a gritty realistic epic](https://softwarepreservation.computerhistory.org/FORTRAN/paper/p25-backus.pdf); and here is mine, [a high fantasy with elves](The-legend-of-the-first-compiler.md). We don't need to go through all the details, except insofar as they bear on the important question: *"What was it for?"*

At that time, STEM-oriented institutions had entered the "open shop" era of computing, where the real users of the computers (scientists, mathematicians, engineers, etc) were allowed to write their own programs with the assistance and advice of the Computer Priests rather than the priesthood writing all the code; a much better use of everyone's time.

Machine code, however, was even more fiddly then than it is now; and the dripfeed acquisition of new models and makes of machines with their own machine code compounded the problem; and data access and output was specific to peripherals. What was needed was a higher-level language, and a performant one: Backus reckoned that if he slowed down excution by a factor of more than two, his potential users would stick with assembly.

Backus saw that the *general* thing we keep doing when we write code is to iterate over a thing doing a thing, and so saw that his users needed three things:

* Variable assignment.
* Indexed arrays.
* `for` loops (which Fortran then as now called `DO`, but I will use `for` loop as a generic term).

To quote Backus himself: ""

And so the problem of making a PL that did math and ran fast could be reduced to making that go fast, plus `IF` statements and `GOTO` (which are a very thin skin over jump statements in object code), and arithmetic.

So the birth of imperative programming, `goto` statements and all, may have contained some mis-steps, but it was not a hack. It was glorious. It is also an early example of a language hitting a target bigger than it was aimed at. Eventually Fortran would be supplanted by later languages for everything but its core competence, but before then it was used for the first adventure game, for payroll software, chess engines, CAD, mainframe utilites, and self-hosting compilers.

In order to understand the achievement of Fortran, consider this. In 1956, at the [Symposium On Advanced Computer Methods](ntrl.ntis.gov/NTRL/dashboard/searchResults/titleDetail/PB121670.xhtml), Robert Serrell of RCA described their system of Associative Machine Languages, which relied essentially on extending assembly by macros. A single sample will give you the picture: operation `11` means "Add the absolute value of all the operands whose addresses follow". 

You will notice how *exactly* this fails to solve the problems that Backus solved.  There are no variables, we still have to fiddle about with addresses. Instead of an indexed array to iterate over, the operation operates over arbitrary memory addresses. And because there are no `for` loops, this operator is one of a curated selection, and if you wanted for example to alternately add and subtract things on the list, you'd have to go back to writing assembly again.

## Lisp

In *****, John McCarthy invented Lisp; his own account of it is [here](http://jmc.stanford.edu/articles/lisp/lisp.pdf). *"What was it for?"*, you ask.

McCarthy had been working for **** on an extension of Fortran for list processing. The people at IBM were very happy with his work; he was frustrated. What he wanted, very reasonably, was to be able to automatically differentiate a function and then evaluate it, which sounds like exactly the sort of thing computers should be making easy for us. Fortran made it impossible despite its math orientation *and* despite the fact that McCarthy had been hired to write his own extensions to the language, a privilege which could surely have overcome any *non-fundamental* obstacle in the semantics of Fortran.

It's obvious how this led to the homoiconicity for which Lisp is famous, but as this clearly isn't essential to functional programming, we can ignore that aspect of the language. A more fundamental innovation from our point of view was *flow-of-control as expressions*, and this, too, stems directly from the specific use-case that McCarthy had in mind. The rules we learned in high-school for calculus, the product rule and chain rule and so forth, can be naturally extended by a rule too obvious to explicitly mention even to schoolchildren: that if we define `h(x)` by saying that it's equal to `f(x)` if some condition `C` holds, and `g(x)` otherwise, then `h'(x)` is equal to `f'(x)` if `C` holds, and `g'(x)` otherwise.

But to work with expressions like that, they do in fact have to be *expressions*; whereas Fortran's `IF` statements are semantically *statements* telling the runtime to *do* something, to jump to this or that part of the code.

## ML



## ISWIM

1962 was an *annus mirabilis* for computer science. It was the year that *****, proved that imperative programming doesn't need `GOTO` statements, and it was the year that Peter Landin invented ISWIM (an abbreviation for "If You See What I Mean").

On this occasion we must waive our usual question *"What's it for?"* because as Landin emphasized again and again, ISWIM is a *family* of languages, within which you can and should choose your primitive operations and your concrete syntax to suit your chosen domain. A particular ISWIM is for the thing you want it to be for, within a very general semantic framework for defining things.

An exchange between ***** summarizes the whole idea:



By "", they just mean McCarthy's discovery that flow-of-control should be expressed as expressions, not statements. And      our variables should be *defined* and not *assigned*.

For example, consider this implementation of Euclid's algorithm:

```
gcd(a, b) :
    b == 0 :
        error "arguments should be strictly positive"
    remainder == 0 :
        b
    else :
        gcd b, remainder
given :
    remainder = a mod b
```

The variable `remainder` is *defined*, it is given a meaning, and so given a value that persists over the lifetime of any particular function call, whereas in imperative code it would be *assigned* a value at a certain point in the execution of the function, and we would have to think about *where*, because if for example we put it before the test for non-positive values, our function wouldn't work as intended.

By using an ISWIM then, we separate two things that should never ever have gotten mixed up in the first place: *giving names to our concepts*, and *flow of control*.

These are the only two ideas that make an ISWIM. As with Fortran, for us to see them as really revolutionary, but consider again the fact that it was only this year people discovered that you didn't need `GOTO`; and Landin spends ********** simply proving that since an ISWIM can be desugared to the lambda calculus, his idea was in fact Turing-complete.

Landin both did and didn't call his idea "functional programming": in his talk explictly considered calling his idea "functional programming", "declarative programming", and "denotative programming", declaring for the latter. (Having made this decision, he then refers to it as "functional programming" *** later.)

If you're wondering why this is considered a seminal paper despite us not getting the 700 languages and "ISWIM" not being a household word, we'll come back to that.

## The birth of "functional progamming"

But we were looking for the birth of the phrase "functional programming", and now we're really getting there.

In ******, a somewhat elderly young firebrand gave a talk at ******** with the puzzling and misleading name "".




Now, about that name, "*********". You can see that what's wrong with it is that the paradigm our elederly young Turk is denouncing has nothing to do with Jonny von Neumann. It is in fact entirely the creation of one John Backus, but the speaker was the only person in the room who couldn't possibly say that, because, in a Massive Plot Twist, the young Turk who invented "functional programming" was John Backus, who invented imperative programming.

It is not, of course, that Backus repented of Fortran. The brief for Fortran was "make big math go vroom", and it does so to this day. But Moore's Law had given people to think of scope for better things, and to dream of languages further from the metal and closer to what they actually wanted.

## The end ... ?

In any case, surely we've reached the end of our quest now. Pipefish is *not* a functional programming language because you can't write point-free code in it. Also, we know know what functional programming *is*, it's that style where you "eschew lambda functions". Maybe one day the emblem of the whole functional programming movement will be a boot stamping on a lambda function for ever!

Oh, wait. Wait. Hold up one moment, that doesn't sound right at all.

It is at this point that we have to confess the weakness of the historical method in tracking down what the phrase "functional programming" means. (It is also a fact that the original recipe for ketchup contained fish and soy beans but no tomatoes: sometimes we have to bow to semantic shift.)

## The growth of functional programming

The sequel to Backus's talk was not that people started making languages in imitation of his own rather horrible APL-inspired language, but rather that everyone who was working with Scheme or ML and probably a lot of other interesting languages now forgotten rushed to their desks and explained that they were *already* doing functional programming and that *their* languages were readable and had standard libraries and could ping the database.

And so a movement was born, or identified, or called into being, or further and incrementally coallesced, or something. There was in fact a hot minute where "functional programming language" could have meant something very precise, because the table stakes for playing Backus's game was having functions as first-class values. But now that the most imperative languages have them too, this is no distinction.

## The rise of laziness

Now we need to go back a bit to Landin's ISWIM, and the `gcd` function that I used to represent it.

```
gcd(a, b) :
    b == 0 :
        error "arguments should be strictly positive"
    remainder == 0 :
        b
    else :
        gcd b, remainder
given :
    remainder = a mod b
```

While ISWIM has liberated the *developer* from deciding where to assign `remainder`, this still does have to happen as an event in time in the messy real world of the object code. And so efficiency requires that things in the `given` block of a function should be evaluated only if needed, *lazily* evaluated.

In the example above the saving is negligible, but consider this rather contrived way of writing the Collatz function:

```
collatz(i) :
    i == 1 :
        "The conjecture stands."
    i mod 2 == 0 :
        evenBranch
    else :
        oddBranch
given :
    evenBranch = collatz(i div 2)
    oddBranch = collatz(3*i + 1)
```

*Now* if we don't evaluate lazily, our function calls explode until we run out of stack.

An ISWIM only needs this small amount of implicit subcucullar laziness to stop it from going crazy. But why stop at a *little laziness? Laziness is powerful juju. Who could resist being able to say "Let P be the list of all primes in order"? We always wanted to.

And so inspired by not only by Landin's talk but by earlier work such as ****** (sketching out a lazy Lisp), and other similar suggestions, people set to work making about a dozen lazy functional programming languages to see what you could really do with the idea, almost each of which besides [Miranda](*****) had no use outside the particular university developing it.

And so the Haskell committe was convened, and toiled mightily to specific one programming language to rule them all, and it was glorious, and behold, it had abstraction and power.

And that is really where our story stops.

## The end

The three classic families of functional language were the Lisps, the MLs, and Haskell and its mostly forgotten predecessors. (Miranda still has fans, and a modernized version called [Admiran](https://github.com/taolson/Admiran) is under active development by one of them.)

A language, I said, takes its character from the use-case and users for which it was first designed. Let's review the classic Big Three:

* Lisp was designed with the thought of doing automatic differentiation.
* ML was designed as the metalanguage for a theorem-prover.
* Haskell was designed as a platform for computer scientists to experiment with functional programming languages.

Meanwhile the imperative people were writing (in many cases) terrible, terrible languages --- *but PHP stands for "Personal Home Page"*. Anyone asking ["Why hasn't functional programming conquered the world?"]() should think about that.

## The next 698 programming languages

You will notice that this story answers my question about re-inventing the wheel. I did! Peter Landin invented Pipefish, along with 699 other programming languages, back in 1962, the same sort of era as so many of my best ideas turn out to have come from. Pipefish is an ISWIM for writing CRUD apps.

And this also explains why we didn't get Landin's 700 programming languages, and in particular why we don't have Pipefish already, despite how [obviously awesome](*****) it is. Haskell sucked all the laziness out of the room, and we got *one* ISWIM; and for anyone else who thought of writing an ISWIM, Haskell was already there, attempting to be all of them while being oriented to no real-world domain.

Now I've written Pipefish, so we have two ISWIMS, and Pipefish is a much better example of one, because it *is* oriented towards an actual domain. We can have more. Laziness in Haskell may well have taken a lot of sweat to implement, but the amount of laziness necessary for an ISWIM can be hacked out in half-an-hour if you're doing a tree-walking interpreter, a little longer for compiled code, and a little more effort if you care about optimization. How to do it should be in beginners' textbooks, and when I have finished writing mine, it will be.

## And what *is* a functional programming language?

And finally, while we're tying up loose ends, I promised you an answer to "what is a functional programming language"?. Let's get back to that. I was listening, as I said, to a conversation about "why do we like functional programming" when one cried up homiconicity, another point-free style, a third the joys of pattern-matching, and a whole lot of other things that are not common to functional languages and which are entirely absent from mine

Into which confusion one wise sage whose name I would certainly credit if I could, spake thus: "We like functional programming because in functional programming there is only one design pattern: The Pipeline." And I recognized this immediately as the right answer because it is an exact description of the joy of coding in Pipefish given by someone who had never heard of it. That's it, *that's* what we like, the soothing feeling of pushing values into a pure function, getting values out, feeding those values into another purpe function, getting values out ... as calming as whitewashing a fence, and with the same inevitability of progress and eventual success.

As a corollary, a functional programming language is one where you code like that. *How* exactly this is done, whether its type system should be algebraic or latticial or as anarchic as Lisp, whether it has homoiconicity, whether it has currying and point-free style, whether it has monads, effect systems, or FC/IS, is a decision, like everything else in language design, that should depend entirely on its primary use-case.

Or to put it another way, the most important question about any language is: *"What is it for?"*
