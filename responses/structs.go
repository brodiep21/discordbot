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
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}
