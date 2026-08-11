If you use the `pf` library for [embedding Pipefish in Go](https://github.com/tim-hardcastle/Pipefish/wiki/Embedding-Pipefish-in-Go), the result you get in the first instance from calling the `Do` method of a `pf.Service` is something of type `pf.Value`. This is a struct having two fields, `T`, of type `pf.Type` (an alias for `int32`) and `V`, of type `any`, containing the actual value of the value, so to speak: the payload.

The service that produces the value also offers three methods (`ToLiteral(v Value)`, `ToString(v Value)`, and `ToGo(v Value, reflect.Type)`) to convert the Pipefish value into something more familiar. If, however, you wish to substitute your own function for handling Pipefish `Value`s, you need to know which `pf.Type`s of `T` correspond to which Golang types of `V`, so that you know what to downcast them to.

Some notes:

In the case of `pf.Error`, `pf.List`, `pf.Map`, and `pf.Set`, these are custom types used by Pipefish and are exposed in the `pf` library along with everything else you're meant to get your hands on. More details of these types will be given at the foot of the page.

The `pf.BLING` type (or pseudotype) in Pipefish is not something you will have seen before. It is the "type" given to infixes and mixfixes when they appear in function calls, and if you poke around in the `pf.Error` type you will sometimes see it returned as part of a list of parameters that caused a runtime error.

The `pf.UNDEFINED_TYPE` is the null element of `pf.Type` and as such is the type of the `Value` returned if the `Do` method of a service returns a non-null error.

At present these types are not entirely stable. Use at your own risk.

| `pf.Type` | Go type |
| --- | --- |
| `UNDEFINED_TYPE` | `nil` |
| `OK` | `nil` |
| `TUPLE` | `[]pf.Value` |
| `ERROR` | `*pf.Error` |
| `NULL` | `nil` |
| `INT` | `int` |
| `BOOL` | `bool` |
| `STRING` | `string` |
| `RUNE` | `rune` |
| `FLOAT` | `float64` |
| `PAIR` | `[]pf.Value` |
| `LIST` | `pf.List` |
| `MAP` | `pf.Map` |
| `SET` | `pf.Set` |
| `LABEL` | `string` |
| any struct | `[]pf.Value` |
| any enum | `int` |
| any clone | the same as its parent |

While the custom types have a number of methods, they are currently unstable, and so we shall mention only those necessary for extracting data from them. For this purpose:

* The `pf.List` type has methods `.Len() int` and `.Index(i int) any`.
* The `pf.Set` type has a method `.ToSlice() []Value`.
* The `pf.Map` type has a method `.ToSlice() [][]Value`.
* The `pf.Error` type is a struct defined as follows:

```
type Error struct {
	ErrorId string      // The error ID.
	Message string      // The error message.
	Args    []any       // Further arguments (if any) from which the error message was constructed.
	Values  []Value     // Usually empty, but may contain the specific values that caused a runtime error.
	Trace   []*Token    // A list of `Token`s tracing the path of a runtime error up the stack.
	Token   *Token      // A `Token` indicting the point in the sourcecode where an error originates.
}
```

The `pf.Token` type is defined as follows:

```
type Token struct {
	Type    TokenType   // A type based on `string`.
	Literal string      // The literal value of the token.
	Line    int         // The line number where it occurred.
	ChStart int         // The position in the line where it started.
	ChEnd   int         // The position in the line where it ended.
	Source  string      // The file it came from.
}
```