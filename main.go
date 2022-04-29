package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/brodiep21/discordbot/responses"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// Variables used for command line parameters
var (
	Token string
)

func main() {
	err := godotenv.Load(".env")
	key := os.Getenv("apikey")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + key)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Close the Discord session.
	dg.Close()

}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.

	if m.Author.ID == s.State.User.ID {
		return
	}

	search := strings.Contains(m.Content, "weather")

	if search {
		//remove the string to only find the city
		city := strings.TrimLeft(m.Content, "whats the weather like in '")
		responses.Weather(city, s, m)

	}

	switch m.Content {
	case "What can you do gopher?", "what can you do gopher?", "what can you do gopher":
		responses.ThingsIcanDo(s, m)
	case "speak gopher", "Speak Gopher":
		responses.SpeakResponse(s, m)
	case "gopher NASA POD", "Gopher NASA POD", "gopher nasa pod":
		responses.NasaResponse(s, m)
	case "Hi gopher", "Hi Gopher", "hi gopher":
		responses.HiGopher(s, m)
		// case "gopher", "Gopher":
		// case :
	}

}
