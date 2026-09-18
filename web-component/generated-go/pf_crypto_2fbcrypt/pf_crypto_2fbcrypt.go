package pf_crypto_2fbcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){}

var PIPEFISH_VALUE_CONVERTER = map[string]any{}

func GoHash(s string, cost int) any {
	r, e := bcrypt.GenerateFromPassword([]byte(s), cost)
	if e == nil {
		return string(r)
	}
	return e
}

func GoCompare(hashedPassword string, password string) bool {
	e := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return e == nil
}
