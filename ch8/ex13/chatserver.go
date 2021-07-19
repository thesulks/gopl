package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	ch chan<- string
	id string
}

var entering = make(chan *client)
var leaving = make(chan *client)
var messages = make(chan string)

func broadcaster() {
	clients := make(map[*client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			cli.ch <- clientsList(clients)
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func clientsList(clients map[*client]bool) string {
	var cliendIds []string
	for cli := range clients {
		cliendIds = append(cliendIds, cli.id)
	}
	return strings.Join(cliendIds, "\n")
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"

	cli := &client{ch, who}
	entering <- cli

	const timeout = 10 * time.Second
	ticker := time.NewTicker(timeout)
	left := func() {
		ticker.Stop()
		leaving <- cli
		messages <- who + " has left"
		conn.Close()
	}

	// if idle for timeout
	closed := make(chan struct{})
	go func() {
		select {
		case <-ticker.C:
			left()
			closed <- struct{}{}
		case <-closed: // do nothing
			fmt.Println("hi")
		}
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		ticker.Reset(timeout)
		messages <- who + ": " + input.Text()
	}

	select {
	case <-closed:
	default:
		closed <- struct{}{}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
