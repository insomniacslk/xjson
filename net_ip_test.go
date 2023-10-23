package xjson

import (
	"encoding/json"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var localhostv4 = net.ParseIP("127.0.0.1")

type tmpIP struct {
	IP IP
}

func TestIPUnmarshal(t *testing.T) {
	data := []byte(`{
"IP": "127.0.0.1"
}`)
	var j tmpIP
	err := json.Unmarshal(data, &j)
	require.NoError(t, err)
	want := IP(localhostv4)
	assert.Equal(t, want, j.IP)
}

func TestIPUnmarshalEmptyString(t *testing.T) {
	data := []byte(`{
    "IP": ""
}`)
	var j tmpIP
	err := json.Unmarshal(data, &j)
	require.Error(t, err)
}

func TestIPUnmarshalBadType(t *testing.T) {
	data := []byte(`{
    "IP": 1
}`)
	var j tmpIP
	err := json.Unmarshal(data, &j)
	require.Error(t, err)
}

func TestIPMarshal(t *testing.T) {
	j := tmpIP{
		IP: IP(localhostv4),
	}
	want := []byte("{\"IP\":\"127.0.0.1\"}")
	b, err := json.Marshal(&j)
	require.NoError(t, err)
	assert.Equal(t, want, b)
}
