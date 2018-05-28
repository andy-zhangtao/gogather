package time

import (
	"time"
	"fmt"
)

const (
	OnDay = iota
	OnHour
	OnMin
)

//OnTimer 定时发生器
//OnDay 每天零时执行callback
//OnHour 每个小时执行callback
//OnMin 每分钟执行callback
//如果执行过程出现error, 则通过chan error获取具体错误原因
/*

##### Example

```go
//定时每天执行callback函数
package main

import (
	"github.com/andy-zhangtao/gogather/zlog"
	"github.com/sirupsen/logrus"
	"github.com/andy-zhangtao/gogather/time"
)

var z *zlog.Zlog

func callback() error {
	logrus.WithFields(z.Fields(logrus.Fields{"from": "callback"})).Info("callback")
	return nil
}

func main() {
	z = zlog.GetZlog()

	time.OnTimer(time.OnDay, 1, callback)
}

```
*/
func OnTimer(kind, duration int, callback func() error) (chan error) {
	var next time.Time
	errChan := make(chan error)

	for {
		now := time.Now()
		switch kind {
		case OnDay:
			next = now.Add(time.Hour * time.Duration(duration*24))
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		case OnHour:
			next = now.Add(time.Minute * time.Duration(duration*60))
			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), 0, 0, 0, next.Location())
		case OnMin:
			next = now.Add(time.Second * time.Duration(duration*60))
			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), 0, 0, next.Location())
		}

		t := time.NewTimer(next.Sub(now))
		fmt.Printf("下次采集时间为[%s]\n", next.Format("200601021504"))

		select {
		case <-t.C:
			if err := callback(); err != nil {
				return errChan
			}
		}
	}
}
