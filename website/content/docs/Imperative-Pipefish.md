In Pipefish, there is a strict distinction between _functions_, which return values, are pure, and are defined in the `def` section of the script; and _commands_, which return no values, are impure, and are defined in the `cmd` section of the script.

It could almost be said that Pipefish is two languages, Functional Pipefish and Imperative Pipefish, united by a common syntax and type system, but otherwise very different.

Imperative Pipefish is there to move and mutate data. It can get data from disc, a database, the system clock, a random number generator, an HTTP query, whatever. (These capabilities are extensible so "whatever" should be taken literally.) In the same way it can write to disc or to a database, seed the random number generator, etc. It can also read and change the global variables of the service. What it cannot do (let's repeat this) is return values. When you want a computation done, you have to turn this over to a function.

This means that between them Imperative Pipefish and Functional Pipefish can do pretty much anything you want to do. But since commands can call functions, whereas functions can't call commands, this means that if you want your service to perform IO at all, the _only_ way to program in Pipefish is to put a thin shell of Imperative Pipefish, performing IO, around business logic expressed in pure Functional Pipefish.

This section of the manual contains the following pages:

* [Global variables](https://github.com/tim-hardcastle/Pipefish/wiki/Global-variables)
* [Basic IO](https://github.com/tim-hardcastle/Pipefish/wiki/Basic-IO)
* [Flow of control](https://github.com/tim-hardcastle/Pipefish/wiki/Flow-of-control)
* [The init and main commands](https://github.com/tim-hardcastle/Pipefish/wiki/The-init-and-main-commands)
* [Imperative error-handling](https://github.com/tim-hardcastle/Pipefish/wiki/Imperative-error-handling)