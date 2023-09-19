package requests

import (
	"e/initializers"
	"e/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetTmdbMovieTitle(title string) models.Response {

	endpoint := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?query=%s&include_adult=false&language=en-US&page=1", url.PathEscape(title))

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Printf("An error occured: %v", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", ("Bearer " + initializers.EnvVars["ACCESS_TOKEN"]))

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("An error occured: %v", err)
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("An error occured: %v", err)
	}

	raw := string(body)

	var resp models.Response
	json.Unmarshal([]byte(raw), &resp)

	return resp
}
