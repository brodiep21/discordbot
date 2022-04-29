package responses

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Nasa struct {
	URL         string `json:"hdurl"`
	Explanation string `json:"explanation"`
}

func SpeakResponse(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, "Gopher reporting for duty!")
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
2.) Tell you the weather of a city in the U.S!  Just say - "what's the weather like in <insert city>"`)
	if err != nil {
		fmt.Println(err)
	}
}

func Weather(city string, s *discordgo.Session, m *discordgo.MessageCreate) {
	type Main struct {
		Temp float64 `json:"temp"`
		High float64 `json:"temp_max"`
		Low  float64 `json:"temp_min"`
	}

	type Weatherinfo struct {
		Main Main
	}

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
