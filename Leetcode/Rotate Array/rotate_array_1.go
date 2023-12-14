// Accepted  29 ms  8 MB

func rotate(nums []int, k int) {
	length := len(nums)
	k = k % length
	numsAux := append(nums[length-k:length], nums[0:length-k]...)

	for i := 0; i < length; i++ {
		nums[i] = numsAux[i]
	}

}