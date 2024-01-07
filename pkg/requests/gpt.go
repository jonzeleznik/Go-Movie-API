package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Choice struct {
	Finish_reason string `json:"finish_reason"`
	Index         int    `json:"index"`
	Message       struct {
		Content string `json:"content"`
		Role    string `json:"role"`
	}
}

type GptResponse struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"Choices"`
}

type GPTRequests struct {
	AccessToken string
	ApiUrl      string
}

func NewGPTRequests() *GPTRequests {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	return &GPTRequests{
		AccessToken: os.Getenv("GPT_ACCESS_TOKEN"),
		ApiUrl:      os.Getenv("GPT_API_URL"),
	}
}

func (t *GPTRequests) GetChatGPT() (GptResponse, error) {
	client := &http.Client{}
	var data = strings.NewReader(`
	{
		"model": "gpt-3.5-turbo",
		"messages": [
			{
				"role": "assistant",
				"content": "You translate english articles to slovenian. Your answer is in Markdown (.md file)."
			},
			{
				"role": "user",
				"content": "article title: 2023 Box Office: Global Revenue Clears Estimated $33.9B in 31 Percent Gain Over Prior Year, article introduction: Leading London-based analytics firm Gower Street has issued its yearly estimate for total movie ticket sales across the globe last year. article content: It was a Barbenheimer and Super Mario world in 2023. Barbie, The Super Mario Bros. Movie and Oppenheimer led the pack of films propelling global box office revenue to $33.9 billion, according to estimates released Thursday by leading UK-based analytics firm Gower Street. That’s a gain of 31 percent over 2022, but was 15 percent behind the average of the three years before the pandemic (2015-2019). Translated, movie ticket sales have yet to match the levels enjoyed before the COVID-19 crisis struck, whether on a global scale, domestically or internationally. Leading U.S.-based firm Comscore said earlier this week that domestic box office revenue hit $9.05 billion in 2023, a 21 percent uptick over 2022 and a post-pandemic best. Gower Street shows international box office revenue making a bigger comeback at $24.8 billion, a year-over-year gain of 35 percent. Excluding China, foreign revenue was $17.1 billion in 2023, a 20 percent year-over-year increase, according to Gower. China exceeded expectations in delivering $7.71 billion — an enormous 83 gain over 2023 and only 6 percent behind the pre-pandemic 2017-2019 average, according to Gower. China’s box office received a powerful boost in recent days from Shining One Thing. Looking ahead to 2024, Gower is already predicting that box office revenue could trail 2023 levels by as much as seven percent because of strike-related delays and massive disruptions to the release calendar. In 2022, worldwide box office revenue was $25.9 billion, versus $21.3 billion in 2021. That pales in comparison to pre-pandemic global revenue, including $42.3 billion in 2019. Two movies crossed $1 billion globally in 2023; Warner Bros.’ Barbie ($1.44 billion) and Universal/Illumination’s The Super Mario Bros. Movie ($1.36 billion), both of which shattered numerous records. Universal’s Oppenheimer followed at No. 3 on the global chart with $952 million, a record for a biographical drama. Disney and Marvel’s Guardians of the Galaxy Vol. 3 placed fourth with $845.6 million — almost as much as the sequel — while Universal’s Fast X rounded out the top five with $704.8 million, fueled by overseas. The Asia Pacific region accounted for 39 percent of global box office revenue in 2023. China contributed the vast majority, or 32 percent, according to Gower. North America contributed 27 percent, as did Europe, Africa and the Middle East (combined). Latin America made up the remaining 8 percent."
			}
		]
	}
	`)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", data)
	if err != nil {
		return GptResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+t.AccessToken)

	resp, err := client.Do(req)
	if err != nil {
		return GptResponse{}, err
	}

	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return GptResponse{}, err
	}

	var r GptResponse
	fmt.Println(string(bodyText))
	err = json.Unmarshal([]byte(bodyText), &r)
	if err != nil {
		return GptResponse{}, err
	}

	return r, nil
}
