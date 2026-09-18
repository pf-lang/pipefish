package pf_markdown

import (
	"github.com/tim-hardcastle/pipefish/source/text"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){}

var PIPEFISH_VALUE_CONVERTER = map[string]any{}

func GoRender(leftMargin string, width int, raw string) string {
	return text.NewMarkdown(leftMargin, width, func(s string) string { return s }).RenderString(raw)
}
