package zReflect

import (
	"reflect"
	"errors"
	"strings"
)

//Write by zhangtao<ztao8607@gmail.com> . In 2018/5/23.

//ReflectStructInfo 获取结构体非空字段, 并将非空字段转换为Map[string]interface{}
//如果struct的字段中存在bson注解,将使用bson标记的字段名称作为key
//当struct中出现内联struct时,会通过 key1.key2的方式进行标记.
/*

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
*/
func ReflectStructInfo(u interface{}, key ...string) (structInfo map[string]interface{}) {

	getType := reflect.TypeOf(u)

	getValue := reflect.ValueOf(u)

	if getType.Kind() == reflect.Ptr {
		return parseStruct(getType.Elem(), getValue.Elem(), false, "bson", key...)
	}

	return parseStruct(getType, getValue, false, "bson", key...)

}

//ReflectStructInfoWithTag 抽取特定标签的非空字段
//使用方式同ReflectStructInfo一致, 只增加了Tag属性
func ReflectStructInfoWithTag(u interface{}, allowEmpty bool, tag string, key ...string) (structInfo map[string]interface{}) {

	getType := reflect.TypeOf(u)

	getValue := reflect.ValueOf(u)

	if getType.Kind() == reflect.Ptr {
		return parseStruct(getType.Elem(), getValue.Elem(), allowEmpty, tag, key...)
	}

	return parseStruct(getType, getValue, allowEmpty, tag, key...)

}

//ExtractValuePtrFromStruct 从给定的结构体中根据给定的字段名称抽取出实际的字段内存地址
//uPtr必须为指针类型
//fields是字段名称,执行抽取操作时会对名称首字母进行大写处理。如: 给定的字段名称是name，实际比对的是Name.
//如果找不到对应的字段，则跳过此字段继续往下查找。
//依靠返回的valPtr是无法得知具体哪些字段查找成功，因此在使用此函数时必须确保fields名称的准确性
/*

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
*/
func ExtractValuePtrFromStruct(uPtr interface{}, fields []string) (valPtr []interface{}, err error) {
	if reflect.TypeOf(uPtr).Kind() != reflect.Ptr {
		err = errors.New("Params uPtr Must Be A Pointer")
		return
	}

	uValue := reflect.ValueOf(uPtr).Elem()

	for _, key := range fields {
		values := uValue.FieldByName(strings.Title(key))
		if values.IsValid() {
			valPtr = append(valPtr, values.Addr().Interface())
		}
	}

	return
}

func parseStruct(u reflect.Type, v reflect.Value, allowEmpty bool, tag string, key ...string) (structInfo map[string]interface{}) {

	structInfo = make(map[string]interface{})
	for i := 0; i < u.NumField(); i++ {
		field := u.Field(i)
		value := v.Field(i)
		zero := reflect.Zero(field.Type)
		bName := field.Tag.Get(tag)
		if bName == "" {
			continue
		}

		if bName == "-" {
			continue
		}

		if len(key) > 0 {
			bName = key[0] + bName
		}

		if !reflect.DeepEqual(zero.Interface(), value.Interface()) || allowEmpty {
			switch reflect.TypeOf(value.Interface()).Kind() {
			case reflect.String:
				fallthrough
			case reflect.Int:
				fallthrough
			case reflect.Bool:
				fallthrough
			case reflect.Array:
				fallthrough
			case reflect.Slice:
				fallthrough
			case reflect.Map:
				fallthrough
			case reflect.Float32:
				fallthrough
			case reflect.Float64:
				structInfo[bName] = value.Interface()
			case reflect.Struct:
				for k, v := range ReflectStructInfo(value.Interface(), bName+".") {
					structInfo[k] = v
				}

			}
		}
	}
	return
}
