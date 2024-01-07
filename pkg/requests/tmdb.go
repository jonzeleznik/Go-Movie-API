package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type genreTmdb struct {
	Id   int    `json:"Id"`
	Name string `json:"name"`
}

type MoviesTmdb struct {
	Adult             bool        `json:"adult"`
	Backdrop_path     string      `json:"backdrop_path"`
	Budget            int         `json:"budget"`
	Genres            []genreTmdb `json:"genres"`
	Homepage          string      `json:"homepage"`
	Id                int         `json:"id"`
	Imdb_id           string      `json:"imdb_id"`
	Original_language string      `json:"original_language"`
	Original_title    string      `json:"original_title"`
	Overview          string      `json:"overview"`
	Popularity        float64     `json:"popularity"`
	Poster_path       string      `json:"poster_path"`
	Release_date      string      `json:"release_date"`
	Runtime           int         `json:"runtime"`
	Title             string      `json:"title"`
	Video             bool        `json:"video"`
	Vote_average      float32     `json:"vote_average"`
	Vote_count        int         `json:"vote_count"`
}

type tmdbResponse struct {
	Page          int
	Results       []MoviesTmdb
	Total_pages   int
	Total_results int
}

type TmdbRequests struct {
	AccessToken string
	ApiUrl      string
}

func NewTmdbRequests() *TmdbRequests {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	return &TmdbRequests{
		AccessToken: os.Getenv("TMDB_ACCESS_TOKEN"),
		ApiUrl:      os.Getenv("TMDB_API_URL"),
	}
}

func (t *TmdbRequests) GetTmdbMovieTitle(title string) (tmdbResponse, error) {

	endpoint := t.ApiUrl + "/search/movie?query=" + url.PathEscape(title) + "&include_adult=false&include_video=false&language=en-US&page=1&sort_by=popularity.desc"

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return tmdbResponse{}, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", ("Bearer " + t.AccessToken))

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return tmdbResponse{}, err
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return tmdbResponse{}, err
	}

	raw := string(body)

	var resp tmdbResponse
	err = json.Unmarshal([]byte(raw), &resp)
	if err != nil {

		return tmdbResponse{}, err
	}
	return resp, nil

}
func (t *TmdbRequests) GetTmdbMovieId(id string) (MoviesTmdb, error) {

	endpoint := t.ApiUrl + "/movie/" + url.PathEscape(id) + "?include_adult=false&language=en-US&page=1"

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return MoviesTmdb{}, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", ("Bearer " + t.AccessToken))

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return MoviesTmdb{}, err
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return MoviesTmdb{}, err
	}

	raw := string(body)

	var resp MoviesTmdb
	err = json.Unmarshal([]byte(raw), &resp)
	if err != nil {

		return MoviesTmdb{}, err
	}
	return resp, nil
}
