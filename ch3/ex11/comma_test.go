package comma

import "testing"

func TestFComma(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"12345", "12,345"},
		{"123456", "123,456"},
		{"1234567", "1,234,567"},
		{"12345678", "12,345,678"},
		{"123456789", "123,456,789"},

		{"+123", "+123"},
		{"-123", "-123"},
		{"+123.4", "+123.4"},
		{"-123.4", "-123.4"},
		{"+1234", "+1,234"},
		{"-1234", "-1,234"},
		{"+1234.56", "+1,234.56"},
		{"-1234.56", "-1,234.56"},
	}

	for _, test := range tests {
		if got := fComma(test.input); got != test.want {
			t.Errorf("comma(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}
