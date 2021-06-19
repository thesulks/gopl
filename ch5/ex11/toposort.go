package main

import (
	"fmt"
	"os"
)

var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": true},
	"calculus":   {"linear algebra": true},
	// cycle
	"linear algebra": {"calculus": true},

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
	courses, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "toposort: %v\n", err)
	}
	for i, course := range courses {
		fmt.Printf("%02d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]bool) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func([]string) error
	visitAll = func(items []string) error {
		for _, item := range items {
			if seen[item] {
				if getIndex(order, item) == -1 {
					return fmt.Errorf("cycle detected at %q", item)
				}
				continue
			}
			seen[item] = true
			var keys []string
			for key := range m[item] {
				keys = append(keys, key)
			}
			if err := visitAll(keys); err != nil {
				return err
			}
			order = append(order, item)
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	if err := visitAll(keys); err != nil {
		return nil, err
	}
	return order, nil
}

func getIndex(s []string, target string) int {
	for i, v := range s {
		if v == target {
			return i
		}
	}
	return -1
}
