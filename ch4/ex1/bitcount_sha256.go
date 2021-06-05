package bitcount

import "crypto/sha256"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// TODO: refactor to pointer version
func DiffBitCountSha256(a, b [sha256.Size]byte) int {
	var count int
	for i := range a {
		count += int(pc[a[i]^b[i]])
	}
	return count
}
