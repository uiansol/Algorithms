func containsNearbyDuplicate(nums []int, k int) bool {
	table := make(map[int][]int)

	for i, n := range nums {
		table[n] = append(table[n], i)
	}

	for _, list := range table {
		if len(list) >= 2 {
			for i := 0; i < len(list)-1; i++ {
				if list[i+1]-list[i] <= k {
					return true
				}
			}
		}
	}

	return false
}