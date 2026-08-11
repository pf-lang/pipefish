In the section we look at the way that imports and modules work in Pipefish.

## Imports

Imports are introduced by the `import` headword.

A line in the `import` section can only have two forms: either simply a path to a library, expressed as a string:

```
import

"filepath/foo.pf"
```

In this case the functions in the library will be in the `foo` namespace, e.g. `foo.qux`.

Alternatively, it can be a pair consisting of a preferred namespace and the path:

```
import

foolib::"filepath/foo.pf" // Has namespace 'foolib'.
NULL::"lib/fmt.pf"        // Has no separate namespace.
```

Use of `NULL` as a namespace means that the imported file isn't put into a separate namespace.

## Imports as modules

An imported script can just be a static library. However, if it has a `var` section then the variables will be initialized; and if the script has an `init` command (as discussed [later in the manual](https://github.com/tim-hardcastle/Pipefish/wiki/The-init-and-main-commands)), it will be executed.

Namespaces can be chained: if your script imports `foo` which imports `bar` which has a function `zort`, then your script can call `foo.bar.zort`. 

## Standard libraries

Producing the standard libraries is usually a matter of wrapping Pipefish functions around Go functions, and so can be done easily and even automatically; and the developers at Google have already tested, optimized, and documented the library functions for us. The Pipefish libraries have been given matching names: `regexp`, `path/filepath`, etc.

Notable differences are:

* Pipefish has a SQL library as standard. This uses *snippets*, Pipefish's general way of embedding DSLs.

* Any form of IO is of course a command rather than a function, and IO commands are given a rather standardized form looking like `get <variable name> from <source>` and `post <value> to <destination>`. This is used for example to get the time from the system clock, or to put data into a SQL database.

* The `reflect` library is entirely different, since Pipefish has a different type system, and since the `type` function already tells you the types of values. Instead, the `reflect` library tells you about the types themselves, allowing you to discover e.g. whether a given type is a clone, and if so what its parent is, etc.

This version of Pipefish comes with the standard libraries `files`, `fmt`, `html`, `math`, `math/big`, `math/cmplx`, `math/rand`, `path`, `path/filepath`, `reflect`, `regexp`, `sql`, `strings`, `terminal`, `time`, and `unicode`. These are imported with their bare name rather than a filepath, e.g:
```
import

"unicode"
"sql"
"math/cmplx`

```
The following pages are in this section:

* [The `files` library](https://github.com/tim-hardcastle/Pipefish/wiki/The-files-library)
* [The `fmt` library](https://github.com/tim-hardcastle/Pipefish/wiki/The-fmt-library)
* [The `html` library](https://github.com/tim-hardcastle/Pipefish/wiki/The-html-library)
* [The `math` library](https://github.com/tim-hardcastle/Pipefish/wiki/The-math-library)
* [The `math/big` library](https://github.com/tim-hardcastle/Pipefish/wiki/The-big-library)
* [The `math/cmplx` library](https://github.com/tim-hardcastle/Pipefish/wiki/The-cmplx-library)
* [The `math/rand` library](https://github.com/tim-hardcastle/Pipefish/wiki/The-rand-library)
* [The `path` library](https://github.com/tim-hardcastle/Pipefish/wiki/The-path-library)
* [The `path/filepath` library](https://github.com/tim-hardcastle/Pipefish/wiki/The-filepath-library)
* [The `reflect` library](https://github.com/tim-hardcastle/Pipefish/wiki/The-reflect-library)
* [The `regexp` library](https://github.com/tim-hardcastle/Pipefish/wiki/The-regexp-library)
* [The `sql` library](https://github.com/tim-hardcastle/Pipefish/wiki/The-sql-library)
* [The `strings` library](https://github.com/tim-hardcastle/Pipefish/wiki/The-strings-library)
* [The `terminal` library](https://github.com/tim-hardcastle/Pipefish/wiki/The-terminal-library)
* [The `time` library](https://github.com/tim-hardcastle/Pipefish/wiki/The-time-library)
* [The `unicode` library](https://github.com/tim-hardcastle/Pipefish/wiki/The-unicode-library)