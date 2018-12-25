package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	stringFields := strings.Fields(s)
	lenStringFields := len(stringFields)
	mapString := make(map[string]int)
	for i := 0; i < lenStringFields; i++ {
		(mapString[stringFields[i]])++
	}
	return mapString
}

func main() {
	wc.Test(WordCount)
}