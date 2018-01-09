package convert

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func TestFindJsonFromStr(t *testing.T) {
	oldStr := "{\"code\":4000,\"message\":\"(-160004)待操作的资源在kubernetes中不存在\",\"codedesc\":\"KubeResourceNotFound\",\"data\":{\"totalcount\":0,\"services\":null}}{\"code\":4000,\"message\":\"(-160004)待操作的资源在kubernetes中不存在\",\"codedesc\":\"KubeResourceNotFound\",\"data\":{\"totalcount\":0,\"services\":null}}{\"code\":0,\"message\":\"\",\"codedesc\":\"Success\",\"data\":{\"totalcount\":0,\"services\":null}}{\"code\":0,\"message\":\"\",\"codedesc\":\"Success\",\"data\":{\"totalcount\":0,\"services\":null}}"
	splJs, err := FindJsonFromStr(oldStr)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, 4, len(splJs), "The size of json error!")

	for _, j := range splJs {
		_, err := json.Marshal(j)
		assert.Equal(t, nil, err, "Conver Json Faile!")
	}
}
