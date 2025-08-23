package piscine

func RevConcatAlternate(slice1, slice2 []int) []int {
	len1 := len(slice1)
	len2 := len(slice2)

	var result []int

	// Indexes to reverse from the end
	i := len1 - 1
	j := len2 - 1

	// If one slice is longer, start adding from it
	if len1 > len2 {
		for ; i >= len2; i-- {
			result = append(result, slice1[i])
		}
	} else if len2 > len1 {
		for ; j >= len1; j-- {
			result = append(result, slice2[j])
		}
	}

	// Now alternate when both slices have equal remaining elements
	for i >= 0 && j >= 0 {
		// If lengths were equal or now equal, start from slice1
		result = append(result, slice1[i])
		result = append(result, slice2[j])
		i--
		j--
	}

	return result
}
