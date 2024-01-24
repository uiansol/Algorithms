func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1

	for i <= j {
		for j >= 0 && nums[j] == val {
			j--
		}

		if i > j {
			break
		}

		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
		}

		i++
	}

	return i
}