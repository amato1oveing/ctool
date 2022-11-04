package ctool

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path"
)

// ReadFile 读取文件
func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

//ReadFileNoErr 读取文件，无错误返回
func ReadFileNoErr(path string) []byte {
	data, _ := ioutil.ReadFile(path)
	return data
}

//ReadFileToStr 读取文件内容到字符串
func ReadFileToStr(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ReadFileByLine 按行读取文件
func ReadFileByLine(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	result := make([]string, 0)
	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()
		l := string(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		result = append(result, l)
	}

	return result, nil
}

// CreateFile 创建文件
func CreateFile(path string) bool {
	file, err := os.Create(path)
	if err != nil {
		return false
	}

	defer file.Close()
	return true
}

// IsExist 判断文件是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

// CreateDir 创建目录
func CreateDir(absPath string) error {
	return os.MkdirAll(path.Dir(absPath), os.ModePerm)
}

// GetAbsPath 获取绝对路径
func GetAbsPath(path string) (string, error) {
	return os.Getwd()
}

// CopyFile 从srcFilePath拷贝到dstFilePath
func CopyFile(srcFilePath string, dstFilePath string) error {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	distFile, err := os.Create(dstFilePath)
	if err != nil {
		return err
	}
	defer distFile.Close()

	var tmp = make([]byte, 1024*4)
	for {
		n, err := srcFile.Read(tmp)
		distFile.Write(tmp[:n])
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

//ClearFile 清空文件
func ClearFile(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString("")
	return err
}
