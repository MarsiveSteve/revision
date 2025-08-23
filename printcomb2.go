package piscine

import "github.com/01-edu/z01"

func printDigit(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else if n == 1 {
		z01.PrintRune('1')
	} else if n == 2 {
		z01.PrintRune('2')
	} else if n == 3 {
		z01.PrintRune('3')
	} else if n == 4 {
		z01.PrintRune('4')
	} else if n == 5 {
		z01.PrintRune('5')
	} else if n == 6 {
		z01.PrintRune('6')
	} else if n == 7 {
		z01.PrintRune('7')
	} else if n == 8 {
		z01.PrintRune('8')
	} else if n == 9 {
		z01.PrintRune('9')
	}
}

func PrintComb2() {
	for a := 0; a <= 98; a++ {
		for b := a + 1; b <= 99; b++ {
			printDigit(a / 10)
			printDigit(a % 10)
			z01.PrintRune(' ')
			printDigit(b / 10)
			printDigit(b % 10)

			if !(a == 98 && b == 99) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
