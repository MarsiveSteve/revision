package piscine

func TrimAtoi(s string) int {
	sign := 1
	result := 0
	signFound := false
	digitFound := false

	for _, c := range s {
		if c == '-' && !digitFound && !signFound {
			sign = -1
			signFound = true
		} else if c == '+' && !digitFound && !signFound {
			signFound = true
		} else if c >= '0' && c <= '9' {
			digitFound = true
			result = result*10 + int(c-'0')
		}
		// ignore any other characters
	}

	if !digitFound {
		return 0
	}
	return sign * result
}
