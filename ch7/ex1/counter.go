package counter

import (
	"bufio"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	n := 0
	for n < len(p) {
		adv, token, err := bufio.ScanWords(p[n:], true)
		n += adv
		if err != nil {
			return n, err
		}
		if token != nil {
			*c++
		}
	}
	return n, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	n := 0
	for n < len(p) {
		adv, token, err := bufio.ScanLines(p[n:], true)
		n += adv
		if err != nil {
			return n, err
		}
		if token != nil {
			*c++
		}
	}
	return n, nil
}
