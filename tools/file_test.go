package tools

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
)

//Write by zhangtao<ztao8607@gmail.com> . In 2018/3/30.

func TestLineCounter(t *testing.T) {
	file, err := os.Open("./file_test.go")
	assert.Empty(t, err)

	num, err := LineCounter(file)
	assert.Empty(t, err)

	assert.Equal(t, num, 32)
}

func TestCopyFile(t *testing.T) {
	ioutil.WriteFile("/tmp/copy.test", []byte("This is a copy test"), 0777)

	err := CopyFile("/tmp/copy.test", "/tmp/copy.test.1")

	except, err := ioutil.ReadFile("/tmp/copy.test.1")
	assert.Nil(t, err)
	assert.EqualValues(t, "This is a copy test", string(except))

	os.Remove("/tmp/copy.test.1")
}
