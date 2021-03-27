// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	occurrence := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, occurrence)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, occurrence)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fileNames := ""
			for key := range occurrence[line] {
				fileNames += key + " "
			}
			fmt.Printf("%d\t%s\t%s\n", n, line, fileNames)
		}
	}
}
func countLines(f *os.File, counts map[string]int, occurrence map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if occurrence[line] == nil {
			occurrence[line] = make(map[string]bool)
		}
		occurrence[line][f.Name()] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}
