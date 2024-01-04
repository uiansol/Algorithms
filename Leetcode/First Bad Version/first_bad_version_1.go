// Accepted  0 ms  1.92 MB

func firstBadVersion(n int) int {
	left, right, mid := 1, n, 0

	for left < right {
		mid = left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}