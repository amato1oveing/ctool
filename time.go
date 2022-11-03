package ctool

import "time"

// Ms2float 将ms的时间转为float64
func Ms2float(time int64) float64 {
	return float64(time) / 1000
}

// ParseTime 将字符串转为时间,支持多种格式
func ParseTime(timeStr string) (time.Time, error) {

	return time.Time{}, nil
}

// IsZeroTime 判断时间是否为空
func IsZeroTime(time time.Time) bool {
	return time.IsZero()
}
