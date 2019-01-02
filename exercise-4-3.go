package main

import "fmt"

func reverse(s *[9]int) {
	lenArray := len(*s)
	for i := 0; i < lenArray/2; i++ {
		temp := (*s)[i]
		(*s)[i] = (*s)[lenArray-1-i]
		(*s)[lenArray-1-i] = temp
	}

}
func main() {
	array := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(&array)
	fmt.Println(array)
}
