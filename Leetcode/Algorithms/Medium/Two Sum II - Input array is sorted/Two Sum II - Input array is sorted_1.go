func twoSum(numbers []int, target int) []int {
	result := make([]int, 2)
	l := len(numbers)

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if numbers[i]+numbers[j] == target {
				result[0], result[1] = i+1, j+1
				return result
			}
		}
	}

	return result
}