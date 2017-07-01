package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func countWordsAndImages(n *html.Node) (words, images int, err error) {
	words = 1
	images = 3
	err = nil
	return
}

func main() {
	words, images, err := CountWordsAndImages("http://www.google.com")
	fmt.Println(words)
	fmt.Println(images)
	fmt.Println(err)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	words, images, err = countWordsAndImages(doc)
	return
}
