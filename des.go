package main

import (
	"bytes"
	"crypto/des"
	"fmt"
)

// func main() {

// 	plainText := []byte("abcdefgafsagasgasdgsdhxz")
// 	key := []byte("12345678")

// 	cipherText, err := Encrypt(plainText, key)

// 	if err == nil {
// 		fmt.Println("plain text is :", string(plainText))
// 		fmt.Println("plain text is :", string(cipherText))
// 	}

// 	decryptedText, err := Decrypt(cipherText, key)
// 	if err == nil {
// 		fmt.Println("cipher text is :", string(cipherText))
// 		fmt.Println("decrypted text is :", string(decryptedText))
// 	}

// }

func Encrypt(plainText, key []byte) ([]byte, error) {

	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	blockSize := block.BlockSize()

	if len(plainText)%blockSize != 0 {
		plainText = AddSpaces(plainText, blockSize)
	}

	cipherText := make([]byte, len(plainText))
	dst := cipherText

	for len(plainText) > 0 {
		block.Encrypt(dst, plainText[:blockSize])
		plainText = plainText[blockSize:]
		dst = dst[blockSize:]
	}

	return cipherText, nil
}

func Decrypt(cipherText, key []byte) ([]byte, error) {

	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	blockSize := block.BlockSize()

	if len(cipherText)%blockSize != 0 {
		cipherText = RemoveSpaces(cipherText, blockSize)
	}

	decryptedText := make([]byte, len(cipherText))
	dst := decryptedText

	for len(cipherText) > 0 {
		block.Decrypt(dst, cipherText[:blockSize])
		cipherText = cipherText[blockSize:]
		dst = dst[blockSize:]
	}

	return decryptedText, nil

}

func AddSpaces(text []byte, size int) []byte {

	spaceSize := size - len(text)%size
	spaces := bytes.Repeat([]byte{0}, spaceSize)
	return append(text, spaces...)
}

func RemoveSpaces(text []byte, size int) []byte {

	return bytes.TrimFunc(text,
		func(r rune) bool {
			return r == rune(0)
		})
}
