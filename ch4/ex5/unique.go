package slice

func Unique(s []string) []string {
	if len(s) == 0 {
		return s
	}

	i := 0
	for _, v := range s {
		if s[i] != v {
			i++
			s[i] = v
		}
	}
	return s[:i+1]
}
