package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	count := 0
	for _, r := range s {
		count++
		if count == n {
			return r
		}
	}
	return 0
}
