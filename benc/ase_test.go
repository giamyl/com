package benc

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	data := []byte("我是加密前的数据")
	// 必须是 32 位的
	key := []byte("12345678901234567890123456789012")

	res, err := AesEncrypt(data, key)
	if err != nil {
		t.Fatalf("加密失败： %v\n", err)
	}

	base64Str := base64.StdEncoding.EncodeToString(res)
	fmt.Println(base64Str)

	unBase64, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		t.Fatal(err)
	}

	b, err := AesDecrypt(unBase64, key)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("解密后输出：", string(b))
}
