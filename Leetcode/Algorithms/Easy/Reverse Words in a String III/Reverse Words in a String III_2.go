func reverseWords(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		words[i] = Reverse(words[i])
	}

	return strings.Join(words, " ")
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}