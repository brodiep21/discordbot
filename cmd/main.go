package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type data struct {
	model    string    `json: "model"`
	messages []message `json: "messages"`
}
type message struct {
	role    string
	content string
}

var key = os.Getenv("apikey")

func PostRequest(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	rw.Header().Add("Authorization", "Bearer "+key)

	if req.Method == "POST" {
		data := Response{Message: "Hello World From - POST"}
		json, _ := json.Marshal(data)
		fmt.Fprint(rw, string(json))
	} else {
		data := Response{Message: "Bad Request"}
		json, _ := json.Marshal(data)
		fmt.Fprint(rw, string(json))
	}
}

func main() {
	PostRequest()
}

// {"id":"chatcmpl-6sHR9kxGsHF7EtR72QKh2RgGpARhJ","object":"chat.completion","created":1678394759,"model":"gpt-3.5-turbo-0301","usage":{"prompt_tokens":9,"completion_tokens":10,"total_tokens":19},
// "choices":[{"message":{"role":"assistant","content":"\n\nHello! How can I assist you?"},"finish_reason":"stop","index":0}]}
