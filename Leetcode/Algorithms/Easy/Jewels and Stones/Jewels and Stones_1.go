func numJewelsInStones(jewels string, stones string) int {
	jewelsTable := make(map[rune]bool)

	for _, c := range jewels {
		jewelsTable[c] = true
	}

	counter := 0
	for _, s := range stones {
		if _, ok := jewelsTable[s]; ok {
			counter++
		}
	}

	return counter
}