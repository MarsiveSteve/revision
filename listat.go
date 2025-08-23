package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}

	current := l
	for i := 0; current != nil; i++ {
		if i == pos {
			return current
		}
		current = current.Next
	}
	return nil
}
