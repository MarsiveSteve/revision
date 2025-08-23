package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n < 0 {
		// As per the instructions, negative numbers are excluded.
		return
	}
	// Special case: if n == 0, just print '0'
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Count digits frequencies: digits[0] through digits[9]
	var digits [10]int

	// Extract digits without converting to int64
	for n > 0 {
		digits[n%10]++
		n /= 10
	}

	// Print digits in ascending order according to frequency
	for digit := 0; digit <= 9; digit++ {
		for count := 0; count < digits[digit]; count++ {
			z01.PrintRune(rune(digit + '0'))
		}
	}
}
