package replace

import (
	"bytes"
	"math/rand"
	"testing"
)

const N = 1_000_000

var buf bytes.Buffer

func init() {
	charSet := []rune{'\t', '\n', '\v', '\f', '\r', ' ', '\u0085', '\u00A0', 'あ'}
	for i := 0; i < N; i++ {
		buf.WriteRune(charSet[rand.Intn(len(charSet))])
	}
}

// '\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).
func TestReplaceSpacesWithSingleSpace(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"\t\n\v\f\r    \u0085\u00A0", " "},
		{"ab안\t\n\v\f\r    \u0085\u00A0", "ab안 "},
		{"\t\na\v\f안\r    \u0085녕\u00A0", " a 안 녕 "},
		{"\t\n\v\f\r    \u0085\u00A0안녕하세요", " 안녕하세요"},
		{"ab안c녕", "ab안c녕"},
		// {"", ""},
	}

	for _, test := range tests {
		bs := []byte(test.input)
		if got := ReplaceSpacesWithSingleSpace(bs); string(got) != test.want {
			t.Errorf("ReplaceSpacesWithSingleSpace(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}
func BenchmarkReplaceSpacesWithSingleSpace(b *testing.B) {
	bs := buf.Bytes()
	for i := 0; i < b.N; i++ {
		ReplaceSpacesWithSingleSpace(bs)
	}
}
