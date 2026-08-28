## Web-facing

Every Pipefish script from `hello world` onwards is *de facto* a webservice: deploying it is a matter of configuration rather than code, and can be done as simply as typing `hub http` into the TUI.

A Pipefish service's public functions and commands are also its web-facing API: if you have a valid username and password for someone else's Pipefish service, your code can use it semantically and syntactically as though it was an library.