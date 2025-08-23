package piscine

func JumpOver(str string) string {
	// If the string is empty, return a newline
	if len(str) == 0 {
		return "\n"
	}

	// Initialize an empty string to hold the result
	result := ""

	// Loop through the string and pick every third character
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	// If result is empty (meaning no third character), return newline
	if len(result) == 0 {
		return "\n"
	}

	return result + "\n"
}
