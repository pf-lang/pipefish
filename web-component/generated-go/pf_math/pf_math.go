package pf_math

import (
	"math"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){}

var PIPEFISH_VALUE_CONVERTER = map[string]any{}

func Abs(x float64) float64 {
	return math.Abs(x)
}

func Acos(x float64) float64 {
	return math.Acos(x)
}

func Acosh(x float64) float64 {
	return math.Acosh(x)
}

func Asin(x float64) float64 {
	return math.Asin(x)
}

func Asinh(x float64) float64 {
	return math.Asinh(x)
}

func Atan(x float64) float64 {
	return math.Atan(x)
}

func Atan_2(y float64, x float64) float64 {
	return math.Atan2(y, x)
}

func Atanh(x float64) float64 {
	return math.Atanh(x)
}

func Cbrt(x float64) float64 {
	return math.Cbrt(x)
}

func Ceil(x float64) float64 {
	return math.Ceil(x)
}

func Copysign(f float64, sign float64) float64 {
	return math.Copysign(f, sign)
}

func Cos(x float64) float64 {
	return math.Cos(x)
}

func Cosh(x float64) float64 {
	return math.Cosh(x)
}

func Dim(x float64, y float64) float64 {
	return math.Dim(x, y)
}

func Erf(x float64) float64 {
	return math.Erf(x)
}

func Erfc(x float64) float64 {
	return math.Erfc(x)
}

func Erfcinv(x float64) float64 {
	return math.Erfcinv(x)
}

func Erfinv(x float64) float64 {
	return math.Erfinv(x)
}

func Exp(x float64) float64 {
	return math.Exp(x)
}

func Exp_2(x float64) float64 {
	return math.Exp2(x)
}

func Expm_1(x float64) float64 {
	return math.Expm1(x)
}

func FMA(x float64, y float64, z float64) float64 {
	return math.FMA(x, y, z)
}

func Floor(x float64) float64 {
	return math.Floor(x)
}

func Frexp(f float64) (float64, int) {
	return math.Frexp(f)
}

func Gamma(x float64) float64 {
	return math.Gamma(x)
}

func Hypot(p float64, q float64) float64 {
	return math.Hypot(p, q)
}

func Inf(sign int) float64 {
	return math.Inf(sign)
}

func IsInf(f float64, sign int) bool {
	return math.IsInf(f, sign)
}

func IsNaN(f float64) bool {
	return math.IsNaN(f)
}

func J_0(x float64) float64 {
	return math.J0(x)
}

func J_1(x float64) float64 {
	return math.J1(x)
}

func J_n(n int, x float64) float64 {
	return math.Jn(n, x)
}

func Ldexp(frac float64, ex int) float64 {
	return math.Ldexp(frac, ex)
}

func Lgamma(x float64) (float64, int) {
	return math.Lgamma(x)
}

func Log(x float64) float64 {
	return math.Log(x)
}

func Log_10(x float64) float64 {
	return math.Log10(x)
}

func Log_1_p(x float64) float64 {
	return math.Log1p(x)
}

func Log_2(x float64) float64 {
	return math.Log2(x)
}

func Log_b(x float64) float64 {
	return math.Logb(x)
}

func Max(x float64, y float64) float64 {
	return math.Max(x, y)
}

func Min(x float64, y float64) float64 {
	return math.Min(x, y)
}

func Mod(x float64, y float64) float64 {
	return math.Mod(x, y)
}

func Modf(x float64) (float64, float64) {
	return math.Modf(x)
}

func Nextafter(x float64, y float64) float64 {
	return math.Nextafter(x, y)
}

func Pow(x float64, y float64) float64 {
	return math.Pow(x, y)
}

func Pow_10(n int) float64 {
	return math.Pow10(n)
}

func Remainder(x float64, y float64) float64 {
	return math.Remainder(x, y)
}

func Round(x float64) float64 {
	return math.Round(x)
}

func RoundToEven(x float64) float64 {
	return math.RoundToEven(x)
}

func Signbit(x float64) bool {
	return math.Signbit(x)
}

func Sin(x float64) float64 {
	return math.Sin(x)
}

func Sincos(x float64) (float64, float64) {
	return math.Sincos(x)
}

func Sinh(x float64) float64 {
	return math.Sinh(x)
}

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Tan(x float64) float64 {
	return math.Tan(x)
}

func Tanh(x float64) float64 {
	return math.Tanh(x)
}

func Trunc(x float64) float64 {
	return math.Trunc(x)
}

func Y_0(x float64) float64 {
	return math.Y0(x)
}

func Y_1(x float64) float64 {
	return math.Y1(x)
}

func Y_n(n int, x float64) float64 {
	return math.Yn(n, x)
}
