## Introduction

In the article I will point out the more important things that need doing on the road between a language that works and a language usable in production.

It is also, as a necessary consequence, a list of things many of which could be made better or easier by contributions of time and/or money.

And it follows from *that* that attracting such contributions is actually part of the roadmap, and that pointing out how I could budget other people's time and money is also part of the roadmap.

## Contributing

So if you're just looking at this document out of curiousity, you should skip to [the next section]().

If you're still with me, I will suppose that you've already been convinced by looking at the language documentation and such articles as [Why Pipefish needs to exist]() that you do by now think that it would be a good thing for the project to be brought to version 1.0, and you're just wondering if you yourself should do more than add a star to the GitHub repo.

### Contributing time

I am the worst person in the world to dogfood Pipefish, because I already know how it works. Your experiences of learning Pipefish and using it as a lightweight scripting language is essential data that I cannot get myself.

For people thinking of contributing code to the project, you will not be the first, and I have produced documentation and tooling to help them. The article [State of the project]() explains to potential contributors what they'd be getting themselves into.

There are a number of things on the roadmap that others could do better or more quickly than I could, either by just being cleverer than me, or because it involves a skill I have never acquired which is more useful to getting the job done than my better knowledge of the project overall. I will point some of them out as we look over the individual items on the roadmap.

### Contributing money

The purpose of the money is (a) infrastructure spending (b) so that I can afford to work on the project full-time (c) to hire other people to do some of the things on the roadmap that they could do better or more quickly than me, or that I actually shouldn't be allowed to do.

"Things I shouldn't be allowed to do", for example, includes security auditing, which I shouldn't do to my own project even if I was otherwise competent and experienced; this seems more like the sort of thing I could collect donations for than expect someone to donate.

People, businesses, or sinister religious cults who want sponsorship rights in exchange for giving me money can have whatever opportunities to promote themselves that they want. I am not proud.

Any tooling or documentation or instructional videos or whatever that I produce auxiliary to the main Pipefish project will also be released under the MIT license to the greater glory of the project. Conversely it seems fair that I should give a sufficiently munificient sponsor a proportionate amount of assistance if they wanted to do that themselves.

## Things that need doing

With that out of the way, let's look at what needs to be done.

### Optimization

Pipefish has no explicit optimization besides constant folding, but is nonetheless [surprisingly fast}(https://github.com/lac-dcc/BenchGen/wiki/Pipefish:-A-functional-and-indentation-sensitive-programming-language-on-BenchGen), and I and the good people at UFMG are interested in knowing why. The basic architecture of the virtual machine was certainly designed to be reasonably efficient considered as an algorithm, but without particular consideration of hardware complexities such as branch prediction and caching. The choices I made must have fortuitously aligned with those considerations, and supposing (as I presently do) that the current VM will be the direct or indirect ancestor of the 1.0 runtime, it would be a good idea to find out why.

Algorithms for optimizing the bytecode after generation have been contributed by Paul C. Anagnostopoulos, and this is an area where other people can very easily contribute, as it's very separable from anything else; as is speeding up any particular part of the VM. Removing the cruft from where it's generated in the compiler is something I would be better qualified to do than anyone else.

### Better IDE support

Pipefish has a syntax highlighter for VSCode and that's it; and of all the features on the roadmap, the one I personally crave the most is red wiggly lines under my syntax errors. Then besides the other standard features, the simplicity of the Pipefish semantics will make it possible to implement some forms of automatic refactoring not possible in a language which isn't referentially transparent.

### Partial recompilation

This will be needed to speed up IDE support and hotcoding. It's not difficulty in principle, but it'll take a bunch of work on my part adding one more concern to each part of the compiler.

### Partial compilation

And because of Pipefish's TUI development, it will sometimes be very useful to a developer if when compilation fails, it does so only for the broken code and its dependencies. Again, this seems like something I need to take charge of.

### Typechecking SQL interop

'Cos it's such a nice feature. Although SQL faces the user as just another DSL in another standard library, the use-case is important enough that there is specialized machinery in the VM to speed up SQL interop, and in the same way the compiler could special-case SQL and type-check it using SQL's extensive capcities for reflection.

### Concurrency etc

As Pipefish naturally works on a request-response model, obviously concurrency should be implemented by being concurrent over the requests, with configuration from the hub as you would expect.

There are some improvements needed in the semantics to support the fact that different users are in fact different people.

### Better documentation

I have peaked as a technical and expository writer. I don't suck; I also don't think this is one of the things I'm going to get any better at by mere persistence. If I had money to throw at the problem, I would hire Al Sweigart to write a book on how to get the best out of Pipefish, while also giving me detailed feedback on the language as he writes it.

### An online sandbox

Ideally one that can be re-used by other people with a minimum of effort. This is not even the *sort* of thing I have ever done; anyone who has done this once would be better at it than I would.

### A formal spec

Other people might well want to implement Pipefish for various purposes, but although it is satisfyingly small it does have its fiddly bits. I need for example to represent the rules for mixfix operators in a form other than my own code.

Producing a formal spec is one of those things like dogfooding where I am too close to the language to be the right person to do it, and so this is the sort of thing that ideally would be done by someone else who thinks of my time as one of their resources rather than vice-versa. This seems like the sort of thing that needs to be paid for.

### More testing

Because of the large number of corner cases, and the things that can lurk in corners, mere code coverage of important parts is not a particularly good metric. It would in fact be a good idea to make a fine-grained list of all the language's features, and then test them systematically in pairs. (That being done, we can go up a dimension.)

Besides just more tests, the project needs a fuzzer. It also seems to me that trying to break the compiler with bad Pipefish is one of those rare cases where an LLM could do much good while being able to do *no harm at all*.

Speaking of AI.

### LLM support

It is likely that there will be reasons for someone to do something that will make LLMs more proficient in Pipefish than they otherwise would be with a new language. It seems a reasonable hope that this would be done by other people in such a way that it doesn't have to become a maintained part of the project itself; but it might be a good idea to be able initially to produce some examples to point at that are not considered part of the project.

### AWS and Google Cloud

Assuming we avoid either a malign or benign technological singularity, we can assume that AWS and Google Cloud will outlive us all, and we can plan for there to be similar way to package up stacks on them; and in the same way, it would be a good idea to have some examples to point at.

### A package registry and version control

Natch.

### Security auditing

This should of course be done by someone other than me, probably twice: once to say what needs doing, and once to check that everything that needs doing has in fact been done.

There are some things that I already know need doing, which have no obvious exploits but which still present an unnecessary surface area for one to occur.

### Various small-ticket items

There are some small tweaks needed to the core language.

Thanks to theft from Go, the production of more standard libraries is relatively trivial, so I put it among the small-ticket items, but we *haven't* quite finished stealing.

"Improve the error messages" is, oddly enough, permanently on my todo list *no matter how often I do it*.

And so on.
