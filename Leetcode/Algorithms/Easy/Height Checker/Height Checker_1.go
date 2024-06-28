func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)

	count := 0
	for i, _ := range heights {
		if heights[i] != expected[i] {
			count++
		}
	}

	return count
}