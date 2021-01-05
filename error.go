package xjson

import (
	"encoding/json"
	"errors"
)

// Error implements a JSON-serializable error
// "Err" for less confusion over func names
type Error struct {
	Err error
}

func NewError(e error) Error { return Error{Err: e} }

func (xe *Error) Error() string { return xe.Err.Error() }

func (xe *Error) Unwrap() error {
	return xe.Err
}

// MarshalJSON implements the json.Marshaler interface.
func (xe Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(xe.Err.Error())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (xe *Error) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	if len(s) == 0 {
		return errors.New("error message must not be empty")
	}
	xe.Err = errors.New(s)
	return nil
}
