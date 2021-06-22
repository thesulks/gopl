package main

import "fmt"

func main() {
	fmt.Println(nonzero())
}

func nonzero() (exitcode int) {
	defer func() {
		switch p := recover(); p {
		case nil:
			exitcode = 0
		case 1:
			exitcode = 1
		default:
			panic(p)
		}
	}()
	panic(1)
}
