package pf_net_2fsmtp

import (
	"errors"
	"net/smtp"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){}

var PIPEFISH_VALUE_CONVERTER = map[string]any{}

func GoCrammd(addr string, username string, password string, sender string, recipients []any, msg string) any {
	auth := smtp.CRAMMD5Auth(username, password)
	to, ok := stringList(recipients)
	if !ok {
		return errors.New("non-string value in list of recipients")
	}
	err := smtp.SendMail(addr, auth, sender, to, []byte(msg))
	if err != nil {
		return err
	}
	return struct{}{}
}

func GoPlain(addr string, identity string, username string, password string, host string, sender string, recipients []any, msg string) any {
	auth := smtp.PlainAuth(identity, username, password, host)
	to, ok := stringList(recipients)
	if !ok {
		return errors.New("non-string value in list of recipients")
	}
	err := smtp.SendMail(addr, auth, sender, to, []byte(msg))
	if err != nil {
		return err
	}
	return struct{}{}
}

func stringList(L []any) ([]string, bool) {
	r := []string{}
	for _, el := range L {
		s, ok := el.(string)
		if !ok {
			return nil, false
		}
		r = append(r, s)
	}
	return r, true
}
