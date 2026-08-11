Given the existence of Golang interop and of reference variables, it is possible to write your own `get` and `post` statements to interact with whatever any Golang library can interact with.

In fact, most of Pipefish's IO is presently implemented that way, in a library called `rsc/pipefish/world.go` which is automatically imported with every Pipefish script. So for example to find out if a file exists, we import the Go `os` library ...

```
import

golang "os"
```

... and we write a tiny Pipefish function ...

```
get (x ref) from (fileAccess FileExists) :
    x = goFileExists(fileAccess[filepath])
```

... which wraps around an almost equally tiny Go function:

```
goFileExists(fname string) : golang {
    _, err := os.Stat(fname)
    return err == nil
}
```

And that's all there is to it.