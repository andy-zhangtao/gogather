package znet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertIP(t *testing.T) {
	iphex := "E702140A"

	ip, err := ConvertIP(iphex)
	assert.Nil(t, err)

	assert.Equal(t, "10.20.2.231", ip)
}
