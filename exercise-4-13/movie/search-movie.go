package movie

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func SearchMovie(keySearch string) (*Movie, error) {
	stringFields := strings.Fields(keySearch)
	stringUrl := strings.Join(stringFields, "+")
	url := BASE_URL + "?t=" + stringUrl + "&apikey=" + API_KEY
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result Movie
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
func GetImage(url string) error {
	fmt.Println(url)
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
		return e
	}
	defer response.Body.Close()

	//open a file for writing
	imgCode := GetMD5Hash(url)
	file, err := os.Create("D:/images/poster_" + imgCode + ".jpg")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Success!")
	return nil
}
