package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func countWordsAndImages(n *html.Node) (words, images int, err error) {
	fmt.Println(n.Data)
	fmt.Println(n.Type)
	for j := range n.Attr {
		fmt.Println(j)
	}

	for m := n.FirstChild; m != nil; m = n.NextSibling {
		words, images, err = countWordsAndImages(m)
	}

	words = 3
	images = 4

	return
}

func main() {
	words, images, err := CountWordsAndImages("http://golang.org")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(words)
	fmt.Println(images)
	fmt.Println(err)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// _, err1 := io.Copy(os.Stdout, resp.Body)
	// if err1 != nil {
	// 	log.Fatal(err1)
	// }
	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	words, images, err = countWordsAndImages(doc)
	return
}
