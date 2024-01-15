// Accepted  327 ms  14.11 MB

func isPalindrome(head *ListNode) bool {
	var vals []int
	current := head

	for current != nil {
		vals = append(vals, current.Val)
		current = current.Next
	}

	l := len(vals)
	for i := 0; i < l/2; i++ {
		if vals[i] != vals[l-1-i] {
			return false
		}
	}

	return true
}