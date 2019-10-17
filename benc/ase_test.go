package benc

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	// 手动初始化
	data := []byte("我是加密前的数据， Hello World!!!!")
	key := []byte("123456")

	// 手动初始化
	nae, err := NewAesEncrypt(data, key)
	if err != nil {
		t.Fatal(err)
	}

	// 此时仍然可以更改
	// nae.SetOriginData([]byte("xxxxxxxxxxxxxxxxx"))
	// nae.SetKey([]byte("xxxxxxxxxxxxxxxxx"))
	// err = nae.SetIV([]byte("qwertyuiopasdfgh"))

	err = nae.CBCEncrypt()
	if err != nil {
		t.Fatal(err)
	}

	base64Str := nae.GetCryptedDataBase64()

	fmt.Println("加密前的数据： ", string(data))
	fmt.Println("加密=========================================================")
	fmt.Println("手动加密后的结果： : ", base64Str)
	fmt.Println("iv: ", string(nae.GetIV()))

	// 使用快捷函数
	base64Str, err = AesEncToBase64(data, key)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("快捷函数加密的结果： ", base64Str)
	fmt.Println("")

	// 将加密后的数据转换回来
	fmt.Println("解密=========================================================")
	res, err := AesDecFromBase64(base64Str, key)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("解密后的数据： ", string(res))

}
