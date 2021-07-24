package popcount

import "sync"

var initPcOnce sync.Once
var pc [256]byte // pc[i] is the population count of i.

func popCount(x byte) byte {
	initPcOnce.Do(initPc)
	return pc[x]
}

func initPc() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(popCount(byte(x>>(0*8))) +
		popCount(byte(x>>(1*8))) +
		popCount(byte(x>>(2*8))) +
		popCount(byte(x>>(3*8))) +
		popCount(byte(x>>(4*8))) +
		popCount(byte(x>>(5*8))) +
		popCount(byte(x>>(6*8))) +
		popCount(byte(x>>(7*8))))
}
