package znginx

import (
	"fmt"
	"strings"
)

//InsertLocation 插入一段新的Location片段
//此函数不检查是否已经存在相同的Location数据
func InsertLocation(nginx, location string) string {

	nginx = strings.TrimSpace(nginx)
	nginx = nginx[:len(nginx)-1]

	return fmt.Sprintf("%s\n%s\n}", nginx, location)
}
