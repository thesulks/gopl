package minmax

func min(val int, vals ...int) int {
	result := val
	for _, v := range vals {
		if result > v {
			result = v
		}
	}
	return result
}

func max(val int, vals ...int) int {
	result := val
	for _, v := range vals {
		if result < v {
			result = v
		}
	}
	return result
}
