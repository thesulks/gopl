package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var npipe = flag.Int("pipe", 2, "number of pipeline")

func main() {
	flag.Parse()
	if *npipe < 2 {
		log.Fatalf("pipeline: invalid number of pipeline: %d", *npipe)
	}

	channels := make([]chan int, *npipe)
	for i := range channels {
		channels[i] = make(chan int)
	}

	for i := 0; i < len(channels)-1; i++ {
		go func(out chan<- int, in <-chan int) {
			for v := range in {
				out <- v
			}
			close(out)
		}(channels[i+1], channels[i])
	}

	channels[0] <- 1
	start := time.Now()
	<-channels[*npipe-1]
	end := time.Since(start)

	fmt.Printf("npipe: %d  %s", *npipe, end)
}
