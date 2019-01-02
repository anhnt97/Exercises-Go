package main

import (
	"crypto/sha256"
	"fmt"
)

func BitCount(x uint8) int {
	counter := 0
	for x != 0 {
		x &= x - 1
		counter++
	}
	return counter
}
func BitsDifference(h1 [32]uint8, h2 [32]uint8) int {
	counter := 0
	for i := range h1 {
		counter += BitCount(h1[i] ^ h2[i])
	}
	return counter
}
func main() {
	s1 := "hello world"
	s2 := "hello golang"
	h1 := sha256.Sum256([]byte(s1))
	h2 := sha256.Sum256([]byte(s2))
	fmt.Printf("H1 : %x\n", h1)
	fmt.Printf("H2 : %x\n", h2)
	fmt.Println(BitsDifference(h1, h2))
}
