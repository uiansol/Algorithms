// Accepted  0 ms  2 MB

func rob(nums []int) int {
	l := len(nums)

	if l == 0 {
		return 0
	}

	value, valuePlusOne, valuePlusTwo := nums[l-1], nums[l-1], 0

	for i := l - 2; i >= 0; i-- {
		value = max(valuePlusTwo+nums[i], valuePlusOne)
		valuePlusTwo = valuePlusOne
		valuePlusOne = value
	}

	return value
}