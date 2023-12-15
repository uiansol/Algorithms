// Accepted  20 ms  6.7 MB

func singleNumber(nums []int) int {
	qnt := make(map[int]int)

	for _, v := range nums {
		qnt[v]++
	}

	for k, v := range qnt {
		if v == 1 {
			return k
		}
	}

	return 0
}