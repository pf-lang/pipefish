package pf_math_2fcmplx

import (
	"math"
	"math/cmplx"
	"strconv"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){
	"Complex": func(t uint32, v any) any { return v },
}

var PIPEFISH_VALUE_CONVERTER = map[string]any{
	"Complex": (*complex128)(nil),
}

func Complex(a float64, b float64) any {
	return complex(a, b)
}

func Abs(x any) float64 {
	return cmplx.Abs(x.(complex128))
}

func Acos(x any) any {
	return cmplx.Acos(x.(complex128))
}

func Acosh(x any) any {
	return cmplx.Acosh(x.(complex128))
}

func Asin(x any) any {
	return cmplx.Asin(x.(complex128))
}

func Asinh(x any) any {
	return cmplx.Asinh(x.(complex128))
}

func Atan(x any) any {
	return cmplx.Atan(x.(complex128))
}

func Atanh(x any) any {
	return cmplx.Atanh(x.(complex128))
}

func Conj(x any) any {
	return cmplx.Conj(x.(complex128))
}

func Cos(x any) any {
	return cmplx.Cos(x.(complex128))
}

func Cosh(x any) any {
	return cmplx.Cosh(x.(complex128))
}

func Cot(x any) any {
	return cmplx.Cot(x.(complex128))
}

func Exp(x any) any {
	return cmplx.Exp(x.(complex128))
}

func Imag(x any) float64 {
	return imag(x.(complex128))
}

func Inf() any {
	return cmplx.Inf()
}

func IsInf(x any) bool {
	return cmplx.IsInf(x.(complex128))
}

func IsNaN(x any) bool {
	return cmplx.IsNaN(x.(complex128))
}

func Log(x any) any {
	return cmplx.Log(x.(complex128))
}

func Log_10(x any) any {
	return cmplx.Log10(x.(complex128))
}

func NaN() any {
	return cmplx.NaN()
}

func Phase(x any) float64 {
	return cmplx.Phase(x.(complex128))
}

func Polar(x any) (float64, float64) {
	return cmplx.Polar(x.(complex128))
}

func Real(x any) float64 {
	return real(x.(complex128))
}

func Rect(r float64, θ float64) any {
	return cmplx.Rect(r, θ)
}

func Sin(x any) any {
	return cmplx.Sin(x.(complex128))
}

func Sinh(x any) any {
	return cmplx.Sinh(x.(complex128))
}

func Sqrt(x any) any {
	return cmplx.Sqrt(x.(complex128))
}

func String(z any) string {
	return strconv.FormatComplex(z.(complex128), 'g', 6, 128)
}

func Tan(x any) any {
	return cmplx.Tan(x.(complex128))
}

func Tanh(x any) any {
	return cmplx.Tanh(x.(complex128))
}

func GoAdd(x any, y any) any {
	return x.(complex128) + y.(complex128)
}

func GoSub(x any, y any) any {
	return x.(complex128) - y.(complex128)
}

func GoMul(x any, y any) any {
	return x.(complex128) * y.(complex128)
}

func GoDiv(x any, y any) any {
	return x.(complex128) / y.(complex128)
}

func GoPow(x any, y any) any {
	return cmplx.Pow(x.(complex128), y.(complex128))
}

func Equals(x, y any) bool {
	switch x := x.(type) {
	case complex128:
		return x == y.(complex128)
	default:
		return false
	}
}

func Literal(x any) string {
	switch x := x.(type) {
	case complex128:
		return "Complex(" + niceFmt(real(x)) + ", " +
			niceFmt(imag(x)) + ")"
	default:
		return "can't serialize value"
	}
}

func niceFmt(f float64) string {
	if f == math.Trunc(f) {
		return strconv.FormatFloat(f, 'f', 1, 64)
	}
	return strconv.FormatFloat(f, 'g', -1, 64)
}
