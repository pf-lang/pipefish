## Introduction

One of the problems with pure functions is seeing what's going on inside them. Fortunately it's only the tiniest blemish on the purity of our functions if we do just that. Logging/instrumentation markup is provided by the `\\` operator, as demonstrated in `examples/log.pf`:

## Manual logging

Let's look at the first function in `examples/log.pf`.

```
foo(x) :            \\ 'foo' called.
    x % 2 == 0  :   \\ Checks if x is even. 
        x           \\ Returns |x|.
    else :          \\ Else branch taken.
        twice x     \\ Returns ||twice x||.
given :             \\
    twice(n) :      \\ 'twice' called.
        two * n     \\ Calculates ||2 * n||.
    two = 2         \\ Lazy assignment to two.
```

Run it in the REPL ...

```
LOG → hub run "examples/log.pf"       
Starting script 'examples/log.pf' as service '#0'.
#0 → foo 8  
Log at line 3:
    'foo' called.
Log at line 4:
    Checks if x is even.
Log at line 5:
    Returns 8.
8 
#0 → foo 9 
Log at line 3:
    'foo' called.
Log at line 4:
    Checks if x is even.
Log at line 6:
    Else branch taken.
Log at line 7:
    Returns 'twice x' = 18.
Log at line 9:
    'twice' called.
Log at line 10:
    Calculates '2 * n' = 18.
Log at line 11:
    Lazy assignment to two.
18     
#0 → 
```

Note that a log statement is executed if the line it is adjoined to is executed. In particular, in a line of the form `<condition> : <return value> \\ <log statement>`, the logging will take place whether or not the condition is met. If you want to avoid that, then you can of course write:

```
<condition> :
    <return value> \\ <log statement>
```

You can adjoin a logging statement to a line saying only `given :` but this will never happen because there is no sense in which the `given` keyword is ever executed / evaluated. Logging statements attached to inner functions will execute when they are called as with normal functions. Logging statements attached to local constants will be executed when the assigned value is actually evaluated, which happens at need.

## Autologging

If you leave the logging statement empty, as in the otherwise identical function `bar` ...

```
bar(x) :            \\ 
    x % 2 == 0  :   \\ 
        x           \\ 
    else :          \\ 
        twice x     \\ 
given :             \\
    twice(n) :      \\ 
        two * n     \\ 
    two = 2         \\ 
```

... then Pipefish will infer what should be logged, as shown in the REPL:

```
#0 → bar 9 
Log at line 13:
    Called 'bar' with 'x' = 9.
Log at line 14:
    (x % 2) is 1, so the condition fails.
Log at line 16:
    The 'else' branch is taken.
Log at line 17:
    Returning 'twice x' = 18.
Log at line 19:
    Called 'twice' with 'n' = 9.
Log at line 20:
    Returning 'two * n' = 18.
18 
#0 →  
```
## `$logTime` and `$logPath`

The formatting and output destination of the logging can be chosen by setting service variables, as shown in the REPL:

```
#0 → $logTime = true  
ok 
#0 → $logPath = "./rsc/test.log" 
ok 
#0 → foo 4 
4 
#0 → os cat ./rsc/test.log  
Log at line 3 @ 2023-11-11 21:31:55.685226 -0800 PST:
    'foo' called.
Log at line 4 @ 2023-11-11 21:31:55.687571 -0800 PST:
    Checks if x is even.
Log at line 5 @ 2023-11-11 21:31:55.687868 -0800 PST:
    Returns 4.
#0 →   
```