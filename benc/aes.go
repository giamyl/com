// Package benc 是用来处理加密的包
package benc

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
)

/*
	参考： https://github.com/polaris1119/myblog_article_code/blob/master/aes/aes.go
*/

// AesStruct 为 Aes 操作定义一个结构
type AesStruct struct {
	originData []byte // 加密或解密的原始数据
	key        []byte // 加密或解密的字符
	iv         []byte // 加密或解密的 iv
	resultData []byte // 加密或解密后的结果
}

// NewAesEncrypt 初始化 aes 的结构
func NewAesEncrypt(origData, key []byte) (*AesStruct, error) {
	as := new(AesStruct)

	if err := as.SetOriginData(origData); err != nil {
		return nil, err
	}

	if err := as.SetKey(key); err != nil {
		return nil, err
	}

	return as, nil
}

// NewAesEncryptFromBase64 从 base64 字符串实例化
func NewAesEncryptFromBase64(base64Str string, key []byte) (*AesStruct, error) {
	originData, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}

	as, err := NewAesEncrypt(originData, key)
	if err != nil {
		return nil, err
	}

	return as, nil
}

// SetOriginData 设置需要加密的数据
func (as *AesStruct) SetOriginData(originData []byte) error {
	if originData == nil {
		return errors.New("origin data can not empty")
	}

	as.originData = originData

	return nil
}

// SetKey 设置加密的 key
// 这里不管传入的 key 是什么， 只要不为空即可， 转换成 md5 去使用
func (as *AesStruct) SetKey(key []byte) error {
	if key == nil {
		return errors.New("key can not empty")
	}

	as.key = []byte(fmt.Sprintf("%x", md5.Sum(key)))

	return nil
}

// SetIV 设置 IV， 需要注意的， IV 必须是 16 位
func (as *AesStruct) SetIV(iv []byte) error {
	if iv == nil {
		return errors.New("key can not empty")
	}

	if len(iv) != 16 {
		fmt.Println("iv length : ", len(iv))
		fmt.Println(string(iv))
		return errors.New("iv length must be 16")
	}

	as.iv = iv

	return nil
}

// CBCEncrypt CBC 加密
func (as *AesStruct) CBCEncrypt() error {
	block, err := aes.NewCipher(as.key)
	// AesEncryptResult 加密后的结果

	if err != nil {
		return err
	}

	blockSize := block.BlockSize()
	origData := PKCS5Padding(as.originData, blockSize)

	// 如果设置了 iv， 使用设置了的， 如果没有， 就从 key 中获取前 16 位
	// 这里是通过 blockSize 中计算出来的, 因为 key 使用的是 32 位的 key
	// blockSize 是 16 位， 取出前面部分正好 16 位
	if as.iv == nil {
		as.iv = as.key[:blockSize]
	}

	blockMode := cipher.NewCBCEncrypter(block, as.iv)

	// 结果
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)

	as.resultData = crypted

	return nil
}

// CBCDecrypt CBC 解密
func (as *AesStruct) CBCDecrypt() error {
	block, err := aes.NewCipher(as.key)
	if err != nil {
		return err
	}

	blockSize := block.BlockSize()
	if as.iv == nil {
		as.iv = as.key[:blockSize]
	}

	blockMode := cipher.NewCBCDecrypter(block, as.iv)
	origData := make([]byte, len(as.originData))

	blockMode.CryptBlocks(origData, as.originData)
	origData = PKCS5UnPadding(origData)

	as.resultData = origData

	return nil
}

// GetCryptedData 获取加密或解密后的结果
func (as *AesStruct) GetCryptedData() []byte {
	return as.resultData
}

// GetCryptedDataBase64 获取 base64 格式的加密结果
func (as *AesStruct) GetCryptedDataBase64() string {
	return base64.StdEncoding.EncodeToString(as.GetCryptedData())
}

// GetKey 获取 key 的内容
func (as *AesStruct) GetKey() []byte {
	return as.key
}

// GetIV 获取 IV 的值
func (as *AesStruct) GetIV() []byte {
	return as.iv
}

// AesEncToBase64 将数据 aes 加密， 并转成 base64 格式
func AesEncToBase64(originData, key []byte) (string, error) {
	nae, err := NewAesEncrypt(originData, key)
	if err != nil {
		return "", err
	}

	err = nae.CBCEncrypt()
	if err != nil {
		return "", err
	}

	return nae.GetCryptedDataBase64(), nil
}

// AesDecFromBase64 数据解密
func AesDecFromBase64(base64Str string, key []byte) ([]byte, error) {
	as, err := NewAesEncryptFromBase64(base64Str, key)
	if err != nil {
		return nil, err
	}

	err = as.CBCDecrypt()
	if err != nil {
		return nil, err
	}

	return as.resultData, nil
}

// PKCS5Padding 填充
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5UnPadding 填充
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
