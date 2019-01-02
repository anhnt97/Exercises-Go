package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func CountDigits(i int) (count int) {
	for i != 0 {
		i /= 10
		count = count + 1
	}
	return count
}

func main() {
	counts := make(map[rune]int) // counts of Unicode characters
	invalid := 0                 // count of invalid UTF-8 characters
	countLetter := 0             // count of unicode letters
	countDigit := 0              // count of unicode digits
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsNumber(r) {
			countDigit++
		}
		if unicode.IsLetter(r) {
			countLetter++
		}
		counts[r]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	if countLetter > 0 {
		fmt.Printf("\n %d unicode letters\n", countLetter)
	}
	if countDigit > 0 {
		fmt.Printf("\n %d unicode digits\n", countDigit)
	}
}
