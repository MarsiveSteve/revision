package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // get all args except program name
	n := len(args)

	// Bubble sort (ASCII order)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if isGreater(args[j], args[j+1]) {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	// Print sorted arguments
	for _, arg := range args {
		for _, ch := range arg {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}

// Compare two strings lexicographically in ASCII order
func isGreater(a, b string) bool {
	lenA, lenB := len(a), len(b)
	minLen := lenA
	if lenB < minLen {
		minLen = lenB
	}

	for i := 0; i < minLen; i++ {
		if a[i] > b[i] {
			return true
		} else if a[i] < b[i] {
			return false
		}
	}
	return lenA > lenB
}
