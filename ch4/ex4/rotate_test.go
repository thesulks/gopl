package rotate

import (
	"math/rand"
	"testing"
)

type input struct {
	s []int
	n int
}

const N = 1000000

var dummy []int

func init() {
	for i := 0; i < N; i++ {
		dummy = append(dummy, rand.Int())
	}
}

func equal(a, b []int) bool {
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

func TestEqual(t *testing.T) {
	tests := []struct {
		a    []int
		b    []int
		want bool
	}{
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}, true},
		{[]int{}, []int{}, true},

		{[]int{1, 2, 3, 4}, []int{1, 2, 3}, false},
		{[]int{1}, []int{1, 2, 3}, false},

		{nil, nil, true},
		{nil, []int{}, true},
		{[]int{}, nil, true},
	}

	for _, test := range tests {
		if got := equal(test.a, test.b); got != test.want {
			t.Errorf("equal(%v, %v) = %t, want %t", test.a, test.b, got, test.want)
		}
	}
}

func TestRotateLeft(t *testing.T) {
	tests := []struct {
		input input
		want  []int
	}{
		{input{[]int{1, 2, 3, 4}, 3}, []int{4, 1, 2, 3}},
		{input{[]int{1, 2, 3, 4}, 0}, []int{1, 2, 3, 4}},
		{input{[]int{1, 2}, 2}, []int{1, 2}},
	}

	for _, test := range tests {
		before := make([]int, len(test.input.s))
		copy(before, test.input.s)
		if got := RotateLeft(test.input.s, test.input.n); !equal(got, test.want) {
			t.Errorf("RotateLeft(%v, %v) = %v, want %v", before, test.input.n, got, test.want)
		}
	}
}

func TestRotateLeftByReverse(t *testing.T) {
	tests := []struct {
		input input
		want  []int
	}{
		{input{[]int{1, 2, 3, 4}, 3}, []int{4, 1, 2, 3}},
		{input{[]int{1, 2, 3, 4}, 0}, []int{1, 2, 3, 4}},
		{input{[]int{1, 2}, 2}, []int{1, 2}},
	}

	for _, test := range tests {
		before := make([]int, len(test.input.s))
		copy(before, test.input.s)
		if got := RotateLeftByReverse(test.input.s, test.input.n); !equal(got, test.want) {
			t.Errorf("RotateLeftByReverse(%v, %v) = %v, want %v", before, test.input.n, got, test.want)
		}
	}
}

func TestRotateRight(t *testing.T) {
	tests := []struct {
		input input
		want  []int
	}{
		{input{[]int{1, 2, 3, 4}, 3}, []int{2, 3, 4, 1}},
		{input{[]int{1, 2, 3, 4}, 0}, []int{1, 2, 3, 4}},
		{input{[]int{1, 2}, 2}, []int{1, 2}},
	}

	for _, test := range tests {
		before := make([]int, len(test.input.s))
		copy(before, test.input.s)
		if got := RotateRight(test.input.s, test.input.n); !equal(got, test.want) {
			t.Errorf("RotateRight(%v, %v) = %v, want %v", before, test.input.n, got, test.want)
		}
	}
}

func BenchmarkRotateLeft(b *testing.B) {
	s := make([]int, N)
	copy(s, dummy)
	n := N / 2
	for i := 0; i < b.N; i++ {
		s = RotateLeft(s, n)
	}
}

func BenchmarkRotateLeftByReverse(b *testing.B) {
	s := make([]int, N)
	copy(s, dummy)
	n := N / 2
	for i := 0; i < b.N; i++ {
		s = RotateLeftByReverse(s, n)
	}
}
