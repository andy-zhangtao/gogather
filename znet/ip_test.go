package znet

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_LocalIP(t *testing.T) {
	_, err := LocallIP()

	assert.Nil(t, err)
}
