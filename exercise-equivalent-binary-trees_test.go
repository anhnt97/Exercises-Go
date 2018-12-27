package main

import (
	"testing"

	"golang.org/x/tour/tree"
)

func TestWalkTree(t *testing.T) {
	chanelTest := make(chan int)
	go Walk(tree.New(1), chanelTest)
	resultCorrect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 10; i++ {
		valueChanel := <-chanelTest
		if valueChanel != resultCorrect[i] {
			t.Errorf("value chanel : %d , value correct : %d", valueChanel, resultCorrect[i])
		}
	}
}
func TestSame(t *testing.T) {
	treeTest1 := tree.New(1)
	treeTest2 := tree.New(2)
	resultSameCorrect := true
	resultNotSameCorrect := false
	resultTestSame := Same(treeTest1, treeTest1)
	resultTestNotSame := Same(treeTest1, treeTest2)
	if resultTestSame != resultSameCorrect {
		t.Errorf("result test : %v , correct : %v", resultTestSame, resultSameCorrect)
	}
	if resultTestNotSame != resultNotSameCorrect {
		t.Errorf("result test : %v , correct : %v", resultTestNotSame, resultNotSameCorrect)
	}
}
