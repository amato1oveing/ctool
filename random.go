package ctool

import (
	"math/rand"
	"time"
)

var strBytes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var numBytes = []byte("0123456789")

// Random32Byte 生成32位以内数字和大小写英文随机字节
func Random32Byte() []byte {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 32)
	for i := range b {
		b[i] = strBytes[rand.Intn(len(strBytes))]
	}
	return b
}

// Random32Str 生成32位以内数字和大小写英文随机数
func Random32Str() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 32)
	for i := range b {
		b[i] = strBytes[rand.Intn(len(strBytes))]
	}
	return string(b)
}

// Random64Byte 生成64位以内数字和大小写英文随机字节
func Random64Byte() []byte {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 64)
	for i := range b {
		b[i] = strBytes[rand.Intn(len(strBytes))]
	}
	return b
}

func Random64Str() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 64)
	for i := range b {
		b[i] = strBytes[rand.Intn(len(strBytes))]
	}
	return string(b)
}

// Random32Num 生成32位以内数字随机数
func Random32Num() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 32)
	for i := range b {
		b[i] = numBytes[rand.Intn(len(numBytes))]
	}
	return string(b)
}
