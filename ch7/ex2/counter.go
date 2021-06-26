package counter

import (
	"bufio"
	"io"
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

type countingWriter struct {
	wrapped io.Writer
	n       int64
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.wrapped.Write(p)
	cw.n += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := countingWriter{w, 0}
	return &cw, &(cw.n)
}
