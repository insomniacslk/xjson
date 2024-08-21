package xjson

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testLocation3339Nano *time.Location
	// using CEST instead of CET because on that date it was summer-time
	testTstring3339Nano = "2024-08-21T11:00:27.107179+00:00"
	testTS3339Nano      time.Time
)

func init() {
	location, err := time.LoadLocation("CET")
	if err != nil {
		panic(fmt.Sprintf("Failed to load location \"CET\": %v", err))
	}
	testLocation = location
	testTS = time.Date(2024, 06, 04, 21, 05, 00, 0, testLocation)
}

type tmpTimeRFC3339Nano struct {
	TimeRFC3339Nano TimeRFC3339Nano
}

func TestTimeRFC3339Unmarshal(t *testing.T) {
	data := []byte("{\"TimeRFC3339Nano\": \"" + testTS.Format(time.RFC3339Nano) + "\"}")
	var j tmpTimeRFC3339Nano
	err := json.Unmarshal(data, &j)
	require.NoError(t, err)
	want := testTS.String()
	assert.Equal(t, want, time.Time(j.TimeRFC3339Nano).String())
}

func TestTimeRFC3339NanoUnmarshalEmptyString(t *testing.T) {
	data := []byte(`{
    "TimeRFC3339Nano": ""
}`)
	var j tmpTimeRFC3339Nano
	err := json.Unmarshal(data, &j)
	require.Error(t, err)
}

func TestTimeRFC3339NanoUnmarshalBadType(t *testing.T) {
	data := []byte(`{
    "TimeRFC3339Nano": 1
}`)
	var j tmpTimeRFC3339Nano
	err := json.Unmarshal(data, &j)
	require.Error(t, err)
}

func TestTimeRFC3339NanoMarshal(t *testing.T) {
	j := tmpTimeRFC3339Nano{
		TimeRFC3339Nano: TimeRFC3339Nano(testTS),
	}
	want := []byte("{\"TimeRFC3339Nano\":\"" + testTstring + "\"}")
	b, err := json.Marshal(&j)
	require.NoError(t, err)
	assert.Equal(t, want, b)
}
