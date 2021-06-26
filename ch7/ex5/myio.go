package myio

import "io"

type limitReader struct {
	wrapped io.Reader
	limit   int64
}

func (lr *limitReader) Read(p []byte) (int, error) {
	if lr.limit == 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > lr.limit {
		p = p[:lr.limit]
	}
	n, err := lr.wrapped.Read(p)
	lr.limit -= int64(n)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r, n}
}
