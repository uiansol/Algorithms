// Accepted  12 ms  6.18 MB

func majorityElement(nums []int) int {
	count, candidate := 0, 0

	for _, v := range nums {
		if count == 0 {
			candidate = v
		}

		if candidate == v {
			count++
		} else {
			count--
		}
	}

	return candidate
}