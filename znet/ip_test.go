package znet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LocalIP(t *testing.T) {
	_, err := LocallIP()

	assert.Nil(t, err)
}

func TestConvertToHex(t *testing.T) {

	hexIp, err := ConvertToHex("127.0.0.1")

	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'7', 'F', '0', '0', '0', '0', '0', '1'}, hexIp)

	hexIp, err = ConvertToHex("192.168.1.1")

	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'C', '0', 'A', '8', '0', '1', '0', '1'}, hexIp)

	hexIp, err = ConvertToHex("192.168.1")

	assert.Error(t, err)
}
