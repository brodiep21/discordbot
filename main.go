package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

const goAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"

// func init() {
// 	flag.StringVar(&Token, "t", "", "Bot Token")
// 	flag.Parse()
// }

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + "")
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

	// Cleanly close down the Discord session.
	dg.Close()
}

type Gopher struct {
	Name string `json:"name"`
}

type Nasa struct {
	URL string `json:"URL"`
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!gopher" {

		//Call the KuteGo API and retrieve our cute Dr Who Gopher
		response, err := http.Get(goAPIURL + "/gopher/" + "dr-who")
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, "dr-who.png", response.Body)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: Can't get dr-who Gopher! :-(")
		}
	}

	if m.Content == "!random" {

		//Call the KuteGo API and retrieve a random Gopher
		response, err := http.Get(goAPIURL + "/gopher/random/")
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
			fmt.Println("Error: Can't get random Gopher! :-(")
		}
	}

	if m.Content == "!gophers" {

		//Call the KuteGo API and display the list of available Gophers
		response, err := http.Get(goAPIURL + "/gophers/")
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			// Transform our response to a []byte
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
			}

			// Put only needed informations of the JSON document in our array of Gopher
			var data []Gopher
			err = json.Unmarshal(body, &data)
			if err != nil {
				fmt.Println(err)
			}

			// Create a string with all of the Gopher's name and a blank line as separator
			var gophers strings.Builder
			for _, gopher := range data {
				gophers.WriteString(gopher.Name + "\n")
			}

			// Send a text message with the list of Gophers
			_, err = s.ChannelMessageSend(m.ChannelID, gophers.String())
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: Can't get list of Gophers! :-(")
		}
	}
	if m.Content == "!speak gopher" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Gopher reporting for duty!")
		if err != nil {
			fmt.Println(err)
		}

		pic, err := os.Open("gopher.jpg")
		if err != nil {
			fmt.Println("Cannot open gopher picture")
		}
		defer pic.Close()

		gopher := io.Reader(pic)

		_, error2 := s.ChannelFileSend(m.ChannelID, "gopher", gopher)
		if error2 != nil {
			fmt.Println("Error io.Reading the gopher jpg")
		}
	}

	if m.Content == "!gopher NASA POD" {
		response, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=e4G3LccD485rgffH5rsvxBjX0bPG2HrH60l0jRXg")
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Gopher couldn't get the picture for you!")
			fmt.Println(err)
		}

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("http.Get -> %v", err)
		}

		//unmarshall NASA get request data to receive "URL" parameter
		// go to URL and read bytes
		ioutil.WriteFile("Planetary_photo.jpg", data, 0666)

		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, "Planetary_photo.jpg", response.Body)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: Can't get nasa photo")
		}
	}
}
