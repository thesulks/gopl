package palindrome

import "testing"

type RuneSlice []rune

func (s RuneSlice) Len() int           { return len(s) }
func (s RuneSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s RuneSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func TestIsPalindrome(t *testing.T) {
	rs := RuneSlice("a")
	if !IsPalindrome(rs) {
		t.Errorf("IsPalindrome reported %v is not palindrome", rs)
	}
	rs = RuneSlice("가나가")
	if !IsPalindrome(rs) {
		t.Errorf("IsPalindrome reported %v is not palindrome", rs)
	}
	rs = RuneSlice("가나다")
	if IsPalindrome(rs) {
		t.Errorf("IsPalindrome reported %v is palindrome", rs)
	}

	is := IntSlice([]int{1, 2, 3})
	if IsPalindrome(is) {
		t.Errorf("IsPalindrome reported %v is palindrome", is)
	}
	is = IntSlice([]int{1, 2, 1})
	if !IsPalindrome(is) {
		t.Errorf("IsPalindrome reported %v is not palindrome", is)
	}
	is = IntSlice([]int{})
	if !IsPalindrome(is) {
		t.Errorf("IsPalindrome reported %v is not palindrome", is)
	}
}
