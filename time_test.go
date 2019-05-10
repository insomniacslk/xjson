package xjson

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type tmpDuration struct {
	Duration Duration
}

func TestDurationUnmarshal(t *testing.T) {
	data := []byte(`{
    "Duration": "1m5s"
}`)
	var j tmpDuration
	err := json.Unmarshal(data, &j)
	require.NoError(t, err)
	want := Duration(time.Minute + 5*time.Second)
	assert.Equal(t, want, j.Duration)
}

func TestDurationUnmarshalEmptyString(t *testing.T) {
	data := []byte(`{
    "Duration": ""
}`)
	var j tmpDuration
	err := json.Unmarshal(data, &j)
	require.Error(t, err)
}

func TestDurationUnmarshalBadType(t *testing.T) {
	data := []byte(`{
    "Duration": 1
}`)
	var j tmpDuration
	err := json.Unmarshal(data, &j)
	require.Error(t, err)
}

func TestDurationMarshal(t *testing.T) {
	j := tmpDuration{
		Duration: Duration(time.Minute + 5*time.Second),
	}
	want := []byte("{\"Duration\":\"1m5s\"}")
	b, err := json.Marshal(&j)
	require.NoError(t, err)
	assert.Equal(t, want, b)
}
