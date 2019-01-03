package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var comics Comics

func GetComic(url string) (*Comic, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Failed: %s", resp.Status)
	}
	result := Comic{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	return &result, nil
}
func GetAllComic() error {
	for i := 1; i < 100; i++ {
		url := BASE_URL + strconv.Itoa(i) + JSON_URL
		comic, _ := GetComic(url)
		comics.Items = append(comics.Items, comic)
	}
	return nil
}
func SearchComic(keyword string) (*Comics, error) {
	resultSearch := Comics{}
	for _, comic := range comics.Items {
		if strings.Contains(comic.TranScript, keyword) {
			resultSearch.Items = append(resultSearch.Items, comic)
		}
	}
	return &resultSearch, nil
}
