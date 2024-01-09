// Accepted  0 ms  2.24 MB

func rob(nums []int) int {
	l := len(nums)
	var getFrom func(i int) int
	memo := make(map[int]int)

	getFrom = func(i int) int {
		if i >= l {
			return 0
		}

		if v, ok := memo[i]; ok {
			return v
		}

		res := max(nums[i]+getFrom(i+2), getFrom(i+1))
		memo[i] = res

		return res
	}

	return getFrom(0)
}