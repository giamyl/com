package bpath

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetRootPath 获取项目根目录
func GetRootPath(key string) string {
	if key == "" {
		key = "ROOT_PATH"
	}

	// 如果设置了， 就用设置的项目根目录
	if r := os.Getenv(key); r != "" {
		return r
	}

	// 如果没有设置，就是用当前可执行文件所在的目录

	// 获取当前正在执行的文件的完整路径
	file, _ := exec.LookPath(os.Args[0])

	// 获取文件的绝对路径
	path, _ := filepath.Abs(file)

	// 获取最后一个文件名所在的位置，如 /home/bro/blog ，
	// 这里可以获取到 blog 的 b 所在的位置，为了下面的 slice 截取字符串
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}
