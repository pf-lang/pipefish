## Comments and continuations

Comments and continuations are illustrated by the example file `examples/wiki/comncon.pf`:

<img width="710" height="290" alt="image" src="https://github.com/user-attachments/assets/855da524-9dc3-4e0a-9387-f5f0a095ccc6" />

As you can see, we distinguish between docstrins and comments. Docstrings are considered part of the API; comments are not. Docstrings can appear at the head of a code file, as here, or before a definition. The `hub api` command, which we met eariler, will render docstrings in markdown.

Continuations must be marked by a `..` at the end of the continued line and a corresponding `..` at the beginning of the continuing line, as in the definitions of `X`, above.

The allowed exception is that the continued line may end in a comma *where this is syntactic*, in which case the continuation must begin with `..` just the same, as in the definition of `Y`, above. The continuations can be placed wherever is most readable: they are exempt from whitespace rules.

## Headwords

We have met four "headwords" so far: `def`, `cmd`, `const`, and `var`. The meaning of a headword is "everything after this until the next headword or the end of file is a `def`/`cmd`/whatever declaration". So after `cmd`, Pipefish expects you to be declaring commands; after `def` you can define functions, after `const` you can define constants, and after `var` you can define variables. Besides these, there are `newtype`, `import`, and `external`, which will be discussed later.