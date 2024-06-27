func findNumbers(nums []int) int {
	count := 0

	for _, v := range nums {
		if (v >= 10 && v <= 99) || (v >= 1000 && v <= 9999) || v == 100000 {
			count++
		}
	}

	return count
}