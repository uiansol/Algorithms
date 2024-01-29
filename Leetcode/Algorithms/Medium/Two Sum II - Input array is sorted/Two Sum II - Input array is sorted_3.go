func twoSum(numbers []int, target int) []int {
	head, tail := 0, len(numbers)-1

	for head < tail {
		sum := numbers[head] + numbers[tail]

		if sum == target {
			return []int{head + 1, tail + 1}
		} else if sum < target {
			head++
		} else {
			tail--
		}
	}

	return []int{0, 0}
}