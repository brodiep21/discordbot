package responses

// provides nasa JSON api structure
type Nasa struct {
	URL         string `json:"hdurl"`
	Explanation string `json:"explanation"`
}

// provides weather JSON api structure
type Main struct {
	Temp float64 `json:"temp"`
	High float64 `json:"temp_max"`
	Low  float64 `json:"temp_min"`
}

// top level weather JSON api structure
type Weatherinfo struct {
	Main Main
}

// joke JSON api structure
type Joke struct {
	Joke string `json:"joke"`
}

// ChatGPT api request/response structure JSON
type ReqRespContent struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatGPT api structure request top level
type Requester struct {
	Model    string           `json:"model"`
	Messages []ReqRespContent `json:"messages"`
}

// top level ChatGPT response from post JSON
type GPT struct {
	Model string       `json:"model"`
	Gpt2  []GptMessage `json:"choices"`
}

// sublevel ChatGPT response from post JSON
type GptMessage struct {
	Message ReqRespContent `json:"message"`
}
