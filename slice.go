package ctool

// Contain 元素是否存在于数组中
func Contain[T comparable](val T, array []T) (exists bool, index int) {
	for i, v := range array {
		if v == val {
			return true, i
		}
	}
	return false, -1
}

// Equal 判断两个数组是否相等
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Remove 根据索引从数组中移除元素
func Remove[T comparable](array []T, index int) []T {
	if index < 0 || index >= len(array) {
		return array
	}
	return append(array[:index], array[index+1:]...)
}

// RemoveByValue 根据值从数组中移除元素
func RemoveByValue[T comparable](array []T, val T) []T {
	_, index := Contain(val, array)
	return Remove(array, index)
}

// RemoveByValues 根据值从数组中移除元素
func RemoveByValues[T comparable](array []T, vals []T) []T {
	for _, val := range vals {
		array = RemoveByValue(array, val)
	}
	return array
}

// RemoveByIndexs 根据索引从数组中移除元素
func RemoveByIndexs[T comparable](array []T, indexs []int) []T {
	for _, index := range indexs {
		array = Remove(array, index)
	}
	return array
}

// RemoveDuplicate 移除数组中重复的元素
func RemoveDuplicate[T comparable](array []T) []T {
	var result []T
	for _, val := range array {
		if _, index := Contain(val, result); index == -1 {
			result = append(result, val)
		}
	}
	return result
}

// RemoveDuplicateByMap 移除数组中重复的元素
func RemoveDuplicateByMap[T comparable](array []T) []T {
	var result []T
	m := make(map[T]bool)
	for _, val := range array {
		if _, ok := m[val]; !ok {
			m[val] = true
			result = append(result, val)
		}
	}
	return result
}

// Reverse 反转数组
func Reverse[T comparable](array []T) []T {
	var result []T
	for i := 0; i < len(array)/2; i++ {
		array[i], array[len(array)-1-i] = array[len(array)-1-i], array[i]
	}
	return result
}

// Merge 合并数组
func Merge[T comparable](array1, array2 []T) []T {
	return append(array1, array2...)
}

// MergeAndRemoveDuplicate 合并数组并去重
func MergeAndRemoveDuplicate[T comparable](array1, array2 []T) []T {
	m := make(map[T]bool)
	for _, val := range array1 {
		m[val] = true
	}
	for _, val := range array2 {
		m[val] = true
	}
	var result []T
	for k := range m {
		result = append(result, k)
	}
	return result
}
