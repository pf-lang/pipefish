package pf_net_2furl

import (
	"errors"
	"net/url"
)

type UserInfo struct {
	Username string
	Password string
}

type Url struct {
	Scheme      string
	Opaque      string
	User        UserInfo
	Host        string
	Path        string
	RawPath     string
	OmitHost    bool
	ForceQuery  bool
	RawQuery    string
	Fragment    string
	RawFragment string
}

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){
	"Url": func(t uint32, v any) any {
		return Url{v.([]any)[0].(string), v.([]any)[1].(string), v.([]any)[2].(UserInfo), v.([]any)[3].(string), v.([]any)[4].(string), v.([]any)[5].(string), v.([]any)[6].(bool), v.([]any)[7].(bool), v.([]any)[8].(string), v.([]any)[9].(string), v.([]any)[10].(string)}
	},
	"UserInfo": func(t uint32, v any) any { return UserInfo{v.([]any)[0].(string), v.([]any)[1].(string)} },
}

var PIPEFISH_VALUE_CONVERTER = map[string]any{
	"Url":      (*Url)(nil),
	"UserInfo": (*UserInfo)(nil),
}

func Encode(values map[any]any) any {
	goMap := url.Values{}
	for key, list := range values {
		if goKey, ok := key.(string); !ok {
			return errors.New("key in map is not of type `string`")
		} else {
			if rawList, ok := list.([]any); ok {
				return errors.New("value in map is not of type `list`")
			} else {
				goList := []string{}
				for _, el := range rawList {
					if goEl, ok := el.(string); !ok {
						return errors.New("element in list in map is not of type `string`")
					} else {
						goList = append(goList, goEl)
					}
				}
				goMap[goKey] = goList
			}
		}
	}
	return goMap.Encode()
}

func EscapedFragment(u Url) string {
	return urlToUrlUrl(u).EscapedFragment()
}

func EscapedPath(u Url) string {
	return urlToUrlUrl(u).EscapedPath()
}

func Hostname(u Url) string {
	return urlToUrlUrl(u).Hostname()
}

func IsAbs(u Url) bool {
	return urlToUrlUrl(u).IsAbs()
}

func JoinPath(u Url, elem ...string) Url {
	return urlUrlToUrl(urlToUrlUrl(u).JoinPath(elem...))
}

func Parse(rawUrl string) any {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return err
	} else {
		return urlUrlToUrl(u)
	}
}

func ParseInContext(v Url, rawUrl string) any {
	u, err := urlToUrlUrl(v).Parse(rawUrl)
	if err != nil {
		return err
	} else {
		return urlUrlToUrl(u)
	}
}

func ParseQuery(query string) any {
	r, err := url.ParseQuery(query)
	if err != nil {
		return err
	} else {
		return r
	}
}

func ParseRequestUri(rawUrl string) any {
	u, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		return err
	} else {
		return urlUrlToUrl(u)
	}
}

func PathEscape(s string) string {
	return url.PathEscape(s)
}

func PathUnescape(s string) any {
	r, e := url.PathUnescape(s)
	if e != nil {
		return e
	}
	return r
}

func Port(u Url) string {
	return urlToUrlUrl(u).Port()
}

func Query(u Url) any {
	return urlToUrlUrl(u).Query()
}

func QueryEscape(s string) string {
	return url.QueryEscape(s)
}

func QueryUnescape(s string) any {
	r, e := url.QueryUnescape(s)
	if e != nil {
		return e
	}
	return r
}

func Redacted(u Url) string {
	return urlToUrlUrl(u).Redacted()
}

func RequestUri(u Url) string {
	return urlToUrlUrl(u).RequestURI()
}

func ResolveReference(u Url, v Url) Url {
	return urlUrlToUrl(urlToUrlUrl(u).ResolveReference(urlToUrlUrl(v)))
}

func String(u Url) string {
	return urlToUrlUrl(u).String()
}

func GoJoinPathToString(base string, elem ...string) any {
	r, e := url.JoinPath(base, elem...)
	if e != nil {
		return e
	}
	return r
}

func urlUrlToUrl(u *url.URL) Url {
	password, _ := u.User.Password()
	return Url{
		Scheme: u.Scheme,
		Opaque: u.Opaque,
		User: UserInfo{Username: u.User.Username(),
			Password: password,
		},
		Host:        u.Host,
		Path:        u.Path,
		RawPath:     u.RawPath,
		OmitHost:    u.OmitHost,
		ForceQuery:  u.ForceQuery,
		RawQuery:    u.RawQuery,
		Fragment:    u.Fragment,
		RawFragment: u.RawFragment,
	}
}

func urlToUrlUrl(u Url) *url.URL {
	return &url.URL{
		Scheme:      u.Scheme,
		Opaque:      u.Opaque,
		User:        url.UserPassword(u.User.Username, u.User.Password),
		Host:        u.Host,
		Path:        u.Path,
		RawPath:     u.RawPath,
		OmitHost:    u.OmitHost,
		ForceQuery:  u.ForceQuery,
		RawQuery:    u.RawQuery,
		Fragment:    u.Fragment,
		RawFragment: u.RawFragment,
	}
}
