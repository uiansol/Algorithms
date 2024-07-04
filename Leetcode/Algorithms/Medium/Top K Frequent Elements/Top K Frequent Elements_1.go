func topKFrequent(nums []int, k int) []int {
	table := make(map[int]int)

	for _, n := range nums {
		table[n]++
	}

	var res []int
	for k > 0 {
		max, next := -1, -1
		for k, v := range table {
			if v > max {
				max = v
				next = k
			}
		}
		res = append(res, next)
		delete(table, next)
		k--
	}

	return res
}