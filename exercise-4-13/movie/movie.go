package movie

const BASE_URL = "http://www.omdbapi.com/"
const API_KEY = "2fb879fd"

type Movie struct {
	Title    string
	Year     string
	Rated    string
	Released string
	Runtime  string
	Director string
	Writer   string
	Actors   string
	Plot     string
	Language string
	Country  string
	Awards   string
	Poster   string
}
