package piscine

func IterativeFactorial(nb int) int64 {
	if nb < 0 {
		return 0
	}
	result := int64(1)
	for i := 1; i <= nb; i++ {
		result *= int64(i)
	}
	return result
}
