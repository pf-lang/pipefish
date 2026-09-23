//go:build js && wasm

package main

import (
	"syscall/js"

	"github.com/tim-hardcastle/pipefish/source/filesystem"
	"github.com/tim-hardcastle/pipefish/source/initializer"
	"github.com/tim-hardcastle/pipefish/source/markdown"
	"github.com/tim-hardcastle/pipefish/source/pf"
	"github.com/tim-hardcastle/pipefish/web-component/generated-go/registry"
)

var service *pf.Service

func compile(this js.Value, args []js.Value) any {
	fs, _ := filesystem.NewVFSFromDirectory(args[1].String())
	service = pf.NewService(fs)
	service.SetFileSystem(fs)
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
	initializer.RegisterWasmGoPackages(registry.Packages)

	js.Global().Set(
		"pipefishCompile",
		js.FuncOf(compile),
	)

	js.Global().Set(
		"pipefishDo",
		js.FuncOf(do),
	)

	js.Global().Set(
		"pipefishHighlight",
		js.FuncOf(func(
			this js.Value,
			args []js.Value,
		) interface{} {
			return markdown.BlockHighlighter(
				args[0].String(),
			)
		}),
	)

	js.Global().Set(
		"pipefishRenderMdAsHtml",
		js.FuncOf(func(
			this js.Value,
			args []js.Value,
		) interface{} {
			return markdown.RenderMdAsHtml(
				args[0].String(),
			)
		}),
	)

	js.Global().Set(
		"pipefishWasmReady",
		true,
	)

	select {}
}