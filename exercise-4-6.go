package main

import (
	"fmt"
	"strings"
)

func remove(slice []byte, i int) []byte {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
func FormatSingleSpace(slices []byte) []byte {
	if len(slices) == 0 {
		return slices
	} else {
		stringFields := strings.Fields(string(slices))
		result := []byte(strings.Join(stringFields, " "))
		return result
	}
}
func main() {
	test := []byte("hello\n \n \t       world \n \t!")
	fmt.Printf("%s", FormatSingleSpace(test))
}
