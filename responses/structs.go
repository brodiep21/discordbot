package responses

type Nasa struct {
	URL         string `json:"hdurl"`
	Explanation string `json:"explanation"`
}
type Main struct {
	Temp float64 `json:"temp"`
	High float64 `json:"temp_max"`
	Low  float64 `json:"temp_min"`
}

type Weatherinfo struct {
	Main Main
}

type Joke struct {
	Joke string `json:"joke"`
}

type ReqRespContent struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type Requester struct {
	Model    string           `json:"model"`
	Messages []ReqRespContent `json:"messages"`
}
type GPT struct {
	Model string       `json:"model"`
	Gpt2  []GptMessage `json:"choices"`
}
type GptMessage struct {
	Message ReqRespContent `json:"message"`
}
