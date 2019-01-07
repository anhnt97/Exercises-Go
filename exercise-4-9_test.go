package main

import "testing"

func TestWordFreq(t *testing.T) {
	resultCorrect := map[string]int{
		"hello":   1,
		"golang":  1,
		"data":    2,
		"group":   1,
		"welcome": 2,
	}
	resultTest := WordFreq()
	for word, count := range resultTest {
		if countCorrect, ok := resultCorrect[word]; !ok || countCorrect != count {
			t.Errorf("Value test : [%s] : %d , ValueCorrect : [%s] : %d", word, count, word, countCorrect)
		}
	}
}
