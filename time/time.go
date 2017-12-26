package time

import (
	"time"
)

// GetTimeStamp 获取当前时间戳
// length 10:获取秒 13:获取毫秒
func GetTimeStamp(length int) string {
	t := time.Now()

	switch length {
	case 10:
		return t.Format("2006010215")
	case 13:
		return t.Format("20060102150405")
	default:
		return t.Format("2006010215")
	}
}
