package bitcount

import (
	"crypto/sha256"
	"testing"
)

func TestDiffBitCountSha256(t *testing.T) {
	zeros := [sha256.Size]byte{}
	ones := [sha256.Size]byte{}
	for i := range zeros {
		ones[i] = 0b00000001
	}
	twos := [sha256.Size]byte{}
	for i := range zeros {
		twos[i] = 0b00000010
	}
	hashDummy := sha256.Sum256([]byte("dummy"))

	type input struct {
		a [sha256.Size]byte
		b [sha256.Size]byte
	}

	tests := []struct {
		in   input
		want int
	}{
		{input{hashDummy, hashDummy}, 0},
		{input{zeros, zeros}, 0},
		{input{ones, ones}, 0},
		{input{twos, twos}, 0},

		{input{zeros, ones}, 32},
		{input{zeros, twos}, 32},
		{input{ones, twos}, 64},
	}

	for _, test := range tests {
		if got := DiffBitCountSha256(test.in.a, test.in.b); got != test.want {
			t.Errorf("DiffBitCountSha256(%x) = %v, want %v", test.in, got, test.want)
		} else {
			t.Logf("DiffBitCountSha256(%x) = %v", test.in, got)
		}
	}
}
