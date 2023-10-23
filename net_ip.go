package xjson

import (
	"fmt"
	"net"
)

// IP implements a JSON-serializable net.IP
type IP net.IP

func (xi IP) String() string {
	h := net.IP(xi)
	return h.String()
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (xi *IP) UnmarshalJSON(b []byte) error {
	s, err := stripQuotes(string(b))
	if err != nil {
		return err
	}
	i := net.ParseIP(s)
	if i == nil {
		return fmt.Errorf("invalid IP")
	}
	*xi = IP(i)
	return nil
}

// MarshalJSON implements the json.Marshaller interface.
func (xi IP) MarshalJSON() ([]byte, error) {
	return []byte("\"" + xi.String() + "\""), nil
}
