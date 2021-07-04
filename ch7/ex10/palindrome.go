package palindrome

import "sort"

func IsPalindrome(s sort.Interface) bool {
	equal := func(i, j int) bool {
		return !s.Less(i, j) && !s.Less(j, i)
	}

	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !equal(i, j) {
			return false
		}
	}
	return true
}
