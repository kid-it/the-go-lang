// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var hist = make(map[string]int)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
	printMap()
}

func visit(n *html.Node) {
	if n.Type == html.ElementNode {
		hist[n.Data]++
	}
	if n.FirstChild != nil {
		visit(n.FirstChild)
	}
	if n.NextSibling != nil {
		visit(n.NextSibling)
	}
	return
}

func printMap() {
	for key, val := range hist {
		fmt.Printf("%s, %d\n", key, val)
	}
}
