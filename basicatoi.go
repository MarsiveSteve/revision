package piscine

func BasicAtoi(s string) int {
	runeStr := []rune(s)
	var num int
	for _, n := range runeStr {
		if n >= '0' && n <= '9' {
			num = num*10 + int(n-48)
		} else {
			return 0
		}
	}
	return num
}
