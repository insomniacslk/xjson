package xjson

import (
	"time"
)

// TimeRFC3339Nano implements a JSON-serializable time.Time with RFC3339Nano layout.
type TimeRFC3339Nano time.Time

func (xt TimeRFC3339Nano) String() string {
	return timeStringer(time.Time(xt))
}

// MarshalJSON implements the json.Marshaller interface.
func (xt TimeRFC3339Nano) MarshalJSON() ([]byte, error) {
	return timeMarshaller(time.Time(xt))
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (xt *TimeRFC3339Nano) UnmarshalJSON(b []byte) error {
	t, err := timeUnmarshaller(b, time.RFC3339Nano)
	if err != nil {
		return err
	}
	*xt = TimeRFC3339Nano(t)
	return nil
}

func timeUnmarshaller3339Nano(b []byte, layout string) (time.Time, error) {
	s, err := stripQuotes(string(b))
	if err != nil {
		return time.Time{}, err
	}
	t, err := time.Parse(layout, s)
	return t, err
}

// MarshalJSON implements the json.Marshaller interface.
func timeMarshaller3339Nano(t time.Time) ([]byte, error) {
	return []byte("\"" + t.Format(time.RFC3339Nano) + "\""), nil
}
