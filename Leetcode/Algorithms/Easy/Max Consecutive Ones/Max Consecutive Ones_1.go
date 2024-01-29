func findMaxConsecutiveOnes(nums []int) int {
	maxNums, currentNums := 0, 0

	for _, v := range nums {
		if v == 1 {
			currentNums++
			if currentNums > maxNums {
				maxNums = currentNums
			}
		} else {
			currentNums = 0
		}
	}

	return maxNums
}