package main

import (
	"fmt"
	"regexp"
)

func main() {
	mapping := map[string]string{
		"1":    "one",
		"2":    "two",
		"test": "TEST",
	}
	mapper := func(s string) string {
		return mapping[s]
	}

	fmt.Println(Expand("$1$2", mapper))
}

var env = regexp.MustCompile(`\$[^\s\$]+`)

func Expand(s string, f func(string) string) string {
	return env.ReplaceAllStringFunc(s, func(sub string) string {
		return f(sub[1:])
	})
}
