package main

import (
	"bufio"
	"bytes"
	"crypto/des"
	"fmt"
	"net"
	"os"
)

var HASSTARTED bool
var key []byte

func main() {

	key = []byte("12345678")

	fmt.Println("THE CHATTER 1.0")
	//Listener
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error lunching the server", err.Error())
		return
	}
	HASSTARTED = false
	/* GLOBAL SERVER CHECK */
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting", err.Error())
			return
		}
		go doServerStuff(connection)
		doClientStuff(connection)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		messageDec, _ := Decrypt(buf, key)
		messageDec = messageDec[:isASCII(messageDec)-1]
		fmt.Printf("POINT TWO SAYS : %v", string(messageDec))
	}
}

func doClientStuff(connect net.Conn) {
	inputReader := bufio.NewReader(os.Stdin)
	message, _ := inputReader.ReadString('\n')
	messageEnc, _ := Encrypt([]byte(message), key)
	_, err2 := connect.Write(messageEnc)
	if err2 != nil {
		fmt.Println("Error sending", err2.Error())
		return
	}
	doClientStuff(connect)
}

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

func isASCII(s []byte) int {
	i := 0
	for _, c := range s {
		i++
		if c > 127 || c < 0 {
			return i
		}
	}
	return -1
}
