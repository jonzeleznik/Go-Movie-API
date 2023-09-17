package models

type Movies struct {
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
