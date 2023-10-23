package xjson

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type tmpHardwareAddr struct {
	HardwareAddr HardwareAddr
}

func TestHardwareAddrUnmarshal(t *testing.T) {
	data := []byte(`{
"HardwareAddr": "aa:bb:cc:dd:ee:ff"
}`)
	var j tmpHardwareAddr
	err := json.Unmarshal(data, &j)
	require.NoError(t, err)
	want := HardwareAddr([]byte{0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff})
	assert.Equal(t, want, j.HardwareAddr)
}

func TestHardwareAddrUnmarshalEmptyString(t *testing.T) {
	data := []byte(`{
    "HardwareAddr": ""
}`)
	var j tmpHardwareAddr
	err := json.Unmarshal(data, &j)
	require.Error(t, err)
}

func TestHardwareAddrUnmarshalBadType(t *testing.T) {
	data := []byte(`{
    "HardwareAddr": 1
}`)
	var j tmpHardwareAddr
	err := json.Unmarshal(data, &j)
	require.Error(t, err)
}

func TestHardwareAddrMarshal(t *testing.T) {
	j := tmpHardwareAddr{
		HardwareAddr: HardwareAddr([]byte{0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff}),
	}
	want := []byte("{\"HardwareAddr\":\"aa:bb:cc:dd:ee:ff\"}")
	b, err := json.Marshal(&j)
	require.NoError(t, err)
	assert.Equal(t, want, b)
}
