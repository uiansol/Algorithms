// Accepted  2 ms  2.4 MB

func longestCommonPrefix(strs []string) string {
	l := 200
	prefix := ""

	for _, v := range strs {
		if len(v) < l {
			l = len(v)
		}
	}

	for i := 0; i < l; i++ {
		c := strs[0][i]
		for _, s := range strs {
			if s[i] != c {
				return prefix
			}
		}
		prefix += string(c)
	}

	return prefix
}