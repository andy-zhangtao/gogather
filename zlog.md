# zlog
--
    import "github.com/andy-zhangtao/gogather/zlog"


## Usage

#### type Zlog

```go
type Zlog struct {
}
```


#### func  GetZlog

```go
func GetZlog() *Zlog
```
GetZlog 获取跟踪ID实例

##### Example

```go

    z := zlog.GetZlog()
    logrus.WithFields(z.Fields(logrus.Fields{"key": value})).Info("main")

```

#### func (*Zlog) AddID

```go
func (this *Zlog) AddID(id string) *Zlog
```
AddID 手动添加跟踪ID.

#### func (*Zlog) Clean

```go
func (this *Zlog) Clean()
```
Clean 清除跟踪ID

#### func (*Zlog) Error

```go
func (this *Zlog) Error(str string) error
```

#### func (*Zlog) Fields

```go
func (this *Zlog) Fields(f map[string]interface{}) map[string]interface{}
```
Fields 在提供的f基础上增加跟踪ID

#### func (*Zlog) MyTrack

```go
func (this *Zlog) MyTrack() string
```
MyTrack 获取当前跟踪ID
