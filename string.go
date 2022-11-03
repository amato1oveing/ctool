package ctool

import (
	"encoding/json"
	"fmt"
	"net"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var isAlphaRegexMatcher = regexp.MustCompile(`^[a-zA-Z]+$`)

// IsAlpha 判断字符串是否全是字母
func IsAlpha(str string) bool {
	return isAlphaRegexMatcher.MatchString(str)
}

// IsNumberStr 判断字符串是否是数字
func IsNumberStr(s string) bool {
	return IsIntStr(s) || IsFloatStr(s)
}

// IsIntStr 判断字符串是否是整数
func IsIntStr(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

// IsFloatStr 判断字符串是否是浮点数
func IsFloatStr(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// IsJSON 判断字符串是否是json
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

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

// IsIp 判断是否是ip
func IsIp(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil
}

// IsIpV4 判断是否是ipv4
func IsIpV4(str string) bool {
	ip := net.ParseIP(str)
	if ip == nil {
		return false
	}
	return strings.Contains(str, ".")
}

// IsIpV6 判断是否是ipv6
func IsIpV6(str string) bool {
	ip := net.ParseIP(str)
	if ip == nil {
		return false
	}
	return strings.Contains(str, ":")
}

// IsPort 判断是否是端口
func IsPort(str string) bool {
	if i, err := strconv.ParseInt(str, 10, 64); err == nil && i > 0 && i < 65536 {
		return true
	}
	return false
}

var isEmailRegexMatcher = regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)

// IsEmail 判断是否是邮箱
func IsEmail(email string) bool {
	return isEmailRegexMatcher.MatchString(email)
}

// IsZeroValue 判断是否是零值
func IsZeroValue(value any) bool {
	if value == nil {
		return true
	}

	rv := reflect.ValueOf(value)
	if !rv.IsValid() {
		return true
	}

	switch rv.Kind() {
	case reflect.String:
		return rv.Len() == 0
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.Ptr, reflect.Chan, reflect.Func, reflect.Interface, reflect.Slice, reflect.Map:
		return rv.IsNil()
	}

	return reflect.DeepEqual(rv.Interface(), reflect.Zero(rv.Type()).Interface())
}
