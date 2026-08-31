## Types 

### `Prolog = wrapper *Interpreter`
## Commands 

### `add(P Prolog, s string)`

Adds a fact or a rule.  

### `add(P Prolog, S snippet)`

Ditto.  
## Functions 

### `Prolog -> Prolog`

Initializes an empty interpreter.  

### `Prolog(s string) -> Prolog`

Initializes the interpreter from a string.  

### `Prolog(S snippet)`

Initializes the interpreter from a snippet.  

### `query(P Prolog, s string / snippet)`

Returns the result of a query.  

### `check(P Prolog, s string / snippet) -> bool / error`

Tests if solutions to a query exist.  

### `count(P Prolog, s string / snippet) -> int / error`

Counts the number of solutions to a query.  

