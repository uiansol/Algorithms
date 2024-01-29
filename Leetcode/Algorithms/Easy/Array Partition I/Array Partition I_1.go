func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0

	for i, j := 0, 1; j < len(nums); i, j = i+2, j+2 {
		sum += min(nums[i], nums[j])
	}

	return sum
}