The `for` loops and piping operators in Pipefish should keep you from having to use recursion unless it is really necessary. But when you do want to, you can. The file `examples/recursion.pf` contains a recursive factorial function and a recursive Fibonacci function:

```
def

factorial (n int) : 
    n == 0 : 
        1
    else n * factorial n - 1 

fib(n) : 
    n == 1 or n == 2 : 
        1
    else : 
        fib(n - 1) + fib(n - 2)
```