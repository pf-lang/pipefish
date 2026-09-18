package initializer

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/tim-hardcastle/pipefish/source/settings"
	"github.com/tim-hardcastle/pipefish/source/values"
)

func GenerateWasmGoFromSource(
	scriptFilepath string,
	sourcecode string,
	outputDirectory string,
) error {
	settings.PipefishHomeDirectory = "."

	iz := NewInitializer(
		NewCommonInitializerBindle(values.Map{}, nil),
	)

	iz.prepareForCompilation(scriptFilepath, sourcecode)

	if iz.errorsExist() {
		return fmt.Errorf(
			"errors while preparing %q:\n%s",
			scriptFilepath,
			iz.P.ReturnErrors(),
		)
	}

	return GenerateWasmGo(iz, outputDirectory)
}

type wasmGoPackageInfo struct {
	source      string
	packageName string
	functions   []string
	hasEquals   bool
	hasLiteral  bool
}

func GenerateWasmGo(
	iz *Initializer,
	outputDirectory string,
) error {
	var packages []wasmGoPackageInfo

	if err := iz.generateWasmGoModules(
		outputDirectory,
		&packages,
	); err != nil {
		return err
	}

	return generateWasmGoRegistry(
		outputDirectory,
		packages,
	)
}

func (iz *Initializer) generateWasmGoModules(
	outputDirectory string,
	packages *[]wasmGoPackageInfo,
) error {
	for pair := iz.initializers.Oldest(); pair != nil; pair = pair.Next() {
		if err := pair.Value.generateWasmGoModules(
			outputDirectory,
			packages,
		); err != nil {
			return err
		}
	}

	iz.collectGo()

	for source := range iz.goBucket.sources {
		packageName := wasmGoPackageName(source)

		goSource, ok := iz.generateGoSource(
			source,
			packageName,
		)
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

		if err := os.MkdirAll(
			packageDirectory,
			0755,
		); err != nil {
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

		if err := os.WriteFile(
			filename,
			formatted,
			0644,
		); err != nil {
			return fmt.Errorf(
				"writing generated Go source %q: %w",
				filename,
				err,
			)
		}

		info, err := inspectWasmGoPackage(
			filename,
			source,
			packageName,
			iz.goBucket.functions[source],
		)
		if err != nil {
			return err
		}

		*packages = append(*packages, info)
	}

	return nil
}

func inspectWasmGoPackage(
	filename string,
	source string,
	packageName string,
	functions []*parsedFunction,
) (wasmGoPackageInfo, error) {
	file, err := parser.ParseFile(
		token.NewFileSet(),
		filename,
		nil,
		0,
	)
	if err != nil {
		return wasmGoPackageInfo{}, fmt.Errorf(
			"parsing generated Go package %q: %w",
			source,
			err,
		)
	}

	var hasEquals bool
	var hasLiteral bool

	for _, declaration := range file.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if !ok {
			continue
		}

		switch function.Name.Name {
		case "Equals":
			hasEquals = true
		case "Literal":
			hasLiteral = true
		}
	}

	functionNames := make([]string, 0, len(functions))

	for _, function := range functions {
		functionNames = append(
			functionNames,
			capitalize(function.op.Literal),
		)
	}

	return wasmGoPackageInfo{
		source:      source,
		packageName: packageName,
		functions:   functionNames,
		hasEquals:   hasEquals,
		hasLiteral:  hasLiteral,
	}, nil
}

func generateWasmGoRegistry(
	outputDirectory string,
	packages []wasmGoPackageInfo,
) error {
	var sb strings.Builder

	fmt.Fprintln(&sb, "package registry")
	fmt.Fprintln(&sb)
	fmt.Fprintln(&sb, "import (")
	fmt.Fprintln(&sb, "\t\"reflect\"")

	for _, pkg := range packages {
		fmt.Fprintf(
			&sb,
			"\t%s %q\n",
			pkg.packageName,
			"github.com/tim-hardcastle/pipefish/web-component/generated-go/"+pkg.packageName,
		)
	}

	fmt.Fprintln(&sb, ")")
	fmt.Fprintln(&sb)

	fmt.Fprintln(&sb, "type Package struct {")
	fmt.Fprintln(
		&sb,
		"\tFunctionConverter map[string](func(t uint32, v any) any)",
	)
	fmt.Fprintln(&sb, "\tValueConverter map[string]any")
	fmt.Fprintln(&sb, "\tEquals func(x any, y any) bool")
	fmt.Fprintln(&sb, "\tLiteral func(x any) string")
	fmt.Fprintln(&sb, "\tFunctions map[string]reflect.Value")
	fmt.Fprintln(&sb, "}")
	fmt.Fprintln(&sb)

	fmt.Fprintln(
		&sb,
		"var Packages = map[string]Package{",
	)

	seen := make(map[string]bool)

	for _, pkg := range packages {
		if seen[pkg.source] {
			continue
		}
		seen[pkg.source] = true

		fmt.Fprintf(&sb, "\t%q: {\n", pkg.source)

		fmt.Fprintf(
			&sb,
			"\t\tFunctionConverter: %s.PIPEFISH_FUNCTION_CONVERTER,\n",
			pkg.packageName,
		)

		fmt.Fprintf(
			&sb,
			"\t\tValueConverter: %s.PIPEFISH_VALUE_CONVERTER,\n",
			pkg.packageName,
		)

		if pkg.hasEquals {
			fmt.Fprintf(
				&sb,
				"\t\tEquals: %s.Equals,\n",
				pkg.packageName,
			)
		}

		if pkg.hasLiteral {
			fmt.Fprintf(
				&sb,
				"\t\tLiteral: %s.Literal,\n",
				pkg.packageName,
			)
		}

		fmt.Fprintln(
			&sb,
			"\t\tFunctions: map[string]reflect.Value{",
		)

		for _, function := range pkg.functions {
			fmt.Fprintf(
				&sb,
				"\t\t\t%q: reflect.ValueOf(%s.%s),\n",
				strings.ToLower(function[:1])+function[1:],
				pkg.packageName,
				function,
			)
		}

		fmt.Fprintln(&sb, "\t\t},")
		fmt.Fprintln(&sb, "\t},")
	}

	fmt.Fprintln(&sb, "}")

	formatted, err := format.Source([]byte(sb.String()))
	if err != nil {
		return fmt.Errorf(
			"formatting generated WASM Go registry: %w",
			err,
		)
	}

	registryDirectory := filepath.Join(
		outputDirectory,
		"registry",
	)

	if err := os.MkdirAll(
		registryDirectory,
		0755,
	); err != nil {
		return fmt.Errorf(
			"creating WASM Go registry directory %q: %w",
			registryDirectory,
			err,
		)
	}

	filename := filepath.Join(
		registryDirectory,
		"registry.go",
	)

	if err := os.WriteFile(
		filename,
		formatted,
		0644,
	); err != nil {
		return fmt.Errorf(
			"writing WASM Go registry %q: %w",
			filename,
			err,
		)
	}

	return nil
}

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