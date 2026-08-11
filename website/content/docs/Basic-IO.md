All Pipefish scripts automatically import a library called `world.pf`. This contains some commands and types that let you perform basic IO.

It does this in ways that are entirely extensible and customizable: you can use Pipefish to talk to whatever applications or devices your OS can talk to. We will discuss how to write libraries for such purposes in [the section on Advanced Pipefish](https://github.com/tim-hardcastle/Pipefish/wiki/Advanced-Pipefish). For now, let's look at the contents of the built-in `world` library.

## `get`

`get` commands perform input. They typically take the form `get <variable name> from <source>`. Apart from the commands for handling [SQL](https://github.com/tim-hardcastle/Pipefish/wiki/SQL-interop) and [microservices](https://github.com/tim-hardcastle/Pipefish/wiki/Microservices), the following commands are available:

* `get x from File(<string>)` sets `x` equal to the contents of the given file.
* `get x from File(<string>, <type>)` sets `x` equal to the contents of the given file as a string if the type is `string`, and broken into a list of strings at the newlines of the file if the type is `list`.
* `get x from FileExists(<string>)` sets `x` equal to `true` or `false` depending on whether the file does in fact exist.
* `get x from Input(<string>)` sets `x` equal to the response we get when prompting the end-user with the given string.
* `get x from Random(<integer>)` sets `x` equal to a random integer less than the supplied integer.
* `get x from Random(<list>)` sets `x` equal to  a random element of the list.
* `get x from UnixClock(<TimeUnit>)` sets `x` equal to the Unix epoch in the given unit, where `TimeUnit` is an enum with values `SECONDS`, `MILLISECONDS`, and `NANOSECONDS`.

## `put`, `post`, and `delete`

These commands are used to perform output. Their meanings are based on HTTP:

* `put` represents output which would have the same effect if you did it several times as if you did it once, such as writing a file to disc.
* `post` represents output which has a distinct effect each time you do it, such as writing text to the terminal.
* `delete` is always implemented so that if the thing to be deleted already doesn't exist, this is counted as a success rather than an error.

Apart from the commands for handling [SQL](https://github.com/tim-hardcastle/Pipefish/wiki/SQL-interop) and [microservices](https://github.com/tim-hardcastle/Pipefish/wiki/Microservices), the following commands are available:

* `post (<tuple>) to Output()` posts the values given to the current output.
* `post (<tuple>)` posts to `Output()` by default.
* `post (<tuple>) to Terminal()` posts the values given to the terminal. For now you may see little difference between that and `post (<tuple>)`, but the difference will become apparent [later](https://github.com/tim-hardcastle/Pipefish/wiki/Microservices#output-versus-terminal).
* `put (<int>) into RandomSeed()` seeds the random number generator.
* `put (<string>) into File(<filename>)` puts the given string into the file, creating it if it doesn't exist.
* `delete File (<filename>)` deletes the given file.