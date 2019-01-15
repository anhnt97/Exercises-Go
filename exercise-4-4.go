package main

import (
	"fmt"
)

func rotate(s []int, position int) []int {
	s1 := s[:position]
	s2 := s[position:]
	newSlice := append(s2, s1...)
	return newSlice
}
func main() {
	s := []int{1, 2, 3, 4, 5}
	newSlice := rotate(s, 2)
	fmt.Println(newSlice)
}
