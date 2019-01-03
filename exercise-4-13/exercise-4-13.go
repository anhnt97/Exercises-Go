package main

import (
	"bufio"
	"fmt"
	"movie"
	"os"
)

func main() {
	fmt.Println("Enter the name movie : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	keyword := scanner.Text()
	resultSearch, _ := movie.SearchMovie(keyword)
	if len(resultSearch.Poster) > 0 {
		movie.GetImage(resultSearch.Poster)
	}
}
