package xjson

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type tmp struct {
	URL URL
}

func TestURLUnmarshaller(t *testing.T) {
	// test valid file://
	data := []byte(`{
    "URL": "file://localhost/some/file.txt"
}`)
	var j tmp
	err := json.Unmarshal(data, &j)
	require.NoError(t, err)
	require.Equal(t, "file", j.URL.Scheme)
	require.Equal(t, "localhost", j.URL.Host)
	require.Equal(t, "/some/file.txt", j.URL.Path)
}

func TestURLUnmarshallerQuotes(t *testing.T) {
	// test valid file://
	data := []byte(`{
    "URL": "file://localh\ost/some/file.txt"
}`)
	var j tmp
	err := json.Unmarshal(data, &j)
	// unquoting is not supported yet
	require.Error(t, err)
}

func TestURLMarshaller(t *testing.T) {
	j := tmp{
		URL: URL{
			Scheme: "https",
			Host:   "insomniac.slackware.it",
		},
	}
	want := []byte("{\"URL\":\"https://insomniac.slackware.it\"}")
	b, err := json.Marshal(&j)
	require.NoError(t, err)
	require.Equal(t, want, b)
}
