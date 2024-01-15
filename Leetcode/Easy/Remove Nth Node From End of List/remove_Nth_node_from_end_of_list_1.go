// Accepted  0 ms  2.23 MB

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodeAux := &ListNode{Val: 0, Next: head}
	first, second := nodeAux, nodeAux

	for i := 1; i <= n+1; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next

	return nodeAux.Next
}