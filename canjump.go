package piscine

func CanJump(input []uint) bool {
	length := len(input)

	if length == 0 {
		return false
	}

	if length == 1 {
		return true
	}

	visited := make([]bool, length)
	var canReach func(index int) bool

	canReach = func(index int) bool {
		if index == length-1 {
			return true
		}

		if index >= length || visited[index] {
			return false
		}

		visited[index] = true

		jump := int(input[index])
		nextIndex := index + jump

		if nextIndex < length {
			return canReach(nextIndex)
		}

		return false
	}

	return canReach(0)
}
