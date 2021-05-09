package popcount_test

import (
	popcount "gopl/ch2/ex4"
	"testing"
)

func BitCountByShift(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		count += int(x & 1)
		x >>= 1
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

func TestBitCountByShift(t *testing.T) {
	var input uint64 = 0xFFFFFFFFFFFFFFFF
	want := 64
	got := BitCountByShift(input)
	if got != 64 {
		t.Errorf("BitCountByShift(%v) = %v, want: %v", input, got, want)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCountByShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCountByShift(0x1234567890ABCDEF)
	}
}
