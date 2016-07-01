package umrest

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strings"
)

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var baseUrl = "https://rest.wsq.umeng.com/0/get_access_token?ak=%s&access_token=%s"

func BuildUrl(aes_key string, data string) string {

	byte_aes_key := make([]byte, len(aes_key))
	strings.NewReader(aes_key).Read(byte_aes_key)

	byte_data := make([]byte, len(data))
	strings.NewReader(data).Read(byte_data)

	access_token := encrypt(byte_aes_key[:16], byte_data)

	fmt.Println("--access_token--", access_token)

	return fmt.Sprintf(baseUrl, aes_key, access_token)
}

func encrypt(key []byte, data []byte) string {

	fmt.Println(string(key))

	block, err := aes.NewCipher(key)

	if err != nil {
		return ""
	}

	blockSize := block.BlockSize()

	data = PKCS7Padding(data, blockSize)

	blockModel := cipher.NewCBCEncrypter(block, key[:blockSize])

	encrypted := make([]byte, len(data))

	blockModel.CryptBlocks(encrypted, data)

	fmt.Println(string(encrypted))

	b64 := base64.NewEncoding(base64Table)

	fmt.Println("--Decryption--", Decryption(key, b64.EncodeToString(encrypted)))

	return b64.EncodeToString(encrypted)

}

func Decryption(key []byte, data string) string {

	b64e := base64.NewEncoding(base64Table)

	b_data, err := b64e.DecodeString(data)

	block, err := aes.NewCipher(key)

	if err != nil {
		return ""
	}

	blockSize := block.BlockSize()

	//	b_data = PKCS7UnPadding(b_data)

	blockModel := cipher.NewCBCDecrypter(block, key[:blockSize])

	origData := make([]byte, len(b_data))

	blockModel.CryptBlocks(origData, b_data)

	origData = PKCS7UnPadding(origData)
	return string(origData)
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	fmt.Println("--padding--", padding)

	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])

	fmt.Println("--unpadding--", unpadding)

	return origData[:(length - unpadding)]
}

const (
	aes_key = "273d7e70c2d115e62e0e45656ff82b39"
)

func main() {

	fmt.Println("--um_rest_demo--")

	data := `{"user_info":{"name":"lmwww","gender":1},"source_uid":"123491239324228","source":"qq"}`

	BuildUrl(aes_key, data)
}
