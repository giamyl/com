package bstr

import (
	"fmt"
	"testing"
)

func TestRandStr(t *testing.T) {
	// 默认的字符
	str := NewRandLetter().SetLen(8).RandProcess()
	fmt.Printf("默认生成的随机数， 类型： %T, 值： %[1]v\n", str)
	fmt.Printf("string 类型的随机数， 类型： %T, 值： %[1]v\n", str.ToString())
	fmt.Printf("大写的随机数， 类型： %T, 值： %[1]v\n", str.ToUpper())
	fmt.Printf("小写随机数， 类型： %T, 值： %[1]v\n", str.ToLower())
	fmt.Printf("快速函数， 类型： %T, 值： %[1]v\n", RandLetter(8))
	fmt.Println("========================================================")

	// 纯字母
	str = NewRandLetter().SetLetters(defaultLetters).SetLen(8).RandProcess()
	fmt.Printf("默认生成的随机数， 类型： %T, 值： %[1]v\n", str.ToString())
	str = NewRandLetter().OnlyLetters().SetLen(8).RandProcess()
	fmt.Printf("默认生成的随机数， 类型： %T, 值： %[1]v\n", str.ToString())
	fmt.Printf("快速函数， 类型： %T, 值： %[1]v\n", RandNumber(8))
	fmt.Println("========================================================")

	// 纯数字
	str = NewRandLetter().SetLetters(defaultNumbers).SetLen(8).RandProcess()
	fmt.Printf("默认生成的随机数， 类型： %T, 值： %[1]v\n", str.ToString())
	str = NewRandLetter().OnlyNumber().SetLen(8).RandProcess()
	fmt.Printf("默认生成的随机数， 类型： %T, 值： %[1]v\n", str.ToString())
	fmt.Printf("快速函数， 类型： %T, 值： %[1]v\n", RandInt(8))

}
