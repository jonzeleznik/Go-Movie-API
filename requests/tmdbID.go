package requests

import (
	"e/initializers"
	"e/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func init() {
	initializers.LoadEnvVariables()
}

func GetTmdbID(id int) models.Movies {

	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%d?language=en-US", id)

	req, err := http.NewRequest("GET", url, nil)
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

	var resp models.Movies
	json.Unmarshal([]byte(raw), &resp)

	return resp
}