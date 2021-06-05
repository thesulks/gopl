package anagram

import (
	"sort"
	"strings"
)

func IsAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	seqa := strings.Split(a, "")
	seqb := strings.Split(b, "")
	sort.Strings(seqa)
	sort.Strings(seqb)

	if len(seqa) != len(seqb) {
		return false
	}

	for i := 0; i < len(seqa); i++ {
		if seqa[i] != seqb[i] {
			return false
		}
	}

	return true
}
