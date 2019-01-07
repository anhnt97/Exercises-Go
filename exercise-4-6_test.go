package main

import "testing"

func TestFormatSingleSpace(t *testing.T) {
	dataTest := []string{"hello\n \n \t       go \n \t!", "\n hello \t", "\t  golang\t\n", "123 \n\n 456\n\n"}
	dataCorrect := []string{"hello go !", "hello", "golang", "123 456"}
	for i, val := range dataTest {
		test := []byte(val)
		resultTest := string(FormatSingleSpace(test))
		if resultTest != dataCorrect[i] {
			t.Errorf("Value test : %s , result test : %s , result correct :  %s", val, resultTest, dataCorrect[i])
		}
	}
}
