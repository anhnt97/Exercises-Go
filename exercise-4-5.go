package main

import (
	"fmt"
)

func eliminateAdjacentDuplicates(strings []string) []string {
	if len(strings) == 0 {
		return strings
	} else {
		result := strings[:1]
		previous := strings[0]
		for i, val := range strings {
			if strings[i] != previous {
				result = append(result, val)
				previous = strings[i]
			}
		}
		return result
	}
}

func main() {
	s := []string{"a", "a", "a", "b", "c", "c", "d", "a"}
	result := eliminateAdjacentDuplicates(s)
	fmt.Println(result)
}
