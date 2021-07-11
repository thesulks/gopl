package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type clock struct {
	tz      string
	addr    string
	curTime string
}

func (c *clock) String() string {
	return fmt.Sprintf("%s %s", c.tz, c.curTime)
}

const Unavailable = "--:--:--"

func main() {
	clocks := make([]*clock, 0)
	for _, arg := range os.Args[1:] {
		s := strings.Split(arg, "=")
		tz, address := s[0], s[1]
		clocks = append(clocks, &clock{tz, address, Unavailable})
	}

	for _, clock := range clocks {
		go startClock(clock)
	}

	for {
		curTimes := make([]string, 0)
		for _, clock := range clocks {
			curTimes = append(curTimes, clock.String())
		}
		fmt.Printf("\r%s", strings.Join(curTimes, "\t\t"))
		time.Sleep(250 * time.Millisecond) // update clock every 100ms
	}
}

func startClock(c *clock) {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		log.Print(err)
		c.curTime = Unavailable
		return
	}
	s := bufio.NewScanner(conn)
	for s.Scan() {
		c.curTime = s.Text()
	}
	if err := s.Err(); err != nil {
		log.Print(err)
		c.curTime = Unavailable
		return
	}
}
