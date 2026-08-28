## Introducing the hub

To run our first example script, we told the TUI `hub run "examples/docs/first.pf"`. That is we ran the script by telling the hub to run the script for us.

The hub is a Pipefish service which is always running inside the TUI and manages other Pipefish services: it can start them, stop them, and supplies you with a lot of useful tooling: too much, in fact, to cover in this introductory article. To get you started here are some of the more essential commands.

### `hub run (filename string)`

As we've seen, this will initialize a service from the given script, and will give it a suitable name. If you want to name the service yourself, use `hub run "&lt;filename>" as "&lt;service name>"`.

You can have more than one service running at a time. The file `examples/docs/second` defines a Fibonnacci function, and so if after running `first` we go ahead and run `second`, it will start up a second service.

```tui
first → hub run "examples/docs/second.pf"
Starting script <span class="service">"second.pf"</span> as service <span class="service">"second"</span>.
second → fib 20
6765
second →
```

This brings us on to the next useful hub command.

### `hub services`

This lists all the services the hub is running.

```tui
second → hub services
The hub is running the following services:

  <span class="bullet">▪</span> Service <span class="service">"first"</span> running script <span class="service">"first.pf"</span>. 
  <span class="bullet">▪</span> Service <span class="service">"second"</span> running script <span class="service">"second.pf"</span>. 

second →
```

### `hub switch (servicename string)`

This changes the current service, i.e. the one you're talking to through the TUI. As you'll have noticed, the name of the current service shows up in the prompt.

```tui
second → hub switch "first"
OK
first → 10!
3628800
first →    
```

### `hub halt (servicename string)`

This halts the named service.

```tui
first → hub halt "second"
OK
first → hub services
The hub is running the following services:

first → hub halt "second"
OK
first → hub services
The hub is running the following services:

  ▪ Service "first" running script "first.pf". 

first →  
```

### `hub api`

This gives the API of the current service.

```tui
first → hub api

≡≡≡≡ first ≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡

════ Commands ═════════════════════════════════════════════════════════════════════════════

• greet :

Says hello to the world. 

════ Functions ════════════════════════════════════════════════════════════════════════════

• (n any?)! :

A factorial function. 
```

### `hub quit`

This shuts down the the hub, the services, and the TUI. but remembers every service you had running at shutdown, and starts them up again when you run Pipefish again.