package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		// Handle negative numbers safely without int64
		if nbr == -nbr { // edge case for MinInt (in Go int32 or int64)
			// can't negate directly, print digits recursively
			printNbrBaseUnsigned(uint(-nbr), base)
			return
		}
		nbr = -nbr
	}
	printNbrBaseUnsigned(uint(nbr), base)
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, c := range base {
		if c == '+' || c == '-' || seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
}

func printNbrBaseUnsigned(n uint, base string) {
	baseLen := uint(len(base))
	if n >= baseLen {
		printNbrBaseUnsigned(n/baseLen, base)
	}
	z01.PrintRune(rune(base[n%baseLen]))
}
