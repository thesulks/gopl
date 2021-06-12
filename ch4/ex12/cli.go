package main

import (
	"fmt"
	"gopl/ch4/ex12/xkcd"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprint(os.Stderr, "xkcd: xkcd needs action(search, index)\n")
		os.Exit(1)
	}

	action := os.Args[1]
	switch action {
	case "index":
		if err := xkcd.Index(); err != nil {
			log.Fatalf("xkcd: %v\n", err)
		}
	case "search":
		index, err := xkcd.GetIndexInMemory()
		if err != nil {
			log.Fatalf("xkcd: %v\n", err)
		}
		for _, word := range os.Args[2:] {
			word = strings.ToLower(word)
			fmt.Printf(">>> keyword: %s\n", word)
			for _, n := range index[word] {
				comic, err := xkcd.GetComic(n)
				if err != nil {
					log.Fatalf("xkcd: %v\n", err)
				}
				fmt.Printf("%s\n", xkcd.StringifyComic(comic))
			}
			fmt.Print("\n")
		}
	default:
		log.Fatalf("xkcd: unknown action: %s\n", action)
	}
}
