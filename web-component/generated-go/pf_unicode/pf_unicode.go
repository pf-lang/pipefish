package pf_unicode

import (
	"unicode"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){}

var PIPEFISH_VALUE_CONVERTER = map[string]any{}

func IsControl(r rune) bool {
	return unicode.IsControl(r)
}

func IsDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func IsGraphic(r rune) bool {
	return unicode.IsGraphic(r)
}

func IsLetter(r rune) bool {
	return unicode.IsLetter(r)
}

func IsLower(r rune) bool {
	return unicode.IsLower(r)
}

func IsMark(r rune) bool {
	return unicode.IsMark(r)
}

func IsNumber(r rune) bool {
	return unicode.IsNumber(r)
}

func IsPrint(r rune) bool {
	return unicode.IsPrint(r)
}

func IsPunct(r rune) bool {
	return unicode.IsPunct(r)
}

func IsSpace(r rune) bool {
	return unicode.IsSpace(r)
}

func IsSymbol(r rune) bool {
	return unicode.IsSymbol(r)
}

func IsTitle(r rune) bool {
	return unicode.IsTitle(r)
}

func IsUpper(r rune) bool {
	return unicode.IsUpper(r)
}

func SimpleFold(r rune) rune {
	return unicode.SimpleFold(r)
}

func ToLower(r rune) rune {
	return unicode.ToLower(r)
}

func ToTitle(r rune) rune {
	return unicode.ToTitle(r)
}

func ToUpper(r rune) rune {
	return unicode.ToUpper(r)
}

func WrapToCase(c int, r rune) rune {
	return unicode.To(c, r)
}
