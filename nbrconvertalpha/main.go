package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	upper := false
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	if len(args) == 0 {
		// No arguments to process, print nothing and exit
		return
	}

	for _, arg := range args {
		n := parseInt(arg)
		if n < 1 || n > 26 {
			z01.PrintRune(' ')
			continue
		}

		var ch rune
		if upper {
			ch = rune('A' + n - 1)
		} else {
			ch = rune('a' + n - 1)
		}
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

// parseInt parses a string into an int, returns -1 if invalid
func parseInt(s string) int {
	if s == "" {
		return -1
	}
	n := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return -1
		}
		n = n*10 + int(c-'0')
	}
	return n
}
