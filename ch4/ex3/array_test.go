package array

import "testing"

func TestReverse(t *testing.T) {
	input := [N]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	want := [N]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	Reverse(&input)
	if input != want {
		t.Errorf("result: %v, want: %v", input, want)
	}
}
