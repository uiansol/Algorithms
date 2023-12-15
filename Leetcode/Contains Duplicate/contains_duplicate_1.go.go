// Accepted  91 ms  14.1 MB

func containsDuplicate(nums []int) bool {
	qnt := make(map[int]int)

	for _, v := range nums {
		qnt[v]++
	}

	if len(qnt) < len(nums) {
		return true
	}
	return false
}