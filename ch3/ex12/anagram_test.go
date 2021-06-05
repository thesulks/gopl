package anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	type Input struct {
		a string
		b string
	}

	tests := []struct {
		input Input
		want  bool
	}{
		{Input{"abcd", "dcba"}, true},
		{Input{"abba", "abba"}, true},
		{Input{"abbacc", "acbcba"}, true},

		{Input{"안녕하세요", "요세하녕안"}, true},
		{Input{"おはよう", "うよはお"}, true},

		{Input{"", ""}, true},

		{Input{"abc", "cbaa"}, false},
		{Input{"おはよう", "うよはaお"}, false},
	}

	for _, test := range tests {
		if got := IsAnagram(test.input.a, test.input.b); got != test.want {
			t.Errorf("IsAnagram(%v) = %v, want %v", test.input, got, test.want)
		} else {
			t.Logf("IsAnagram(%v) = %v", test.input, got)
		}
	}
}
