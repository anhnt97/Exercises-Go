package main

import (
	"fmt"
	"math"
	"errors"
)

func Sqrt(x float64) (float64,error) {
	z := float64(1)
	s := float64(0)
	if x < 0 {
		return float64(math.NaN()),
			   errors.New("Value must be a positive number")
	}else if x == 0 {
		return 0,nil
	}else {
		for i := 0; i < 11 ; i++ {
			z = z - (z*z - x)/(2*z)
			s = z
		}
		return s,nil
	}
}

func main() {
	fmt.Println(Sqrt(0.875428))
	fmt.Println(math.Sqrt(0.875428))
}