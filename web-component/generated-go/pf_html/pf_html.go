package pf_html

import (
	"html"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){}

var PIPEFISH_VALUE_CONVERTER = map[string]any{}

func EscapeString(s string) string {
	return html.EscapeString(s)
}

func UnescapeString(s string) string {
	return html.UnescapeString(s)
}
