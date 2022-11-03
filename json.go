package ctool

import "encoding/json"

// MarshalNoErr 序列化，无错误返回
func MarshalNoErr(any interface{}) string {
	data, _ := json.Marshal(any)
	return string(data)
}

// Unmarshal 反序列化
func Unmarshal(data []byte, obj any) error {
	return json.Unmarshal(data, obj)
}

// UnmarshalNoErr 反序列化，无错误返回
func UnmarshalNoErr(data []byte, obj any) {
	_ = json.Unmarshal(data, obj)
}
