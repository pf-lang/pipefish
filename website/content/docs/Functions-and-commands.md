Let's look again at our first example script.

```
cmd

greet :
    post "Hello world!"

def

(n)! :
    n == 0 : 
        1
    n > 0 :
        n * (n - 1)!
    else :
        error "can't take the factorial of a negative number"
```

It has a command, introduced by the heading `cmd`, and a function, introduced by the heading `def`. These are two very different things.

A command *does things*: it moves data in and out of the service: in this case posting data to output. You would also use commands to interact with the file system, a database, the system clock, etc. Commands don't return values: they succeed or fail.

A function *computes things*. It returns a result dependent on the arguments it was called on, and this is all. It has no effect on the outside world, and indeed doesn't know that there *is* an outside world.

Commands can call functions and other commands: functions can call functions but can't call commands.

This distinction is fundamental to Pipefish. It means that your code must end up as an *imperative shell* of interaction with the outside world, and a *functional core* of pure functions expressing business logic.