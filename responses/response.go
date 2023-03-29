package responses

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	googlesearch "github.com/rocketlaunchr/google-search"
)

func SpeakResponse(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, `Gopher reporting for duty! Type "gopher help" if you'd like more information`)
	if err != nil {
		fmt.Println(err)
	}

	pic, err := os.Open("gopher.jpg")
	if err != nil {
		fmt.Println("Cannot open gopher picture")
	}
	defer pic.Close()

	_, err = s.ChannelFileSend(m.ChannelID, "gopher.jpg", io.Reader(pic))
	if err != nil {
		fmt.Println("Error io.Reading the gopher jpg")
	}
}

func NasaResponse(s *discordgo.Session, m *discordgo.MessageCreate) {

	nasakey := os.Getenv("nasakey")

	var n Nasa

	response, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=" + nasakey)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Gopher couldn't get the picture for you!")
		fmt.Println(err)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("http.Get -> %v", err)
	}

	//unmarshall data from NASA API to receive hdurl
	err = json.Unmarshal(data, &n)
	if err != nil {
		fmt.Printf("Could not unmarshal %v", err)
	}

	CreateFilePhoto(n.URL, "NasaPod.jpg")

	defer response.Body.Close()

	pic, err := os.Open("NasaPod.jpg")
	if err != nil {
		fmt.Println("Cannot open Nasa picture")
	}
	defer pic.Close()

	if response.StatusCode == 200 {
		_, err = s.ChannelFileSendWithMessage(m.ChannelID, n.Explanation, "NasaPod.jpg", io.Reader(pic))
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Error: Can't get nasa photo")
	}
}
func HiGopher(s *discordgo.Session, m *discordgo.MessageCreate) {

	_, err := s.ChannelMessageSend(m.ChannelID, "Hi "+m.Author.Username)
	if err != nil {
		fmt.Println(err)
	}
}

func ThingsIcanDo(s *discordgo.Session, m *discordgo.MessageCreate) {

	_, err := s.ChannelMessageSend(m.ChannelID, "Well "+m.Author.Username+" let me give you the lowdown.")
	if err != nil {
		fmt.Println(err)
	}
	_, err = s.ChannelMessageSend(m.ChannelID, `So far, I can't do a ton, but I'm working on it!. I'm currently able to:
---------------------------------------------------------------
1.) Send you nasa's picture of the day! Just say - "gopher NASA POD"
2.) Tell you the weather of a city in the U.S!  Just say - "what's the weather like in <insert city>"
3.) I can search google for you and provide the top 3 results! Just say - "gopher google search", I will respond that I'm listening and you can then type in what you want to search!
4.)I can roll a regular die, or a DnD die. Just say - "roll the die gopher" or "die roll", I will ask which one you want to roll. I can also roll 2 regular dice for you. Just say - "roll the dice"
5.)I can give you a random joke.  Just say - "joke", "gopher joke", "tell me a joke"
6.)I can give you an Andy Dwyer gif. Just say - "andy", or "surprised andy"`)
	if err != nil {
		fmt.Println(err)
	}
}

func Weather(city string, s *discordgo.Session, m *discordgo.MessageCreate) {

	var w Weatherinfo

	weatherapi := os.Getenv("weatherapi")

	weatherinputstring := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + weatherapi + "&units=imperial"

	client := &http.Client{Timeout: 3 * time.Second}

	req, err := client.Get(weatherinputstring)
	if err != nil {
		fmt.Println(err)
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(body, &w)

	//alter json temps from float to string
	temp := strconv.Itoa(int(w.Main.Temp))
	h := strconv.Itoa(int(w.Main.High))
	l := strconv.Itoa(int(w.Main.Low))

	_, err = s.ChannelMessageSend(m.ChannelID, "It is currently "+temp+" with a high of "+h+", and a low of "+l)
	if err != nil {
		fmt.Println(err)
	}
}

// responses to the Person requesting to search google
func GoogleSearch(s *discordgo.Session, m *discordgo.MessageCreate) error {
	listenTo := m.Author.Username
	_, err := s.ChannelMessageSend(m.ChannelID, "Type what you'd like me to search "+m.Author.Username)
	if err != nil {
		fmt.Println(err)
	}

	//this is a secondary listener to look for the follow up response to the google search func initiation.
	newlistener := func(s *discordgo.Session, m *discordgo.MessageCreate) {
		//verify the author we are now listening to for the google search is the one who initiated the search.
		if m.Author.Username == listenTo {
			google, err := googlesearch.Search(context.TODO(), m.Content)
			if err != nil {
				_, err = s.ChannelMessageSend(m.ChannelID, "I'm sorry, I couldn't search google because I encountered an error ")
				fmt.Println(err)
			}
			_, err = s.ChannelMessageSend(m.ChannelID, "Thanks for responding! Here's your top 3 search results!")
			if err != nil {
				_, err = s.ChannelMessageSend(m.ChannelID, "I'm sorry, I couldn't return your google results because of an error: ")
				if err != nil {
					fmt.Println(err)
				}
			}
			//range over the google search and only find the first 3 search results. (We can limit or add to search results up to 100 results per search)
			for _, v := range google {
				if v.Rank < 4 {
					_, err = s.ChannelMessageSend(m.ChannelID, strconv.Itoa(v.Rank)+"\n"+v.URL)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}

	s.Open()
	s.AddHandlerOnce(newlistener)
	return nil
}

func DiceRoll(s *discordgo.Session, m *discordgo.MessageCreate) {
	listenTo := m.Author.Username
	if m.Content == "roll the dice" {
		ddice := rand.Intn(6)
		ddice2 := rand.Intn(6)

		if ddice == 0 {
			ddice = 1
		}
		if ddice2 == 0 {
			ddice2 = 1
		}
		_, err := s.ChannelMessageSend(m.ChannelID, "You rolled a "+strconv.Itoa(ddice)+", and a "+strconv.Itoa(ddice2))
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	_, err := s.ChannelMessageSend(m.ChannelID, "What type of Dice? Regular or DnD?")
	if err != nil {
		fmt.Println("Couldn't roll the dice", err)
	}

	//this is a secondary listener to look for the follow up response to the dice request func initiation if they aren't calling for two dice.
	newlistener := func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.Username == listenTo {
			if m.Content == "dnd" || m.Content == "DnD" || m.Content == "DND" {
				ddice := rand.Intn(20)
				if ddice == 0 {
					ddice = 1
				}
				_, err = s.ChannelMessageSend(m.ChannelID, "You rolled a "+strconv.Itoa(ddice)+", "+Quirkresponse(ddice, "DnD"))
				if err != nil {
					fmt.Println(err)
				}
			} else if m.Content == "Regular" || m.Content == "regular" {
				ddice := rand.Intn(6)
				_, err = s.ChannelMessageSend(m.ChannelID, "You rolled a "+strconv.Itoa(ddice)+Quirkresponse(ddice, "regular"))
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	s.Open()
	s.AddHandlerOnce(newlistener)
}

func Jokes(s *discordgo.Session, m *discordgo.MessageCreate) {
	var d Joke
	client := &http.Client{Timeout: 3 * time.Second}

	req, err := client.Get("https://v2.jokeapi.dev/joke/Any?type=single")
	if err != nil {
		fmt.Println(err)
	}
	req.Header = http.Header{
		"Accept":     []string{"application/json"},
		"User-Agent": []string{"discordbotGopher"},
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body, &d)

	if req.StatusCode == 200 {
		_, err = s.ChannelMessageSend(m.ChannelID, `Rando joke!
`+d.Joke)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func Andy(s *discordgo.Session, m *discordgo.MessageCreate) {
	pic, err := os.Open("surprised-andy.gif")
	if err != nil {
		fmt.Println("Cannot open andy gif")
	}
	defer pic.Close()

	_, err = s.ChannelFileSend(m.ChannelID, "surprised-andy.gif", io.Reader(pic))
	if err != nil {
		fmt.Println("Error io.Reading the gopher jpg")
	}
}
