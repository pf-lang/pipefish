package pf_encoding_2fbase_5f32

import (
	"encoding/base32"
)

type Encoding string

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){
	"Encoding": func(t uint32, v any) any { return Encoding(v.(string)) },
}

var PIPEFISH_VALUE_CONVERTER = map[string]any{
	"Encoding": (*Encoding)(nil),
}

func Encode(encoding Encoding, plaintext string) string {
	enc := base32.NewEncoding(string(encoding))
	return enc.EncodeToString([]byte(plaintext))
}

func Decode(encoding Encoding, codedText string) any {
	enc := base32.NewEncoding(string(encoding))
	bytes, err := enc.DecodeString(codedText)
	if err != nil {
		return err
	}
	return string(bytes)
}
