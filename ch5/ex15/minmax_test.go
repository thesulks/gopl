package minmax

import "testing"

func TestMin(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4}, 1},
		{[]int{1}, 1},
		{[]int{1, 2, 3, -4}, -4},
	}

	for _, test := range tests {
		if got := min(test.input[0], test.input[1:]...); got != test.want {
			t.Errorf("min(%v) = %d, want %d", test.input, got, test.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4}, 4},
		{[]int{1}, 1},
		{[]int{1, 2, 3, -4}, 3},
	}

	for _, test := range tests {
		if got := max(test.input[0], test.input[1:]...); got != test.want {
			t.Errorf("min(%v) = %d, want %d", test.input, got, test.want)
		}
	}
}
