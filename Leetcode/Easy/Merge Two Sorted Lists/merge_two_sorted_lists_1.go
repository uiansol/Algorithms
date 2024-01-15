// Accepted  0 ms  2.4 MB

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	prehead := &ListNode{Val: -1, Next: nil}

	current := prehead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	return prehead.Next
}