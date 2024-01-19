// Accepted  16 ms  6.4 MB

func moveZeroes(nums []int) {
	length := len(nums)
	i, j := 0, 0

	for i < length {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
		i++
	}
}