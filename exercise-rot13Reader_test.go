package main

import (
	"testing"
	"strings"
	"bytes"
)
func TestRot13Reader(t *testing.T)  {
	dataTest := []string{
		"Lbh penpxrq gur pbqr!",
		"Jul qvq gur puvpxra pebff gur ebnq?",
	}
	dataCorrect := []string{
		"You cracked the code!",
		"Why did the chicken cross the road?",
	}
	for i := 0 ; i < len(dataTest) ; i++ {
		stringTest := strings.NewReader(dataTest[i])
		resultTest := rot13Reader{stringTest}
		buf := new(bytes.Buffer)
		buf.ReadFrom(resultTest)
		stringResultTest := buf.String()
		if stringResultTest != dataCorrect[i] {
			t.Errorf("result test: %+v , result correct : %s", resultTest, dataCorrect[i])
		}
	}
}