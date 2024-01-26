func pivotIndex(nums []int) int {
	s, leftSum := 0, 0

	for _, v := range nums {
		s += v
	}

	for i, v := range nums {
		if leftSum == (s - leftSum - v) {
			return i
		}
		leftSum += v
	}

	return -1
}