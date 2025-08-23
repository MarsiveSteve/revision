package piscine

func Capitalize(s string) string {
	result := []rune(s)
	n := len(result)
	inWord := false

	for i := 0; i < n; i++ {
		c := result[i]

		if isAlnum(c) {
			if !inWord {
				// Capitalize first letter of the word if it's lowercase letter
				if c >= 'a' && c <= 'z' {
					result[i] = c - ('a' - 'A')
				}
				inWord = true
			} else {
				// Lowercase other letters if uppercase
				if c >= 'A' && c <= 'Z' {
					result[i] = c + ('a' - 'A')
				}
			}
		} else {
			// Not alphanumeric — end of a word
			inWord = false
		}
	}
	return string(result)
}

func isAlnum(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
