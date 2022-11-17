package ctool

type mapUtil struct {
	lock   chan struct{}
	bucket map[string]interface{}
}

//NewMapUtil 创建一个map工具
func NewMapUtil() *mapUtil {
	return &mapUtil{
		lock:   make(chan struct{}, 1),
		bucket: make(map[string]interface{}),
	}
}

// Lock 加锁
func (m *mapUtil) Lock() {
	m.lock <- struct{}{}
}

// Unlock 解锁
func (m *mapUtil) Unlock() {
	<-m.lock
}

// Has 判断key是否存在
func (m *mapUtil) Has(key string) bool {
	m.Lock()
	_, ok := m.bucket[key]
	m.Unlock()
	return ok
}

// GetOrSet 存在则返回已设值,不存在则设置值
func (m *mapUtil) GetOrSet(key string, value interface{}) interface{} {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.bucket[key]; !ok {
		m.bucket[key] = value
	}
	return m.bucket[key]
}

// GetOrSetFunc 存在则返回已设值,不存在则设置值,并且使用函数返回值
func (m *mapUtil) GetOrSetFunc(key string, f func() interface{}) interface{} {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.bucket[key]; !ok {
		m.bucket[key] = f()
	}
	return m.bucket[key]
}

// Remove 移除key
func (m *mapUtil) Remove(key string) {
	m.Lock()
	delete(m.bucket, key)
	m.Unlock()
}

// Clear 清空map
func (m *mapUtil) Clear() {
	m.Lock()
	m.bucket = make(map[string]interface{})
	m.Unlock()
}

// Size 获取map长度
func (m *mapUtil) Size() int {
	m.Lock()
	size := len(m.bucket)
	m.Unlock()
	return size
}

// Keys 获取所有key
func (m *mapUtil) Keys() []string {
	m.Lock()
	keys := make([]string, 0, len(m.bucket))
	for k := range m.bucket {
		keys = append(keys, k)
	}
	m.Unlock()
	return keys
}

// Values 获取所有value
func (m *mapUtil) Values() []interface{} {
	m.Lock()
	values := make([]interface{}, 0, len(m.bucket))
	for _, v := range m.bucket {
		values = append(values, v)
	}
	m.Unlock()
	return values
}

//Set 设置值
func (m *mapUtil) Set(key string, value interface{}) {
	m.Lock()
	m.bucket[key] = value
	m.Unlock()
}

//Get 获取值
func (m *mapUtil) Get(key string) interface{} {
	m.Lock()
	value := m.bucket[key]
	m.Unlock()
	return value
}

//GetBool 获取bool值
func (m *mapUtil) GetBool(key string) bool {
	m.Lock()
	value := m.bucket[key]
	m.Unlock()
	return ToBool(value)
}

//GetFloat64 获取float64值
func (m *mapUtil) GetFloat64(key string) float64 {
	m.Lock()
	value := m.bucket[key]
	m.Unlock()
	return ToFloat64(value)
}

//GetInt 获取int值
func (m *mapUtil) GetInt(key string) int {
	m.Lock()
	value := m.bucket[key]
	m.Unlock()
	return ToInt(value)
}

//GetInt64 获取int64值
func (m *mapUtil) GetInt64(key string) int64 {
	m.Lock()
	value := m.bucket[key]
	m.Unlock()
	return ToInt64(value)
}

//GetString 获取string值
func (m *mapUtil) GetString(key string) string {
	m.Lock()
	value := m.bucket[key]
	m.Unlock()
	return ToString(value)
}
