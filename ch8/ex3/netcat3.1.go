package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	tcpConn := conn.(*net.TCPConn)
	done := make(chan struct{})
	go func() {
		if _, err := io.Copy(os.Stdout, tcpConn); err != nil {
			log.Print(err)
		}
		tcpConn.CloseRead()
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(tcpConn, os.Stdin)
	tcpConn.CloseWrite()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
