package time

import (
	"time"
	"strings"
	)

type Ztime struct {
	time time.Time
	err  error
}

func (this *Ztime) Now() (*Ztime) {
	this.time = time.Now()
	return this
}

func (this *Ztime) SetLocation(timezone string) (*Ztime) {
	if this.err != nil {
		return this
	}

	t, err := time.LoadLocation(timezone)
	if err != nil {
		this.err = err
		return this
	}

	this.time.In(t)
	return this
}

//Format 设定时间格式
//兼容官方Format格式. 同时支持YYYY-MM-DD hh:mm:ss格式
/*
##### Example

```go
	new(Ztime).Now().UTC().AddHour(7).Format("YYYY-MM-DDThh:mm")
```

*/
func (this *Ztime) Format(format string) (string, error) {
	if strings.Contains(format, "YYYY") {
		format = strings.Replace(format, "YYYY", "2006", -1)
	}

	if strings.Contains(format, "MM") {
		format = strings.Replace(format, "MM", "01", -1)
	}

	if strings.Contains(format, "DD") {
		format = strings.Replace(format, "DD", "02", -1)
	}

	if strings.Contains(format, "hh") {
		format = strings.Replace(format, "hh", "15", -1)
	}

	if strings.Contains(format, "mm") {
		format = strings.Replace(format, "mm", "04", -1)
	}

	if strings.Contains(format, "ss") {
		format = strings.Replace(format, "ss", "05", -1)
	}

	return this.time.Format(format), this.err
}

func (this *Ztime) String() (string, error) {
	return this.time.String(), this.err
}

//AddHour 调整时间
/*

##### Example

```go
	new(Ztime).Now().SetLocation("Asia/Shanghai").AddHour(-1).Format("2006-01-02T15:04")
```
*/
func (this *Ztime) AddHour(n int) (*Ztime) {
	this.time = this.time.Add(time.Duration(n) * time.Hour)
	return this
}

//UTC 返回UTC时间
//在当前没有时区文件的场景中可以通过UTC+AddHour计算指定时区的时间
/*
##### Example

```go
	new(Ztime).Now().UTC()
```
*/
func (this *Ztime) UTC() (*Ztime) {
	this.time = this.time.UTC()
	return this
}
