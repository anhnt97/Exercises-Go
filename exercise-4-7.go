package main

import "fmt"

func reverse(s []byte) {
	if len(s) == 0 {
		return
	} else {
		for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}
}
func main() {
	stringS := []byte("hello golang")
	reverse(stringS)
	fmt.Printf("%s", stringS)
}
