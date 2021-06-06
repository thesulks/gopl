package rotate

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func RotateLeftByReverse(s []int, n int) []int {
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
	return s
}

func RotateLeft(s []int, n int) []int {
	buf := make([]int, n)
	copy(buf, s[:n])
	copy(s, s[n:])
	copy(s[len(s)-n:], buf)
	return s
}

func RotateRight(s []int, n int) []int {
	buf := make([]int, n)
	copy(buf, s[len(s)-n:])
	copy(s[n:], s)
	copy(s, buf)
	return s
}
