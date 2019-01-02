package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func GetHashInByte(stringHash []byte, flag int) {
	switch flag {
	case 0:
		sha384Bytes := sha512.Sum384(stringHash)
		fmt.Printf("SHA 384 : %x", sha384Bytes)
	case 1:
		sha512Bytes := sha512.Sum512(stringHash)
		fmt.Printf("SHA 512 : %x", sha512Bytes)
	default:
		sha256Bytes := sha256.Sum256(stringHash)
		fmt.Printf("SHA 256 : %x", sha256Bytes)
	}

}
func main() {
	s := []byte("Hello world")
	GetHashInByte(s, 1)
}
