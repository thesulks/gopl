package intset

import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	type test struct {
		input IntSet
		want  int
	}

	input0 := IntSet{}
	test0 := test{input0, 0}

	input1 := IntSet{}
	input1.Add(1)
	test1 := test{input1, 1}

	input4 := IntSet{}
	for i := 0; i < 4; i++ {
		input4.Add(i)
	}
	test4 := test{input4, 4}

	tests := []test{test0, test1, test4}
	for _, test := range tests {
		if got := test.input.Len(); got != test.want {
			t.Errorf("(*IntSet).Len(%v) = %d, want %d", &test.input, got, test.want)
		}
	}
}

func TestRemove(t *testing.T) {
	set := IntSet{}
	set.Add(1)
	set.Add(2)
	set.Add(3)

	set.Remove(4)
	if set.Len() != 3 || set.String() != "{1 2 3}" {
		t.Errorf("{1 2 3}.Remove(4) != {1 2 3}")
	}

	set.Remove(2)
	if set.Len() != 2 || set.String() != "{1 3}" {
		t.Errorf("{1 2 3}.Remove(2) != {1 3}")
	}
}

func TestClear(t *testing.T) {
	set := IntSet{}
	set.Add(1)
	set.Add(2)
	set.Add(3)

	set.Clear()
	if set.Len() != 0 || set.String() != "{}" {
		t.Errorf("{1 2 3}.Clear() != {}")
	}
}

func TestCopy(t *testing.T) {
	set := IntSet{}
	set.Add(1)
	set.Add(2)
	set.Add(3)

	copy := set.Copy()
	if copy.Len() != set.Len() || copy.String() != set.String() {
		t.Errorf("{1 2 3}.Copy() != {1 2 3}")
	}
}

func TestAddAdd(t *testing.T) {
	s1, s2 := IntSet{}, IntSet{}
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2.AddAll(1, 2, 3)
	if s1.String() != s2.String() {
		t.Errorf("%v != %v", &s1, &s2)
	}

	s2.AddAll()
	if s1.String() != s2.String() {
		t.Errorf("%v != %v", &s1, &s2)
	}

	s1.Add(4)
	s2.AddAll(3, 3, 3, 3, 3, 1, 2, 1, 4)
	if s1.String() != s2.String() {
		t.Errorf("%v != %v", &s1, &s2)
	}

}

func TestIntersectWith(t *testing.T) {
	a, b := IntSet{}, IntSet{}
	a.AddAll(1, 999, 1829)
	b.AddAll(0, 1, 23, 92)

	s1 := a.Copy()
	s2 := b.Copy()
	s1.IntersectWith(s2)
	want := &IntSet{}
	want.AddAll(1)
	if s1.String() != want.String() {
		t.Errorf("want %v, got %v", want, s1)
	}
}

func TestDifferenceWith(t *testing.T) {
	a, b := IntSet{}, IntSet{}
	a.AddAll(1, 999, 1829)
	b.AddAll(0, 1, 23, 92)

	s1 := a.Copy()
	s2 := b.Copy()
	s1.DifferenceWith(s2)
	want := &IntSet{}
	want.AddAll(999, 1829)
	if s1.String() != want.String() {
		t.Errorf("want %v, got %v", want, s1)
	}
}

func TestSymmetricDifferenceWith(t *testing.T) {
	a, b := IntSet{}, IntSet{}
	a.AddAll(1, 999, 1829)
	b.AddAll(0, 1, 23, 92)

	s1 := a.Copy()
	s2 := b.Copy()
	s1.SymmetricDifferenceWith(s2)
	want := &IntSet{}
	want.AddAll(0, 23, 92, 999, 1829)
	if s1.String() != want.String() {
		t.Errorf("want %v, got %v", want, s1)
	}
}
func TestElems(t *testing.T) {
	s := IntSet{}
	s.AddAll(1, 0, 1, 0, 34, 444, 52, 1233, 999991)

	if fmt.Sprintf("%v", s.Elems()) != "[0 1 34 52 444 1233 999991]" {
		t.Errorf("s.Elems(): %v\n", s.Elems())
	}
}
