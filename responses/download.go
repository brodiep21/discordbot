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
	fmt.Println(URL)

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

// func DownloadFile() {

// }
