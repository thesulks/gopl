package mystrings

import "io"

type StringReader string

func (sr *StringReader) Read(p []byte) (int, error) {
	if len(*sr) == 0 {
		return 0, io.EOF
	}
	n := copy(p, *sr)
	*sr = (*sr)[n:]
	return n, nil
}

func NewReaderFromString(s string) io.Reader {
	sr := StringReader(s)
	return &sr
}
