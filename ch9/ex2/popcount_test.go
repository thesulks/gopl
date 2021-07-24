package popcount_test

import (
	popcount "gopl/ch2/ex5"
	"testing"
)

func TestPopCount(t *testing.T) {
	var input uint64 = 0xFFFFFFFFFFFFFFFF
	want := 64
	got := popcount.PopCount(input)
	if got != 64 {
		t.Errorf("PopCount(%v) = %v, want: %v", input, got, want)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}
