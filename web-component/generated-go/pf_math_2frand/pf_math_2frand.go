package pf_math_2frand

import (
	"math/rand"
)

type Zipf struct {
	S    float64
	V    float64
	Imax int
}

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){
	"Zipf": func(t uint32, v any) any {
		return Zipf{v.([]any)[0].(float64), v.([]any)[1].(float64), v.([]any)[2].(int)}
	},
}

var PIPEFISH_VALUE_CONVERTER = map[string]any{
	"Zipf": (*Zipf)(nil),
}

func GoIntn(i int) int {
	return RNG.Intn(i)
}

func GoInt() int {
	return RNG.Int()
}

func GoFloat() float64 {
	return RNG.Float64()
}

func GoExponential() float64 {
	return RNG.ExpFloat64()
}

func GoNormal() float64 {
	return RNG.NormFloat64()
}

func GoZipf(z Zipf) int {
	return int(rand.NewZipf(RNG, z.S, z.V, uint64(z.Imax)).Uint64())
}

func GoShuffle(L []any) []any {
	result := make([]any, len(L))
	perm := RNG.Perm(len(L))
	for i, el := range L {
		result[perm[i]] = el
	}
	return result
}

func GoSeed(i int) any {
	RNG = rand.New(rand.NewSource(int64(i)))
	return struct{}{}
}

var RNG = rand.New(rand.NewSource(rand.Int63()))
