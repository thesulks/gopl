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
		fmt.Fprintf(os.Stderr, "findtext: %v\n", err)
		os.Exit(1)
	}
	for _, t := range visit(nil, doc) {
		fmt.Println(t)
	}
}

func visit(text []string, n *html.Node) []string {
	if n == nil {
		return text
	}
	if n.Type == html.ElementNode && (n.Data == "style" || n.Data == "script") {
		return text
	}
	if n.Type == html.TextNode {
		if trimed := strings.TrimSpace(n.Data); len(trimed) != 0 {
			text = append(text, trimed)
		}
	}
	text = visit(text, n.FirstChild)
	return visit(text, n.NextSibling)
}
