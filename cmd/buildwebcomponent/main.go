package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/tim-hardcastle/pipefish/source/initializer"
)

func main() {
	outputDirectory := filepath.Join("web-component", "generated-go")

	// Make sure we start clean.
	if err := os.RemoveAll("web-component"); err != nil {
		panic(fmt.Errorf("removing old web-component directory: %w", err))
	}

	// The generated Go packages are needed only while the WASM binary
	// is being compiled.
	defer func() {
		if err := os.RemoveAll("web-component"); err != nil {
			fmt.Fprintf(os.Stderr, "warning: couldn't remove web-component: %v\n", err)
		}
	}()

	if err := initializer.GenerateWasmGoStandardLibraries(outputDirectory); err != nil {
		panic(fmt.Errorf("generating WASM Go libraries: %w", err))
	}

	cmd := exec.Command(
		"go",
		"build",
		"-o",
		filepath.Join("web", "assets", "pipefish.wasm"),
		"./cmd/generatewebcomponent",
	)

	cmd.Env = append(os.Environ(),
		"GOOS=js",
		"GOARCH=wasm",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(fmt.Errorf("building WASM web component: %w", err))
	}
}