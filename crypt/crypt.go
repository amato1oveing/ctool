package crypt

import (
	"crypto/md5"
	"fmt"
)

//Md5 md5加密
func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

//Md5WithPrefix md5加密,并且加前缀
func Md5WithPrefix(str, prefix string) string {
	return prefix + Md5(str)
}

//Md5WithSalt md5加密,并且加盐
func Md5WithSalt(str, salt string) string {
	return Md5(str + salt)
}

//Md5WithTimes md5加密,并且加密次数
func Md5WithTimes(str string, times int) string {
	md5Str := Md5(str)
	for i := 0; i < times; i++ {
		md5Str = Md5(md5Str)
	}
	return md5Str
}

//Md5WithSaltAndTimes md5加密,并且加盐,并且加密次数
func Md5WithSaltAndTimes(str, salt string, times int) string {
	md5Str := Md5WithSalt(str, salt)
	for i := 0; i < times; i++ {
		md5Str = Md5(md5Str)
	}
	return md5Str
}
