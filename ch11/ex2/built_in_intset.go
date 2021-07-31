package intset

import (
	"bytes"
	"fmt"
	"sort"
)

type IntSetByMap map[int]bool

func (s *IntSetByMap) Has(x int) bool {
	return (*s)[x]
}

func (s *IntSetByMap) Add(x int) {
	(*s)[x] = true
}

func (s *IntSetByMap) UnionWith(t *IntSetByMap) {
	for e := range *t {
		(*s)[e] = true
	}
}

func (s *IntSetByMap) String() string {
	var elems []int
	for e := range *s {
		elems = append(elems, e)
	}
	sort.Ints(elems)

	var buf bytes.Buffer
	buf.WriteByte('{')
	for _, e := range elems {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", e)
	}
	buf.WriteByte('}')
	return buf.String()
}
