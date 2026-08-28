## Comments and continuations

Comments and continuations are illustrated by the example file `examples/docs/comncon.pf`:

```
~~ This is a docstring.

// This is a comment.

// And below, the two types of continuation.

const

HELLO = "hello " + ..
     .. "world"

LIST = [1, 2, 3, 
     .. 4, 5, 6]
```

As you can see, we distinguish between docstrins and comments. Comments work just as you'd expect. Docstrings are like comments, but are considered part of the API and are used to generate documentation, which will render markdown if you use it. Docstrings and comments can be interleaved if you wish. 

Continuations must be marked by a `..` at the end of the continued line and a corresponding `..` at the beginning of the continuing line, as in the definitions of `X`, above.

The allowed exception is that the continued line may end in a comma *where this is syntactic*, in which case the continuation must begin with `..` just the same, as in the definition of `Y`, above. The continuations can be placed wherever is most readable: they are exempt from whitespace rules.

## Headwords

We have met four "headwords" so far: `def`, `cmd`, `const`, and `var`. The meaning of a headword is "everything after this until the next headword or the end of file is a `def`/`cmd`/whatever declaration". So after `cmd`, Pipefish expects you to be declaring commands; after `def` you can define functions, after `const` you can define constants, and after `var` you can define variables. Besides these, there are `newtype`, `import`, and `external`, which will be discussed later.

Headwords can be used inline.

```
const MONTHS_IN_A_YEAR = 12

def square(i) :
    i * i
```
This is not idiomatic, but it's useful if you want to do code generation.