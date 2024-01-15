// Time Limit Exceeded

func maxSubArray(nums []int) int {
	l := len(nums)
	maxSubarray := nums[0]

	for i := 0; i < l; i++ {
		currentSubarray := 0
		for j := i; j < l; j++ {
			currentSubarray += nums[j]
			maxSubarray = max(maxSubarray, currentSubarray)
		}
	}

	return maxSubarray
}