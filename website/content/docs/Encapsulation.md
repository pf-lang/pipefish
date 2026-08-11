## Introduction

In this page we will explain the use and semantics of the `private` keyword.

## Usage

`private` works like a headword that modifies other headwords, in that if you write for example:

```
def

<some functions>

private

<some more functions>
```

... then the functions below `private` will be private. The `private` declaration has scope until the next headword.

Anything you can declare can be declared `private`: variables, constants, functions, commands, imports, types, and external services (which we haven't met yet). If it has a headword, you can make it private.

Unlike in OO languages, there is no way and no real need to declare individual fields of a struct private.

## Semantics

`private` means, first of all, *private from the REPL*. An end-user, using Pipefish as its own front-end, can only call public functions and commands, and only see and mutate public variables.

Second, it means *private when imported as a library*. An application using a Pipefish library is in just the same condition as a human using the same service through its REPL, and is faced with the same API.

Third, it means *private when called as an external service*. But as we haven't dealt with microservices yet, we'll come back to that.

