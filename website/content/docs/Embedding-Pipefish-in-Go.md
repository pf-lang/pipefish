## Introduction

By importing the appropriate part of the Pipefish distribution (`github.com/tim-hardcastle/Pipefish/source/pf`) as a library, you can embed Pipefish in your own applications, rather than accessing it via the TUI or a client.

The various methods etc are documented formally [by godev](https://pkg.go.dev/github.com/tim-hardcastle/Pipefish/source/pf), so here we will give a tutorial on how to use them in practice.

Many of the available features are only necessary for the most sophisticated users, and you may just want to skip down to [the example](https://github.com/tim-hardcastle/Pipefish/wiki/Embedding-Pipefish-in-Go#example) at the foot of the page and copy that.

## Initialization

Having imported the `pf` library, we request a new service using a function with the signature `func NewService() *Service`.

At this point, optionally:

* If you want the service to use a database, this is the point at which you should assign it, since it might be used by `init` commands in the code. You do this with a method having signature `func (sv *Service) SetDatabase(db *sql.DB)`.

* If you want the service to reference *local* external services (as though they were on the same hub) you can do this by passing it a map of strings to services using a method with signature `func (sv *Service) SetLocalExternalServices(svs map[string]*Service)`; again, if you want to do this, you should do so at this stage, before any `init` command can assume that these services exist. Remote external services work as usual.

Having done either of those things if necessary, we can now compile the service. You can do this either using `func (sv *Service) InitializeFromCode(code string) error` or `func (sv *Service) InitializeFromFilepath(scriptFilepath string) error`, the names of which are more or less self-explanatory.

`func (sv *Service) NeedsUpdate() (bool, error)` will return `true` if the source code files for a service have been changed since the service was first initialized.

## Evaluating expressions, and what to do with the values

Then either you can perform `func (s *Service) CallMain() (values.Value, error)`, the name of which explains itself); or to evaluate an expression or perform any other command, you use the method `func (sv *Service) Do(line string) (Value, error)`. 

This will return an non-nil error if there are *compile-time* errors. If intead the line compiles successfully but there is a *runtime error*, then that will be returned as the `Value` together with a `nil` error.

As you can see, what it returns is a *Pipefish* value, of type `pf.Value`, which is a struct having two fields: `T` of type `pf.Type`, and `V` of type `any`. You have four choices of what to do with it besides leaving it as it is.

* Use the method `func (sv *Service) ToLiteral(v Value) string` to convert it into a Pipefish literal.

* Use the method `func (sv *Service) ToString(v Value) string` to apply the `string` method of the service to it.

* Use the method `func (sv *Service) ToGo (pfValue Value, goType reflect.Type) (any, error)` to convert it to a Go value of the given type, if possible.

* Deal with it yourself. To this end we have supplied constants `UNDEFINED_TYPE`, `BLING`, `OK`, `TUPLE`, `ERROR`, `NULL`, `INT`, `BOOL`, `STRING`, `RUNE`, `TYPE`, `FUNC`, `PAIR`, `LIST`, `MAP`, `SET`, `LABEL`. There are methods `func (sv *Service) IsClone(v Value) bool`, `func (sv *Service) IsEnum(v Value) bool`, and `func (sv *Service) IsStruct(v Value) bool`, which do what you think they would do, and a function `func (sv *Service) UnderlyingType(v Value) Type` which returns the parent type of a type if it's a clone, and its own type otherwise. Also, the `pf` library exposes the internal types used to represent Pipefish `Values`. See [Appendix C](https://github.com/tim-hardcastle/Pipefish/wiki/Appendix-C:-Internal-representation-of-Pipefish-types) of this wiki for further details of the internal representation of Pipefish values.

## Errors

To deal with error conditions, we have the following methods:

* `func (sv *Service) IsBroken() bool` will return `true` if no script has been compiled yet, or if the compilation failed.

* `func (sv *Service) ErrorsExist() bool` will return `true` if the last thing the service did produced errors, whether this was a result of trying to compile the script; of trying to compile a line using `Do`; or of a runtime error caused by executing the line.

* `func (sv *Service) GetErrors() []*Error` will get a list of errors.

* `func (sv *Service) GetErrorReport() (string, error)` gets the errors summarized in a string.

* `func ExplainError(es []*Error, i int) (string, error)` will get you the same explanation you'd get from `hub why`.

* `func GetTraceReport(e *err.Error) string` extracts the trace from a runtime error.

* `func (sv *Service) GetTrackingReport() (string, error)` gets the tracking data, if any.

* `func PrettyString(s string, left, right int) string` converts the strings returned by the previous four methods from markup to highlighted text.

## InHandlers and OutHandlers

An `InHandler` is an interface to a type defining what happens when the Pipefish code performs `get x from Input("<prompt>"). It has one method, `Get() string`.

An `OutHandler` is an interface to a type defining what happens when the Pipefish code performs `post x to Output()`. It has two methods: `Out(v Value)` and `Write(s string)`, the first of which will typically serialize the `Value` in some way before writing it to some output, and the second of which will typically write the string to the same output.

You can set them using the methods `func (sv *Service) SetInHandler(in InHandler) error` and `func (sv *Service) SetOutHandler(out OutHandler) error`.

To help construct them, there are the following methods and functions:

* `func MakeSimpleInHandler(in io.Reader) *SimpleInHandler`. Returns an InHandler that will get its string from the `io.Reader` and do nothing else.

* `func MakeTerminalInHandler(prompt string) *TerminalInHandler`. Returns an `Inhandler` that will get its string from the terminal, prompting the end-user with the given prompt.

* `func (sv *Service) MakeLiteralWritingOutHandler(out io.Writer) *SimpleOutHandler`. Returns an `OutHandler` that will write the literal of the passed value to the given `io.Writer`.

* `func (sv *Service) MakeStringWritingOutHandler(out io.Writer) *SimpleOutHandler`. Returns an `OutHandler` that will write the string representation of the passed value to the given `io.Writer`.

* `func (sv *Service) MakeLiteralTerminalOutHandler() *SimpleOutHandler`. Returns an `OutHandler` that will write the literal of the passed value to the terminal.

* `func (sv *Service) MakeStringTerminalOutHandler() *SimpleOutHandler`. Returns an `OutHandler` that will write the string representation of the passed value to the terminal.

## Miscellaneous

* `func (sv *Service) GetVariable(vname string) (values.Value, error)` gets the value of a global variable. Unlike doing this using the `Do` method, `GetVariable` gives you access to private variables.

* `func (sv *Service) SetVariable(vname string, ty values.ValueType, v any) error` sets the value of a global variable. Unlike doing this using the `Do` method, `SetVariable` gives you access to private variables.

* `func (sv *Service) GetFilepath() (string, error)` gets the path to the root file of the service.

* `func (sv *Service) GetSources() (map[string][]string, error)` gets the source code of the service as a map from filenames to lists of strings (i.e. lines of source code).

* `func (sv *Service) TypeToTypeName(t Type) (string, error)`: does what it says: converts something of type `pf.Type` to a Go string naming the type, e.g. converting `pf.BOOL` to the string `"bool"`.

* `func (sv *Service) TypeNameToType(s string) (Type, error)` reversed the previous process and gets a `pf.Type` from a string representing the type name.

## Example

As a complete if simple example, the following Go code will tell you that the eighth Fibonacci number is 21.

```
package main

import (
	"reflect"
	"github.com/tim-hardcastle/Pipefish/source/pf"
)

const fibCode = `// Pipefish code for computing the nth Fibonnaci number.

def

fib(n int) :
    first from a, b = 0, 1 for i = 0; i < n; i + 1 :
        b, a + b	
`
func main() {
	fibService := pf.NewService()
	fibService.InitializeFromCode(fibCode)
	pfResult, _ := fibService.Do(`fib 8`)
	goResult, _ := fibService.ToGo(pfResult, reflect.TypeFor[int]())
	println("The eighth Fibonacci number is", goResult.(int))
}
```