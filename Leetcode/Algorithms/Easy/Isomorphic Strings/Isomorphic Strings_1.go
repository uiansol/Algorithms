func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	table := make(map[byte]byte)
	used := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		val1, ok1 := table[s[i]]
		_, ok2 := used[t[i]]

		if !ok1 && !ok2 {
			table[s[i]] = t[i]
			used[t[i]] = true
		} else if val1 != t[i] {
			return false
		}
	}

	return true
}