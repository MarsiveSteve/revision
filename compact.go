package piscine

// Compact removes zero values from the slice and returns the number of non-zero values.
func Compact(ptr *[]string) int {
	// Get the slice
	slice := *ptr

	// This will keep track of where the next non-zero value should be placed
	newLength := 0

	// Iterate through the slice
	for i := 0; i < len(slice); i++ {
		// If the element is non-zero (non-empty string), move it to the new position
		if slice[i] != "" {
			slice[newLength] = slice[i]
			newLength++
		}
	}

	// Slice the array to remove zero-value elements (elements beyond newLength)
	*ptr = slice[:newLength]

	// Return the number of non-zero elements
	return newLength
}
