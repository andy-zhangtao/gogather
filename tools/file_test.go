package tools

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)

//Write by zhangtao<ztao8607@gmail.com> . In 2018/3/30.

func TestLineCounter(t *testing.T) {
	file, err := os.Open("./file_test.go")
	assert.Empty(t, err)

	num, err := LineCounter(file)
	assert.Empty(t, err)

	assert.Equal(t, num, 18)
}