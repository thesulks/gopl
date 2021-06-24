package intset

import "testing"

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
