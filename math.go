package ctool

// Max 返回两数中大的数
func Max[T int64 | int32 | int | float32 | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min 返回两数中小的数
func Min[T int64 | int32 | int | float32 | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}
