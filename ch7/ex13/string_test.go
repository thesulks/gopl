package eval

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		expr string
		want string
	}{
		{"sqrt(A / pi)", ""},
		{"pow(x, 3) + pow(y, 3)", ""},
		{"pow(x, 3) + pow(pow(2 * 5, 4), 3)", ""},
		{"5 / 9 * (F - 32)", ""},
		{"-1 + -x", ""},
		{"-1 - x", ""},
	}

	for _, test := range tests {
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := fmt.Sprint(expr)
		fmt.Printf("\n%s\n", test.expr)
		fmt.Printf("%s\n", got)

		// TODO: Test Automation
		// if got != test.want {
		// 	t.Errorf("%q.String() = %q, want: %q\n", test.expr, got, test.want)
		// 	fmt.Printf("%s\n", got)
		// } else {
		// 	fmt.Printf("%s\n", test.expr)
		// 	fmt.Printf("%s\n", got)
		// }
	}
}
