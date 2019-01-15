package movie

import "testing"

func TestSearchMovie(t *testing.T) {
	dataTest := []string{
		"bat man",
		"spider man",
		"superman",
	}
	dataCorrect := []string{
		"https://m.media-amazon.com/images/M/MV5BNDNhYmFjZGQtNjE3ZC00N2VmLWI0MWItODkxMmE2MWNiM2M5XkEyXkFqcGdeQXVyNzU1OTYxNzU@._V1_SX300.jpg",
		"https://m.media-amazon.com/images/M/MV5BM2EwYzA2YjMtNDdhYi00OTI1LWE2ODUtOWViODk4YjRjNzVmXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
		"https://m.media-amazon.com/images/M/MV5BMzA0YWMwMTUtMTVhNC00NjRkLWE2ZTgtOWEzNjJhYzNiMTlkXkEyXkFqcGdeQXVyNjc1NTYyMjg@._V1_SX300.jpg",
	}
	for i, val := range dataTest {
		resultTest, err := SearchMovie(val)
		if err != nil {
			t.Errorf("Error : %s", err)
		}
		poster := resultTest.Poster
		if poster != dataCorrect[i] {
			t.Errorf("Keyword : %s , Result test : %s , Result correct : %s", val, resultTest, dataCorrect[i])
		}
	}
}
