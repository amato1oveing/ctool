package ctool

import "fmt"

//NewError 创建一个新的错误
func NewError(value string) error {
	return fmt.Errorf(value)
}

//NewErrorf 创建一个新的错误,并且可以使用fmt.Sprintf的方式传参
func NewErrorf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
