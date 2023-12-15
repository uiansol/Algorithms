// Accepted  8 ms  6.1 MB

func singleNumber(nums []int) int {
	alone := 0

	for _, v := range nums {
		alone ^= v
	}

	return alone
}