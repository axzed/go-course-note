package aes

import (
	"crypto/aes"
	"crypto/cipher"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func Encrypt(plainText string, keyText string) (cipherByte []byte, err error) {
	// 转换成字节数据, 方便加密
	plainByte := []byte(plainText)
	keyByte := []byte(keyText)

	// 创建加密算法aes
	c, err := aes.NewCipher(keyByte)
	if err != nil {
		return nil, err
	}

	//加密字符串
	// dst 返回密文
	// src 明文
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	cipherByte = make([]byte, len(plainByte))
	cfb.XORKeyStream(cipherByte, plainByte)

	return
}

func Decrypt(cipherByte []byte, keyText string) (plainText string, err error) {
	// 转换成字节数据, 方便加密
	keyByte := []byte(keyText)
	// 创建加密算法aes
	c, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plainByte := make([]byte, len(cipherByte))
	cfbdec.XORKeyStream(plainByte, cipherByte)
	plainText = string(plainByte)
	return
}

//加密
func AESCtrEncrypt(plainText []byte, key string) ([]byte, error) {
	//1. 创建cipher.Block接口
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	//2. 计数器模式, 初始化接收器
	// 具体过程参考: https://blog.csdn.net/weixin_42940826/article/details/83687007
	// iv := bytes.Repeat([]byte("1"), block.BlockSize())
	iv := []byte("abcd123123123122")
	stream := cipher.NewCTR(block, iv)

	//3. 采用CTR进行异或运算, 明文 --异或--> 秘文 --异或--> 明文
	dst := make([]byte, len(plainText))
	stream.XORKeyStream(dst, plainText)

	return dst, nil
}

//解密
func AESCtrDecrypt(encryptData []byte, key string) ([]byte, error) {
	return AESCtrEncrypt(encryptData, key)
}
