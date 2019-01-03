package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

//	Function count the frequency of each word in an input file.
func WordFreq() {
	dat, err := ioutil.ReadFile("input.txt")
	checkError(err)
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// count the frequency of each word.
	counts := make(map[string]int)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
		os.Exit(1)
	}
	fmt.Printf("word\tcount\n")
	for word, count := range counts {
		fmt.Printf("%q\t%d\n", word, count)
	}
}

func main() {
	WordFreq()
}
