// Package bjson 序列化使用的包
package bjson

import (
	"fmt"
)

// JSONFloat64One 将 float64 格式化成保留小数点后 1 位的字符串
type JSONFloat64One float64

// MarshalJSON JSONFloat64Two 序列化方法
func (j JSONFloat64One) MarshalJSON() ([]byte, error) {
	// 保留两位再截取最后一位，是为了方式 printf 四舍五入
	s := fmt.Sprintf("%.2f", j)
	b := []byte(s)

	return b[:len(b)-1], nil
}

// JSONFloat64Two 将 float64 格式化成保留小数点后 2 位的字符串
type JSONFloat64Two float64

// MarshalJSON JSONFloat64Two 序列化方法
func (j JSONFloat64Two) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%.3f", j)
	b := []byte(s)

	return b[:len(b)-1], nil
}

// JSONFloat64Six 将 float64 格式化成保留小数点后 2 位的字符串
type JSONFloat64Six float64

// MarshalJSON JSONFloat64Six 序列化方法
func (j JSONFloat64Six) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%.7f", j)
	b := []byte(s)

	return b[:len(b)-1], nil
}
