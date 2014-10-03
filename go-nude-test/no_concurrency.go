package main

import (
	"fmt"
	"github.com/koyachi/go-nude"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

// Download a file and write to disk
func downloadFile(url string) (string, error) {
	splittedFileName := strings.Split(url, "/")
	fileName := splittedFileName[len(splittedFileName)-1]
	fmt.Print("Downloading ", fileName, " ... ")
	output, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer output.Close()
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	n, err := io.Copy(output, response.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(n, " Bytes")
	return fileName, nil
}

func main() {
	/*
		First image is Doge, second is not a picture, three next are naked pictures, fourth is a test image
		that comes from the original repo of the creator of go-nude package.
	*/
	totalStart := time.Now()
	testString := []string{
		"http://img1.wikia.nocookie.net/__cb20140807211510/bayonetta/images/0/05/Doge.png",
		"http://notapicture.com",
		"http://i.imgur.com/BsUmG5H.jpg",
		"http://i.imgur.com/WeXPbWf.jpg",
		"http://i.imgur.com/FxhJf9Jh.jpg",
		"https://raw.githubusercontent.com/koyachi/go-nude/master/example/images/test2.jpg",
	}
	r, _ := regexp.Compile("^https?:.*(jpg|png|gif)$")
	for _, url := range testString {
		if r.MatchString(url) {
			fileName, err := downloadFile(url)
			if err != nil {
				fmt.Println("Error while downloading file ", url, " - ", err)
				return
			}
			defer os.Remove(fileName)
			start := time.Now()
			isNude, err := nude.IsNude(fileName)
			elapsed := time.Since(start)
			if err != nil {
				fmt.Println("Error while checking for nudity -", err)
			}
			fmt.Println(isNude, " Test took ", elapsed)
		}
	}
	fmt.Println("Program took : ", time.Since(totalStart))
}
