func thirdMax(nums []int) int {
	first, second, third := math.MinInt, math.MinInt, math.MinInt

	for _, v := range nums {
		if v > first {
			third = second
			second = first
			first = v
		} else if v > second && v != first {
			third = second
			second = v
		} else if v > third && v != first && v != second {
			third = v
		}
	}

	if third == math.MinInt {
		return first
	}

	return third
}