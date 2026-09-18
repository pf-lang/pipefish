package pf_strings

import (
	"errors"
	"strings"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){}

var PIPEFISH_VALUE_CONVERTER = map[string]any{}

func Compare(a string, b string) int {
	return strings.Compare(a, b)
}

func Contains(s string, substr string) bool {
	return strings.Contains(s, substr)
}

func ContainsAny(s string, chars string) bool {
	return strings.ContainsAny(s, chars)
}

func Count(s string, substr string) int {
	return strings.Count(s, substr)
}

func Cut(s string, sep string) (string, string, bool) {
	return strings.Cut(s, sep)
}

func EqualFold(s string, t string) bool {
	return strings.EqualFold(s, t)
}

func Fields(s string) any {
	return strings.Fields(s)
}

func HasPrefix(s string, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func HasSuffix(s string, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

func Index(s string, substr string) int {
	return strings.Index(s, substr)
}

func IndexAny(s string, chars string) int {
	return strings.IndexAny(s, chars)
}

func Join(elems []any, sep string) any {
	strL := make([]string, len(elems))
	for i, elem := range elems {
		strElem, ok := elem.(string)
		if !ok {
			return errors.New("non-string element in list")
		}
		strL[i] = strElem
	}
	return strings.Join(strL, sep)
}

func LastIndex(s string, substr string) int {
	return strings.LastIndex(s, substr)
}

func LastIndexAny(s string, chars string) int {
	return strings.LastIndexAny(s, chars)
}

func Repeat(s string, c int) string {
	return strings.Repeat(s, c)
}

func Replace(s string, old string, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

func ReplaceAll(s string, old string, new string) string {
	return strings.ReplaceAll(s, old, new)
}

func Split(s string, sep string) any {
	return strings.Split(s, sep)
}

func SplitAfter(s string, sep string) any {
	return strings.SplitAfter(s, sep)
}

func SplitAfterN(s string, sep string, n int) any {
	return strings.SplitAfterN(s, sep, n)
}

func SplitN(s string, sep string, n int) any {
	return strings.SplitN(s, sep, n)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func ToTitle(s string) string {
	return strings.ToTitle(s)
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToValidUTF_8(s string, replacementString string) string {
	return strings.ToValidUTF8(s, replacementString)
}

func Trim(s string, cutset string) string {
	return strings.Trim(s, cutset)
}

func TrimLeft(s string, cutset string) string {
	return strings.TrimLeft(s, cutset)
}

func TrimPrefix(s string, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

func TrimRight(s string, cutset string) string {
	return strings.TrimRight(s, cutset)
}

func TrimSuffix(s string, prefix string) string {
	return strings.TrimSuffix(s, prefix)
}
