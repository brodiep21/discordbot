package messages

import (
	"strings"

	"github.com/brodiep21/discordbot/responses"
	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	search := strings.Contains(m.Content, "weather")

	if search {
		//remove the string to only find the city
		city := strings.TrimLeft(m.Content, "whats the weather like in '")

		responses.Weather(city, s, m)

	}

	content := strings.ToLower(m.Content)
	switch content {
	case "what can you do gopher?", "gopher help", "what can you do gopher":
		responses.ThingsIcanDo(s, m)
	case "speak gopher":
		responses.SpeakResponse(s, m)
	case "gopher nasa pod":
		responses.NasaResponse(s, m)
	case "hi gopher":
		responses.HiGopher(s, m)
	case "gopher google search":
		responses.GoogleSearch(s, m)
	case "roll the die gopher", "die roll", "roll the dice":
		responses.DiceRoll(s, m)
	case "joke", "gopher joke", "tell me a joke":
		responses.Jokes(s, m)
	case "andy", "surprised andy":
		responses.Andy(s, m)
	case "!chatgpt":
		responses.ChatGptMessage(s, m)
	}

}
