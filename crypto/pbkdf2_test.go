package crypto

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

//Write by zhangtao<ztao8607@gmail.com> . In 2018/5/7.
func TestGeneratePasswd(t *testing.T) {
	password := GeneratePasswd("admin", "yhYWUEo4DNqj", 36000)

	assert.Equal(t, "pbkdf2_sha256$36000$yhYWUEo4DNqj$SpxtdIOm9nwRG+X76jUUlGvdDcLaMBl7Z+rJ8sfSMcU=", password, "Generate Password Error")
}
