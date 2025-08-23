package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	asc := true
	desc := true

	for i := 0; i < len(a)-1; i++ {
		comp := f(a[i], a[i+1])
		if comp > 0 {
			asc = false
		}
		if comp < 0 {
			desc = false
		}
	}
	return asc || desc
}
