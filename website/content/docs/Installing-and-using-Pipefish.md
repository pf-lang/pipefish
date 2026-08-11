## Installation

To install Pipefish, download and unzip [the Pipefish repository](https://github.com/tim-hardcastle/pipefish), then [install Go](https://go.dev/doc/install), and then in your terminal go into the main `pipefish` folder and `go build`. You now have a working copy of Pipefish. Start using it with `./pipefish tui` to open up the TUI, the Text User Interface; the first time you do this it will take a second or two to build some auxiliary files.

And then you will see something like this in your terminal:

<img width="710" height="159" alt="image" src="https://github.com/user-attachments/assets/8b6a6357-d916-4785-b3fd-c3f92d541407" />

This is a REPL. Even though no service is running yet, it will still evaluate expressions using the built-in features of Pipefish.

<img width="710" height="115" alt="image" src="https://github.com/user-attachments/assets/83d3241f-cd07-41a4-85da-67c708fa96c2" />

You can stop it with `hub quit`, or Ctrl+C. We will discuss what the hub is later on.

## How best to use Pipefish

We have supplied Pipefish with a VSCode extension which highlights Pipefish code and takes care of indentation for you. You can add this to VSCode with `cp -r <Main folder of the Pipefish distribution>/pipefish-highlighter ~/.vscode/extensions`. If VSCode is already running, it may be necessary to restart it.

We intend to add support for other editors, but until then, the best way to code in Pipefish is probably to edit your scripts in VSCode while running the Pipefish TUI in a terminal in the bottom panel, like this. We will use this format to demonstrate Pipefish's features in the rest of this wiki.

<img width="710" height="378" alt="image" src="https://github.com/user-attachments/assets/6e4b8ac0-a578-4744-985d-c3a1bab3f74b" />

If you are using VSCode for the first time, note that it ships with autosave turned off. We recommend that you turn it on.