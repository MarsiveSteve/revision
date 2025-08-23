package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	// Find the last '/' to get the program name
	lastSlash := -1
	for i := 0; i < len(name); i++ {
		if name[i] == '/' {
			lastSlash = i
		}
	}
	if lastSlash != -1 {
		name = name[lastSlash+1:]
	}

	// Print each character using z01.PrintRune
	for _, ch := range name {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
