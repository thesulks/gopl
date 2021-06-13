package elemcount

import (
	"golang.org/x/net/html"
)

func CountElements(count map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return count
	}

	if count == nil {
		count = make(map[string]int)
	}
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	count = CountElements(count, n.FirstChild)
	return CountElements(count, n.NextSibling)
}
