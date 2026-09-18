package pf_strconv

import (
	"strconv"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){}

var PIPEFISH_VALUE_CONVERTER = map[string]any{}

func FormatFloat(f float64, fmt rune, prec int) string {
	return strconv.FormatFloat(f, byte(fmt), prec, 64)
}

func FormatInt(i int, base int) string {
	return strconv.FormatInt(int64(i), base)
}

func GoParseInt(s string, base int) any {
	r, e := strconv.ParseInt(s, base, 64)
	if e == nil {
		return int(r)
	}
	return e
}
