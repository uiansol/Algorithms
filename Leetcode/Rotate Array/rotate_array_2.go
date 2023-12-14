// Accepted  20 ms  7.5 MB

func rotate(nums []int, k int) {
	length := len(nums)
	k = k % length
	aux := 0

	for i := 0; i < length/2; i++ {
		aux = nums[i]
		nums[i] = nums[length-1-i]
		nums[length-1-i] = aux
	}

	for i := 0; i < k/2; i++ {
		aux = nums[i]
		nums[i] = nums[k-1-i]
		nums[k-1-i] = aux
	}

	for i := k; i < k+((length-k)/2); i++ {
		aux = nums[i]
		nums[i] = nums[length-1-i+k]
		nums[length-1-i+k] = aux
	}
}