## Installation

To install Pipefish, download and unzip <a href="https://github.com/tim-hardcastle/pipefish">the Pipefish repository</a>, then install <a href="https://go.dev/doc/install">Go</a>, and then in your terminal go into the main `pipefish` folder and `go build`. You now have a working copy of Pipefish. Start using it with `./pipefish tui` to open up the TUI, the Text User Interface; the first time you do this it will take a second or two to build some auxiliary files.

And then you will see something like this in your terminal:

```tui

  ╔═══════════════════════════╗
  ║ 🧿 Pipefish version 0.7.0 ║
  ╚═══════════════════════════╝

Pipefish is running hub "user".

No services are running on this hub.

→
```

This is a REPL. Even though no service is running yet, it will still evaluate expressions using the built-in features of Pipefish, as you'll see if you try it out.

```tui
→ 2 + 2
4
→ len "aardvark"
8
→
```

You can stop the TUI with `hub quit`, or with Ctrl+C. We'll talk about what the "hub" is in the next page of this documentation.

## How best to use Pipefish

We have supplied Pipefish with a VSCode extension which highlights Pipefish code and takes care of indentation for you. You can add this to VSCode with `cp -r &ltMain folder of the Pipefish distribution>/pipefish-highlighter ~/.vscode/extensions`. If VSCode is already running, it may be necessary to restart it.

We intend to add support for other editors, but until then, the best way to code in Pipefish is probably to edit your scripts in VSCode while running the Pipefish TUI in a terminal; the one at the bottom of the VSCode window if you have a small screen.

If you are using VSCode for the first time, note that it ships with autosave turned off. We recommend that you turn it on.