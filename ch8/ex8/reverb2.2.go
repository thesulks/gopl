package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

const period = 10 * time.Second

func handleConn(c net.Conn) {
	var wg sync.WaitGroup

	lines := make(chan string)
	input := bufio.NewScanner(c)
	go func() {
		for input.Scan() {
			lines <- input.Text()
		}
	}()
	// NOTE: ignoring potential errors from input.Err()

	ticker := time.NewTicker(period)
	for {
		expired := false
		select {
		case line := <-lines:
			ticker.Reset(period)
			wg.Add(1)
			go func(text string) {
				echo(c, text, 1*time.Second)
				wg.Done()
			}(line)
		case <-ticker.C:
			ticker.Stop()
			expired = true
		}
		if expired {
			break
		}
	}

	wg.Wait()
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
