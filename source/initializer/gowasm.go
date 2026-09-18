package initializer

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
)

func GenerateWasmGo(iz *Initializer, outputDirectory string) error {
	return iz.generateWasmGoModules(outputDirectory)
}

func (iz *Initializer) generateWasmGoModules(
	outputDirectory string,
) error {
	// Generate dependencies first, matching compileGoModules().
	for pair := iz.initializers.Oldest(); pair != nil; pair = pair.Next() {
		if err := pair.Value.generateWasmGoModules(outputDirectory); err != nil {
			return err
		}
	}

	// Collect the Go declarations and functions belonging to this
	// initializer.
	iz.collectGo()

	for source := range iz.goBucket.sources {
		packageName := wasmGoPackageName(source)

		goSource, ok := iz.generateGoSource(source, packageName)
		if !ok {
			return fmt.Errorf(
				"could not generate Go source for %q",
				source,
			)
		}

		formatted, err := format.Source([]byte(goSource))
		if err != nil {
			return fmt.Errorf(
				"formatting generated Go for %q: %w",
				source,
				err,
			)
		}

		packageDirectory := filepath.Join(
			outputDirectory,
			packageName,
		)

		if err := os.MkdirAll(packageDirectory, 0755); err != nil {
			return fmt.Errorf(
				"creating generated Go package directory %q: %w",
				packageDirectory,
				err,
			)
		}

		filename := filepath.Join(
			packageDirectory,
			packageName+".go",
		)

		if err := os.WriteFile(filename, formatted, 0644); err != nil {
			return fmt.Errorf(
				"writing generated Go source %q: %w",
				filename,
				err,
			)
		}
	}

	return nil
}

// wasmGoPackageName converts a Pipefish Go source name into a
// deterministic valid Go package identifier.
func wasmGoPackageName(source string) string {
	var b strings.Builder

	b.WriteString("pf_")

	for _, r := range source {
		switch {
		case r >= 'a' && r <= 'z',
			r >= 'A' && r <= 'Z',
			r >= '0' && r <= '9':
			b.WriteRune(r)

		default:
			fmt.Fprintf(&b, "_%02x", r)
		}
	}

	return b.String()
}
