package main

import (
	"testing"
	"math/rand"
	"math"
)
func TestSqrtValues(t *testing.T) {
	table := [10]float64{}
	result := [10]float64{}
	for i := 0 ; i < 10 ; i++ {
		table[i] = rand.Float64()*float64(i-1)
		result[i] = math.Sqrt(table[i])
	}
	for index, test := range table {
		sqrtResultTest,_ := Sqrt(test)
		resultValue := result[index]
		if sqrtResultTest != result[index] && math.IsNaN(sqrtResultTest) != math.IsNaN(resultValue) {
			t.Errorf("Fail test : %f,  value test : %f , value result : %f", test, sqrtResultTest, resultValue)
		}
	}
}