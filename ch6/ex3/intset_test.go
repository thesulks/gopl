package intset

import "testing"

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
