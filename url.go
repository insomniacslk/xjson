package xjson

import (
	"log"
	"net/url"
)

// URL implements a JSON-serializable url.URL.
type URL url.URL

// UnmarshalJSON implements the json.Unmarshaller interface.
func (xu *URL) UnmarshalJSON(b []byte) error {
	log.Printf("Parsing %s", b)
	u, err := url.Parse(stripQuotes(string(b)))
	if err != nil {
		return err
	}
	*xu = URL(*u)
	return nil
}

func stripQuotes(s string) string {
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return s
	}
	// TODO unquote backslashes etc
	return s[1 : len(s)-1]
}
