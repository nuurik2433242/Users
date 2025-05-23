package encripter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encripter struct {
	Key string
}

func NewEncripter() *Encripter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("Не передaн параметр KEY для в переменой окружения")
	}
	return &Encripter{
		Key: key,
	}
}

func (enc *Encripter) Encripter(plainStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err)
	}
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
        panic(err.Error())
	}
	nonce := make([]byte,aesGSM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
      panic(err.Error())
	}
	return aesGSM.Seal(nonce, nonce, plainStr, nil)
}

func (enc *Encripter) Decryter(encryptedStr []byte) []byte {
    block,err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	aesGSM,err := cipher.NewGCM(block)
	if err != nil {
        panic(err.Error())
	}
	nonceSize := aesGSM.NonceSize()
	nonce, cipherText := encryptedStr[:nonceSize], encryptedStr[nonceSize:]
	plainText, err := aesGSM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return plainText
}