Besides the ordinary "vanilla" functions we met on the previous page, Pipefish functions can have pretty much whatever syntax you like. Examples are given in `examples/fancy.pf`:

```
def

(x) squared : x * x

(x) times (y) : x * y

(x) divides (y) :
    y % x == 0

(x) is even:
    2 divides x

say (x) nicely :
    "*~*" + x + "*~*"

divide (x) by (y) : 
    x / y
```

If we demonstrate this in the REPL, it does what you'd think it would do:

```
→ hub run "examples/fancy.pf" as "Fancy"
Starting script 'examples/fancy.pf' as service 'Fancy'.
Fancy → 4 squared
16
Fancy → 6 times 7
42
Fancy → 7 divides 8
false
Fancy → 6 is even 
true
Fancy → 6 is even and 7 divides 8 
false
Fancy → say "hello" nicely 
*~*hello*~*
Fancy → divide 12 by 4 
3
Fancy →  
```

Such functions should be used with care. They are there to clarify code: don't use them to obfuscate it!