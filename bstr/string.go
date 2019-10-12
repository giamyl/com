package bstr

import (
	"bytes"
	"strings"
)

// ShowLess 字符串隐藏部分内容， 只显示较少的内容
func ShowLess(str string, opts ...int) string {
	start := 3  // 前面开始的位置
	end := 4    // 结束的位置
	method := 0 // 隐藏方式， 0 隐藏前后， 1 隐藏中间

	if str == "" {
		return str
	}

	if len(opts) == 3 {
		start, end, method = opts[0], opts[1], opts[2]
	} else if len(opts) == 2 {
		start, end = opts[0], opts[1]
	}

	if len(str) < start+end {
		return str
	}

	strBytes := []byte(str)

	buf := bytes.Buffer{}
	if method == 0 {
		buf.Write(strBytes[:start])
		buf.WriteString(strings.Repeat("*", len(strBytes)-start-end))
		buf.Write(strBytes[len(str)-end:])
	} else {
		buf.WriteString(strings.Repeat("*", start))
		buf.Write(strBytes[start : len(strBytes)-end])
		buf.WriteString(strings.Repeat("*", end))
	}

	return buf.String()
}
