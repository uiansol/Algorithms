// Accepted  0 ms  2.7 MB

func isPalindrome(s string) bool {
	l := len(s)
	i, j := 0, l-1
	for i < j {
		for i < l && !isAlphanumeric(s[i]) {
			i++
		}

		for j >= 0 && !isAlphanumeric(s[j]) {
			j--
		}

		if i >= j {
			break
		}

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}

		i++
		j--
	}

	return true
}

func isAlphanumeric(b byte) bool {
	return ((b >= '0' && b <= '9') || (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z'))
}