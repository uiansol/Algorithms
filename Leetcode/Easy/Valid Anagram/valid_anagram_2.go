// Accepted  5 ms  2.79 MB

func isAnagram(s string, t string) bool {
	counter := make(map[byte]int)

	if len(s) != len(t) {
		return false
	}

	for i := range s {
		counter[s[i]]++
		counter[t[i]]--
	}

	for _, v := range counter {
		if v != 0 {
			return false
		}
	}

	return true
}