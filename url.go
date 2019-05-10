package xjson

import (
	"net/url"
)

// URL implements a JSON-serializable url.URL.
type URL url.URL

func (xu URL) String() string {
	u := url.URL(xu)
	return u.String()
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (xu *URL) UnmarshalJSON(b []byte) error {
	s, err := stripQuotes(string(b))
	if err != nil {
		return err
	}
	u, err := url.Parse(s)
	if err != nil {
		return err
	}
	*xu = URL(*u)
	return nil
}

// MarshalJSON implements the json.Marshaller interface.
func (xu URL) MarshalJSON() ([]byte, error) {
	return []byte("\"" + xu.String() + "\""), nil
}
