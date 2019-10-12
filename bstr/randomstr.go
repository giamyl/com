// Package bstr 包是用来处理字符串相关的操作
package bstr

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	defaultLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	defaultNumbers = "0123456789"
)

// LetterStruct Letter 结构
type LetterStruct struct {
	letters string
	length  uint
}

// ResultLetter 返回的数据类型
type ResultLetter string

// ToString 返回值转换成 string 类型
func (rl ResultLetter) ToString() string {
	return string(rl)
}

// ToInt 返回值转换成 int 类型
func (rl ResultLetter) ToInt() (int, error) {
	return strconv.Atoi(string(rl))
}

// ToUpper 返回值转换成 string 类型， 并转转成大写
func (rl ResultLetter) ToUpper() string {
	return strings.ToUpper(string(rl))
}

// ToLower 返回值转换成 string 类型， 并转转成小写
func (rl ResultLetter) ToLower() string {
	return strings.ToLower(string(rl))
}

// NewRandLetter 初始化 rLetter
func NewRandLetter() *LetterStruct {
	return &LetterStruct{
		letters: defaultLetters + defaultNumbers,
	}
}

// SetLetters 设置生成随机数使用的字母
func (ls *LetterStruct) SetLetters(letters string) *LetterStruct {
	ls.letters = letters

	return ls
}

// OnlyLetters 设置生成随机数使用纯字母
func (ls *LetterStruct) OnlyLetters() *LetterStruct {
	ls.letters = defaultLetters

	return ls
}

// OnlyNumber 设置生成随机数使用纯数字
func (ls *LetterStruct) OnlyNumber() *LetterStruct {
	ls.letters = defaultNumbers

	return ls
}

// SetLen 设置生成随机字符串的长度
func (ls *LetterStruct) SetLen(length uint) *LetterStruct {
	ls.length = length

	return ls
}

// RandProcess 执行生成随机数的操作
func (ls *LetterStruct) RandProcess() ResultLetter {
	if ls.length == 0 || ls.letters == "" {
		return ""
	}
	n := int(ls.length)
	letters := ls.letters

	var src = rand.NewSource(time.Now().UnixNano())

	var letterIdxBits uint = 6
	var letterIdxMask int64 = 1<<letterIdxBits - 1
	var letterIdxMax = 63 / letterIdxBits

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return ResultLetter(b)
}

// ===== 下面是使用默认参数的快速函数 =====

// RandLetter 获取纯字符的随机字符串
func RandLetter(length uint) string {
	return NewRandLetter().SetLen(length).RandProcess().ToString()
}

// RandNumber 获取纯数字的随机字符串
func RandNumber(length uint) string {
	return NewRandLetter().OnlyNumber().SetLen(length).RandProcess().ToString()
}

// RandInt 获取纯数字的随机数
func RandInt(length uint) (n int) {
	for i := 0; i < 10; i++ {
		n, _ = NewRandLetter().OnlyNumber().SetLen(length).RandProcess().ToInt()
		fmt.Println("==========", i)
		if n > 10*int(length) {
			return
		}
	}

	return
}
