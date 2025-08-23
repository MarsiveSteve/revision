package piscine

func ListReverse(l *List) {
	if l == nil || l.Head == nil {
		return
	}

	var prev *NodeL = nil
	current := l.Head
	l.Tail = l.Head // After reverse, old Head becomes Tail

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev
}
