package xkcd

const BASE_URL = "https://xkcd.com/"
const JSON_URL = "/info.0.json"

type Comic struct {
	Num        int    `json: "num"`
	Link       string `json: "link"`
	Month      string `json: "month"`
	Year       string `json: "year"`
	News       string `json: news`
	Title      string `json: "title"`
	SafeTitle  string `json: "safe_title"`
	TranScript string `json: "transcript"`
	Alt        string `json: "alt"`
	Img        string `json: "img"`
	Day        string `json: "day"`
}
type Comics struct {
	Items []*Comic
}
