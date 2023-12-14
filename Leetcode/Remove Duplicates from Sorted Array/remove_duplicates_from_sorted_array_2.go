// Accepted  9 ms  4.5 MB
// This solution is also fast and consumes less memory. More noticed for bigger inputs.

func removeDuplicates(nums []int) int {
	size := len(nums)
	n := -200
	head := 0

	for i := 0; i < size; i++ {
		if nums[i] > n {
			n = nums[i]
			nums[head] = n
			head++
		}
	}

	return head
}