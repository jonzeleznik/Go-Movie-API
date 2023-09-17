package models

type genre struct {
	Id   int64
	Name string
}

type Movies struct {
	Adult                 bool
	Backdrop_path         string
	Belongs_to_collection []string
	Budget                int
	Genres                []genre
	Homepage              string
	Id                    int
	Imdb_id               string
	Original_language     string
	Original_title        string
	Overview              string
	Popularity            float64
	Poster_path           string
	Release_date          string
	Title                 string
	Video                 bool
	Vote_average          float32
	Vote_count            int
}
