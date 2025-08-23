package main

import "github.com/01-edu/z01"

func main() {
	for abriton := 'z'; abriton >= 'a'; abriton-- {
		z01.PrintRune(abriton)
	}
	z01.PrintRune('\n')
}
