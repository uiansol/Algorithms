func findDisappearedNumbers(nums []int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		nums[Abs(nums[i])-1] = -Abs(nums[Abs(nums[i])-1])
	}

	for i, v := range nums {
		if v > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}