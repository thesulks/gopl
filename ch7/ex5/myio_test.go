package myio

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	const limit = 4
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, limit)

	var buf bytes.Buffer
	n, err := io.Copy(&buf, lr)
	if err != nil {
		t.Fatal(err)
	}

	if n != limit || buf.String() != "some" {
		t.Errorf("buf: %s, n: %d", buf.String(), n)
	}
}
