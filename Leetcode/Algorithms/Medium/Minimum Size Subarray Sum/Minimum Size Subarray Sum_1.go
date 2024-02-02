func minSubArrayLen(target int, nums []int) int {
	l, ans, left, sum := len(nums), math.MaxInt, 0, 0

	for i := 0; i < l; i++ {
		sum += nums[i]
		for sum >= target {
			ans = min(ans, i+1-left)
			sum -= nums[left]
			left++
		}
	}

	if ans != math.MaxInt {
		return ans
	}
	return 0
}