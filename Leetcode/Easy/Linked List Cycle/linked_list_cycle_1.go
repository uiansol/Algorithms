// Accepted  7 ms  4.5 MB

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	current := head
	for current.Val != 10000000 {
		if current.Next == nil {
			return false
		}
		current.Val = 10000000
		current = current.Next
	}

	return true
}