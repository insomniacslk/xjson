package xjson

import (
	"net"
)

// HardwareAddr implements a JSON-serializable net.HardwareAddr
type HardwareAddr net.HardwareAddr

func (xh HardwareAddr) String() string {
	h := net.HardwareAddr(xh)
	return h.String()
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (xh *HardwareAddr) UnmarshalJSON(b []byte) error {
	s, err := stripQuotes(string(b))
	if err != nil {
		return err
	}
	h, err := net.ParseMAC(s)
	if err != nil {
		return err
	}
	*xh = HardwareAddr(h)
	return nil
}

// MarshalJSON implements the json.Marshaller interface.
func (xh HardwareAddr) MarshalJSON() ([]byte, error) {
	return []byte("\"" + xh.String() + "\""), nil
}
