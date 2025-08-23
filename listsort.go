package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

// ListSort sorts the linked list in ascending order and returns the new head.
func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	// Split the list into two halves
	mid := getMiddle(l)
	nextToMid := mid.Next
	mid.Next = nil

	// Recursively sort the two halves
	left := ListSort(l)
	right := ListSort(nextToMid)

	// Merge the sorted halves
	return sortedMerge(left, right)
}

// Helper to merge two sorted linked lists
func sortedMerge(a, b *NodeI) *NodeI {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	var result *NodeI
	if a.Data <= b.Data {
		result = a
		result.Next = sortedMerge(a.Next, b)
	} else {
		result = b
		result.Next = sortedMerge(a, b.Next)
	}
	return result
}

// Helper to find the middle node of the linked list (using slow and fast pointers)
func getMiddle(head *NodeI) *NodeI {
	if head == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
