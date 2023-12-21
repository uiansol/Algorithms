// Accepted  11 ms  6.9 MB

func moveZeroes(nums []int) {
	aux := nums[:]

	j := 0
	for _, v := range aux {
		if v != 0 {
			nums[j] = v
			j++
		}
	}

	for j < len(nums) {
		nums[j] = 0
		j++
	}
}