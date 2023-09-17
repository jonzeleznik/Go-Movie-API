package models

type Response struct {
	Page          int
	Results       []results
	Total_pages   int
	Total_results int
}

type results struct {
	Adult             bool
	Backdrop_path     string
	Genre_ids         []string
	Id                int
	Original_language string
	Original_title    string
	Overview          string
	Popularity        float64
	Poster_path       string
	Release_date      string
	Title             string
	Video             bool
	Vote_average      float32
	Vote_count        int
}
