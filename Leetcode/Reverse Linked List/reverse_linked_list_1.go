// Accepted  0 ms  2.54 MB

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	prev = nil
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	head = prev

	return head
}