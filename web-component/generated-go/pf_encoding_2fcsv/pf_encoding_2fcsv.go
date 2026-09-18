package pf_encoding_2fcsv

import (
	"encoding/csv"
	"errors"
	"strings"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){}

var PIPEFISH_VALUE_CONVERTER = map[string]any{}

func Encode(L []any) any {
	stringLists := [][]string{}
	for _, shouldBeList := range L {
		if shouldBeList, ok := shouldBeList.([]any); !ok {
			return errors.New("non-list in list of things to be encoded")
		} else {
			stringList := []string{}
			for _, shouldBeString := range shouldBeList {
				if shouldBeString, ok := shouldBeString.(string); !ok {
					return errors.New("non-string element in row of things to be encoded")
				} else {
					stringList = append(stringList, shouldBeString)
				}
			}
			stringLists = append(stringLists, stringList)
		}
	}
	var b = strings.Builder{}
	w := csv.NewWriter(&b)
	w.WriteAll(stringLists)
	return b.String()
}

func Decode(s string) any {
	rd := csv.NewReader(strings.NewReader(s))
	r, e := rd.ReadAll()
	if e != nil {
		return e
	}
	return r
}
