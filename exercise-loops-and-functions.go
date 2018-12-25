package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	s := float64(1)
	for i := 0; i < 11 ; i++ {
		z = z - (z*z - x)/(2*z)
		s = z
	}
	return s
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}