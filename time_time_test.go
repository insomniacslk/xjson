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
	testLocation *time.Location
	// using CEST instead of CET because on that date it was summer-time
	testTstring = "04 Jun 24 21:05 CEST"
	testTS      time.Time
)

func init() {
	location, err := time.LoadLocation("CET")
	if err != nil {
		panic(fmt.Sprintf("Failed to load location \"CET\": %v", err))
	}
	testLocation = location
	testTS = time.Date(2024, 06, 04, 21, 05, 00, 0, testLocation)
}

type tmpTimeRFC822 struct {
	TimeRFC822 TimeRFC822
}

func TestTimeRFC822Unmarshal(t *testing.T) {
	data := []byte("{\"TimeRFC822\": \"" + testTS.Format(time.RFC822) + "\"}")
	var j tmpTimeRFC822
	err := json.Unmarshal(data, &j)
	require.NoError(t, err)
	want := testTS.String()
	assert.Equal(t, want, time.Time(j.TimeRFC822).String())
}

func TestTimeRFC822UnmarshalEmptyString(t *testing.T) {
	data := []byte(`{
    "TimeRFC822": ""
}`)
	var j tmpTimeRFC822
	err := json.Unmarshal(data, &j)
	require.Error(t, err)
}

func TestTimeRFC822UnmarshalBadType(t *testing.T) {
	data := []byte(`{
    "TimeRFC822": 1
}`)
	var j tmpTimeRFC822
	err := json.Unmarshal(data, &j)
	require.Error(t, err)
}

func TestTimeRFC822Marshal(t *testing.T) {
	j := tmpTimeRFC822{
		TimeRFC822: TimeRFC822(testTS),
	}
	want := []byte("{\"TimeRFC822\":\"" + testTstring + "\"}")
	b, err := json.Marshal(&j)
	require.NoError(t, err)
	assert.Equal(t, want, b)
}
