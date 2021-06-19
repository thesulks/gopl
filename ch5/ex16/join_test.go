package join

import "testing"

func TestVariadicJoin(t *testing.T) {
	tests := []struct {
		elems []string
		sep   string
		want  string
	}{
		{[]string{"a", "b", "c"}, "-", "a-b-c"},
		{nil, "-", ""},
		{[]string{}, "-", ""},
		{[]string{"a"}, "-", "a"},
	}

	for _, test := range tests {
		if got := variadicJoin(test.sep, test.elems...); got != test.want {
			t.Errorf("variadicJoin(%s, %v) = %s, want %s", test.sep, test.elems, got, test.want)
		}
	}
}
