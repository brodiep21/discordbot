package responses

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

// get env key for chatGPT
var key = os.Getenv("key")

func ChatGptMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	listenTo := m.Author.Username
	_, err := s.ChannelMessageSend(m.ChannelID, "I'm listening, go ahead and type what you'd like me to send to chatgpt, "+m.Author.Username+" (please allow a few seconds for it to respond back and I'll relay the message!) ")
	if err != nil {
		fmt.Println(err)
	}
	//this is a secondary listener to look for the follow up response to send over to chatGPT
	newlistener := func(s *discordgo.Session, m *discordgo.MessageCreate) {
		//verify the author we are now listening to for the google search is the one who initiated the search.
		if m.Author.Username == listenTo {
			msg, err := RequestToChatGpt(m.Content)
			if err != nil {
				return
			}
			_, err = s.ChannelMessageSend(m.ChannelID, "ChatGPT 3.5Turbo Bot says \n"+msg)
			if err != nil {
				fmt.Println(err)
			}

		}

	}
	//open a new websocket connection in discord to allow for another handler
	s.Open()
	s.AddHandlerOnce(newlistener)
	return nil
}
func RequestToChatGpt(m string) (string, error) {

	endmsg := "tell Brodie to fix it"
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", nil)
	if err != nil {
		return "Sorry the request didn't go through to ChatGPT, looks like an error on their end.", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	d := Requester{
		Model: "gpt-3.5-turbo",
		Messages: []ReqRespContent{
			ReqRespContent{
				Role:    "user",
				Content: m,
			},
		},
	}

	msg, err := json.Marshal(d)
	if err != nil {
		return "Sorry, I couldn't JSON marshall your response, that's a golang issue maybe? " + endmsg, err
	}
	// fmt.Println(string(data))
	req.Body = ioutil.NopCloser(bytes.NewReader(msg))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "Sorry, I could not properly make an api request to ChatGPT, " + endmsg, err
	}

	defer resp.Body.Close()

	respMsg := GPT{}
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Sorry, I couldn't read the response from ChatGPT, " + endmsg, err
	}

	err = json.Unmarshal(response, &respMsg)
	if err != nil {
		return "I couldn't parse ChatGPT's response, " + endmsg, err
	}
	fmt.Println(respMsg.Gpt2[0].Message.Content)
	return respMsg.Gpt2[0].Message.Content, nil
}
