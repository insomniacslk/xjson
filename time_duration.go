package xjson

import (
	"time"
)

// Duration implements a JSON-serializable time.Duration.
type Duration time.Duration

func (xd Duration) String() string {
	d := time.Duration(xd)
	return d.String()
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (xd *Duration) UnmarshalJSON(b []byte) error {
	s, err := stripQuotes(string(b))
	if err != nil {
		return err
	}
	d, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	*xd = Duration(d)
	return nil
}

// MarshalJSON implements the json.Marshaller interface.
func (xd Duration) MarshalJSON() ([]byte, error) {
	return []byte("\"" + xd.String() + "\""), nil
}
