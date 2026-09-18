package pf_source_2finitializer_2flibraries_2fmath_2fbig_2epf

import (
	"errors"
	"math/big"
	"strconv"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){
	"Float": func(t uint32, v any) any { return v },
	"Int":   func(t uint32, v any) any { return v },
	"Rat":   func(t uint32, v any) any { return v },
}

var PIPEFISH_VALUE_CONVERTER = map[string]any{
	"Int":   (**big.Int)(nil),
	"Rat":   (**big.Rat)(nil),
	"Float": (**big.Float)(nil),
}

func Float(x float64) any {
	return big.NewFloat(x)
}

func IsInf(z any) bool {
	return z.(*big.Float).IsInf()
}

func IsInt(z any) bool {
	return z.(*big.Float).IsInt()
}

func GoTextF(z any, format int, prec int) string {
	return z.(*big.Float).Text(byte(format), prec)
}

func GoAbsF(x any) any {
	z := big.NewFloat(0.0)
	z.Abs(x.(*big.Float))
	return z
}

func GoAddF(x any, y any) any {
	z := big.NewFloat(0.0)
	z.Add(x.(*big.Float), y.(*big.Float))
	return z
}

func GoSubF(x any, y any) any {
	z := big.NewFloat(0.0)
	z.Sub(x.(*big.Float), y.(*big.Float))
	return z
}

func GoMulF(x any, y any) any {
	z := big.NewFloat(0.0)
	z.Mul(x.(*big.Float), y.(*big.Float))
	return z
}

func GoNegF(x any) any {
	z := big.NewFloat(0.0)
	z.Neg(x.(*big.Float))
	return z
}

func GoQuoF(x any, y any) any {
	z := big.NewFloat(0.0)
	z.Quo(x.(*big.Float), y.(*big.Float))
	return z
}

func GoCmpF(z any, y any) int {
	return z.(*big.Float).Cmp(y.(*big.Float))
}

func GoSignF(z any) int {
	return z.(*big.Float).Sign()
}

func GoSqrtF(x any) any {
	z := big.NewFloat(0.0)
	z.Sqrt(x.(*big.Float))
	return z
}

func BigFTof(z any) float64 {
	r, _ := z.(*big.Float).Float64()
	return r
}

func BigItoF(x any) any {
	z := big.NewFloat(0.0)
	z.SetInt(x.(*big.Int))
	return z
}

func BigFtobigI(x any) any {
	z := big.NewInt(0)
	z, _ = x.(*big.Float).Int(z)
	return z
}

func Itof(z any) int {
	r, _ := z.(*big.Float).Uint64()
	return int(r)
}

func Rtof(x any) any {
	z := big.NewFloat(0.0)
	z.SetRat(x.(*big.Rat))
	return z
}

func Ftoi(z any) int {
	r, _ := z.(*big.Float).Uint64()
	return int(r)
}

func Atof(s string) any {
	var (
		ok bool
	)
	bf := big.NewFloat(0)
	bf, ok = bf.SetString(s)
	if !ok {
		return errors.New("Can't convert string \"" + s + "\" to Float type")
	}
	return bf
}

func Ftoa(z any) string {
	return z.(*big.Float).String()
}

func Int(i int) any {
	return big.NewInt(int64(i))
}

func Abs(x any) any {
	z := big.NewInt(0)
	z.Abs(x.(*big.Int))
	return z
}

func AndInt(x any, y any) any {
	z := big.NewInt(0)
	z.And(x.(*big.Int), y.(*big.Int))
	return z
}

func AndNotInt(x any, y any) any {
	z := big.NewInt(0)
	z.AndNot(x.(*big.Int), y.(*big.Int))
	return z
}

func Binomial(n int, k int) any {
	z := big.NewInt(0)
	z.Binomial(int64(n), int64(k))
	return z
}

func Bit(z any, i int) int {
	return int(z.(*big.Int).Bit(i))
}

func BitLen(z any) int {
	return z.(*big.Int).BitLen()
}

func DivMod(x any, y any) (any, any) {
	z := big.NewInt(0)
	m := big.NewInt(0)
	z.DivMod(x.(*big.Int), y.(*big.Int), m)
	return z, m
}

func FromString(s string, base int) any {
	z := big.NewInt(0)
	z, ok := z.SetString(s, base)
	if ok {
		return z
	}
	return errors.New("string is not well-formed to convert to base " + strconv.Itoa(base))
}

func Gcd(a any, b any) any {
	z := big.NewInt(0)
	z.GCD(nil, nil, a.(*big.Int), b.(*big.Int))
	return z
}

func IsSmallInt(z any) bool {
	return z.(*big.Int).IsInt64()
}

func Lsh(x any, n int) any {
	z := big.NewInt(0)
	z.Lsh(x.(*big.Int), uint(n))
	return z
}

func ModInverse(g any, n any) any {
	z := big.NewInt(0)
	z.ModInverse(g.(*big.Int), n.(*big.Int))
	return z
}

func ModSqrt(x any, p any) any {
	z := big.NewInt(0)
	z.ModSqrt(x.(*big.Int), p.(*big.Int))
	return z
}

func MulRange(a int, b int) any {
	z := big.NewInt(0)
	z.MulRange(int64(a), int64(b))
	return z
}

func NotInt(x any) any {
	z := big.NewInt(0)
	z.Not(x.(*big.Int))
	return z
}

func OrInt(x any, y any) any {
	z := big.NewInt(0)
	z.Or(x.(*big.Int), y.(*big.Int))
	return z
}

func ProbablyPrime(z any, n int) bool {
	return z.(*big.Int).ProbablyPrime(n)
}

func Quo(x any, y any) any {
	z := big.NewInt(0)
	z.Quo(x.(*big.Int), y.(*big.Int))
	return z
}

func QuoRem(x any, y any) (any, any) {
	z := big.NewInt(0)
	r := big.NewInt(0)
	z.QuoRem(x.(*big.Int), y.(*big.Int), r)
	return z, r
}

func Rem(x any, y any) any {
	z := big.NewInt(0)
	z.Rem(x.(*big.Int), y.(*big.Int))
	return z
}

func Rsh(x any, n int) any {
	z := big.NewInt(0)
	z.Rsh(x.(*big.Int), uint(n))
	return z
}

func Sign(z any) int {
	return z.(*big.Int).Sign()
}

func Sqrt(x any) any {
	z := big.NewInt(0)
	z.Sqrt(x.(*big.Int))
	return z
}

func Text(z any, base int) string {
	return z.(*big.Int).Text(base)
}

func TrailingZeroBits(z any) int {
	return int(z.(*big.Int).TrailingZeroBits())
}

func XorInt(x any, y any) any {
	z := big.NewInt(0)
	z.Xor(x.(*big.Int), y.(*big.Int))
	return z
}

func GoAdd(x any, y any) any {
	z := big.NewInt(0)
	return z.Add(x.(*big.Int), y.(*big.Int))
}

func GoSub(x any, y any) any {
	z := big.NewInt(0)
	return z.Sub(x.(*big.Int), y.(*big.Int))
}

func GoMul(x any, y any) any {
	z := big.NewInt(0)
	return z.Mul(x.(*big.Int), y.(*big.Int))
}

func GoDiv(x any, y any) any {
	z := big.NewInt(0)
	return z.Div(x.(*big.Int), y.(*big.Int))
}

func GoMod(x any, y any) any {
	z := big.NewInt(0)
	return z.Mod(x.(*big.Int), y.(*big.Int))
}

func GoNeg(x any) any {
	z := big.NewInt(0)
	return z.Neg(x.(*big.Int))
}

func GoInt(z any) int {
	return int(z.(*big.Int).Int64())
}

func GoCmp(x any, y any) int {
	return x.(*big.Int).Cmp(y.(*big.Int))
}

func GoExp(x any, y any, m any) any {
	z := big.NewInt(0)
	z.Exp(x.(*big.Int), y.(*big.Int), m.(*big.Int))
	return z
}

func Atoi(s string) any {
	var (
		ok bool
	)
	bi := big.NewInt(0)
	bi, ok = bi.SetString(s, 10)
	if !ok {
		return errors.New("Can't convert string \"" + s + "\" to Int type")
	}
	return bi
}

func BigItolittlef(z any) float64 {
	r, _ := z.(*big.Int).Float64()
	return r
}

func Itoa(z any) string {
	return (z.(*big.Int)).String()
}

func Denom(z any) any {
	return z.(*big.Rat).Denom()
}

func FloatString(z any, prec int) string {
	return z.(*big.Rat).FloatString(prec)
}

func Inv(x any) any {
	z := big.NewRat(0, 0)
	z.Inv(x.(*big.Rat))
	return z
}

func Num(z any) any {
	return z.(*big.Rat).Num()
}

func RatString(z any) string {
	return z.(*big.Rat).RatString()
}

func GoAbsR(x any) any {
	z := big.NewRat(0, 0)
	z.Abs(x.(*big.Rat))
	return z
}

func GoAddR(x any, y any) any {
	z := big.NewRat(0, 0)
	z.Add(x.(*big.Rat), y.(*big.Rat))
	return z
}

func GoCmpR(z any, y any) int {
	return z.(*big.Rat).Cmp(y.(*big.Rat))
}

func GoNegR(x any) any {
	z := big.NewRat(0, 0)
	z.Neg(x.(*big.Rat))
	return z
}

func GoQuoR(x any, y any) any {
	z := big.NewRat(0, 0)
	z.Quo(x.(*big.Rat), y.(*big.Rat))
	return z
}

func GoMulR(x any, y any) any {
	z := big.NewRat(0, 0)
	z.Mul(x.(*big.Rat), y.(*big.Rat))
	return z
}

func GoSubR(x any, y any) any {
	z := big.NewRat(0, 0)
	z.Sub(x.(*big.Rat), y.(*big.Rat))
	return z
}

func GoSignR(z any) int {
	return z.(*big.Rat).Sign()
}

func GoisIntR(z any) bool {
	return z.(*big.Rat).IsInt()
}

func BigRtolittlef(z any) float64 {
	x, _ := z.(*big.Rat).Float64()
	return x
}

func LittleftoR(x float64) any {
	z := big.NewRat(0, 0)
	z.SetFloat64(x)
	return z
}

func BigIItoR(a any, b any) any {
	z := big.NewRat(0, 0)
	z.SetFrac(a.(*big.Int), b.(*big.Int))
	return z
}

func BigItoR(x any) any {
	z := big.NewRat(0, 0)
	z.SetInt(x.(*big.Int))
	return z
}

func LittleitoR(x int) any {
	z := big.NewRat(0, 0)
	z.SetInt64(int64(x))
	return z
}

func Ator(s string) any {
	var (
		ok bool
	)
	br := big.NewRat(0, 0)
	br, ok = br.SetString(s)
	if !ok {
		return errors.New("Can't convert string \"" + s + "\" to Rat type")
	}
	return br
}

func Ftor(x any) any {
	r := big.NewRat(0, 1)
	r, _ = x.(*big.Float).Rat(r)
	return r
}

func Ijtor(i int, j int) any {
	return big.NewRat(int64(i), int64(j))
}

func Rtoa(z any) string {
	return z.(*big.Rat).String()
}

func Equals(x, y any) bool {
	switch x := x.(type) {
	case *big.Float:
		return x.Cmp(y.(*big.Float)) == 0
	case *big.Int:
		return x.Cmp(y.(*big.Int)) == 0
	case *big.Rat:
		return x.Cmp(y.(*big.Rat)) == 0
	default:
		return false
	}
}

func Literal(x any) string {
	switch x := x.(type) {
	case *big.Float:
		return "Float(`" + x.String() + "`)"
	case *big.Int:
		return "Int(`" + x.String() + "`)"
	case *big.Rat:
		return "Rat(`" + x.String() + "`)"
	default:
		return "can't serialize value"
	}
}
