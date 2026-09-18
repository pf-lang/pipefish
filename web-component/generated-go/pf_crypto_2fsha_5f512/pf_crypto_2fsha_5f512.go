package pf_crypto_2fsha_5f512

import (
	"crypto/sha512"
	"fmt"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){}

var PIPEFISH_VALUE_CONVERTER = map[string]any{}

func Hash(s string) string {
	r := sha512.Sum512([]byte(s))
	return fmt.Sprintf("%x", r)
}
