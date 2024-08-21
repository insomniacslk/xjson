package xjson

import (
	"time"
)

// TimeRFC822 implements a JSON-serializable time.Time with RFC822 layout.
type TimeRFC822 time.Time

func (xt TimeRFC822) String() string {
	return timeStringer(time.Time(xt))
}

// MarshalJSON implements the json.Marshaller interface.
func (xt TimeRFC822) MarshalJSON() ([]byte, error) {
	return timeMarshaller(time.Time(xt))
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (xt *TimeRFC822) UnmarshalJSON(b []byte) error {
	t, err := timeUnmarshaller(b, time.RFC822)
	if err != nil {
		return err
	}
	*xt = TimeRFC822(t)
	return nil
}

func timeStringer(t time.Time) string {
	return time.Time(t).String()
}

func timeUnmarshaller(b []byte, layout string) (time.Time, error) {
	s, err := stripQuotes(string(b))
	if err != nil {
		return time.Time{}, err
	}
	t, err := time.Parse(layout, s)
	return t, err
}

// MarshalJSON implements the json.Marshaller interface.
func timeMarshaller(t time.Time) ([]byte, error) {
	return []byte("\"" + t.Format(time.RFC822) + "\""), nil
}
