package responses

import (
	"strings"

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

		Weather(city, s, m)

	}
	content := strings.ToLower(m.Content)
	switch content {
	case "what can you do gopher?", "gopher help":
		ThingsIcanDo(s, m)
	case "speak gopher":
		SpeakResponse(s, m)
	case "gopher nasa pod":
		NasaResponse(s, m)
	case "hi gopher":
		HiGopher(s, m)
	case "gopher google search":
		GoogleSearch(s, m)
	case "roll the die gopher", "die roll", "roll the dice":
		DiceRoll(s, m)
	case "joke", "gopher joke", "tell me a joke":
		Jokes(s, m)
	case "andy", "surprised andy":
		Andy(s, m)
	}

}
