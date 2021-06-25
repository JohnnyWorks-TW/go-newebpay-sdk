package newebpay

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func TradeInfoEncrypt(src string, key string, iv string) string {
	encryptedData := AESEncrypt(src, []byte(key), []byte(iv))
	return hex.EncodeToString(encryptedData)
}

func TradeInfoDecrypt(encrypted string, key string, iv string) string {
	encryptedData, _ := hex.DecodeString(encrypted)
	decrypted := AESDecrypt(encryptedData, []byte(key), []byte(iv))
	return string(decrypted)
}

func TradeInfoHash(plainText string, key string, iv string) string {
	encrypted := TradeInfoEncrypt(plainText, key, iv)
	hashStr := "HashKey=" + key + "&" + encrypted + "&HashIV=" + iv
	h := sha256.New()
	h.Write([]byte(hashStr))
	return strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
}

func AESEncrypt(src string, key []byte, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, iv)
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted
}

func AESDecrypt(crypt []byte, key []byte, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if len(crypt) == 0 {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return PKCS5Trimming(decrypted)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
