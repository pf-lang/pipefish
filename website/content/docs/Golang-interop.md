Pipefish is implemented in Go, and so has features for interoperating with Go that are different from the more generic ways in which you can interoperate with other languages.

The basic idea is that you can write the signature of a function in Pipefish, and then use the `gocode` keyword to write the body of the function/command in Go. The same keyword allows you to import a Go library.



## A Pipefish function with a Golang body.

Consider for example the following function:

```
def

fib(n int) : golang {
	var a, b int
	b = 1
	for i := 0; i < n; i++ {
		a += b
		a, b = b, a
	}
	return a
}
```

This has its signature written in Pipefish, but its body written in Go. Think of the `golang` keyword as a `:`-to-`{` converter.

Such a function needs to have the types of its parameters specified. If `int` was left out of the example above, it wouldn't work. Go is a static language: it needs to know what type it's getting.

There is no need to specify return types (though as usual you can). What the function returns is a Pipefish value, and Pipefish is a dynamic language.

## Pure Go functions

In addition, you may want to write helper functions that are entirely in Go. You can do this by writing a block beginning with `golang` and delimited by braces:

```
golang {
    func double(x int) int {
        return 2 * x
    }
}
```

## Importing from Go

We can import from Go in the `import` section of a script just by prefixing the `golang` keyword to whatever we want to import. For example Pipefish's own `strings` library begins with:

```
import

golang "strings"
```

Having done that, we can then write: 

```
contains(haystack, needle string) : golang {
    return strings.Contains(haystack, needle)
}
```

You will see from this example that it is very easy to wrap a Pipefish library around a Golang library.

## Automagical type sharing

If a function with a Golang body refers to a user-defined type in its signature, then the Go code will be able to use this type. This is true even if the reference is implicit, i.e. if the type is of a field of a struct mentioned in the signature, etc. To explain by example, if we have types declared as follows:

```
newtype

Dragon = struct(name string, color Color, temperature Temperature)
Color = enum RED, GREEN, GOLD, BLACK
Temperature = clone int
```

... then the following function will work: the Go code will know about the `Dragon` and `Color` and `Temperature` types, and the result of the function will be a Pipefish value of type `Color`.

```
def

// Returns the color of the hottest dragon.
dragonFight(x, y Dragon) : golang {
	if x.Temperature >= y.Temperature {
		return x.Color
	}
	return y.Color
}
```

Note that the names of the fields of the `Dragon` struct are capitalized in the Go code, so that they can be exported.