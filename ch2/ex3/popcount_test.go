package popcount_test

import (
	popcount "gopl/ch2/ex3"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByLoop(0x1234567890ABCDEF)
	}
}
