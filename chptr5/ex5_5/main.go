package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"strings"

	"time"

	"golang.org/x/net/html"
)

var wordCount int = 0
var imageCount int = 0

func wordsAndImages(n *html.Node) {
	switch n.Type {
	case html.TextNode:
		input := bufio.NewScanner(strings.NewReader(n.Data))
		input.Split(bufio.ScanWords)
		for input.Scan() {
			wordCount++
		}
	case html.ElementNode:
		if n.Data == "img" {
			fmt.Println(n.Data)
			imageCount++
		}
	}
}

func countWordsAndImages(n *html.Node) (words, images int, err error) {
	wordsAndImages(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countWordsAndImages(c)
	}
	return wordCount, imageCount, nil
}

func main() {
	start := time.Now()
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: PROG URL")
	}
	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Words: %d Images: %d\n", words, images)
	fmt.Printf("%.4f elapse\n", time.Since(start).Seconds())
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	words, images, err = countWordsAndImages(doc)
	return
}
