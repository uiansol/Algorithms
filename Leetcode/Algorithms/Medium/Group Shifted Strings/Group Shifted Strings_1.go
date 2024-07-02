func groupStrings(strings []string) [][]string {
	table := make(map[string][]string)

	for _, str := range strings {
		key := transformToStartA(str)
		table[key] = append(table[key], str)
	}

	var result [][]string
	for _, value := range table {
		result = append(result, value)
	}

	return result
}

func transformToStartA(str string) string {
	dist := str[0] - 'a'

	result := ""
	for i := 0; i < len(str); i++ {
		newChar := str[i] - dist
		if newChar < 97 {
			newChar += 26
		}
		result += string(newChar)
	}

	return result
}