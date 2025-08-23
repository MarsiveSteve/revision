package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // skip the program name

	if len(args) == 0 {
		return // no args? do nothing, exit safely
	}

	for i := len(args) - 1; i >= 0; i-- {
		for _, ch := range args[i] {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
