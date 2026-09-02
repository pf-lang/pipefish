//go:build js && wasm

package main

import (
    "syscall/js"

    "github.com/tim-hardcastle/pipefish/source/pf"
)

var service *pf.Service

func compile(this js.Value, args []js.Value) any {
    service = pf.NewService()

    if err := service.InitializeFromCode(args[0].String()); err != nil {
        return err.Error()
    }

    return nil
}

func do(this js.Value, args []js.Value) any {
    result, err := service.Do(args[0].String())
    if err != nil {
        return err.Error()
    }

    return service.ToString(result)
}

func main() {
    js.Global().Set("pipefishCompile", js.FuncOf(compile))
    js.Global().Set("pipefishDo", js.FuncOf(do))

    select {}
}