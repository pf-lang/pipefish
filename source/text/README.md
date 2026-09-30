This module supplies ways of rendering markdown and doing code highlighting for use in the terminal, the browser, etc.

* `ast` contains the types for the AST of the markdown parser.
* `box_drawing` contains functions for turning box drawings made with regular ASCII into box drawings using Unicode's box drawing characters.
* `highlighter` provides syntax highlighting for Pipefish code.
* `parser` parses markdown into an AST.
* `renderer` renders markdown from the AST.
* `text_native` defines functions like `Yellow` and `Bold` so that they apply terminal control codes to the strings passed to them.
* `text_wasm` defines the same functions to wrap text in HTML.
* `text` defines a few basic text-handling functions.

