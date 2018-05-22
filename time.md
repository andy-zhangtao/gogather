# time
--
    import "github.com/andy-zhangtao/gogather/time"


## Usage

#### func  GetCurrentDayOfMonth

```go
func GetCurrentDayOfMonth(t time.Time) (day int)
```
GetCurrentDayOfMonth 获取给定的时间是当月的第几天

#### func  GetCurrentDayOfYear

```go
func GetCurrentDayOfYear(t time.Time) (day int, err error)
```

#### func  GetTimeStamp

```go
func GetTimeStamp(length int) string
```
GetTimeStamp 获取当前时间戳 length 10:获取秒 13:获取毫秒
