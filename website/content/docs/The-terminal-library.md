## Overview 

The `terminal` library supplies functions which read from the keyboard, and write to the terminal, *of the machine the hub is running on*. It will *not* post to the terminal of a remote user: for that you use the built-in `post ... to Output()` method. And *nothing* will actively solicit input from a remote user. 

`get ... from Keyboard()` is therefore a poison pill; you can only use it for desktop apps, you can't expose it to the web.  

## Types 

### `Keyboard = struct (prompt string)`

A data structure wrapping a prompt to ask the user for keyboard input.  

### `Terminal = struct `

A data structure to pass to `post ... to` to say that the destination is the terminial.  
## Commands 

### `post(x any?) to (s Terminal)`

Posts the given value to the terminal.  

### `get(x ref) from (k Keyboard)`

Prompts the user for input using `k[prompt]` and then uses the result to populate the reference variable.  

### `get(x ref) from masked (k Keyboard)`

Prompts the user for input using `k[prompt]`, masking the input for privacy, and then uses the result to populate the reference variable.  

