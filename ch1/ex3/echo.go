package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(echoJoin(os.Args[1:]))
}

func echoNaive(args []string) string {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}

func echoJoin(args []string) string {
	return strings.Join(args, " ")
}
