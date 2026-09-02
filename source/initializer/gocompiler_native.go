//go:build !js && !wasm

package initializer

import (
	"os"
	"plugin"
	"reflect"
	"strconv"

	"github.com/tim-hardcastle/pipefish/source/parser"
	"github.com/tim-hardcastle/pipefish/source/settings"
	"github.com/tim-hardcastle/pipefish/source/text"
	"github.com/tim-hardcastle/pipefish/source/token"
)

// This will if necessary compile or recompile the relevant .so files, and will extract from them
// the functions and converter data needed by the compiler and vm and put it into its proper place.
func (iz *Initializer) compileGo() {
	// The purpose of putting timestamps on the .so files is not that we ever read the timestamps
	// from the filenames (we look either at the OS metadata of the file or at the 'gotimes' file),
	// but simply because you can't re-use the names of .so files in the same Go runtime and since
	// we're looking up times anyway this is a reasonable way to achieve that.

	// We get the blocks of pure Go, if any, and put them in the appropriate place in the goBucket.
	for _, tc := range iz.tokenizedCode[golangDeclaration] {
		golang := tc.(*tokenizedGolangDeclaration)
		iz.goBucket.sources.Add(golang.goCode.Source)
		iz.goBucket.pureGo[golang.goCode.Source] = append(iz.goBucket.pureGo[golang.goCode.Source],
			golang.goCode.Literal)
	}

	// And the Go types declared by `wrapper` in the `newtype` section.
	for _, tc := range iz.tokenizedCode[wrapperDeclaration] {
		wrapper := tc.(*tokenizedWrapperDeclaration)
		iz.goBucket.sources.Add(wrapper.op.Source)
		iz.goBucket.types[wrapper.op.Source] = append(iz.goBucket.types[wrapper.op.Source],
			iz.cp.ConcreteTypeNow(wrapper.op.Literal))
	}

	for j := functionDeclaration; j <= commandDeclaration; j++ {
		for _, pc := range iz.parsedCode[j] {
			fn := pc.(*parsedFunction)
			if fn.body.GetToken().Type == token.GOLANG {
				iz.goBucket.sources.Add(fn.op.Source)
				iz.goBucket.functions[fn.op.Source] = append(iz.goBucket.functions[fn.op.Source], fn)
			}
		}
	}

	timeMap := iz.getGoTimes() // We slurp a map from sources to times from the `gotimes` file.

	for source := range iz.goBucket.sources {
		sourceToken := &token.Token{Source: source}
		f, err := os.Stat(MakeFilepath(source))
		if err != nil {
			iz.throw("golang/file", sourceToken, source, err.Error())
			break
		}
		var plugins *plugin.Plugin
		sourceCodeModified := f.ModTime().UnixMilli()
		objectCodeModified, ok := timeMap[source]
		if !ok || sourceCodeModified != int64(objectCodeModified) {
			plugins = iz.makeNewSoFile(source, sourceCodeModified)
		} else {
			soFile := settings.PipefishHomeDirectory + "source/initializer/gobucket/" + text.Flatten(source) + "_" + strconv.Itoa(int(sourceCodeModified)) + ".so"
			plugins, err = plugin.Open(soFile)
			if err != nil {
				iz.throw("golang/open.b", sourceToken, err.Error())
				return
			}
		}
		if plugins == nil { // Then the Go has failed to compile.
			iz.throw("golang/compile", sourceToken)
			return
		}

		// We extract the conversion data from the object code, reformat it, and store the results
		// in the vm.
		newGoConverter := make([](func(t uint32, v any) any), len(iz.cp.Vm.ConcreteTypeInfo))
		copy(newGoConverter, iz.cp.Vm.GoConverter)
		functionConverterSymbol, _ := plugins.Lookup("PIPEFISH_FUNCTION_CONVERTER")
		functionConverter := *functionConverterSymbol.(*map[string](func(t uint32, v any) any))
		if equalsFunctionSymbol, err := plugins.Lookup("Equals"); err == nil {
			iz.cp.Vm.GoEquals = equalsFunctionSymbol.(func(x any, y any) bool)
		}
		if literalFunctionSymbol, err := plugins.Lookup("Literal"); err == nil {
			iz.cp.Vm.GoLiteral = literalFunctionSymbol.(func(x any) string)
		}
		for k, v := range BUILTIN_FUNCTION_CONVERTER {
			functionConverter[k] = v
		}
		for typeName, constructor := range functionConverter {
			typeNumber := iz.cp.ConcreteTypeWithNamespaceNow(typeName)
			newGoConverter[typeNumber] = constructor
		}
		iz.cp.Vm.GoConverter = newGoConverter
		valueConverterSymbol, _ := plugins.Lookup("PIPEFISH_VALUE_CONVERTER")
		valueConverter := *valueConverterSymbol.(*map[string]any)
		for k, v := range BUILTIN_VALUE_CONVERTER {
			valueConverter[k] = v
		}
		for typeName, goValue := range valueConverter {
			iz.cp.Vm.GoToPipefishTypes[reflect.TypeOf(goValue).Elem()] = iz.cp.ConcreteTypeWithNamespaceNow(typeName)
		}
		//We attach the compiled functions to the (pointers to) the functions, which are
		// also pointed to by the function table and by the list of common functions
		// in the common parser bindle. I.e. we are returning our result by mutating the
		// functions.
		for _, function := range iz.goBucket.functions[source] {
			goFunction, _ := plugins.Lookup(capitalize(function.op.Literal))
			function.body.(*parser.GolangExpression).GoFunction = reflect.ValueOf(goFunction)
		}
	}
}