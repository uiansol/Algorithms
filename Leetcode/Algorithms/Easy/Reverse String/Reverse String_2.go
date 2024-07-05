func reverseString(s []byte) {
	var helper func(start int, end int, ls []byte)

	helper = func(start int, end int, ls []byte) {
		if start > end {
			return
		}

		ls[start], ls[end] = ls[end], ls[start]
		helper(start+1, end-1, ls)
	}

	helper(0, len(s)-1, s)
}