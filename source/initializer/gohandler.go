package initializer

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"strconv"
	"strings"

	"github.com/tim-hardcastle/pipefish/source/dtypes"
	"github.com/tim-hardcastle/pipefish/source/settings"
	"github.com/tim-hardcastle/pipefish/source/text"
	"github.com/tim-hardcastle/pipefish/source/token"
	"github.com/tim-hardcastle/pipefish/source/values"
	"github.com/tim-hardcastle/pipefish/source/vm"
)

// This allows the compiler to extract functions and converter data from the relevant `.so` files,
// rebuilding them if necessary.
//
//The code for generating the .go source code when a build/rebuild is necessary is kept in the
// `gogen.go` file in this same `service` package. This file is mainly devoted to housekeeping.
//
// When multiple sources are NULL-imported into a module, we must still treat each source with Go
// in it separately, otherwise for example we would have to produce a separate compilation of the
// `strings` library every time it was null-namespaed. Or `world`, which would be every time you
// run Pipefish.

// While the Pipefish needs to be recompiled each time the app it's in is recompiled, the Go doesn't,
// because the .so file hooks in just the same. Hence we can use the `gotimes.dat` file to keep track
// of when a given file was created. TODO --- this (or the practical equivalent is already metadata.
// The only additional purpose the `gotimes` file serves is that the time gives you a unique way of
// mangling the filename and the the file allows you to save the nae you came up with. But it seems like
// there should be a better way. (Ideally involving Google fixing the problem with .so files but I'm
// not holding my breath.))

var counter int // This variable is used to make a unique filename for each gocode_<counter>.go file.

// This struct type is used to accumulate the various data encountered during parsing that we need to
// build a `.go` file or files. There is one per compiler. The "sources" are as given in the `Source`
// field of any GOCODE token encountered. The sources may be plural because of NULL-imports.
// The `imports`, `functions`, `pureGo`, and `types` maps are indexed by these sources.
type goBucket struct {
	sources   dtypes.Set[string]
	imports   map[string][]string
	functions map[string][]*parsedFunction
	pureGo    map[string][]string
	types     map[string][]values.ValueType
}

func (iz *Initializer) newGoBucket() {
	gb := goBucket{
		sources:   make(dtypes.Set[string]),
		imports:   make(map[string][]string),
		functions: make(map[string][]*parsedFunction),
		pureGo:    make(map[string][]string),
		types:     make(map[string][]values.ValueType),
	}
	iz.goBucket = &gb
}

// Maps for converting base types.
var BUILTIN_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){
	"bool":   func(t uint32, v any) any { return v.(bool) },
	"float":  func(t uint32, v any) any { return v.(float64) },
	"int":    func(t uint32, v any) any { return v.(int) },
	"rune":   func(t uint32, v any) any { return v.(rune) },
	"string": func(t uint32, v any) any { return v.(string) },
}

var BUILTIN_VALUE_CONVERTER = map[string]any{
	"bool":   (*bool)(nil),
	"float":  (*float64)(nil),
	"int":    (*int)(nil),
	"rune":   (*rune)(nil),
	"string": (*string)(nil),
}

type types = dtypes.Set[values.ValueType]

// This makes a new .so file, opens it, and returns the plugins.
// Most of the code generation is in the `gogen.go` file in this same `initializer` package.
func (iz *Initializer) makeNewSoFile(source string, newTime int64) *plugin.Plugin {
	sourceToken := &token.Token{Source: source}
	iz.cmG("Making golang from source '"+source+"'\n\n", source)
	var StringBuilder strings.Builder
	sb := &StringBuilder
	// We emit the package declaration and builtins.
	fmt.Fprint(sb, "package main\n\n")
	if len(iz.goBucket.imports) > 0 {
		fmt.Fprint(sb, "import (\n")
		for _, v := range iz.goBucket.imports[source] {
			fmt.Fprint(sb, "    \""+v+"\"\n")
		}
		fmt.Fprint(sb, ")\n\n")
	}
	// We extract all the types we're going to need to declare.
	// TODO --- do we need guards here on the types?
	userDefinedTypes := make(types)
	for _, function := range iz.goBucket.functions[source] {
		for _, v := range function.sig {
			name := withoutDots(v.VarType.String())
			// Note: casting the type info to `BuiltinType`` won't have the same effect,
			// since it will include $_ types.
			if name == "any" || name == "any?" {
				continue
			}
			abType := iz.cp.GetAbstractTypeFromAstType(v.VarType)
			for _, conc := range abType.Types {
				if _, ok := iz.cp.Vm.ConcreteTypeInfo[conc].(vm.BuiltinType); ok || iz.cp.Vm.ConcreteTypeInfo[conc].GetName(vm.LITERAL) == "Hub" {
					continue
				}
				userDefinedTypes.Add(conc)
			}
		}
		for _, pair := range function.callInfo.ReturnTypes {
			abType := iz.cp.GetAbstractTypeFromAstType(pair.VarType)
			name := withoutDots(pair.VarType.String())
			if name == "any" || name == "any?" {
				continue
			}
			for _, conc := range abType.Types {
				if _, ok := iz.cp.Vm.ConcreteTypeInfo[conc].(vm.BuiltinType); ok {
					continue
				}
				userDefinedTypes.Add(conc)
			}
		}
	}
	for _, v := range iz.goBucket.types[source] {
		userDefinedTypes.Add(v)
	}
	iz.transitivelyCloseTypes(userDefinedTypes)
	if iz.errorsExist() {
		return nil
	}
	// We emit the type declarations and converters.
	iz.generateDeclarations(sb, userDefinedTypes)
	// And the functions.
	for _, function := range iz.goBucket.functions[source] {
		iz.generateGoFunctionCode(sb, function)
	}
	// And any blocks of pure Go.
	for _, pureGo := range iz.goBucket.pureGo[source] {
		fmt.Fprint(sb, pureGo)
	}
	counter++ // The number of the gocode_<counter>.go source file we're going to write.
	soFile := filepath.Join(settings.PipefishHomeDirectory, filepath.FromSlash("source/initializer/gobucket/"+text.Flatten(source)+"_"+strconv.Itoa(int(newTime))+".so"))
	timeMap := iz.getGoTimes()
	if oldTime, ok := timeMap[source]; ok {
		os.Remove(filepath.Join(settings.PipefishHomeDirectory, filepath.FromSlash("source/initializer/gobucket/"+text.Flatten(source)+"_"+strconv.Itoa(int(oldTime))+".so")))
	}
	goFile := filepath.Join(settings.PipefishHomeDirectory, "gocode_"+strconv.Itoa(counter)+".go")
	iz.cmG("Creating goFile with filepath '"+goFile+"'\n\n", source)
	file, err := os.Create(goFile)
	if err != nil {
		iz.throw("golang/create", sourceToken, err.Error())
		return nil
	}
	file.WriteString(sb.String())
	iz.cmG("*************GENERATED GO IS*************\n\n"+sb.String()+"*****************************************\n\n", source)
	file.Close()
	if settings.SHOW_GOLANG && !(settings.MandatoryImportSet()).Contains(source) {
		println("Creating soFile with filepath '" + soFile + "'\n\n")
	}
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", soFile, goFile) // Version to use running from terminal.
	// cmd := exec.Command("go", "build", "-gcflags=all=-N -l", "-buildmode=plugin", "-o", soFile, goFile) // Version to use with debugger.
	output, err := cmd.Output()
	if err != nil {
		iz.throw("golang/build", sourceToken, err.Error()+": "+string(output))
		return nil
	}
	plugins, err := plugin.Open(soFile)
	if err != nil {
		iz.throw("golang/open.a", sourceToken, err.Error())
		return nil
	}
	// We do this here and not earlier with defer because a .go file that doesn't compile should
	// be visible for debugging.
	os.Remove(goFile)
	timeMap[source] = newTime
	iz.recordGoTimes(timeMap)
	return plugins
}

// This makes sure that if  we're generating declarations for a struct type,
// we're also generating declarations for the types of its fields if need be, and so on recursively. We do
// a traditional non-recursive breadth-first search.
func (iz *Initializer) transitivelyCloseTypes(userDefinedTypes types) {
	structsToCheck := types{}
	for ty := range userDefinedTypes {
		if iz.cp.Vm.ConcreteTypeInfo[ty].IsStruct() {
			structsToCheck.Add(ty)
		}
	}
	for len(structsToCheck) > 0 {
		newStructsToCheck := make(types)
		for ty := range structsToCheck {
			structInfo := iz.cp.Vm.ConcreteTypeInfo[ty].(vm.StructType)
			for i, fieldType := range structInfo.AbstractStructFields {
				if fieldType.Len() != 1 {
					iz.throw("golang/concrete", INTEROP_TOKEN, iz.cp.Vm.DescribeAbstractType(fieldType, vm.LITERAL, 0), structInfo.Name, i)
					structsToCheck = newStructsToCheck
					continue
				}
				typeOfField := fieldType.Types[0]
				switch iz.cp.Vm.ConcreteTypeInfo[typeOfField].(type) {
				case vm.CloneType, vm.EnumType:
					userDefinedTypes.Add(typeOfField)
				case vm.StructType:
					if !userDefinedTypes.Contains(typeOfField) {
						newStructsToCheck.Add(typeOfField)
						userDefinedTypes.Add(typeOfField)
					}
				}
			}
			structsToCheck = newStructsToCheck
		}
	}
}

func (iz *Initializer) recordGoTimes(timeMap map[string]int64) {
	f, err := os.Create(filepath.Join(settings.PipefishHomeDirectory + "source/initializer/gobucket/gotimes.dat"))
	if err != nil {
		panic("Can't create file gotimes.dat")
	}
	defer f.Close()
	for k, v := range timeMap {
		f.WriteString(k + "\n")
		f.WriteString(strconv.Itoa(int(v)) + "\n")
	}
}

func (iz *Initializer) getGoTimes() map[string]int64 {
	timeMap := make(map[string]int64)
	pathToGoResourceDirectory := filepath.Join(settings.PipefishHomeDirectory + "source/initializer/gobucket/")
	os.Mkdir(pathToGoResourceDirectory, os.ModePerm) // We may be using Pipefish as a library and this needs creating. Will do nothing if the directory exists.
	pathToGoTimes := filepath.Join(pathToGoResourceDirectory, "gotimes.dat")
	file, err := os.Open(pathToGoTimes)
	if err != nil {
		if os.IsNotExist(err) {
			os.Create(pathToGoTimes)
			return timeMap
		}
		panic("Can't open file '" + pathToGoTimes + "'.")
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < (len(lines) / 2); i++ {
		time, _ := strconv.Atoi(lines[(2*i)+1])
		timeMap[lines[2*i]] = int64(time)
	}
	return timeMap
}

// Creates comments on the gohandler and gogen when settings.SHOW_GOLANG is true.
func (iz *Initializer) cmG(text, source string) {
	if settings.SHOW_GOLANG && !(settings.MandatoryImportSet()).Contains(source) {
		println(text)
	}
}

func withoutDots(s string) string {
	if text.Head(s, "...") {
		return s[3:]
	} else {
		return s
	}
}
