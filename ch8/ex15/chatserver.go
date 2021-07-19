package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
				select {
				case cli.ch <- msg: /* do nothing */
				default: /* non-blocking */
				}
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
	return "\nUser List:\n" + strings.Join(cliendIds, "\n") + "\n"
}

func handleConn(conn net.Conn) {
	const bufSize = 20
	ch := make(chan string, bufSize)
	go clientWriter(conn, ch)

	ch <- "Enter your name: "
	input := bufio.NewScanner(conn)
	input.Scan()

	who := input.Text()
	ch <- "You are " + who
	messages <- who + " has arrived"

	cli := &client{ch, who}
	entering <- cli

	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
