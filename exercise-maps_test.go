package main

import (
	"testing"
	"reflect"
)
func TestWordCount(t *testing.T) {
	stringTest := []string{"I am learning Go!", "The quick brown fox jumped over the lazy dog.", "I ate a donut. Then I ate another donut.", "A man a plan a canal panama."}
	mapResult := map[int]map[string]int{}
	mapResult[0] = map[string]int{"am":1, "learning":1, "Go!":1, "I":1}
	mapResult[1] = map[string]int{"brown":1, "fox":1, "jumped":1, "the":1, "dog.":1, "The":1, "quick":1, "over":1, "lazy":1}
	mapResult[2] = map[string]int{"I":2, "ate":2, "a":1, "donut.":2, "Then":1, "another":1}
	mapResult[3] = map[string]int{"a":2, "plan":1, "canal":1, "panama.":1, "A":1, "man":1}
	for i:= 0 ; i < len(stringTest) ; i++ {
		resultTest := WordCount(stringTest[i])
		if !reflect.DeepEqual(resultTest, mapResult[i]) {
			t.Errorf("Word count error!, result test: %+v , result correct: %+v", resultTest, mapResult[i])
		}
	}
}