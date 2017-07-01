// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

func visit(n *html.Node) {

	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}
	if n.Type == html.TextNode {
		if len(n.Data) > 0 && !strings.HasPrefix(n.Data, "\n  ") {
			fmt.Printf("%s", n.Data)
		}
	}

	if n.FirstChild != nil {
		visit(n.FirstChild)
	}
	if n.NextSibling != nil {
		visit(n.NextSibling)
	}
	return
}
