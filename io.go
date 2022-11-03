package ctool

import "io/ioutil"

// ReadFile 读取文件
func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

//ReadFileNoErr 读取文件，无错误返回
func ReadFileNoErr(path string) []byte {
	data, _ := ioutil.ReadFile(path)
	return data
}
