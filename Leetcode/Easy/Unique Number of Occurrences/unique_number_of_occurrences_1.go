// Accepted  2 ms  2.46 MB

func uniqueOccurrences(arr []int) bool {
	freq := make(map[int]int)
	occurrences := make(map[int]int)

	for _, v := range arr {
		freq[v]++
	}

	for _, v := range freq {
		occurrences[v]++
	}

	for _, v := range occurrences {
		if v > 1 {
			return false
		}
	}

	return true
}