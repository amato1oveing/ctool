package ctool

import (
	"fmt"
	"strconv"
)

// StrToInt64NoErr string转int64
func StrToInt64NoErr(str string) int64 {
	parseInt, _ := strconv.ParseInt(str, 10, 64)
	return parseInt
}

// ToString interface转string
func ToString(val interface{}) string {
	return fmt.Sprintf("%v", val)
}

// Bool2String bool转string
func Bool2String(is bool) string {
	if is {
		return "true"
	}
	return "false"
}
