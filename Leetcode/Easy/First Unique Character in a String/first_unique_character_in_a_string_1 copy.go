// Accepted  25 ms  5.54 MB

func firstUniqChar(s string) int {
	letters := make(map[rune]int)

	for _, v := range s {
		letters[v]++
	}

	for i, v := range s {
		if letters[v] == 1 {
			return i
		}
	}

	return -1
}