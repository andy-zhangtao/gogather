package znet

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_GetFreePort(t *testing.T) {
	_, err := GetFreePort()

	assert.Nil(t, err)

}
