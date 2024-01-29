func removeElement(nums []int, val int) int {
	j, l := 0, len(nums)

	for i := 0; i < l; i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}