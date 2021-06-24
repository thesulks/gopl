package intset

import "testing"

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
