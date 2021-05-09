package popcount_test

import (
	popcount "gopl/ch2/ex5"
	"testing"
)

// 10100
// 10011
// -----
// 10000
// 01111
// -----
// 00000

func BitCountByClearing(x uint64) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

func TestPopCount(t *testing.T) {
	var input uint64 = 0xFFFFFFFFFFFFFFFF
	want := 64
	got := popcount.PopCount(input)
	if got != 64 {
		t.Errorf("PopCount(%v) = %v, want: %v", input, got, want)
	}
}

func TestBitCountByClearing(t *testing.T) {
	var input uint64 = 0xFFFFFFFFFFFFFFFF
	want := 64
	got := BitCountByClearing(input)
	if got != 64 {
		t.Errorf("BitCountByClearing(%v) = %v, want: %v", input, got, want)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCountByClearing(0x1234567890ABCDEF)
	}
}
