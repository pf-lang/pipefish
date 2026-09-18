package pf_path

import (
	"errors"
	"path"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){}

var PIPEFISH_VALUE_CONVERTER = map[string]any{}

func Base(p string) string {
	return path.Base(p)
}

func Clean(p string) string {
	return path.Clean(p)
}

func Dir(p string) string {
	return path.Dir(p)
}

func Ext(p string) string {
	return path.Ext(p)
}

func IsAbs(p string) bool {
	return path.IsAbs(p)
}

func Join(t []any) any {
	strings := []string{}
	for _, v := range t {
		switch v := v.(type) {
		case string:
			strings = append(strings, v)
		default:
			return errors.New("'join' function in 'path' library passed non-string type")
		}
	}
	return path.Join(strings...)
}

func Match(pattern string, name string) any {
	m, err := path.Match(pattern, name)
	if err != nil {
		return err
	}
	return m
}

func Split(p string) (string, string) {
	return path.Split(p)
}
