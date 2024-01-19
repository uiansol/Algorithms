// Accepted  9 ms  4.44 MB
// This solution is also fast and consumes less memory. More noticed for bigger inputs.

func removeDuplicates(nums []int) int {
	size := len(nums)
	head := 1

	for i := 1; i < size; i++ {
		if nums[i] != nums[i-1] {
			nums[head] = nums[i]
			head++
		}
	}

	return head
}