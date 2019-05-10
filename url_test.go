package xjson

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type tmpURL struct {
	URL URL
}

func TestURLUnmarshal(t *testing.T) {
	// test valid file://
	data := []byte(`{
    "URL": "file://localhost/some/file.txt"
}`)
	var j tmpURL
	err := json.Unmarshal(data, &j)
	require.NoError(t, err)
	assert.Equal(t, "file", j.URL.Scheme)
	assert.Equal(t, "localhost", j.URL.Host)
	assert.Equal(t, "/some/file.txt", j.URL.Path)
}

func TestURLUnmarshalEmptyString(t *testing.T) {
	data := []byte(`{
    "URL": ""
}`)
	var j tmpURL
	err := json.Unmarshal(data, &j)
	require.NoError(t, err)
	assert.Equal(t, URL{}, j.URL)
}

func TestURLUnmarshalBadType(t *testing.T) {
	data := []byte(`{
    "URL": 1
}`)
	var j tmpURL
	err := json.Unmarshal(data, &j)
	require.Error(t, err)
}

func TestURLUnmarshalBadURL(t *testing.T) {
	data := []byte(`{
    "URL": "://blah"
}`)
	var j tmpURL
	err := json.Unmarshal(data, &j)
	require.Error(t, err)
}

func TestURLMarshal(t *testing.T) {
	j := tmpURL{
		URL: URL{
			Scheme: "https",
			Host:   "insomniac.slackware.it",
		},
	}
	want := []byte("{\"URL\":\"https://insomniac.slackware.it\"}")
	b, err := json.Marshal(&j)
	require.NoError(t, err)
	assert.Equal(t, want, b)
}
