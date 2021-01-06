package xjson

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type tmpError struct {
	Error *Error
}

func TestErrorUnmarshal(t *testing.T) {
	data := []byte(`{
    "Error": "example error message"
}`)
	var j tmpError
	err := json.Unmarshal(data, &j)
	require.NoError(t, err)
	want := tmpError{Error: &Error{Err: errors.New("example error message")}}

	assert.Equal(t, want, j)
}

func TestErrorUnmarshalEmptyString(t *testing.T) {
	data := []byte(`{
    "Error": ""
}`)
	var j tmpError
	err := json.Unmarshal(data, &j)
	require.Error(t, err)
}

func TestErrorMarshal(t *testing.T) {
	j := tmpError{
		Error: &Error{
			Err: errors.New("example error message"),
		},
	}

	want := []byte(`{"Error":"example error message"}`)
	b, err := json.Marshal(&j)
	require.NoError(t, err)
	assert.Equal(t, want, b)
}
