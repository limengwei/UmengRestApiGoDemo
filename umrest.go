package umrest

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strings"
)

var baseUrl = "https://rest.wsq.umeng.com/api1?ak=%s&access_token=%s"

func BuildUrl(ak string, data string) string {

	b_ak := make([]byte, len(ak))
	strings.NewReader(ak).Read(b_ak)

	b_data := make([]byte, len(data))
	strings.NewReader(data).Read(b_data)

	access_token := encrypt(b_ak[:16], b_data)

	return fmt.Sprintf(baseUrl, ak, access_token)
}

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

func encrypt(key []byte, data []byte) string {

	fmt.Println(string(key))

	block, err := aes.NewCipher(key)

	if err != nil {
		return ""
	}

	data = PKCS7Padding(data, block.BlockSize())

	blockModel := cipher.NewCBCDecrypter(block, key)

	encrypted := make([]byte, len(data))

	blockModel.CryptBlocks(encrypted, data)

	fmt.Println(string(encrypted))

	return base64.NewEncoding(base64Table).EncodeToString(encrypted)

}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
