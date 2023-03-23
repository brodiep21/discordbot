package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/brodiep21/discordbot/responses"
	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func main() {
	//automatically load ".env files"
	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println("problem loading .env file")
	// }

	key := os.Getenv("apikey")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + key)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	// dg.AddHandler(responses.SignOnFunc())
	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(responses.MessageCreate)

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
