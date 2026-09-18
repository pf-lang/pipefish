//go:build js && wasm

package initializer

import (
	"embed"
	"io/fs"
	"os"
	"reflect"
	"strings"

	"github.com/tim-hardcastle/pipefish/source/parser"
	"github.com/tim-hardcastle/pipefish/source/compiler"
	"github.com/tim-hardcastle/pipefish/source/token"
)

var wasmGoPackages map[string]WasmGoPackage

func RegisterWasmGoPackages(
	packages map[string]WasmGoPackage,
) {
	wasmGoPackages = packages
}

func (iz *Initializer) compileGo() {
	iz.collectGo()

	for source := range iz.goBucket.sources {
		sourceToken := &token.Token{Source: source}

		pkg, ok := wasmGoPackages[source]
		if !ok {
			iz.throw(
				"golang/compile",
				sourceToken,
				"no precompiled WASM Go package for "+source,
			)
			continue
		}

		newGoConverter := make(
			[](func(t uint32, v any) any),
			len(iz.cp.Vm.ConcreteTypeInfo),
		)

		copy(
			newGoConverter,
			iz.cp.Vm.GoConverter,
		)

		functionConverter := make(
			map[string](func(t uint32, v any) any),
			len(pkg.FunctionConverter),
		)

		for k, v := range pkg.FunctionConverter {
			functionConverter[k] = v
		}

		for k, v := range BUILTIN_FUNCTION_CONVERTER {
			functionConverter[k] = v
		}

		for typeName, constructor := range functionConverter {
			typeNumber := iz.cp.ConcreteTypeWithNamespaceNow(typeName)
			newGoConverter[typeNumber] = constructor
		}

		iz.cp.Vm.GoConverter = newGoConverter

		if pkg.Equals != nil {
			iz.cp.Vm.GoEquals = pkg.Equals
		}

		if pkg.Literal != nil {
			iz.cp.Vm.GoLiteral = pkg.Literal
		}

		valueConverter := make(
			map[string]any,
			len(pkg.ValueConverter),
		)

		for k, v := range pkg.ValueConverter {
			valueConverter[k] = v
		}

		for k, v := range BUILTIN_VALUE_CONVERTER {
			valueConverter[k] = v
		}

		for typeName, goValue := range valueConverter {
			iz.cp.Vm.GoToPipefishTypes[
				reflect.TypeOf(goValue).Elem(),
			] = iz.cp.ConcreteTypeWithNamespaceNow(typeName)
		}

		for _, function := range iz.goBucket.functions[source] {
			goFunction, ok := pkg.Functions[function.op.Literal]

			if !ok {
				iz.throw(
					"golang/function",
					sourceToken,
					"no precompiled function for "+function.op.Literal,
				)
				continue
			}

			function.body.(*parser.GolangExpression).GoFunction = goFunction
		}
	}
}

func GetSourceCode(scriptFilepath string) (string, error) {
	var sourcebytes []byte
	var err error

	if scriptFilepath != "" {
		if sourcecode, ok := wasmStandardLibrarySources[scriptFilepath]; ok {
			sourcebytes = []byte(sourcecode)
		} else if len(scriptFilepath) >= 11 && scriptFilepath[:11] == "test-files/" {
			sourcebytes, err = compiler.TestFolder.ReadFile(scriptFilepath)
		} else {
			sourcebytes, err = os.ReadFile(MakeFilepath(scriptFilepath))
		}

		if err != nil {
			return "", err
		}
	}

	sourcebytes = append(sourcebytes, '\n')
	return string(sourcebytes), nil
}

//go:embed libraries
var wasmStandardLibraryFiles embed.FS

var wasmStandardLibrarySources = make(map[string]string)

func init() {
	err := fs.WalkDir(
		wasmStandardLibraryFiles,
		"libraries",
		func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if entry.IsDir() || !strings.HasSuffix(path, ".pf") {
				return nil
			}

			data, err := wasmStandardLibraryFiles.ReadFile(path)
			if err != nil {
				return err
			}

			source := strings.TrimPrefix(path, "libraries/")
			wasmStandardLibrarySources[
				"/source/initializer/libraries/"+source,
			] = string(data)

			return nil
		},
	)

	if err != nil {
		panic(err)
	}
}