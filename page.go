package ctool

// Page 分页
func Page[T any, U int64 | int32 | int](data []T, page, size U) []T {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 1
	}
	start := (page - 1) * size
	end := page * size
	if start >= U(len(data)) {
		return nil
	}
	if end > U(len(data)) {
		end = U(len(data))
	}
	return data[start:end]
}

// PageInfo 获取分页信息,size默认为10
func PageInfo[T int64 | int32 | int](page, size, total T) (totalPage, currentPage, pageSize T) {
	if total == 0 {
		return
	}
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	totalPage = total / size
	if total%size != 0 {
		totalPage++
	}
	currentPage = page
	pageSize = size
	return
}
