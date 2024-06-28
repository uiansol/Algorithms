func findMaxConsecutiveOnes(nums []int) int {
	left, right, max, zeros := 0, 0, 0, 0

	for right < len(nums) {
		if nums[right] == 0 {
			zeros++
		}

		for zeros == 2 {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}

		if right-left+1 > max {
			max = right - left + 1
		}
		right++
	}

	return max
}