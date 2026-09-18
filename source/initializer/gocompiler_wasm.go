//go:build js && wasm

package initializer

import (
	"reflect"

	"github.com/tim-hardcastle/pipefish/source/parser"
	"github.com/tim-hardcastle/pipefish/source/token"
	//"github.com/tim-hardcastle/pipefish/web-component/generated-go/registry"
)

func (iz *Initializer) compileGo() {
	iz.collectGo()

	for source := range iz.goBucket.sources {
		pkg, ok := registry.Packages[source]
		if !ok {
			sourceToken := &token.Token{Source: source}
			iz.throw(
				"golang/wasm",
				sourceToken,
				"no precompiled Go package for source "+source,
			)
			return
		}

		newGoConverter := make(
			[]func(t uint32, v any) any,
			len(iz.cp.Vm.ConcreteTypeInfo),
		)

		copy(
			newGoConverter,
			iz.cp.Vm.GoConverter,
		)

		functionConverter := pkg.FunctionConverter

		for k, v := range BUILTIN_FUNCTION_CONVERTER {
			functionConverter[k] = v
		}

		for typeName, constructor := range functionConverter {
			typeNumber := iz.cp.ConcreteTypeWithNamespaceNow(
				typeName,
			)

			newGoConverter[typeNumber] = constructor
		}

		iz.cp.Vm.GoConverter = newGoConverter

		valueConverter := pkg.ValueConverter

		for k, v := range BUILTIN_VALUE_CONVERTER {
			valueConverter[k] = v
		}

		for typeName, goValue := range valueConverter {
			iz.cp.Vm.GoToPipefishTypes[
				reflect.TypeOf(goValue).Elem(),
			] = iz.cp.ConcreteTypeWithNamespaceNow(
				typeName,
			)
		}

		if pkg.Equals != nil {
			iz.cp.Vm.GoEquals = pkg.Equals
		}

		if pkg.Literal != nil {
			iz.cp.Vm.GoLiteral = pkg.Literal
		}

		for _, function := range iz.goBucket.functions[source] {
			goFunction, ok := pkg.Functions[function.op.Literal]
			if !ok {
				sourceToken := &token.Token{
					Source: source,
				}

				iz.throw(
					"golang/wasm",
					sourceToken,
					"no precompiled Go function for "+function.op.Literal,
				)

				return
			}

			function.body.(*parser.GolangExpression).GoFunction =
				goFunction
		}
	}
}