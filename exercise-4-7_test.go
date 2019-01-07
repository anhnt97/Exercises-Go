package main

import "testing"

func TestReverse(t *testing.T) {
	dataTest := []string{
		"hello golang",
		"123456",
		"data \n test",
		"data123\t",
		"  data  \t data",
	}
	dataCorrect := []string{
		"gnalog olleh",
		"654321",
		"tset \n atad",
		"\t321atad",
		"atad \t  atad  ",
	}
	for i, val := range dataTest {
		resultTest := []byte(val)
		reverse(resultTest)
		if string(resultTest) != dataCorrect[i] {
			t.Errorf("Value test : %s , result test : %s, result correct : %s", val, resultTest, dataCorrect[i])
		}
	}
}
