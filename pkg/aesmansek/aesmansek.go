package aesmansek

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"errors"
)

// Credential ...
type Credential struct {
	Key string
}

// AESEncrypt ...
func (cred *Credential) AESEncrypt(src string) (string, error) {
	if src == "" {
		return "", errors.New("empty_parameter")
	}

	hashedKey := sha1.Sum([]byte(cred.Key))
	key := hashedKey[:16]

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ecb := cipher.NewCBCEncrypter(block, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

// AESDecrypt ...
func (cred *Credential) AESDecrypt(cryptString string) (string, error) {
	hashedKey := sha1.Sum([]byte(cred.Key))
	key := hashedKey[:16]

	crypt, _ := base64.StdEncoding.DecodeString(cryptString)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(crypt) == 0 {
		return "", errors.New("empty_parameter")
	}
	ecb := cipher.NewCBCDecrypter(block, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return string(PKCS5Trimming(decrypted)), nil
}

// PKCS5Padding ...
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5Trimming ...
func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
