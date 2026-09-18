package pf_crypto_2faes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

type AesEncode struct {
	Key       string
	Plaintext string
}

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){
	"AesEncode": func(t uint32, v any) any { return AesEncode{v.([]any)[0].(string), v.([]any)[1].(string)} },
}

var PIPEFISH_VALUE_CONVERTER = map[string]any{
	"AesEncode": (*AesEncode)(nil),
}

func Decode(key string, ciphertext string) any {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, encrypted := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, encrypted, nil)
	if err == nil {
		return string(plaintext)
	}
	return err
}

func GoEncode(a AesEncode) any {
	block, err := aes.NewCipher([]byte(a.Key))
	if err != nil {
		return err
	}
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	ciphertext := gcm.Seal(nonce, nonce, []byte(a.Plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext)
}
