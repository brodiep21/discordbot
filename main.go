package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var Token string

const goapiURL = "https://kutego-api-xxxxx-ew.a.run.app"

func init() {
	flag.StringVar(&Token, "t", "", "Bot 966488243895947264")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}
	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	//Open a websocket Connect to Discord and start listening
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Connection to Discord,", err)
		return
	}

	fmt.Println("Bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

type Gopher struct {
	Name string `json: "name"`
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!gopher" {

		//call the goapiURL and retrieve the Dr. Who gopher
		response, err := http.Get(goapiURL + "/gopher/" + "dr-who")
		if err != nil {
			fmt.Println(err)
		}

		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, "dr-who.png", response.Body)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Error: Can't get Dr.Who Gopher!")
			}
		}
	}

	if m.Content == "!random" {

		response, err := http.Get(goapiURL + "/gopher/random/")
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, "random-gopher.png", response.Body)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: Can't get a random gopher!")
		}
	}

	if m.Content == "!gophers" {

		response, err := http.Get(goapiURL + "/gophers/")
		if err != nil {
			fmt.Println(err)
		}

		defer response.Body.Close()

		if response.StatusCode == 200 {
			//Transform response to []byte
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
			}

			var data []Gopher
			err = json.Unmarshal(body, &data)
			if err != nil {
				fmt.Println(err)
			}

			var gophers strings.Builder
			for _, gopher := range data {
				gophers.WriteString(gopher.Name + "\n")
			}

			//Send a text message with the list of Gophers
			_, err = s.ChannelMessageSend(m.ChannelID, gophers.String())
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: Can't get the list of Gophers!")
		}
	}
}
