package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tim-hardcastle/pipefish/source/initializer"
)

func main() {
	outputDirectory := filepath.Join(
		"web-component",
		"generated-go",
	)

	if err := os.RemoveAll(outputDirectory); err != nil {
		panic(fmt.Errorf(
			"removing old generated Go: %w",
			err,
		))
	}

	if err := os.MkdirAll(outputDirectory, 0755); err != nil {
		panic(fmt.Errorf(
			"creating generated Go directory: %w",
			err,
		))
	}

	if err := initializer.GenerateWasmGoStandardLibraries(
		outputDirectory,
	); err != nil {
		panic(err)
	}

	fmt.Printf(
		"Generated WASM Go packages in %s\n",
		outputDirectory,
	)
}