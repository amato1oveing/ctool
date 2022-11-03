package ctool

import "reflect"

// InArray 元素是否存在于数组中
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)    //获取数组的reflect.Value
		for i := 0; i < s.Len(); i++ { //遍历数组
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true { //比较数组的值
				exists = true
				index = i
				return
			}
		}
	}
	return
}
