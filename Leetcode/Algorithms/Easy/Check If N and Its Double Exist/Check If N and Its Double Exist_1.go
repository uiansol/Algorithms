func checkIfExist(arr []int) bool {
	table := make(map[int]bool)

	for _, v := range arr {
		_, ok1 := table[v*2]
		if ok1 {
			return true
		}

		if v%2 == 0 {
			_, ok2 := table[v/2]
			if ok2 {
				return true
			}
		}

		table[v] = true
	}

	return false
}