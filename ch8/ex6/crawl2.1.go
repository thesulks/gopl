package main

import (
	"flag"
	"fmt"
	"gopl/links"
	"log"
	"strings"
)

type List struct {
	links []string
	depth int
}

func (l *List) String() string {
	return strings.Join(l.links, "\n")
}

var depthLimit = flag.Int("depth", 1, "depth limiting to crawler")

func main() {
	flag.Parse()

	worklist := make(chan *List)
	var n int

	n++
	go func() { worklist <- &List{flag.Args(), 1} }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		if list.depth == *depthLimit {
			fmt.Println(list)
			continue
		}
		for _, link := range list.links {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- &List{crawl(link), list.depth + 1}
				}(link)
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
