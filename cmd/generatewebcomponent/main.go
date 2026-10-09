//go:build js && wasm

package main

import (
	"path/filepath"
	"sort"
	"syscall/js"

	"github.com/tim-hardcastle/pipefish/source/err"
	"github.com/tim-hardcastle/pipefish/source/filesystem"
	"github.com/tim-hardcastle/pipefish/source/initializer"
	"github.com/tim-hardcastle/pipefish/source/pf"
	"github.com/tim-hardcastle/pipefish/source/text"
	"github.com/tim-hardcastle/pipefish/source/values"
	"github.com/tim-hardcastle/pipefish/source/vm"
	"github.com/tim-hardcastle/pipefish/web-component/generated-go/registry"
)

var (
	service *pf.Service
	fs      *filesystem.VFS
	oH      *vm.CapturingOutHandler
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

	service = pf.NewService().Update(pf.Dependencies{FileSystem: fs})

	if err := service.InitializeFromCode(string(main)); err != nil {
		errorReport, _ := service.GetErrorReport()
		return errorReport
	}
	service.Update(pf.Dependencies{OutHandler: service.MakeCapturingOutHandler()})
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
	service.Update(pf.Dependencies{OutHandler: service.MakeCapturingOutHandler()})
	return nil
}

func do(this js.Value, args []js.Value) any {
	var result values.Value
	if args[0].String() == "" {
		compileMain(this, []js.Value{})
	} else {
		result, _ = service.Do(args[0].String())
	}
	if errorsExist, _ := service.ErrorsExist(); errorsExist {
		errors, _ := service.GetErrorReport()
		return text.GetTuiRenderer(-1)(errors)
	}
	if args[0].String() == "" {
		return ""
	}
	if result.T == pf.ERROR {
		e := result.V.(*pf.Error)
		if e.Message == "" {
			e = err.CreateErr(e.ErrorId, e.Token, e.Args...)
		}
		return text.GetTuiRenderer(-1)("[0] " + text.ERROR + e.Message + err.DescribePos(e.Token) + ".")
	}
	if service.PostHappened() {
		dump, _ := service.Dump()
		return dump
	}
	return service.ToString(result)
}

func getFileTree(this js.Value, args []js.Value) any {
	if fs == nil {
		return "filesystem has not been initialized"
	}

	return buildDirectoryTree(".")
}

func buildDirectoryTree(directory string) (result js.Value) {
	result = js.Global().Get("Object").New()
	result.Set("type", "folder")

	name := directory
	if directory == "." {
		name = ""
	} else {
		name = filepath.Base(directory)
	}
	result.Set("name", name)

	children := js.Global().Get("Array").New()

	directories, err := fs.GetDirectoryNames(directory, false)
	if err != nil {
		return result
	}
	sort.Strings(directories)

	for _, child := range directories {
		children.Call("push", buildDirectoryTree(child))
	}

	files, err := fs.GetFilenames(directory, false)
	if err != nil {
		return result
	}
	sort.Strings(files)

	for _, file := range files {
		child := js.Global().Get("Object").New()
		child.Set("type", "file")
		child.Set("name", filepath.Base(file))
		children.Call("push", child)
	}

	result.Set("children", children)
	return result
}

func readFile(this js.Value, args []js.Value) any {
	result := js.Global().Get("Object").New()

	if fs == nil {
		result.Set("ok", false)
		result.Set("error", "filesystem has not been initialized")
		return result
	}

	data, err := fs.ReadFile(args[0].String())
	if err != nil {
		result.Set("ok", false)
		result.Set("error", err.Error())
		return result
	}

	result.Set("ok", true)
	result.Set("data", string(data))
	return result
}

func createDirectory(this js.Value, args []js.Value) any {
	if fs == nil {
		return "filesystem has not been initialized"
	}

	if err := fs.CreateDirectory(args[0].String()); err != nil {
		return err.Error()
	}
	return nil
}

func deleteFile(this js.Value, args []js.Value) any {
	if fs == nil {
		return "filesystem has not been initialized"
	}

	if err := fs.DeleteFile(args[0].String()); err != nil {
		return err.Error()
	}
	return nil
}

func deleteDirectory(this js.Value, args []js.Value) any {
	if fs == nil {
		return "filesystem has not been initialized"
	}

	if err := fs.DeleteDirectory(args[0].String()); err != nil {
		return err.Error()
	}
	return nil
}

func renamePath(this js.Value, args []js.Value) any {
	if fs == nil {
		return "filesystem has not been initialized"
	}

	oldPath := args[0].String()
	newPath := args[1].String()

	if err := fs.Rename(oldPath, newPath); err != nil {
		return err.Error()
	}
	return nil
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
			return text.BlockHighlighter(
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
			return text.RenderMdAsBookHtml(
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

	js.Global().Set(
		"pipefishGetFileTree",
		js.FuncOf(getFileTree),
	)

	js.Global().Set(
		"pipefishReadFile",
		js.FuncOf(readFile),
	)

	js.Global().Set(
		"pipefishCreateDirectory",
		js.FuncOf(createDirectory),
	)

	js.Global().Set(
		"pipefishDeleteFile",
		js.FuncOf(deleteFile),
	)

	js.Global().Set(
		"pipefishDeleteDirectory",
		js.FuncOf(deleteDirectory),
	)

	js.Global().Set(
		"pipefishRenamePath",
		js.FuncOf(renamePath),
	)

	select {}
}
