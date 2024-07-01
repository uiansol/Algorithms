func findRestaurant(list1 []string, list2 []string) []string {
	table := make(map[string]int)

	for i, s := range list1 {
		table[s] = i
	}

	common := make(map[string]int)
	for i, s := range list2 {
		if v, ok := table[s]; ok {
			common[s] = v + i
		}
	}

	min := 2000
	for _, val := range common {
		if val < min {
			min = val
		}
	}

	var result []string
	for key, val := range common {
		if val == min {
			result = append(result, key)
		}
	}

	return result
}