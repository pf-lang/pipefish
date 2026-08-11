In this page we will catalog the basic hub functions for starting and stopping services, listing them by their type signatures:

#### `hub run(filename string) as (srv service)`

This starts up the given file as a service with the given name. E.g:

```
→ hub run "examples/foo.pf" as "foo"
foo → 
```

The service name `foo` now appears in the prompt to remind you that it's now the "current service", i.e. the service that your input will be passed to by default.

When you start using Pipefish, hotcoding is turned on. That is, when you change the script and then interact with the REPL again, Pipefish will reinitialize the script without you telling it to. So you don't need to keep rerunning your code.

#### `hub run(filename string)`

If you omit the name of the service, and just do `hub run(filename string)`, then the hub will supply a name for the service from the series `#0`, `#1`, `#2` ...

```
→ hub run "examples/bar.pf"
#0 → 
```

#### `hub services`

Starting up one service doesn't stop the other services from running. `hub services` provides you with a list of all the ones that are.

#### `hub switch(srv string)`

This changes the current service to the service given.

#### `hub halt(srv string)`

This halts the given service. (If this is the service you're using, you will find yourself booted off it and instead using the "empty service", i.e. the service you'd get if you compiled a completely empty script.)

#### `hub quit`

This shuts down Pipefish, but remembers every service you had running at shutdown which you gave a specific name to (by using `hub run ... as`) and starts them up again when you restart Pipefish.

#### `hub hot on`

This turns [hotcoding](https://github.com/tim-hardcastle/Pipefish/wiki/Hotcoding) on.

#### `hub hot off`

This turns [hotcoding](https://github.com/tim-hardcastle/Pipefish/wiki/Hotcoding) off.

#### `hub rerun`

This reruns the last thing you ran. Like `hub reset` (below) it is therefore useful only if you have hotcoding turned off.

#### `hub reset`

This reruns the current service. Like `hub reset` (above) it is therefore useful only if you have hotcoding turned off.

These two commands should be rationalized into one sensible one.