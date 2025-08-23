package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Convert nbr from baseFrom to int
	num := AtoiBase(nbr, baseFrom)
	// Convert int to baseTo string
	return ItoaBase(num, baseTo)
}

// Converts a string number in given base to int (same as previous AtoiBase)
func AtoiBase(s string, base string) int {
	baseLen := len(base)
	value := 0

	for _, c := range s {
		digit := indexOf(base, c)
		if digit == -1 {
			return 0 // invalid digit
		}
		value = value*baseLen + digit
	}
	return value
}

// Converts an int to string in given base (like PrintNbrBase but returns string)
func ItoaBase(n int, base string) string {
	if n == 0 {
		return string(base[0])
	}

	baseLen := len(base)
	result := ""

	for n > 0 {
		digit := n % baseLen
		result = string(base[digit]) + result
		n = n / baseLen
	}
	return result
}

func indexOf(s string, c rune) int {
	for i, r := range s {
		if r == c {
			return i
		}
	}
	return -1
}
