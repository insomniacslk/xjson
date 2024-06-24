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
	testUnixTs   = int64(1717527900)
	testUnixTime time.Time
)

func init() {
	location, err := time.LoadLocation("CET")
	if err != nil {
		panic(fmt.Sprintf("Failed to load location \"CET\": %v", err))
	}
	testLocation = location
	testUnixTime = time.Date(2024, 06, 04, 21, 05, 00, 0, testLocation)
}

type tmpTimeUnix struct {
	TimeUnix TimeUnix
}

func TestTimeUnixUnmarshal(t *testing.T) {
	data := []byte(fmt.Sprintf(`{"TimeUnix": "%d"}`, testUnixTs))
	var j tmpTimeUnix
	err := json.Unmarshal(data, &j)
	require.NoError(t, err)
	want := testUnixTs
	assert.Equal(t, want, time.Time(j.TimeUnix).Unix())
}

func TestTimeUnixUnmarshalEmptyString(t *testing.T) {
	data := []byte(`{
    "TimeUnix": ""
}`)
	var j tmpTimeUnix
	err := json.Unmarshal(data, &j)
	require.Error(t, err)
}

func TestTimeUnixUnmarshalBadType(t *testing.T) {
	data := []byte(`{
    "TimeUnix": 2.5
}`)
	var j tmpTimeUnix
	err := json.Unmarshal(data, &j)
	require.Error(t, err)
}

func TestTimeUnixMarshal(t *testing.T) {
	j := tmpTimeUnix{
		TimeUnix: TimeUnix(testUnixTime),
	}
	want := []byte(fmt.Sprintf(`{"TimeUnix":"%d"}`, testUnixTs))
	b, err := json.Marshal(&j)
	require.NoError(t, err)
	assert.Equal(t, want, b)
}
