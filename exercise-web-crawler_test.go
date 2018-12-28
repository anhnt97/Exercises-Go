package main

import (
	"strings"
	"testing"
)

func TestCraw(t *testing.T) {
	dataTest := []string{"https://golang.org/", "https://golang.org/pkg/", "https://golang.org/pkg/fmt/", "https://golang.org/pkg/os/"}
	dataResult := map[string]string{
		"https://golang.org/":         "The Go Programming Language",
		"https://golang.org/pkg/":     "Packages",
		"https://golang.org/pkg/fmt/": "Package fmt",
		"https://golang.org/pkg/os/":  "Package os",
	}
	for _, data := range dataTest {
		resultTest := make(chan string)
		go Crawl(data, 4, fetcher, resultTest)
		for s := range resultTest {
			stringResult := s
			stringFields := strings.Fields(stringResult)
			if stringFields[0] == "found" {
				url := stringFields[1]
				name := stringFields[2]
				if name != dataResult[url] {
					t.Errorf("Result test : %v , Result correct: %v", name, dataResult[url])
				}
			}
		}
	}
}
