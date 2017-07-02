// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"os"
	"strings"

	"log"

	"golang.org/x/net/html"
)

func main() {
	file, err := os.Open("test.html")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(file)
	forEachNode(doc, startElement, endElement)
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node) {
	if n.Type == html.TextNode && !strings.HasPrefix(n.Data, "\n  ") {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}

	if n.Type == html.ElementNode {
		if n.Data == "a" {
			fmt.Printf("%*s<%s>", depth*2, "", n.Data)
			for j := range n.Attr {
				fmt.Printf(" %v ", j)
			}
			fmt.Printf("\n")
		} else if n.FirstChild == nil {
			fmt.Printf("%*s<%s/> \n", depth*2, "", n.Data)
		} else {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		}
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
		depth--
	}

}

//!-startend
