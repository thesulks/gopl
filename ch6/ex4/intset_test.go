package intset

import (
	"fmt"
	"testing"
)

func TestElems(t *testing.T) {
	s := IntSet{}
	s.AddAll(1, 0, 1, 0, 34, 444, 52, 1233, 999991)

	if fmt.Sprintf("%v", s.Elems()) != "[0 1 34 52 444 1233 999991]" {
		t.Errorf("s.Elems(): %v\n", s.Elems())
	}
}
