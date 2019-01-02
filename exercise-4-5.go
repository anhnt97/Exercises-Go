package main

import (
	"fmt"
)

func eliminateAdjacentDuplicates(strings []string) []string {
	if len(strings) == 0 {
		return strings
	} else {
		result := []string{}
		previous := strings[0]
		result = append(result, previous)
		for i, val := range strings {
			if strings[i] != previous {
				result = append(result, val)
				previous = strings[i]
			}
		}
		return result
	}
}
func remove(slice []int, i int) []int {
	// slice[i] = slice[len(slice)-1]
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	// s := []int{5, 6, 7, 8, 9}
	// result := remove(s, 2)
	s := []string{"a", "a", "a", "b", "c", "c", "d", "a"}
	result := eliminateAdjacentDuplicates(s)
	fmt.Println(result)
}
