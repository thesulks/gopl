package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		fmt.Println(url)
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "count: %v\n", err)
			continue
		}
		fmt.Printf("words: %d\timages: %d\n\n", words, images)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode {
		if n.Data == "style" || n.Data == "script" {
			return
		}
		if n.Data == "img" {
			images++
		}
	}
	if n.Type == html.TextNode {
		s := bufio.NewScanner(strings.NewReader(n.Data))
		s.Split(bufio.ScanWords)
		for s.Scan() {
			words++
		}
	}
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		childWords, childImages := countWordsAndImages(child)
		words += childWords
		images += childImages
	}
	return
}
