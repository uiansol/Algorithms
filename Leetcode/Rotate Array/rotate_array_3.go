// Accepted  20 ms  7.65 MB

func rotate(nums []int, k int) {
	length := len(nums)
	k = k % length
	var start, count int = 0, 0

	for count < length {
		current := start
		prev := nums[start]
		for {
			nextIdx := (current + k) % length
			nums[nextIdx], prev = prev, nums[nextIdx]
			current = nextIdx
			count++

			if start == current {
				break
			}
		}

		start++
	}
}