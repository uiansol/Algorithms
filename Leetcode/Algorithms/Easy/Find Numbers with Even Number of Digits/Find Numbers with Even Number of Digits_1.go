func findNumbers(nums []int) int {
	count := 0

	for _, v := range nums {
		n := strconv.Itoa(v)
		if len(n)%2 == 0 {
			count++
		}
	}

	return count
}