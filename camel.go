package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	// Validate input: only letters allowed
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if (ch < 'a' || ch > 'z') && (ch < 'A' || ch > 'Z') {
			return s
		}
	}

	// Check for consecutive uppercase letters
	for i := 1; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' && s[i-1] >= 'A' && s[i-1] <= 'Z' {
			return s
		}
	}

	// Check if last character is uppercase
	last := len(s)-1
	if last >= 'A' && last <= 'Z' {
		return s
	}

	result := ""

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'A' && ch <= 'Z' {
			// add underscore if not the first character
			if i != 0 {
				result += "_"
			}
		
		}
		result += string(ch)
	}

	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
