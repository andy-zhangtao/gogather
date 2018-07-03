# zReflect
--
    import "github.com/andy-zhangtao/gogather/zReflect"


## Usage

#### func  ExtractValuePtrFromStruct

```go
func ExtractValuePtrFromStruct(uPtr interface{}, fields []string) (valPtr []interface{}, err error)
```
ExtractValuePtrFromStruct 从给定的结构体中根据给定的字段名称抽取出实际的字段内存地址 uPtr必须为指针类型
fields是字段名称,执行抽取操作时会对名称首字母进行大写处理。如: 给定的字段名称是name，实际比对的是Name.
如果找不到对应的字段，则跳过此字段继续往下查找。 依靠返回的valPtr是无法得知具体哪些字段查找成功，因此在使用此函数时必须确保fields名称的准确性

Example:

```go

    type User struct {
    	Name string `json:"name" bw:"name"`
    	Age  int    `json:"age" bw:"age"`
    	Addr string `json:"addr"`
    }

    var fileds = []string{
    	"name", "age", "addr",
    }
    u := new(User)
    vals, err := ExtractValuePtrFromStruct(u, fileds)

```

#### func  ReflectStructInfo

```go
func ReflectStructInfo(u interface{}, key ...string) (structInfo map[string]interface{})
```
ReflectStructInfo 获取结构体非空字段, 并将非空字段转换为Map[string]interface{}
如果struct的字段中存在bson注解,将使用bson标记的字段名称作为key 当struct中出现内联struct时,会通过
key1.key2的方式进行标记.

##### Example

```go

    u := User{
    	Name:     "andy@gmail.com",
    	Password: "pbkdf2_sha256$12000$sYPLrXcUlw0r$lNZsiNWBHS/9DUNsYvKYtL1UjxUPv+IKaYJ5JMJtz9U=",
    	Projects: Project{
    		ID: []string{
    			"iddd",
    		},
    	},
    	Statis: UserStatis{
    		DeployFailed: 4,
    	},
    	CurrentAuthority: "dev",
    }

    structInfo := ReflectStructInfo(u)

    assert.Equal(t, "andy@gmail.com", structInfo["name"])
    assert.Equal(t, "pbkdf2_sha256$12000$sYPLrXcUlw0r$lNZsiNWBHS/9DUNsYvKYtL1UjxUPv+IKaYJ5JMJtz9U=", structInfo["password"])
    assert.Equal(t, "iddd", structInfo["projects.id"].([]string)[0])
    assert.Equal(t, 4, structInfo["statis.deployfailed"].(int))
    assert.Equal(t, "dev", structInfo["currentauthority"])

```

User中的Projects同样是一个struct, 因此ID会被解析为 projects.id

#### func  ReflectStructInfoWithTag

```go
func ReflectStructInfoWithTag(u interface{}, allowEmpty bool, tag string, key ...string) (structInfo map[string]interface{})
```
ReflectStructInfoWithTag 抽取特定标签的非空字段 使用方式同ReflectStructInfo一致, 只增加了Tag属性
