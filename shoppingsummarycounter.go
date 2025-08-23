package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)

	i := 0
	n := len(str)
	start := 0

	for i <= n {
		if i == n || str[i] == ' ' {
			item := str[start:i]
			result[item]++
			i++
			start = i
		} else {
			i++
		}
	}
	return result
}
