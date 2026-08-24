## Web-facing

Every Pipefish script from `hello world` onwards is *de facto* a webservice: deploying it is a matter of configuration rather than code, and a service's public functions and commands are also its web-facing API, so that if you have a valid username and password for another Pipefish service, your code can use it semantically and syntactically as though it was an imported library.