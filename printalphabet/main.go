package main

import "github.com/01-edu/z01"

func main() {
	for chacha := 'a'; chacha <= 'z'; chacha++ {
		z01.PrintRune(chacha)
	}
	z01.PrintRune('\n')
}
