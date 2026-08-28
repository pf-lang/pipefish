Let's look at an example script. It has a command `greet`, for greeting people, and a factorial function.

```
cmd

~~ Says hello to the world.
greet :
    post "Hello world!"

def

~~ A factorial function.
(n)! :
    n == 0 : 
        1
    n > 0 :
        n * (n - 1)!
    else :
        error "can't take the factorial of a negative number"
```

We can see it working in the TUI. The script has no `main` command, but when we run the script it will declare `greet` and `!` and we'll be able to use them from the REPL.

```tui
→ hub run "examples/docs/first.pf"
Starting script <span class="service">"first.pf"</span> as service <span class="service">"first"</span>. 
first → greet
Hello world!
first → 5!
120
first → (-5)!
[0] <span class="error-message">Error</span>: can't take the factorial of a negative number at line <span class="line-no">16:8-13</span> of <span class="service">"examples/docs/first.pf"</span>.  
first →
```

Hopefully that's just what you expected it to do. But you may be wondering what exactly is meant by `hub run`, as we haven't mentioned the hub yet. This will be dealt with on the next page.