func dominantIndex(nums []int) int {
	first, second, idx := -1, -1, 0

	for i, v := range nums {
		if v > first {
			second = first
			first = v
			idx = i
		} else if v > second {
			second = v
		}
	}

	if first >= 2*second {
		return idx
	}

	return -1
}