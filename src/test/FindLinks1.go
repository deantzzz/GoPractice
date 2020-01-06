package main

import (
	"fmt"
	html2 "gitgo/src/test/html"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func main() {
	featch := html2.Featch()
	doc, err := html.Parse(strings.NewReader(string(featch)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
