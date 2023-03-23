package main

import "github.com/brodiep21/discordbot/responses"

func main() {
	responses.RequestToChatGpt("hello!")
}

// {"id":"chatcmpl-6sHR9kxGsHF7EtR72QKh2RgGpARhJ","object":"chat.completion","created":1678394759,"model":"gpt-3.5-turbo-0301","usage":{"prompt_tokens":9,"completion_tokens":10,"total_tokens":19},
// "choices":[{"message":{"role":"assistant","content":"\n\nHello! How can I assist you?"},"finish_reason":"stop","index":0}]}

// package responses

// check api

// todo list
// https://openai.com/blog/introducing-chatgpt-and-whisper-apis
