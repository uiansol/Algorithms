// Accepted  1 ms  2 MB

func rob(nums []int) int {
	l := len(nums)

	if l == 0 {
		return 0
	}

	var maxRobbedAmount []int = make([]int, l+1)
	maxRobbedAmount[l], maxRobbedAmount[l-1] = 0, nums[l-1]

	for i := l - 2; i >= 0; i-- {
		maxRobbedAmount[i] = max(maxRobbedAmount[i+2]+nums[i], maxRobbedAmount[i+1])
	}

	return maxRobbedAmount[0]
}