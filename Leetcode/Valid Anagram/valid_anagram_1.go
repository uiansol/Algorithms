// Accepted  4 ms  2.8 MB

func isAnagram(s string, t string) bool {
	mapS, mapT := make(map[byte]int), make(map[byte]int)

	if len(s) != len(t) {
		return false
	}

	for i := range s {
		mapS[s[i]]++
		mapT[t[i]]++
	}

	for k, v := range mapS {
		if v != mapT[k] {
			return false
		}
	}

	return true
}