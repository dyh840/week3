package getDecryptedPaper

import (
	"encoding/base64"
)

// XOR 解密
func xorEncryptDecrypt(input, key string) string {
	keyLen := len(key)
	output := make([]byte, len(input))

	for i := range input {
		output[i] = input[i] ^ key[i%keyLen]
	}

	return string(output)
}

// 解密函数
func GetDecryptedPaper(encodedPaper, key string) string {
	data, _ := base64.StdEncoding.DecodeString(encodedPaper)
	return xorEncryptDecrypt(string(data), key)
}
