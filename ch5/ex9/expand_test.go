package main

import "testing"

func TestExpand(t *testing.T) {
	mapping := map[string]string{
		"1":    "one",
		"2":    "two",
		"test": "TEST",
	}
	mapper := func(s string) string {
		return mapping[s]
	}

	tests := []struct {
		input string
		want  string
	}{
		{"$1$2", "onetwo"},
		{"$1 test$2", "one testtwo"},
		{"$test Driven Development", "TEST Driven Development"},
		{"$te st$2 1", " sttwo 1"},
		{"$test$2", "TESTtwo"},
		{"$testt", ""},
	}

	for _, test := range tests {
		if got := Expand(test.input, mapper); got != test.want {
			t.Errorf("Expand(%s, mapper) = %s, want %s", test.input, got, test.want)
		}
	}
}
