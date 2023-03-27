package responses

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type info struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type requester struct {
	Model    string `json:"model"`
	Messages []info `json:"messages"`
}
type GPT struct {
	Model string       `json:"model"`
	Gpt2  []gptMessage `json:"choices"`
}
type gptMessage struct {
	Message info `json:"message"`
}

var key = os.Getenv("key")

// func ChatGptMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
// 	listenTo := m.Author.Username
// 	_, err := s.ChannelMessageSend(m.ChannelID, "I'm listening, go ahead and type what you'd like me to send to chatgpt, "+m.Author.Username+" (please allow a few seconds for it to respond back and I'll relay the message!) ")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	//this is a secondary listener to look for the follow up response to send over to chatGPT
// 	newlistener := func(s *discordgo.Session, m *discordgo.MessageCreate) {
// 		//verify the author we are now listening to for the google search is the one who initiated the search.
// 		if m.Author.Username == listenTo {
// 			msg, err := RequestToChatGpt(m.Content)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 		_, err = s.ChannelMessageSend(m.ChannelID, `ChatGPT 3.5Turbo Bot says`)
// 		if err != nil {
// 			fmt.Println(err)
// 		}

//		}
//		//open a new websocket connection in discord to allow for another handler
//		s.Open()
//		s.AddHandlerOnce(newlistener)
//	}
func RequestToChatGpt(m string) error {
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	d := requester{
		Model: "gpt-3.5-turbo",
		Messages: []info{
			info{
				Role:    "user",
				Content: m,
			},
		},
	}

	msg, err := json.Marshal(d)
	if err != nil {
		return err
	}
	// fmt.Println(string(data))
	req.Body = ioutil.NopCloser(bytes.NewReader(msg))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	respMsg := GPT{}
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	err = json.Unmarshal(response, &respMsg)
	if err != nil {
		return err
	}

	fmt.Println(respMsg.Gpt2[0].Message.Content)
	return nil
}

// {"id":"chatcmpl-6ylTIq30Cn3sjlKi8rC4Fbyd7qLTi","object":"chat.completion","created":1679940180,"model":"gpt-3.5-turbo-0301",
// "usage":{"prompt_tokens":10,"completion_tokens":10,"total_tokens":20},
// "choices":[{"message":{"role":"assistant","content":"Hello there! How can I assist you today?"},"finish_reason":"stop","index":0}]}
