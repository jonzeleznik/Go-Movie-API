package models

type Movies struct {
	Adult             bool    `json:"adult"`
	Backdrop_path     string  `json:"backdrop_path"`
	Budget            int     `json:"budget"`
	Genres            []Genre `json:"genres"`
	Homepage          string  `json:"homepage"`
	Id                int     `json:"id"`
	Imdb_id           string  `json:"imdb_id"`
	Original_language string  `json:"original_language"`
	Original_title    string  `json:"original_title"`
	Overview          string  `json:"overview"`
	Popularity        float64 `json:"popularity"`
	Poster_path       string  `json:"poster_path"`
	Release_date      string  `json:"release_date"`
	Runtime           int     `json:"runtime"`
	Title             string  `json:"title"`
	Video             bool    `json:"video"`
	Vote_average      float32 `json:"vote_average"`
	Vote_count        int     `json:"vote_count"`
}
