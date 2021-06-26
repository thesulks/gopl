package intset

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint
}

// const uintSize = 32 << (^uint(0) >> 32 & 1)
const bitSize = 32 << (^uint(0) >> 63)

func (s *IntSet) Len() int {
	var len int
	for _, word := range s.words {
		for i := 0; i < bitSize; i++ {
			if word&(1<<uint(i)) != 0 {
				len++
			}
		}
	}
	return len
}

func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		for j := 0; j < bitSize; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, i*bitSize+j)
			}
		}
	}
	return elems
}

func (s *IntSet) Remove(x int) {
	word, bit := x/bitSize, uint(x%bitSize)
	if word >= len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	copy := &IntSet{make([]uint, len(s.words))}
	for i, word := range s.words {
		copy.words[i] = word
	}
	return copy
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/bitSize, uint(x%bitSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/bitSize, uint(x%bitSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(list ...int) {
	for _, x := range list {
		s.Add(x)
	}
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i := range s.words {
		if i >= len(t.words) {
			s.words = s.words[:i]
			return
		}
		s.words[i] &= t.words[i]
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range s.words {
		if i >= len(t.words) {
			return
		}
		s.words[i] &^= t.words[i]
	}
}

func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bitSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
