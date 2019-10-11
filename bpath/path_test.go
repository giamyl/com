package bpath

import (
	"fmt"
	"testing"
)

func TestRootPath(t *testing.T) {
	// 测试使用命令1： go test -v -test.run=TestRootPath
	// 测试使用命令2： ROOT_PATH=$(pwd) go test -v -test.run=TestRootPath

	// 不传入环境变量的 key
	root := GetRootPath("")
	fmt.Println(root)

	// 传入和默认相同的 key
	root = GetRootPath("ROOT_PATH")
	fmt.Println(root)

	// 传入不存在的 key
	root = GetRootPath("NO_PATH")
	fmt.Println(root)
}
