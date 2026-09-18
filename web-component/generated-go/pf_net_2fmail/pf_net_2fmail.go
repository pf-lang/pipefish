package pf_net_2fmail

import (
	"errors"
	"io"
	"net/mail"
	"strings"
	"time"
)

type Time struct {
	Year       int
	Month      int
	Day        int
	Hour       int
	Minute     int
	Second     int
	Nanosecond int
	Location   string
}

type Address struct {
	Name    string
	Address string
}

type Message struct {
	Header map[any]any
	Body   string
}

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){
	"Address": func(t uint32, v any) any { return Address{v.([]any)[0].(string), v.([]any)[1].(string)} },
	"Message": func(t uint32, v any) any { return Message{v.([]any)[0].(map[any]any), v.([]any)[1].(string)} },
	"Time": func(t uint32, v any) any {
		return Time{v.([]any)[0].(int), v.([]any)[1].(int), v.([]any)[2].(int), v.([]any)[3].(int), v.([]any)[4].(int), v.([]any)[5].(int), v.([]any)[6].(int), v.([]any)[7].(string)}
	},
}

var PIPEFISH_VALUE_CONVERTER = map[string]any{
	"Address": (*Address)(nil),
	"Message": (*Message)(nil),
	"Time":    (*Time)(nil),
}

func AddressList(header map[any]any, key string) any {
	g, err1 := pfToGoHeader(header)
	if err1 != nil {
		return err1
	}
	goAddresses, err2 := g.AddressList(key)
	if err2 != nil {
		return err2
	}
	pfAddresses := []Address{}
	for _, ad := range goAddresses {
		pfAddresses = append(pfAddresses, goToPfAddress(ad))
	}
	return pfAddresses
}

func Date(h map[any]any) any {
	g, err1 := pfToGoHeader(h)
	if err1 != nil {
		return err1
	}
	t, err2 := g.Date()
	if err2 != nil {
		return err2
	}
	return goToPfTime(t)
}

func ParseAddress(address string) any {
	ad, err := mail.ParseAddress(address)
	if err != nil {
		return err
	}
	return Address{ad.Name, ad.Address}
}

func ParseDate(s string) any {
	t, err := mail.ParseDate(s)
	if err != nil {
		return err
	}
	return goToPfTime(t)
}

func ReadMessage(raw string) any {
	m, err := mail.ReadMessage(strings.NewReader(raw))
	if err != nil {
		return err
	}
	body, e := io.ReadAll(m.Body)
	if e != nil {
		return e
	}
	return Message{goToPfHeader(m.Header), string(body)}
}

func String(ad Address) string {
	return (&mail.Address{ad.Name, ad.Address}).String()
}

func IotaM(m Message) Message {
	return m
}

func IotaT(m Time) Time {
	return m
}

func pfToGoHeader(header map[any]any) (mail.Header, error) {
	goMap := mail.Header{}
	for key, list := range header {
		if goKey, ok := key.(string); !ok {
			return nil, errors.New("key in map is not of type `string`")
		} else {
			if rawList, ok := list.([]any); ok {
				return nil, errors.New("value in map is not of type `list`")
			} else {
				goList := []string{}
				for _, el := range rawList {
					if goEl, ok := el.(string); !ok {
						return nil, errors.New("element in list in map is not of type `string`")
					} else {
						goList = append(goList, goEl)
					}
				}
				goMap[goKey] = goList
			}
		}
	}
	return goMap, nil
}

func goToPfHeader(m mail.Header) map[any]any {
	pfMap := map[any]any{}
	for key, list := range m {
		pfList := []any{}
		for _, el := range list {
			pfList = append(pfList, el)
		}
		pfMap[key] = pfList
	}
	return pfMap
}

func goToPfAddress(a *mail.Address) Address {
	return Address{a.Name, a.Address}
}

func goToPfTime(t time.Time) Time {
	return Time{t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location().String()}
}
