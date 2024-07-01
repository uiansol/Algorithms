func groupAnagrams(strs []string) [][]string {
	table := make(map[string][]string)
	for _, str := range strs {
		jmap, _ := json.Marshal(makeMap(str))
		key := string(jmap)
		table[key] = append(table[key], str)
	}

	var result [][]string
	for _, value := range table {
		result = append(result, value)
	}

	return result
}

func makeMap(str string) map[rune]int {
	table := make(map[rune]int)

	for _, s := range str {
		table[s]++
	}

	return table
}