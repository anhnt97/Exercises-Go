package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"xkcd"
)

func main() {
	xkcd.GetAllComic()
	fmt.Println("Enter the key search: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	keyword := scanner.Text()
	resultSearch, _ := xkcd.SearchComic(keyword)
	fmt.Println("Result search:")
	for _, comic := range resultSearch.Items {
		url := xkcd.BASE_URL + strconv.Itoa(comic.Num) + xkcd.JSON_URL
		fmt.Printf("Url : %s , Transcript: %s\n", url, comic.TranScript)
	}
}
