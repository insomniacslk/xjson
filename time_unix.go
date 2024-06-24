package xjson

import (
	"fmt"
	"strconv"
	"time"
)

// TimeUnix implements a JSON-serializable time.Time with Unix layout.
type TimeUnix time.Time

func (xt TimeUnix) String() string {
	return time.Time(time.Time(xt)).String()
}

// MarshalJSON implements the json.Marshaller interface.
func (xt TimeUnix) MarshalJSON() ([]byte, error) {
	return timeUnixMarshaller(time.Time(xt))
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (xt *TimeUnix) UnmarshalJSON(b []byte) error {
	t, err := timeUnixUnmarshaller(b)
	if err != nil {
		return err
	}
	*xt = TimeUnix(t)
	return nil
}

func timeUnixUnmarshaller(b []byte) (time.Time, error) {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(ts, 0), err
}

// MarshalJSON implements the json.Marshaller interface.
func timeUnixMarshaller(t time.Time) ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.Unix())), nil
}
