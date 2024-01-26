func pivotIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if sum(nums, 0, i) == sum(nums, i+1, len(nums)) {
			return i
		}
	}

	return -1
}

func sum(nums []int, start, end int) int {
	result := 0

	for i := start; i < end; i++ {
		result += nums[i]
	}

	return result
}