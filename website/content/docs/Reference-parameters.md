You have already met reference parameters as an end-user. When in a command you issue an instruction like `get data from File("filename")`, then `data` is not playing the ordinary role of a parameter in a command, to be evaluated and then passed. Instead, the variable `data` itself is being passed. The signature of the command is `get (x ref) from (file File)`, where the `ref` in the signature tells Pipefish to take the variable itself, and not its value, as its argument.

Let's have some simple examples as found in the first few lines of `examples/ref.pf`:

```
var

z = 42

cmd

zero (x ref) :
    x = 0

(x ref) ++ :
    x = x + 1
```

If we try it out in the REPL:

```
→ hub run "examples/ref.pf" as "REF"                                                
Starting script 'examples/ref.pf' as service 'REF'.
REF → z
42
REF → zero z 
ok
REF → z 
0
REF → z ++ 
ok
REF → z 
1 
```

People experienced with this sort of idiom might ask how, if we wanted to, we would _prevent_ the `x`s in the example commands from dereferencing themselves. You can't: why would you ever want to? The only case in which they don't is when the reference parameter is passed to another command which is expecting a reference, as shown here.

```
doThing(x ref) :
    zero x 
    x ++
    post x
```

Only commands have reference parameters; functions don't because the only reason you would want to pass a reference instead of its value is so you could mutate it: but functions have no way of mutating anything, so a `ref` in a function definition would be superfluous and misleading.