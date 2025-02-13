package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: web-downloader <url> <output_file>")
		os.Exit(1)

	}

	url := os.Args[1]
	outputFile := os.Args[2]
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	fmt.Printf("Downloading %s to %s\n", url, outputFile)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the http request", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("There was an error. the http status code is: ", resp.StatusCode)
		os.Exit(1)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating file: \n", err.Error())
		os.Exit(1)
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("error copying the content of the body to the file", err.Error())
		os.Exit(1)
	}
	fmt.Println("Download complete!")
}
