package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	sign := 1
	start := 0
	result := 0

	// Handle optional sign
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		result = result*10 + int(s[i]-'0')
	}

	return sign * result
}
