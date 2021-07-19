package main

import (
	"flag"
	"fmt"
	"gopl/links"
	"log"
	"os"
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

var done = make(chan struct{})

func main() {
	flag.Parse()

	worklist := make(chan *List)
	var n int

	n++
	go func() { worklist <- &List{flag.Args(), 1} }()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	seen := make(map[string]bool)
	for n > 0 {
		select {
		case list := <-worklist:
			n--
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
		case <-done:
			for range worklist {
				n-- // drain
				if n == 0 {
					break
				}
			}
			return
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)

	select {
	case tokens <- struct{}{}:
	case <-done:
		return nil
	}

	list, err := links.CancelableExtract(done, url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
