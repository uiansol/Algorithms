// Accepted  95 ms  9 MB

func maxSubArray(nums []int) int {
	l := len(nums)
	currentSubArray, maxSubArray := nums[0], nums[0]

	for i := 1; i < l; i++ {
		currentSubArray = max(nums[i], currentSubArray+nums[i])
		maxSubArray = max(maxSubArray, currentSubArray)
	}

	return maxSubArray
}