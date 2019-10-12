package bstr

import (
	"fmt"
	"testing"
)

func TestCreateUUID(t *testing.T) {
	res := CreateUUID()
	fmt.Println(res)

	res = ManualUUID()
	fmt.Println(res)

	res = Md5UnixNanoString()
	fmt.Println(res)
}
