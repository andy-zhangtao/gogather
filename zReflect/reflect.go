package zReflect

import (
	"reflect"
)

//Write by zhangtao<ztao8607@gmail.com> . In 2018/5/23.

//ReflectStructInfo 获取结构体非空字段, 并将非空字段转换为Map[string]interface{}
//如果struct的字段中存在bson注解,将使用bson标记的字段名称作为key
func ReflectStructInfo(u interface{}) (structInfo map[string]interface{}) {

	getType := reflect.TypeOf(u)

	getValue := reflect.ValueOf(u)

	return parseStruct(getType, getValue)

}

func parseStruct(u reflect.Type, v reflect.Value) (structInfo map[string]interface{}) {
	structInfo = make(map[string]interface{})
	for i := 0; i < u.NumField(); i++ {
		field := u.Field(i)
		value := v.Field(i)
		zero := reflect.Zero(field.Type)

		if !reflect.DeepEqual(zero.Interface(), value.Interface()) {
			bName := field.Tag.Get("bson")
			if bName == "" {
				bName = field.Name
			}
			structInfo[bName] = value.Interface()
		}
	}
	return
}
