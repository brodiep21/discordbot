package responses

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//Gets an image provided from the URL, then creates a filename and saves the image to it.
func CreateFilePhoto(URL, filename string) {

	//GET CALL URL
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	//create a file based on provided name
	newfile, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}

	defer newfile.Close()

	//writes the bytes into the file and closes.
	_, err = io.Copy(newfile, resp.Body)
	if err != nil {
		fmt.Println("Could not copy the file into memory.")
	}

}

//creates a response to what they rolled if it's a DND die, just for fun!
func Quirkresponse(ddice int, style string) string {
	var response string
	if style == "DnD" {
		switch ddice {
		case 20:
			response = "Perfect roll, critical hit!"
			return response
		case 19, 18, 17, 16, 15, 14, 13, 12, 11, 10:
			response = "great roll!"
			return response
		case 9, 8, 7, 6:
			response = "not bad!"
			return response
		case 5, 4, 3, 2:
			response = "could still be worse."
			return response
		case 1:
			response = "oof. Better luck next time!"
			return response
		}
	} else if style == "regular" {
		response = ""
		return response
	}
	return response
}

// func Scraper (link string) {

// }
