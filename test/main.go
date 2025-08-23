package main

import (
	"fmt"
)

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

func Atoi(s string) (int, error) {
	result := 0
	sign := 1
	start := 0

	// Handle optional '+' or '-' sign
	if len(s) > 0 {
		if s[0] == '-' {
			sign = -1
			start = 1
		} else if s[0] == '+' {
			start = 1
		}
	}

	for i := start; i < len(s); i++ {
		ch := s[i]

		// Check if character is a digit
		if ch < '0' || ch > '9' {
			return 0, fmt.Errorf("invalid character: %c", ch)
		}

		// Convert character to digit value
		digit := int(ch - '0')
		result = result*10 + digit
	}

	return result * sign, nil
}
