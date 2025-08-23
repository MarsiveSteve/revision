package piscine

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	for player := 0; player < 4; player++ {
		// Print "Player "
		z01.PrintRune('P')
		z01.PrintRune('l')
		z01.PrintRune('a')
		z01.PrintRune('y')
		z01.PrintRune('e')
		z01.PrintRune('r')
		z01.PrintRune(' ')

		// Print player number (1 to 4)
		if player == 0 {
			z01.PrintRune('1')
		} else if player == 1 {
			z01.PrintRune('2')
		} else if player == 2 {
			z01.PrintRune('3')
		} else if player == 3 {
			z01.PrintRune('4')
		}

		// Print ": "
		z01.PrintRune(':')
		z01.PrintRune(' ')

		for i := 0; i < 3; i++ {
			card := deck[player*3+i]

			// Handle two-digit cards 10, 11, 12
			if card == 10 {
				z01.PrintRune('1')
				z01.PrintRune('0')
			} else if card == 11 {
				z01.PrintRune('1')
				z01.PrintRune('1')
			} else if card == 12 {
				z01.PrintRune('1')
				z01.PrintRune('2')
			} else {
				// Single digit cards 1-9
				switch card {
				case 1:
					z01.PrintRune('1')
				case 2:
					z01.PrintRune('2')
				case 3:
					z01.PrintRune('3')
				case 4:
					z01.PrintRune('4')
				case 5:
					z01.PrintRune('5')
				case 6:
					z01.PrintRune('6')
				case 7:
					z01.PrintRune('7')
				case 8:
					z01.PrintRune('8')
				case 9:
					z01.PrintRune('9')
				}
			}

			if i < 2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}

		z01.PrintRune('\n')
	}
}
