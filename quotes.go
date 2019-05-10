package xjson

import (
	"errors"
)

func stripQuotes(s string) (string, error) {
	if len(s) < 2 || (s[0] != '"' && s[len(s)-1] != '"') {
		return s, errors.New("not a properly double-quoted string")
	}
	return s[1 : len(s)-1], nil
}
