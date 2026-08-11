## Introducing the hub

To run our first example script, we told the TUI `hub run "examples/wiki/first.pf"`. That is we ran the script by telling the hub to run the script for us.

The hub is a Pipefish service which is always running inside the TUI and manages other Pipefish services: it can start them, stop them, and supplies you with a lot of useful tooling: too much, in fact, to cover in this introductory article. To get you started here are some of the more essential commands.

### `hub run (filename string)`

As we've seen, this will initialize a service from the given script, and will give it a suitable name. If you want to name the service yourself, use `hub run "<filename>" as "<service name>"`.

You can have more than one service running at a time. If, having run our first service, we now wish to run a second, we could for example do this:

<img width="710" height="365" alt="image" src="https://github.com/user-attachments/assets/094de55f-2a50-48b9-b022-6f8482743ba9" />

This brings us on to the next useful hub command.

### `hub services`

This lists all the services the hub is running.

<img width="710" height="145" alt="image" src="https://github.com/user-attachments/assets/fb337589-91cd-4d7e-9e37-e7daf6912a4e" />

### `hub switch (servicename string)`

This changes the current service, i.e. the one you're talking to through the TUI. As you'll have noticed, the name of the current service shows up in the prompt.

<img width="710" height="110" alt="image" src="https://github.com/user-attachments/assets/40eb0a11-eb60-4478-b50f-bbc1ad1263d4" />

### `hub halt (servicename string)`

This halts the named service.

<img width="710" height="169" alt="image" src="https://github.com/user-attachments/assets/939b70ea-e23d-47ba-91cd-4ece1b23ff04" />

### `hub api`

This gives the API of the current service.

<img width="710" height="226" alt="image" src="https://github.com/user-attachments/assets/171393b3-b181-4751-bbd0-4c1f7e974d24" />

### `hub quit`

This shuts down the hub and the TUI, but remembers every service you had running at shutdown, and starts them up again when you restart the TUI.