package zlog

import (
	"github.com/petermattis/goid"
	"github.com/andy-zhangtao/gogather/random"
	"errors"
	"fmt"
)

type Zlog struct {
	idMap map[int64]string
}

//GetZlog 获取跟踪ID实例
/*

##### Example

```go

	z := zlog.GetZlog()
	logrus.WithFields(z.Fields(logrus.Fields{"key": value})).Info("main")

```

*/
func GetZlog() (*Zlog) {
	return &Zlog{
		idMap: make(map[int64]string),
	}
}

//Fields 在提供的f基础上增加跟踪ID
func (this *Zlog) Fields(f map[string]interface{}) (map[string]interface{}) {
	_id := goid.Get()
	id := this.idMap[_id]
	if id == "" {
		id = random.GetString(12)
		this.idMap[_id] = id
	}

	f["_track"] = id
	return f
}

//AddID 手动添加跟踪ID.
func (this *Zlog) AddID(id string) (*Zlog) {
	this.idMap[goid.Get()] = id
	return this
}

//MyTrack 获取当前跟踪ID
func (this *Zlog) MyTrack() (string) {
	_id := goid.Get()
	if id := this.idMap[_id]; id == "" {
		this.idMap[_id] = random.GetString(12)
		return this.idMap[_id]
	} else {
		return id
	}

}

//Clean 清除跟踪ID
func (this *Zlog) Clean() {
	delete(this.idMap, goid.Get())
}

func (this *Zlog) Error(str string) error {
	return errors.New(fmt.Sprintf("_track [%s] %s", this.idMap[goid.Get()], str))
}
