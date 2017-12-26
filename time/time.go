package time

import (
	"time"
	"strconv"
)

// GetTimeStamp 获取当前时间戳
// length 10:获取秒 13:获取毫秒
func GetTimeStamp(length int) string {
	return strconv.FormatInt(time.Now().UTC().UnixNano(), 10)[:length]
}
