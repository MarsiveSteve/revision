package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// os.Args[1:] contains all arguments except the program name (os.Args[0])
	for _, arg := range os.Args[1:] {
		for _, ch := range arg {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
