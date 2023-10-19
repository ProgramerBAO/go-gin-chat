package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// 加密密码,这样数据库管理员也看不到明文密码
// 唯一文档也可以使用
// 可以增加算法

// Md5EncodeToLower 转小写
func Md5EncodeToLower(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// Md5EncodeToUpper 转大写
func Md5EncodeToUpper(data string) string {
	return strings.ToUpper(Md5EncodeToLower(data))
}

// MakePassword 加密
func MakePassword(plainpwd, salt string) string {
	return Md5EncodeToLower(plainpwd + salt)
}
