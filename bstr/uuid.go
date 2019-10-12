package bstr

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

// CreateUUID 创建 UUID
func CreateUUID() string {
	u, err := uuid.NewV4()
	if err == nil {
		return u.String()
	}

	return ManualUUID()
}

// ManualUUID 当 uuid 包返回错误时的救援方案， 只是一个简单实现， 正常不要用到
// 将当前 nano 格式的时间 sha256 加密， 然后在 md5 加密， 然组合成和 uuid 包相同的格式
func ManualUUID() string {
	s5 := Md5UnixNanoString()

	// 将结果拼装返回
	return strings.Join(
		[]string{s5[:8], s5[8:12], s5[12:16], s5[16:20], s5[20:]}, "-")
}

// Md5UnixNanoString 使用 unix nano 时间戳生成 md5 字符串
func Md5UnixNanoString() string {
	t := time.Now().UnixNano()

	s256 := fmt.Sprintf("%x", sha256.Sum256([]byte(strconv.FormatInt(t, 10))))
	s5 := fmt.Sprintf("%x", md5.Sum([]byte(s256)))

	return s5
}
