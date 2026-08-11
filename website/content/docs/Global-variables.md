## Commands and global variables

Earlier in this wiki we introduced global variables and saw how we could change them from the REPL.

We can also access and change global variables from commands, but if we wish to do so we must bring them into the scope of the command using the `global` keyword. An example is given in `examples/globals.pf`:

```
var

animal = "aardvark"

cmd

setAnimal(s string) :
    global animal
    animal = s 

getAnimal :
    global animal
    post animal
```

This does what you think it would do, as we can demonstrate in the REPL:

```
→ hub run "examples/globals.pf" as "Glob"
Starting script 'examples/globals.pf' as service 'Glob'.
Glob → getAnimal
aardvark
Glob → setAnimal "zebra" 
ok
Glob → getAnimal
zebra
Glob → 
```

## Encapsulation

As usual we can make variables private from the REPL by using the word `private`. This affects everything from there on down until either the next headword or the end of the script. Let's change the script above by inserting `private` before the variable declaration.

```
var

private

animal = "aardvark"

cmd

setAnimal(s string) :
    global animal
    animal = s 

getAnimal :
    global animal
    post animal
```

And now we can access the variable only through the commands:

```
Glob → animal 

[0] Error: attempt to access the value of a private or non-existent variable or
constant 'animal' at line 1:0-6 of REPL input

Glob → animal = "whale"                                                                              

[0] Error: attempt to assign the value of a private or non-existent variable or
constant 'animal' at line 1:7 of REPL input

Glob → setAnimal "whale"
ok
Glob → getAnimal 
whale
Glob → 
```