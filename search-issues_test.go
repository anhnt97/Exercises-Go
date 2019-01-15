package github

import "testing"

func TestSearchIssues(t *testing.T) {
	dataTest := []string{
		"bat man",
		"spider man",
		"superman",
	}
	dataCorrect := []string{
		"ComicConnect June 2018",
		"Consistently Source Director/Writer Metadata for Series",
	}
	resultTest, _ := SearchIssues(dataTest)
	for i, val := range resultTest.Items {
		if val.Title != dataCorrect[i] {
			t.Errorf("Value test : %s , value correct : %s ", val.Title, dataCorrect[i])
		}
	}

}
