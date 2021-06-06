package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countWords(os.Stdin, counts)
	} else {
		for _, filepath := range files {
			f, err := os.Open(filepath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
				continue
			}
			countWords(f, counts)
			f.Close()
		}
	}
	for word, n := range counts {
		fmt.Printf("%s\t%d\n", word, n)
	}
}

func countWords(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
}
