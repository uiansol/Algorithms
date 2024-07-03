func lengthOfLongestSubstring(s string) int {
	max, i, j := 0, 0, 0
	table := make(map[byte]bool)

	for j < len(s) {
		if _, ok := table[s[j]]; !ok {
			table[s[j]] = true
			if j-i+1 > max {
				max = j - i + 1
			}
			j++
		} else {
			delete(table, s[i])
			i++
		}
	}

	return max
}