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

var (
	service *pf.Service
	fs      *filesystem.VFS
)

func compile(this js.Value, args []js.Value) any {
	files := args[0]

	fs = filesystem.NewVFS()

	for i := 0; i < files.Length(); i++ {
		file := files.Index(i)

		path := file.Get("path").String()
		dataJS := file.Get("data")

		data := make([]byte, dataJS.Get("byteLength").Int())
		js.CopyBytesToGo(data, dataJS)

		if err := fs.WriteFile(path, data); err != nil {
			return err.Error()
		}
	}

	main, err := fs.ReadFile("main.pf")
	if err != nil {
		return err.Error()
	}

	service = pf.NewService(fs)
	service.SetFileSystem(fs)

	if err := service.InitializeFromCode(string(main)); err != nil {
		return err.Error()
	}

	return nil
}

func updateFile(this js.Value, args []js.Value) any {
	path := args[0].String()
	data := []byte(args[1].String())

	if err := fs.WriteFile(path, data); err != nil {
		return err.Error()
	}

	return nil
}

func compileMain(this js.Value, args []js.Value) any {
	main, err := fs.ReadFile("main.pf")
	if err != nil {
		return err.Error()
	}

	if err := service.InitializeFromCode(string(main)); err != nil {
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

	js.Global().Set(
		"pipefishUpdateFile",
		js.FuncOf(updateFile),
	)

	js.Global().Set(
		"pipefishCompileMain",
		js.FuncOf(compileMain),
	)

	select {}
}