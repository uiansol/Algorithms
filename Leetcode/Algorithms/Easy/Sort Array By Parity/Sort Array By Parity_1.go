func sortArrayByParity(nums []int) []int {
	i, j, l := 0, len(nums)-1, len(nums)

	for {
		for nums[i]%2 == 0 && i < l-1 {
			i++
		}
		for nums[j]%2 != 0 && j > 0 {
			j--
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	return nums
}