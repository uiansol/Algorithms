// Accepted  0 ms  4.7 MB
// This solution is fast but consumes more memory

func removeDuplicates(nums []int) int {
	var orderedNums []int

	result := 1
	n := nums[0]
	orderedNums = append(orderedNums, n)
	for _, v := range nums {
		if v > n {
			n = v
			orderedNums = append(orderedNums, n)
			result++
		}
	}

	for i, v := range orderedNums {
		nums[i] = v
	}

	return result
}