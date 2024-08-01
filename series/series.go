package series

func All(n int, s string) (series []string) {
	for i := 0; i <= len(s)-n; i++ {
		series = append(series, s[i:i+n])
	}

	return
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return
	}

	return UnsafeFirst(n, s), true
}
