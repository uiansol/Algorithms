// Accepted  0 ms  2 MB

func strStr(haystack string, needle string) int {
	h, n, result := len(haystack), len(needle), -1

	for idx := range haystack {
		result = idx
		i, j := idx, 0

		for i < h && j < n {
			if haystack[i] != needle[j] {
				break
				result = -1
			}
			i++
			j++
		}

		if i == result+n {
			return result
		}
	}

	return -1
}