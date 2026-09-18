package pf_regexp

import (
	"errors"
	"regexp"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){}

var PIPEFISH_VALUE_CONVERTER = map[string]any{}

func Match(pattern string, text string) any {
	regObj, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	return regObj.MatchString(text)
}

func Find(pattern string, text string) any {
	regObj, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	return regObj.FindString(text)
}

func FindAllString(pattern string, text string, start int) any {
	regObj, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	return regObj.FindAllString(text, start)
}

func FindAllIndex(pattern string, text string, start int) any {
	regObj, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	indices := regObj.FindAllStringIndex(text, start)
	return indices
}

func ReplaceAll(pattern [2]any, text string) any {
	var (
		lhs string
		rhs string
		ok  bool
	)
	if lhs, ok = pattern[0].(string); !ok {
		return errors.New("lhs of pattern must be string")
	}
	if rhs, ok = pattern[1].(string); !ok {
		return errors.New("rhs of pattern must be string")
	}
	regObj, err := regexp.Compile(lhs)
	if err != nil {
		return err
	}
	return regObj.ReplaceAllString(text, rhs)
}
