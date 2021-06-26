package counter

import (
	"fmt"
	"testing"
)

func TestWordCounter(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"interface types", 2},
		{"that interface.\n\nThe io.Writer \t", 4},
		{"            \t\n", 0},
	}

	for _, test := range tests {
		var c WordCounter
		fmt.Fprint(&c, test.input)
		if c != WordCounter(test.want) {
			t.Errorf("%s, got %d, want %d", test.input, c, test.want)
		}
	}
}

func TestLineCounter(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"interface types", 1},
		{"that interface.\n\nThe io.Writer \t", 3},
		{"            \t\n", 1}, // 1 vs. 2
	}

	for _, test := range tests {
		var c LineCounter
		fmt.Fprint(&c, test.input)
		if c != LineCounter(test.want) {
			t.Errorf("%s, got %d, want %d", test.input, c, test.want)
		}
	}
}
