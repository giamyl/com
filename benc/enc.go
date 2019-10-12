package benc

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

// PwdSlat 给密码通过盐加密
func PwdSlat(password, slat string) string {

	// 如果传入的时空字符串， 就不做处理
	if password == "" {
		return password
	}

	if slat == "" {
		slat = "default"
	}

	md5Slat := md5.Sum([]byte(slat))
	sha256Pass := sha256.Sum256([]byte(password))

	lenSlat := len(md5Slat)
	lenPass := len(sha256Pass)

	md5Bytes := make([]byte, lenSlat+lenPass)

	copy(md5Bytes[:lenPass], md5Slat[:])
	copy(md5Bytes[lenPass:], sha256Pass[:])

	encrypt := md5.Sum(md5Bytes)

	return fmt.Sprintf("%x", encrypt)
}

// StrToMd5 将 []byte 加密成 md5
func StrToMd5(b string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(b)))
}

// StrToSha256 将 []byte 加密成 md5
func StrToSha256(b string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(b)))
}
