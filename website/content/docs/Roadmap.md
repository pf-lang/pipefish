## Introduction

In the article I will point out the more important things that need doing on the road between a language that works and a language usable in production. They are not in the order in which they should be done, since this is mostly optional, except the first one, which needs doing now, and which I am in fact doing now.

## Attract contributions of time and money

The roadmap has many items that could be done better or made easier by contributions of time and/or money. I have a [plausible promise](https://sar.informatik.hu-berlin.de/teaching/neu-etc/oss/The_Cathedral_and_the_Bazaar_.pdf) of success, and the advantages of solo development no longer outweigh its disadvantages.

It is therefore time to go around banging a big drum and trying to get noticed and seek contributions, if not by direct begging, then by a publicity campaign that will lead people to, in fact, this particular page of this website, where they can learn how they can help. Here you are. It worked!

### Contributing time

Just using Pipefish and complaining about it would be a very important contribution. I am the worst person in the world to dogfood Pipefish, because I already know how it works. Your experiences of learning Pipefish (thirty minutes well spent) and using it as a lightweight scripting language is essential data that I *cannot* get myself.

For people thinking of contributing code to the project, you will not be the first, and so the documentation and tooling has already been improved to meet the needs of people other than me. The article [For contributors]() explains what sort of shape the project is in considered as a codebase, which bits are gnarly, which bits are good, the tooling for people working on Pipefish, the tests the internal documentation and where to find it, etc.

There are a number of things on the roadmap that others could do better or more quickly than I could, either by just being cleverer than me, or because it involves a skill I have never acquired which is more useful to getting the job done than my better knowledge of the project overall. I will point some of them out as we look over the individual items on the roadmap.

And the project has necessarily had little in the way of code review except me staring in mute horror at the abomination I perpetrated six months earlier. I call this the Aghast Methodology.

### Contributing money

The purpose of the money is (a) infrastructure spending (b) so that I can afford to work on the project full-time (c) to hire other people to do some of the things on the roadmap that they could do better or more quickly than me, or that I actually shouldn't be allowed to do.

"Things I shouldn't be allowed to do", for example, includes security auditing, which I shouldn't do to my own project even if I was otherwise competent and experienced; this seems more like the sort of thing I could collect donations for than expect someone to donate.

People, businesses, or sinister religious cults who want sponsorship rights in exchange for giving me money can have whatever opportunities to promote themselves that they want. I am not proud.

Any tooling or documentation or instructional videos or whatever that I produce auxiliary to the main Pipefish project will also be released under the MIT license to the greater glory of the project. Conversely it seems fair that I should give a sufficiently munificient sponsor a proportionate amount of assistance if they wanted to do that themselves.

### Please put a star on the repo

Each star you add to [this page]() is a tiny crumb of credibility from which I will build a Plausibility Cake. Thank you.

## Optimization

Pipefish has no explicit optimization besides constant folding, but is nonetheless [surprisingly fast](https://github.com/lac-dcc/BenchGen/wiki/Pipefish:-A-functional-and-indentation-sensitive-programming-language-on-BenchGen), and I and the good people at UFMG are interested in knowing why. The basic architecture of the virtual machine was certainly designed to be reasonably efficient considered as an algorithm, but without particular consideration of hardware complexities such as branch prediction and caching. The choices I made must have fortuitously aligned with those considerations, and supposing (as I presently do) that the current VM will be the direct or indirect ancestor of the 1.0 runtime, it would be a good idea to find out why.

Algorithms for optimizing the bytecode after generation have been contributed by Paul C. Anagnostopoulos, and this is an area where other people can very easily contribute, as it's very separable from anything else; as is speeding up any particular part of the VM. Removing the cruft from where it's generated in the compiler is something I would be better qualified to do than anyone else.

## Better IDE support

Pipefish has a syntax highlighter for VSCode and that's it; and of all the features on the roadmap, the one I personally crave the most is red wiggly lines under my syntax errors. Then besides the other standard features, the simplicity of the Pipefish semantics will make it possible to implement some forms of automatic refactoring not possible in a language which isn't referentially transparent.

## Partial recompilation

This will be needed to speed up IDE support and hotcoding. It's not difficulty in principle, but it'll take a bunch of work on my part adding one more concern to each part of the compiler.

## Partial compilation

And because of Pipefish's TUI development, it will sometimes be very useful to a developer if when compilation fails, it does so only for the broken code and its dependencies. Again, this seems like something I need to take charge of.

## Typechecking SQL interop

'Cos it's such a nice feature. Although SQL faces the user as just another DSL in another standard library, the use-case is important enough that there is specialized machinery in the VM to speed up SQL interop, and in the same way the compiler could special-case SQL and type-check it using SQL's extensive capcities for reflection.

## Concurrency etc

As Pipefish naturally works on a request-response model, obviously concurrency should be implemented by being concurrent over the requests, with configuration from the hub as you would expect.

There are some improvements needed in the semantics to support the fact that different users are in fact different people.

## Better documentation

I have peaked as a technical and expository writer. I don't suck; I also don't think this is one of the things I'm going to get any better at by mere persistence. If I had money to throw at the problem, I would hire Al Sweigart to write a book on how to get the best out of Pipefish, while also giving me detailed feedback on the language as he writes it.

## An online playground

Ideally one that can be re-used by other people with a minimum of effort. OTOH, this is not even the *sort* of thing I have ever done; anyone who has done this once would be better at it than I would. OTOH, it might well be such a generic task that one good example will suffice, or ChatGPT can talk me through the basics.

## A formal spec

Other people might well want to implement Pipefish for various purposes, but although it is satisfyingly small it does have its fiddly bits. I need for example to represent the rules for mixfix operators in a form other than my own code.

Producing a formal spec is one of those things like dogfooding where I am too close to the language to be the right person to do it, and so this is the sort of thing that ideally would be done by someone else who thinks of my time as one of their resources rather than vice-versa. This seems like the sort of thing that needs to be paid for.

## More testing

Because of the large number of corner cases, and the things that can lurk in corners, mere code coverage of important parts is not a particularly good metric. It would in fact be a good idea to make a fine-grained list of all the language's features, and then test them systematically in pairs, with a publicly viewable grid on the repo to display progress. (That being done, we can go up a dimension.) This is something that could be massively parallelized between people who each want to contribute a little code. 

Besides just more tests, the project needs a fuzzer.

Writing the tests is something that could potentially be done by AI, though I wouldn't like it. I would however be completely fine with spending money on the most powerful AI possible to supplement the fuzzer by serving as an adversarial tester trying to break the compiler with terrible Pipefish.

Speaking of AI.

## LLM support

It is likely that there will be reasons for someone to do something that will make LLMs more proficient in Pipefish than they otherwise would be with a new language. It seems a reasonable hope that this would be done by other people in such a way that it doesn't have to become a maintained part of the project itself; but it might be a good idea to be able initially to produce some examples to point at that are not considered part of the project.

## AWS and Google Cloud

Assuming we avoid either a malign or benign technological singularity, we can assume that AWS and Google Cloud will outlive us all, and we can plan for there to be similar way to package up stacks on them; and in the same way, it would be a good idea to have some examples to point at.

## An SDK for something that isn't Go

Again, one good one as an example would be fine: something that allows e.g. Rust to embed a Pipefish service with the same function and method signatures as the `pf` library.

## The semantics of `private` types

This one's one me, but being able to talk about my options with some clever people might be helpful.

It's very easy to see what `private` means for anything else you want to declare, from variables to imports to functions.

`private` for types on the other hand is only partially implemented and not stable because it isn't clear what exactly `private` should mean when talking about a type.

And yet users will want to be able to do it because in many cases is will be quite clear to them, and quite clear in fact, what they're trying to do by saying a type is `private`.

## A package registry and version control

Natch.

## More comptime

You should be able to manufacture declarations at compile time, not just compute values. This can be done very nicely using snippets plus a `compile <headword> :` control structure . I could do this, but it's sufficiently separable that someone relatively new to the project could also do it and enjoy doing it. The semantics will take some thought on my part.

## Security auditing

This should of course be done by someone other than me, probably twice: once to say what needs doing, and once to check that everything that needs doing has in fact been done.

There are some things that I already know need doing, which have no obvious exploits but which still present an unnecessary surface area for one to occur.

## Various small-ticket items

There are some small tweaks needed to the core language.

Thanks to theft from Go, the production of more standard libraries is relatively trivial, so I put it among the small-ticket items, but we *haven't* quite finished stealing.

The API of the `pf` library is awkward and ugly.

"Improve the error messages" is, oddly enough, permanently on my todo list *no matter how often I do it*. We must imagine Sisyphus happy.