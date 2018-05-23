# zReflect
--
    import "github.com/andy-zhangtao/gogather/zReflect"


## Usage

#### func  ReflectStructInfo

```go
func ReflectStructInfo(u interface{}) (structInfo map[string]interface{})
```
ReflectStructInfo 获取结构体非空字段, 并将非空字段转换为Map[string]interface{}
如果struct的字段中存在bson注解,将使用bson标记的字段名称作为key
