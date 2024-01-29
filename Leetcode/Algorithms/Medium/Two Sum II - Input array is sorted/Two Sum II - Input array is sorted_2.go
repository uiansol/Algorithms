func twoSum(numbers []int, target int) []int {
	result := make([]int, 2)
	searched := make(map[int]int)

	for i, v := range numbers {
		if idx, ok := searched[target-v]; ok {
			result[0], result[1] = idx+1, i+1
			return result
		}
		searched[v] = i
	}

	return result
}