package pf_path_2ffilepath

import (
	"path/filepath"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){}

var PIPEFISH_VALUE_CONVERTER = map[string]any{}

func Abs(path string) any {
	r, e := filepath.Abs(path)
	if e == nil {
		return r
	}
	return e
}

func Base(path string) string {
	return filepath.Base(path)
}

func Clean(path string) string {
	return filepath.Clean(path)
}

func Dir(path string) string {
	return filepath.Dir(path)
}

func EvalSymlinks(path string) any {
	r, e := filepath.EvalSymlinks(path)
	if e == nil {
		return r
	}
	return e
}

func Ext(path string) string {
	return filepath.Ext(path)
}

func FromSlash(path string) string {
	return filepath.FromSlash(path)
}

func IsAbs(path string) bool {
	return filepath.IsAbs(path)
}

func IsLocal(path string) bool {
	return filepath.IsLocal(path)
}

func Join(elem ...string) string {
	return filepath.Join(elem...)
}

func Localize(path string) any {
	r, e := filepath.Localize(path)
	if e == nil {
		return r
	}
	return e
}

func Match(pattern string, name string) any {
	r, e := filepath.Match(pattern, name)
	if e == nil {
		return r
	}
	return e
}

func Rel(basepath string, targpath string) any {
	r, e := filepath.Rel(basepath, targpath)
	if e == nil {
		return r
	}
	return e
}

func Split(path string) (string, string) {
	return filepath.Split(path)
}

func SplitList(path string) []any {
	p := filepath.SplitList(path)
	r := make([]any, len(p), len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

func ToSlash(path string) string {
	return filepath.ToSlash(path)
}
