package pf_crypto_2frsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"math/big"
)

type PrivateKey struct {
	PublicKey PublicKey
	Exponent  *big.Int
	PrimeA    *big.Int
	PrimeB    *big.Int
}

type EncryptPkcs struct {
	PublicKey PublicKey
	Plaintext string
}

type EncryptOaep struct {
	PublicKey PublicKey
	Plaintext string
	OaepLabel string
}

type PublicKey struct {
	Modulus  *big.Int
	Exponent int
}

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){
	"PrivateKey": func(t uint32, v any) any {
		return PrivateKey{v.([]any)[0].(PublicKey), v.([]any)[1].(*big.Int), v.([]any)[2].(*big.Int), v.([]any)[3].(*big.Int)}
	},
	"EncryptPkcs": func(t uint32, v any) any { return EncryptPkcs{v.([]any)[0].(PublicKey), v.([]any)[1].(string)} },
	"EncryptOaep": func(t uint32, v any) any {
		return EncryptOaep{v.([]any)[0].(PublicKey), v.([]any)[1].(string), v.([]any)[2].(string)}
	},
	"PublicKey": func(t uint32, v any) any { return PublicKey{v.([]any)[0].(*big.Int), v.([]any)[1].(int)} },
}

var PIPEFISH_VALUE_CONVERTER = map[string]any{
	"EncryptPkcs": (*EncryptPkcs)(nil),
	"EncryptOaep": (*EncryptOaep)(nil),
	"PublicKey":   (*PublicKey)(nil),
	"PrivateKey":  (*PrivateKey)(nil),
}

func DecryptPkcs(k PrivateKey, ciphertext string) any {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return err
	}
	goPk := convertPrivateKey(k)
	plaintext, err := rsa.DecryptPKCS1v15(
		nil,
		goPk,
		data,
	)
	if err != nil {
		return err
	}
	return string(plaintext)
}

func GoPkcs(o EncryptPkcs) any {
	ciphertext, err := rsa.EncryptPKCS1v15(
		rand.Reader,
		convertPublicKey(o.PublicKey),
		[]byte(o.Plaintext),
	)
	if err == nil {
		return base64.StdEncoding.EncodeToString(ciphertext)
	}
	return err
}

func GoGetPrivateKey(bits int) any {
	pk, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	return PrivateKey{PublicKey{pk.N, pk.E}, pk.D, pk.Primes[0], pk.Primes[1]}
}

func GoDecryptOaep(k PrivateKey, ciphertext string, oaepLabel string) any {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return err
	}
	goPk := convertPrivateKey(k)
	plaintext, err := rsa.DecryptOAEP(
		sha256.New(),
		nil,
		goPk,
		data,
		[]byte(oaepLabel),
	)
	if err != nil {
		return err
	}
	return string(plaintext)
}

func GoOAEP(o EncryptOaep) any {
	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		convertPublicKey(o.PublicKey),
		[]byte(o.Plaintext),
		[]byte(o.OaepLabel),
	)
	if err == nil {
		return base64.StdEncoding.EncodeToString(ciphertext)
	}
	return err
}

func convertPublicKey(k PublicKey) *rsa.PublicKey {
	return &rsa.PublicKey{k.Modulus, k.Exponent}
}

func convertPrivateKey(k PrivateKey) *rsa.PrivateKey {
	return &rsa.PrivateKey{rsa.PublicKey{k.PublicKey.Modulus, k.PublicKey.Exponent},
		k.Exponent, []*big.Int{k.PrimeA, k.PrimeB}, rsa.PrecomputedValues{}}
}
