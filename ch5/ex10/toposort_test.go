package main

import (
	"log"
	"testing"
)

func TestTopoSort(t *testing.T) {
	const repeat = 1_000_000
	for i := 0; i < repeat; i++ {
		courses := topoSort(prereqs)
		if !isValidResult(courses, prereqs) {
			t.Error("results are not valid: see above logs")
		}
	}
}

func isValidResult(result []string, m map[string]map[string]bool) bool {
	for course, prereqs := range m {
		pos := getIndex(result, course)
		if pos == -1 {
			log.Printf("%q not found in %v\n", course, result)
			return false
		}
		for prereq := range prereqs {
			posPre := getIndex(result, prereq)
			if posPre == -1 {
				log.Printf("%q not found in %v\n", prereq, result)
				return false
			}
			if posPre >= pos {
				log.Printf("invalid topological ordering: (%d, %s), (%d, %s)\n", pos, course, posPre, prereq)
				return false
			}
		}
	}
	return true
}

func getIndex(s []string, target string) int {
	for i, v := range s {
		if v == target {
			return i
		}
	}
	return -1
}
