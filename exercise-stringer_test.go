package main

import (
	"testing"
	"fmt"
)
func TestString(t *testing.T)  {
	dataTest := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	resultCorrect := map[string]string{
		"loopback": "loopback: 127.0.0.1",
		"googleDNS": "googleDNS: 8.8.8.8",
	}
	for name, ip := range dataTest {
		resultTest := fmt.Sprintf("%v: %v", name, ip)
		if resultTest != resultCorrect[name] {
			t.Errorf("Result test: %v , Result correct: %v", resultTest, resultCorrect[name])
		}
	}
}