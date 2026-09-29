//go:build !js && !wasm

package markdown 

func GetTuiRenderer(width int) func(string)string {
	return NewTerminalRenderer(Iota, width).Render
}

func NewTerminalRenderer(highlighter func(string) string, width int) Renderer {
	return NewRenderer(MakeRenderFunction(getTerminalRenderer(92), highlighter))
}