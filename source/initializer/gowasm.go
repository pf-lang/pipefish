package initializer

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/tim-hardcastle/pipefish/source/settings"
	"github.com/tim-hardcastle/pipefish/source/values"
)

type wasmGoPackageInfo struct {
	source         string
	packageName    string
	functions      []string
	hasEquals      bool
	hasLiteral     bool
}

func GenerateWasmGoFromSource(
	scriptFilepath string,
	sourcecode string,
	outputDirectory string,
) error {
	settings.PipefishHomeDirectory = "."

	iz := NewInitializer(
		NewCommonInitializerBindle(values.Map{}, nil),
	)

	iz.prepareForCompilation(scriptFilepath, sourcecode, nil)

	if iz.errorsExist() {
		return fmt.Errorf(
			"errors while preparing %q:\n%s",
			scriptFilepath,
			iz.P.ReturnErrors(),
		)
	}

	return iz.generateWasmGoModules(outputDirectory, nil)
}

func GenerateWasmGo(iz *Initializer, outputDirectory string) error {
	var packages []wasmGoPackageInfo

	if err := iz.generateWasmGoModules(outputDirectory, &packages); err != nil {
		return err
	}

	return generateWasmGoRegistry(outputDirectory, packages)
}

func (iz *Initializer) generateWasmGoModules(
	outputDirectory string,
	packages *[]wasmGoPackageInfo,
) error {
	for pair := iz.initializers.Oldest(); pair != nil; pair = pair.Next() {
		if err := pair.Value.generateWasmGoModules(outputDirectory, packages); err != nil {
			return err
		}
	}

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

		if packages != nil {
			functions := make([]string, 0, len(iz.goBucket.functions[source]))

			for _, function := range iz.goBucket.functions[source] {
				functions = append(
					functions,
					function.op.Literal,
				)
			}

			*packages = append(
				*packages,
				wasmGoPackageInfo{
					source:      wasmGoSourcePath(source),
					packageName: packageName,
					functions:   functions,
					hasEquals:   strings.Contains(goSource, "func Equals("),
					hasLiteral:  strings.Contains(goSource, "func Literal("),
				},
			)
		}
	}

	return nil
}

func wasmGoSourcePath(source string) string {
	path := filepath.ToSlash(MakeFilepath(source))

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return path
}

func GenerateWasmGoStandardLibraries(
	outputDirectory string,
) error {
	settings.PipefishHomeDirectory = "."

	var packages []wasmGoPackageInfo

	for source := range StandardLibraries {

		filename := filepath.Join(
			"source",
			"initializer",
			"libraries",
			source+".pf",
		)

		sourcecode, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf(
				"reading standard library %q from %q: %w",
				source,
				filename,
				err,
			)
		}

		iz := NewInitializer(
			NewCommonInitializerBindle(values.Map{}, nil),
		)

		iz.prepareForCompilation(source, string(sourcecode), nil)

		if iz.errorsExist() {
			return fmt.Errorf(
				"errors while preparing %q:\n%s",
				source,
				iz.P.ReturnErrors(),
			)
		}

		if err := iz.generateWasmGoModules(
			outputDirectory,
			&packages,
		); err != nil {
			return fmt.Errorf(
				"generating Go for standard library %q: %w",
				source,
				err,
			)
		}
	}

	return generateWasmGoRegistry(
		outputDirectory,
		packages,
	)
}

func generateWasmGoRegistry(
	outputDirectory string,
	packages []wasmGoPackageInfo,
) error {
	// The same package can be encountered more than once because
	// standard-library initializers recursively generate their
	// dependencies. Keep only one registry entry for each source.
	unique := make(map[string]wasmGoPackageInfo)

	for _, packageInfo := range packages {
		unique[packageInfo.source] = packageInfo
	}

	packages = packages[:0]

	for _, packageInfo := range unique {
		packages = append(packages, packageInfo)
	}

	sort.Slice(
		packages,
		func(i, j int) bool {
			return packages[i].source < packages[j].source
		},
	)

	var sb strings.Builder

	fmt.Fprintln(&sb, "package registry")
	fmt.Fprintln(&sb)
	fmt.Fprintln(&sb, "import (")

	fmt.Fprintln(
		&sb,
		"\t\"reflect\"",
	)

	fmt.Fprintln(
		&sb,
		"\t\"github.com/tim-hardcastle/pipefish/source/initializer\"",
	)

	for _, packageInfo := range packages {
		fmt.Fprintf(
			&sb,
			"\t%s %q\n",
			packageInfo.packageName,
			"github.com/tim-hardcastle/pipefish/web-component/generated-go/"+packageInfo.packageName,
		)
	}

	fmt.Fprintln(&sb, ")")
	fmt.Fprintln(&sb)

	fmt.Fprintln(
		&sb,
		"var Packages = map[string]initializer.WasmGoPackage{",
	)

	for _, packageInfo := range packages {
		fmt.Fprintf(
			&sb,
			"\t%q: {\n",
			packageInfo.source,
		)

		fmt.Fprintf(
			&sb,
			"\t\tFunctionConverter: %s.PIPEFISH_FUNCTION_CONVERTER,\n",
			packageInfo.packageName,
		)

		fmt.Fprintf(
			&sb,
			"\t\tValueConverter: %s.PIPEFISH_VALUE_CONVERTER,\n",
			packageInfo.packageName,
		)

		if packageInfo.hasEquals {
			fmt.Fprintf(
				&sb,
				"\t\tEquals: %s.Equals,\n",
				packageInfo.packageName,
			)
		}

		if packageInfo.hasLiteral {
			fmt.Fprintf(
				&sb,
				"\t\tLiteral: %s.Literal,\n",
				packageInfo.packageName,
			)
		}

		fmt.Fprintln(
			&sb,
			"\t\tFunctions: map[string]reflect.Value{",
		)

		for _, function := range packageInfo.functions {
			fmt.Fprintf(
				&sb,
				"\t\t\t%q: reflect.ValueOf(%s.%s),\n",
				function,
				packageInfo.packageName,
				capitalize(function),
			)
		}

		fmt.Fprintln(&sb, "\t\t},")
		fmt.Fprintln(&sb, "\t},")
	}

	fmt.Fprintln(&sb, "}")
	fmt.Fprintln(&sb)

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

	if err := os.MkdirAll(registryDirectory, 0755); err != nil {
		return fmt.Errorf(
			"creating registry directory %q: %w",
			registryDirectory,
			err,
		)
	}

	filename := filepath.Join(
		registryDirectory,
		"registry.go",
	)

	if err := os.WriteFile(filename, formatted, 0644); err != nil {
		return fmt.Errorf(
			"writing generated WASM Go registry %q: %w",
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