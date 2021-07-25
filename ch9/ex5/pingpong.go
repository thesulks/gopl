package main

import (
	"fmt"
	"time"
)

func main() {
	chA, chB := make(chan string), make(chan string)
	var ping, pong int
	const nsec = 10

	tick := time.Tick(time.Second * nsec)
	go func() {
		for msg := range chA {
			ping++
			chB <- msg
		}
	}()
	go func() {
		for msg := range chB {
			pong++
			chA <- msg
		}
	}()
	go func() { chA <- "hey" }()
	<-tick

	fmt.Printf("%.2f\n", float64(ping+pong)/nsec)
}
