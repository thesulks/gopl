package slice

import "testing"

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestUnique(t *testing.T) {
	tests := []struct {
		input []string
		want  []string
	}{
		{[]string{"a", "a", "a", "b", "a"}, []string{"a", "b", "a"}},
		{[]string{"a", "a", "a", "a", "a"}, []string{"a"}},
		{[]string{"a", "a", "b", "b", "b"}, []string{"a", "b"}},
		{[]string{"a", "b", "c", "d"}, []string{"a", "b", "c", "d"}},
		{[]string{"a"}, []string{"a"}},

		{[]string{}, []string{}},
		{nil, nil},
	}

	for _, test := range tests {
		input := make([]string, len(test.input))
		copy(input, test.input)
		if got := Unique(test.input); !equal(got, test.want) {
			t.Errorf("Unique(%q) = %q, want %q", input, got, test.want)
		} else {
			t.Logf("Unique(%q) = %q", input, got)
		}
	}
}
