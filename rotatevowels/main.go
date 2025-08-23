package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	// If no arguments, print just a newline and exit
	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	// Collect vowels from all args
	vowels := []rune{}
	for _, arg := range args {
		for _, ch := range arg {
			if isVowel(ch) {
				vowels = append(vowels, ch)
			}
		}
	}

	// If no vowels, print arguments as-is
	if len(vowels) == 0 {
		printArgs(args)
		return
	}

	// Reverse vowels slice to mirror vowels
	for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
		vowels[i], vowels[j] = vowels[j], vowels[i]
	}

	// Replace vowels in arguments with mirrored vowels
	idx := 0
	for i, arg := range args {
		runes := []rune(arg)
		for j, ch := range runes {
			if isVowel(ch) {
				runes[j] = vowels[idx]
				idx++
			}
		}
		args[i] = string(runes)
	}

	// Print all arguments separated by space, no trailing space
	for i, arg := range args {
		for _, ch := range arg {
			z01.PrintRune(ch)
		}
		if i != len(args)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func printArgs(args []string) {
	for i, arg := range args {
		for _, ch := range arg {
			z01.PrintRune(ch)
		}
		if i != len(args)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
