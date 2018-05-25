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

#### type Ztime

```go
type Ztime struct {
}
```


#### func (*Ztime) AddHour

```go
func (this *Ztime) AddHour(n int) *Ztime
```
AddHour 调整时间

##### Example

```go

    new(Ztime).Now().SetLocation("Asia/Shanghai").AddHour(-1).Format("2006-01-02T15:04")

```

#### func (*Ztime) Format

```go
func (this *Ztime) Format(format string) (string, error)
```
Format 设定时间格式 兼容官方Format格式. 同时支持YYYY-MM-DD hh:mm:ss格式

##### Example

```go

    new(Ztime).Now().UTC().AddHour(7).Format("YYYY-MM-DDThh:mm")

```

#### func (*Ztime) Now

```go
func (this *Ztime) Now() *Ztime
```

#### func (*Ztime) SetLocation

```go
func (this *Ztime) SetLocation(timezone string) *Ztime
```

#### func (*Ztime) String

```go
func (this *Ztime) String() (string, error)
```

#### func (*Ztime) UTC

```go
func (this *Ztime) UTC() *Ztime
```
UTC 返回UTC时间 在当前没有时区文件的场景中可以通过UTC+AddHour计算指定时区的时间

##### Example

```go

    new(Ztime).Now().UTC()

```
