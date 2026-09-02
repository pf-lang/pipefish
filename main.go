//go:build !js && !wasm

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tim-hardcastle/pipefish/source/hub"
	"github.com/tim-hardcastle/pipefish/source/settings"
	"github.com/tim-hardcastle/pipefish/source/text"
)

func main() {
	if len(os.Args) == 1 {
		showhelp()
		return
	}
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-h", "--help", "help":
			showhelp()
			return
		case "-v", "--version", "version":
			os.Stdout.WriteString("\nPipefish version " + text.VERSION + ".\n\n")
			return
		case "-r", "--run", "run":
			hub.StartServiceFromCli()
		case "-t", "--tui", "tui": // Left blank to avoid the default.
		case "-w", "--w", "wiki":  
			hub.GetWiki()
		default:
			os.Stdout.WriteString("\nPipefish doesn't recognize the command '" + os.Args[1] + "'.\n")
			println()
			showhelp()
			os.Exit(1)
		}
	}

	fmt.Print(text.Logo())
	bytes, _ := os.ReadFile(filepath.Join(settings.PipefishHomeDirectory, ("user/hub.dat")))
	filename := string(bytes)
	if filepath.IsLocal(filename) {
		filepath.Join(settings.PipefishHomeDirectory, filename)
	}
	h := hub.New(filename, os.Stdout)
	h.Repl()
}

func showhelp() {
	os.Stdout.WriteString(hub.HELP)
}
