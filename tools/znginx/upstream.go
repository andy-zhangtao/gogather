package znginx

import (
	"fmt"
	"strings"
)

//InsertUpstream 插入一段新的Upstream片段
//此函数不检查是否已经存在相同的Upstream数据
func InsertUpstream(nginx, upstream string) string {

	nginx = strings.TrimSpace(nginx)

	return fmt.Sprintf("%s\n%s", upstream, nginx)
}
