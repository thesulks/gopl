package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// 1. https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string
// 2. https://stackoverflow.com/questions/46365221/fill-os-stdin-for-function-that-reads-from-it
func TestMain(t *testing.T) {
	input := "aaaaa"
	want := "rune\tcount\n'a'\t5\n\nlen\tcount\n1\t5\n2\t0\n3\t0\n4\t0\n"

	mockStdin, toStdin, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	rescueStdin := os.Stdin
	os.Stdin = mockStdin

	if _, err := toStdin.Write([]byte(input)); err != nil {
		toStdin.Close()
		log.Fatal(err)
	}
	toStdin.Close()

	fromStdout, mockStdout, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	rescueStdout := os.Stdout
	os.Stdout = mockStdout

	main()
	mockStdout.Close()
	got, err := ioutil.ReadAll(fromStdout)
	if err != nil {
		t.Fatal(err)
	}

	if want != string(got) {
		t.Errorf("charcount: in: %q, got: %q, want %q", input, got, want)
	}

	os.Stdout = rescueStdout
	os.Stdin = rescueStdin
}
