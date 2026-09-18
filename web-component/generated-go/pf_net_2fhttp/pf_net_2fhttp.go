package pf_net_2fhttp

import (
	"io"
	"net/http"
	"strings"
)

type Response struct {
	Body string
}

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){
	"Response": func(t uint32, v any) any { return Response{v.([]any)[0].(string)} },
}

var PIPEFISH_VALUE_CONVERTER = map[string]any{
	"Response": (*Response)(nil),
}

func GoGet(url string) any {
	r, err := http.Get(url)
	if err == nil {
		body, err2 := io.ReadAll(r.Body)
		r.Body.Close()
		if err2 == nil {
			return Response{string(body)}
		}
		return err2
	} else {
		return err
	}
}

func GoPost(url string, contentType string, s string) any {
	_, err := http.Post(url, contentType, strings.NewReader(s))
	if err == nil {
		return struct{}{}
	} else {
		return err
	}
}
