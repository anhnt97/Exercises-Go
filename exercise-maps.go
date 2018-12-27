package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	if len(s) == 0 {
		fmt.Println("S is empty!")
		return nil
	}else {
		stringFields := strings.Fields(s)
		mapString := make(map[string]int)
		for i := 0; i < len(stringFields); i++ {
			mapString[stringFields[i]] = mapString[stringFields[i]] + 1
		}
		return mapString
	}
}

func main() {
	 wc.Test(WordCount)
}