package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	// Print help if no args OR if -h or --help anywhere
	if len(args) == 0 {
		printHelp()
		return
	}
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			printHelp()
			return
		}
	}

	var insert string
	order := false
	var input string

	// Parse flags and input string
	for _, arg := range args {
		if arg == "-o" || arg == "--order" {
			order = true
		} else if len(arg) >= 9 && arg[:9] == "--insert=" {
			insert = arg[9:]
		} else if len(arg) >= 3 && arg[:3] == "-i=" {
			insert = arg[3:]
		} else {
			// Assume it's the string to process
			input = arg
		}
	}

	result := input + insert

	if order {
		result = sortASCII(result)
	}

	for _, ch := range result {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func printHelp() {
	writeLine("--insert")
	writeLine("  -i")
	writeLine("	 This flag inserts the string into the string passed as argument.")
	writeLine("--order")
	writeLine("  -o")
	writeLine("	 This flag will behave like a boolean, if it is called it will order the argument.")
}

func writeLine(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func sortASCII(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes)
}
