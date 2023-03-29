package responses

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func SignOnFunc(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	_, err := s.ChannelMessageSend(m.ChannelID, `Gopherbot has signed on! Here are some things I can do:
---------------------------------------------------------------
1.) Send you nasa's picture of the day! Just say - "gopher NASA POD"
2.) Tell you the weather of a city in the U.S!  Just say - "what's the weather like in <insert city>"
3.) I can search google for you and provide the top 3 results! Just say - "gopher google search", I will respond that I'm listening and you can then type in what you want to search!
4.)I can roll a regular die, or a DnD die. Just say - "roll the die gopher" or "die roll", I will ask which one you want to roll. I can also roll 2 regular dice for you. Just say - "roll the dice"
5.)I can give you a random joke.  Just say - "joke", "gopher joke", "tell me a joke"
6.)I can give you an Andy Dwyer gif. Just say - "andy", or "surprised andy"
7.)I can talk to chatGPT for you, just say "!chatgpt"`)
	if err != nil {
		fmt.Println(err)
	}
}
