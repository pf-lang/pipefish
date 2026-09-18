package pf_encoding_2fbase_5f64

import (
	"encoding/base64"
)

type Encoding struct {
	Padding bool
	Key     string
}

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){
	"Encoding": func(t uint32, v any) any { return Encoding{v.([]any)[0].(bool), v.([]any)[1].(string)} },
}

var PIPEFISH_VALUE_CONVERTER = map[string]any{
	"Encoding": (*Encoding)(nil),
}

func Encode(encoding Encoding, plaintext string) string {
	pad := base64.NoPadding
	if encoding.Padding {
		pad = base64.StdPadding
	}
	enc := base64.NewEncoding(encoding.Key).WithPadding(pad)
	return enc.EncodeToString([]byte(plaintext))
}

func Decode(encoding Encoding, codedText string) any {
	pad := base64.NoPadding
	if encoding.Padding {
		pad = base64.StdPadding
	}
	enc := base64.NewEncoding(encoding.Key).WithPadding(pad)
	bytes, err := enc.DecodeString(codedText)
	if err != nil {
		return err
	}
	return string(bytes)
}
