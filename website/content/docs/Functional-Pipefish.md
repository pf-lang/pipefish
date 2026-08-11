Pipefish functions are perhaps different from the functions you are used to.

(1) They are **pure**. They have no side-effects (writing to the terminal, or disc, or database); they have no side-causes (reading from the disc, or the database, or the system clock, or the random number generator); they are unaware of the global variables, but evaluate an expression dependent only on the arguments passed and, perhaps, global _constants_.

Another way of looking at this is to say that in principle every pure function could be replaced by a lookup table (maybe a very large one!) in which the keys are all the possible combinations of arguments and the values are the corresponding return values of the function.

(2) They are **referentially transparent**. That is, any given symbol means the same thing throughout the course of the function. This means that there can be no local variables: if they varied, that would spoil referential transparency.  This means that functions don't contain the sort of loops that you're used to, since for example the index variable of an imperative-style `for` loop would break this rule.

It follows that the _only_ thing a Pipefish function can do is return a value, and that the body of a Pipefish function is just one big (or tiny) expression to be evaluated. If you want to do messy things like performing IO and mutating variables, you need to use, not a function, but a command, as will be explained in the section on Imperative Pipefish. This makes functions in Pipefish at once peculiarly simple and peculiarly powerful in ways that you may not quite be able to imagine if you aren't used to programming with pure functions.

The following pages tell you pretty much everything about functions in Pipefish.

* [Introducing functions](https://github.com/tim-hardcastle/Pipefish/wiki/Introducing-functions)
* [Infixes, postfixes, etc](https://github.com/tim-hardcastle/Pipefish/wiki/Infixes,-postfixes,-etc)
* [Function signatures and overloading](https://github.com/tim-hardcastle/Pipefish/wiki/Function-signatures-and-overloading)
* [Local constants and inner functions](https://github.com/tim-hardcastle/Pipefish/wiki/Local-constants-and-inner-functions)
* [First-class functions, lambdas, closures](https://github.com/tim-hardcastle/Pipefish/wiki/First%E2%80%90class-functions,-lambdas,-closures)
* [Iterators and piping](https://github.com/tim-hardcastle/Pipefish/wiki/Iterators-and-piping)
* [Unsatisfied conditionals](https://github.com/tim-hardcastle/Pipefish/wiki/Unsatisfied-conditionals)
* [Functions and tuples](https://github.com/tim-hardcastle/Pipefish/wiki/Functions-and-tuples)
* [Error handling](https://github.com/tim-hardcastle/Pipefish/wiki/Error-handling)
* [Recursion](https://github.com/tim-hardcastle/Pipefish/wiki/Recursion)