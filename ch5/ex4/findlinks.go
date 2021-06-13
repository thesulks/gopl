package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Printf("%q\n", link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a", "link":
			if link, err := getAttrVal(n, "href"); err == nil {
				links = append(links, link)
			}
		case "img", "script":
			if link, err := getAttrVal(n, "src"); err == nil {
				links = append(links, link)
			}
		}
	}
	links = visit(links, n.FirstChild)
	return visit(links, n.NextSibling)
}

func getAttrVal(n *html.Node, key string) (string, error) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, nil
		}
	}
	return "", fmt.Errorf("getAttrVal: %q not found in <%s>", key, n.Data)
}
