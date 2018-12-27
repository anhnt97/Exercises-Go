package main

import (
	"testing"
)
func TestFibonacciResult(t *testing.T) {
	dataTest := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	resultTest := Fibonacci()
	for i:= 0 ; i < 10 ; i++ {
		value := resultTest()
		if value != dataTest[i] {
			t.Errorf("result test : %d , result correct : %d", value, dataTest[i])
		}
	}
}