package element

import (
	"golang.org/x/net/html"
)

func ElementByID(doc *html.Node, id string) *html.Node {
	return ForEachNode(doc, finder(id), nil)
}

func finder(id string) func(*html.Node) bool {
	return func(n *html.Node) (cont bool) {
		if n.Type != html.ElementNode {
			return true
		}
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return false
			}
		}
		return true
	}
}

func ForEachNode(n *html.Node, pre, post func(*html.Node) bool) *html.Node {
	if pre != nil {
		if cont := pre(n); !cont {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if n := ForEachNode(c, pre, post); n != nil {
			return n
		}
	}

	if post != nil {
		if cont := post(n); !cont {
			return n
		}
	}

	return nil
}
