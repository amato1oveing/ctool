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

//ToBool interface转bool
func ToBool(val interface{}) bool {
	if val == nil {
		return false
	}
	switch v := val.(type) {
	case bool:
		return v
	default:
		return false
	}
}

// ToInt64 interface转int64
func ToInt64(val interface{}) int64 {
	if val == nil {
		return 0
	}
	switch v := val.(type) {
	case int, int8, int16, int32, uint, uint8, uint16, uint32, uint64, float32, float64:
		return reflect.ValueOf(v).Int()
	case int64:
		return v
	default:
		return 0
	}
}

// ToInt interface转int
func ToInt(val interface{}) int {
	return int(ToInt64(val))
}

// ToInt32 interface转int32
func ToInt32(val interface{}) int32 {
	return int32(ToInt64(val))
}

// ToInt16 interface转int16
func ToInt16(val interface{}) int16 {
	return int16(ToInt64(val))
}

// ToUint64 interface转uint64
func ToUint64(val interface{}) uint64 {
	if val == nil {
		return 0
	}
	switch v := val.(type) {
	case int, int8, int16, int32, uint, uint8, uint16, uint32, float32, float64:
		return reflect.ValueOf(v).Uint()
	case uint64:
		return v
	default:
		return 0
	}
}

// ToUint interface转uint
func ToUint(val interface{}) uint {
	return uint(ToUint64(val))
}

// ToUint32 interface转uint32
func ToUint32(val interface{}) uint32 {
	return uint32(ToUint64(val))
}

// ToUint16 interface转uint16
func ToUint16(val interface{}) uint16 {
	return uint16(ToUint64(val))
}

// ToFloat64 interface转float64
func ToFloat64(val interface{}) float64 {
	if val == nil {
		return 0
	}
	switch v := val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32:
		return float64(reflect.ValueOf(v).Int())
	case float64:
		return v
	default:
		return 0
	}
}

// ToFloat32 interface转float32
func ToFloat32(val interface{}) float32 {
	return float32(ToFloat64(val))
}

// ToFloat interface转float
func ToFloat(val interface{}) float64 {
	return ToFloat64(val)
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
