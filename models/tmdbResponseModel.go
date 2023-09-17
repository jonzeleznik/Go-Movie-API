package models

type Response struct {
	Page          int
	Results       []Movies
	Total_pages   int
	Total_results int
}
