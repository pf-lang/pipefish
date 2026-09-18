package pf_fmt

import (
	"fmt"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){}

var PIPEFISH_VALUE_CONVERTER = map[string]any{}

func Errorf(format string, a ...any) any {
	return fmt.Errorf(format, a...)
}

func Sprint(a ...any) string {
	return fmt.Sprint(a...)
}

func Sprintf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

func Sprintln(a ...any) string {
	return fmt.Sprintln(a...)
}
