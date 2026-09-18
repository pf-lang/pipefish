package pf_crypto_2frand

import (
	"crypto/rand"
	"math/big"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){
	"big.Int": func(t uint32, v any) any { return v },
}

var PIPEFISH_VALUE_CONVERTER = map[string]any{
	"big.Int": (**big.Int)(nil),
}

func GoRandomInt(max any) any {
	x, e := rand.Int(rand.Reader, max.(*big.Int))
	if e == nil {
		return x
	}
	return e
}

func GoRandomPrime(bits int) any {
	x, e := rand.Prime(rand.Reader, bits)
	if e == nil {
		return x
	}
	return e
}

func GoRandomText() string {
	return rand.Text()
}
