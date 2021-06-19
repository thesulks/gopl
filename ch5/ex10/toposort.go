package main

import (
	"fmt"
)

var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": true},
	"calculus":   {"linear algebra": true},

	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":       {"discrete math": true},
	"databases":             {"data structures": true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math": true},
	"networks":              {"operating systems": true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%02d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]bool) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func([]string)
	visitAll = func(items []string) {
		for _, item := range items {
			if seen[item] {
				continue
			}
			seen[item] = true
			var keys []string
			for key := range m[item] {
				keys = append(keys, key)
			}
			visitAll(keys)
			order = append(order, item)
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	visitAll(keys)
	return order
}
