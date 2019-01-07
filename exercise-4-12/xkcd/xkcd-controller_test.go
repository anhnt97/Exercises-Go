package xkcd

import "testing"

func TestSearchComic(t *testing.T) {
	dataTest := []string{
		"A boy sits in a barrel which is floating in an ocean",
		"Two trees are growing on opposite sides of a sphere",
		"A sketch of an Island",
		"A sketch of a landscape with sun on the horizon",
		"A black number 70 sees a red package",
	}
	dataCorrect := []int{1, 2, 3, 4, 5}
	GetAllComic()
	for i, keyword := range dataTest {
		resultTest, error := SearchComic(keyword)
		if len(resultTest.Items) == 0 {
			t.Errorf("result search is empty, keyword : %s", keyword)
		} else if error == nil {
			if resultTest.Items[0].Num != dataCorrect[i] {
				t.Errorf("Search error , key search : %s ", keyword)
			}
		}
	}
}
