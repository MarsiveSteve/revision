package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}

func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil || l.Head == nil {
		return
	}

	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	if l.Head == nil {
		l.Tail = nil
		return
	}

	curr := l.Head
	for curr.Next != nil {
		if curr.Next.Data == data_ref {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	l.Tail = l.Head
	for l.Tail != nil && l.Tail.Next != nil {
		l.Tail = l.Tail.Next
	}
}
