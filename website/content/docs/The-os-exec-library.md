## Overview 

The `exec` library supplies access to OS commands.  

## Modules 

### ` ("os/exec")`

### ` ("github.com/google/shlex")`
## Types 

### `Os = clone snippet`

A snippet for containing fragments of OS commands and values to inject into them.  
## Commands 

### `get(x ref) from (command Os)`

Populates the reference variable with the result of executing the command.  

### `post to (command Os)`

Executes the command.  
## Functions 

### `string(command Os)`

